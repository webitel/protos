#!/usr/bin/env node
// Build a public, OpenAPI 3.0 spec from the merged Swagger 2.0 bundle.
//
// Pipeline: swagger/api.json (Swagger 2.0, produced by the swagger-mix workflow)
//   -> convert 2.0 -> 3.0 (swagger2openapi)
//   -> set servers from PUBLIC_API_HOST / basePath
//   -> optionally keep only whitelisted tags (INCLUDE_TAGS)
//   -> write swagger/api.public.json (OpenAPI 3.0)
//
// Postman Spec Hub only accepts OpenAPI 3.x, hence the conversion.
//
// Env (all optional):
//   SOURCE_SPEC     input file           (default: swagger/api.json)
//   OUTPUT_SPEC     output file          (default: swagger/api.public.json)
//   PUBLIC_API_HOST public host          (default: source `host`, e.g. dev.webitel.com)
//   PUBLIC_BASE_PATH url base path        (default: source `basePath`, e.g. /api)
//   PUBLIC_TITLE    info.title override  (default: keep source title)
//   INCLUDE_TAGS    comma-separated tag whitelist; when set, paths whose
//                   operations carry none of these tags are dropped
//                   (default: keep everything)

import { readFileSync, writeFileSync, readdirSync } from "node:fs";
import { dirname, join } from "node:path";
import converter from "swagger2openapi";

const SOURCE = process.env.SOURCE_SPEC || "swagger/api.json";
const OUTPUT = process.env.OUTPUT_SPEC || "swagger/api.public.json";

const source = JSON.parse(readFileSync(SOURCE, "utf8"));

const host = process.env.PUBLIC_API_HOST || source.host || "";
const basePath = process.env.PUBLIC_BASE_PATH ?? source.basePath ?? "";
if (!host) {
  console.error(
    "No public host: set PUBLIC_API_HOST or ensure the source spec has `host`.",
  );
  process.exit(1);
}
const serverUrl = `https://${host}${basePath}`;

const options = {
  patch: true, // fix up minor spec issues instead of failing
  warnOnly: true, // don't throw on non-fatal conversion warnings
  refSiblings: "preserve",
};

const { openapi } = await converter.convertObj(source, options);

const HTTP_METHODS = new Set([
  "get",
  "put",
  "post",
  "delete",
  "patch",
  "options",
  "head",
  "trace",
]);

// --- re-tag by protos module ------------------------------------------------
// gRPC emits verbose per-service tags (webitel.im.api.gateway.v1.Message), and
// one service's endpoints get split across URL prefixes (/im, /v1). To make
// Postman group requests into one folder per protos module, re-tag every
// operation with the module it came from — derived from which per-service
// swagger/<module>.swagger.json file declares that path. This mirrors the repo
// layout (im, cases, storage, engine, wfm, ...).
function moduleNameFromFile(f) {
  const base = f.replace(/\.swagger\.json$/, "");
  return base === "im-gateway" ? "im" : base;
}
function retagByModule(spec, dir) {
  // (METHOD PATH) -> module, built from the individual swagger files.
  const map = new Map();
  const files = readdirSync(dir)
    .filter((f) => f.endsWith(".swagger.json"))
    .sort();
  for (const f of files) {
    let doc;
    try {
      doc = JSON.parse(readFileSync(join(dir, f), "utf8"));
    } catch {
      continue;
    }
    const mod = moduleNameFromFile(f);
    for (const [p, item] of Object.entries(doc.paths || {})) {
      for (const method of Object.keys(item)) {
        if (!HTTP_METHODS.has(method)) continue;
        const key = `${method} ${p}`;
        if (!map.has(key)) map.set(key, mod); // first file (sorted) wins
      }
    }
  }

  const used = new Set();
  let unmapped = 0;
  for (const [p, item] of Object.entries(spec.paths || {})) {
    for (const [method, op] of Object.entries(item)) {
      if (!HTTP_METHODS.has(method)) continue;
      const mod = map.get(`${method} ${p}`);
      if (mod) {
        op.tags = [mod];
        used.add(mod);
      } else {
        unmapped++;
      }
    }
  }
  // Stable, alphabetical folder order in Postman.
  spec.tags = [...used].sort().map((name) => ({ name }));
  console.error(
    `Re-tagged operations into ${used.size} module folders: ${[...used].sort().join(", ")}` +
      (unmapped ? ` (${unmapped} ops had no source file, kept original tags)` : ""),
  );
}

retagByModule(openapi, dirname(SOURCE));

// --- normalize colliding path templates ------------------------------------
// gRPC HTTP annotations across services use inconsistent path-param names, so
// e.g. `/devices/{id}` and `/devices/{device.id}` normalize to the same
// OpenAPI path template and collide (invalid OpenAPI 3.0). Merge each colliding
// group under one canonical path, picking the cleanest name per segment slot
// and remapping each operation's path parameters to match.
function normalizePathTemplates(spec) {
  const paths = spec.paths || {};
  const templateOf = (p) => p.replace(/\{[^}]+\}/g, "{}");
  const paramsOf = (p) => (p.match(/\{([^}]+)\}/g) || []).map((s) => s.slice(1, -1));
  // best name = no dot preferred, then shortest, then lexicographically first
  const better = (a, b) => {
    const dotA = a.includes(".") ? 1 : 0;
    const dotB = b.includes(".") ? 1 : 0;
    if (dotA !== dotB) return dotA < dotB ? a : b;
    if (a.length !== b.length) return a.length < b.length ? a : b;
    return a < b ? a : b;
  };

  const groups = {};
  for (const p of Object.keys(paths)) (groups[templateOf(p)] ||= []).push(p);

  const collisions = [];
  const conflicts = [];

  for (const [tpl, members] of Object.entries(groups)) {
    if (members.length < 2) continue;

    const slotCount = (tpl.match(/\{\}/g) || []).length;
    const canonical = Array.from({ length: slotCount }, () => "");
    for (const p of members) {
      paramsOf(p).forEach((name, i) => {
        canonical[i] = canonical[i] ? better(canonical[i], name) : name;
      });
    }
    let ci = 0;
    const canonicalPath = tpl.replace(/\{\}/g, () => `{${canonical[ci++]}}`);

    // Start empty: the canonical path (if it exists) is itself a member and
    // gets folded in by the loop below — pre-seeding would double-count it.
    const merged = {};
    for (const p of members) {
      const rename = new Map(paramsOf(p).map((old, i) => [old, canonical[i]]));
      const item = paths[p];
      for (const [method, op] of Object.entries(item)) {
        if (method === "parameters") {
          for (const prm of op) if (prm.in === "path") prm.name = rename.get(prm.name) ?? prm.name;
          merged.parameters = op;
          continue;
        }
        if (!HTTP_METHODS.has(method)) {
          merged[method] = op;
          continue;
        }
        for (const prm of op.parameters || [])
          if (prm.in === "path") prm.name = rename.get(prm.name) ?? prm.name;
        if (merged[method]) {
          conflicts.push(`${method.toUpperCase()} ${canonicalPath} (from ${p})`);
          continue; // path+method must be unique; keep the first, drop the rest
        }
        merged[method] = op;
      }
      if (p !== canonicalPath) delete paths[p];
    }
    paths[canonicalPath] = merged;
    collisions.push(`${canonicalPath}  <-  ${members.join(", ")}`);
  }

  if (collisions.length) {
    console.error(
      `Normalized ${collisions.length} colliding path template(s) (fix at proto HTTP annotations for a permanent solution):`,
    );
    for (const c of collisions) console.error(`  ${c}`);
  }
  if (conflicts.length) {
    console.error(
      `WARNING: dropped ${conflicts.length} operation(s) that collided on the same method+path:`,
    );
    for (const c of conflicts) console.error(`  ${c}`);
  }
}

normalizePathTemplates(openapi);

// --- servers ---------------------------------------------------------------
openapi.servers = [{ url: serverUrl, description: "Webitel public API" }];

// --- title -----------------------------------------------------------------
if (process.env.PUBLIC_TITLE) {
  openapi.info = openapi.info || {};
  openapi.info.title = process.env.PUBLIC_TITLE;
}

// --- optional tag whitelist ------------------------------------------------
const includeTags = (process.env.INCLUDE_TAGS || "")
  .split(",")
  .map((t) => t.trim())
  .filter(Boolean);

if (includeTags.length > 0) {
  const keep = new Set(includeTags);
  const methods = HTTP_METHODS;
  let droppedOps = 0;
  let droppedPaths = 0;

  for (const [pathKey, pathItem] of Object.entries(openapi.paths || {})) {
    for (const [method, op] of Object.entries(pathItem)) {
      if (!methods.has(method)) continue;
      const tags = op.tags || [];
      if (!tags.some((t) => keep.has(t))) {
        delete pathItem[method];
        droppedOps++;
      }
    }
    if (!Object.keys(pathItem).some((m) => methods.has(m))) {
      delete openapi.paths[pathKey];
      droppedPaths++;
    }
  }
  console.error(
    `Tag whitelist [${includeTags.join(", ")}]: dropped ${droppedOps} operations, ${droppedPaths} paths.`,
  );
}

writeFileSync(OUTPUT, JSON.stringify(openapi, null, 2));

const pathCount = Object.keys(openapi.paths || {}).length;
console.error(
  `Wrote ${OUTPUT}: OpenAPI ${openapi.openapi}, ${pathCount} paths, server ${serverUrl}`,
);
