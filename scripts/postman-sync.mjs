#!/usr/bin/env node
// Push an OpenAPI 3.0 spec to Postman Spec Hub via the Specs API.
//
// Create-or-update:
//   - POSTMAN_SPEC_ID set  -> PATCH the spec's ROOT file with the new content
//   - POSTMAN_SPEC_ID unset -> POST a new spec, then print the id to save
//                              as the POSTMAN_SPEC_ID repo variable
//
// Env (required):
//   POSTMAN_API_KEY        Postman API key (X-Api-Key)
//   SPEC_FILE              path to the OpenAPI 3.0 JSON (e.g. swagger/api.public.json)
// Env (one of):
//   POSTMAN_SPEC_ID        existing spec id -> update in place
//   POSTMAN_WORKSPACE_ID   target workspace  -> create a new spec
// Env (optional):
//   POSTMAN_SPEC_NAME      name for a newly created spec (default: "Webitel API")
//
// Docs: Postman Specs API (Spec Hub). Base https://api.getpostman.com,
// auth via X-Api-Key. A spec file update accepts exactly one of
// { content | name | type } and files must be <= 10 MB.

import { readFileSync } from "node:fs";

const BASE = "https://api.getpostman.com";
const KEY = req("POSTMAN_API_KEY");
const SPEC_FILE = req("SPEC_FILE");
const SPEC_ID = process.env.POSTMAN_SPEC_ID || "";
const WORKSPACE_ID = process.env.POSTMAN_WORKSPACE_ID || "";
const SPEC_NAME = process.env.POSTMAN_SPEC_NAME || "Webitel API";

const content = readFileSync(SPEC_FILE, "utf8");
if (Buffer.byteLength(content) > 10 * 1024 * 1024) {
  fail(`${SPEC_FILE} exceeds Postman's 10 MB per-file limit.`);
}

function req(name) {
  const v = process.env[name];
  if (!v) fail(`Missing required env ${name}`);
  return v;
}
function fail(msg) {
  console.error(`postman-sync: ${msg}`);
  process.exit(1);
}
async function api(method, path, body) {
  const res = await fetch(`${BASE}${path}`, {
    method,
    headers: { "X-Api-Key": KEY, "Content-Type": "application/json" },
    body: body ? JSON.stringify(body) : undefined,
  });
  const text = await res.text();
  let json;
  try {
    json = text ? JSON.parse(text) : {};
  } catch {
    json = { raw: text };
  }
  if (!res.ok) {
    fail(`${method} ${path} -> ${res.status}\n${JSON.stringify(json, null, 2)}`);
  }
  return json;
}

if (SPEC_ID) {
  // Update: find the ROOT file, then PATCH its content.
  const files = await api("GET", `/specs/${SPEC_ID}/files`);
  const list = files.files || files.data || files;
  const root = Array.isArray(list)
    ? list.find((f) => (f.type || "").toUpperCase() === "ROOT") || list[0]
    : null;
  if (!root) fail(`No files found for spec ${SPEC_ID}`);
  const filePath = encodeURIComponent(root.path || root.name);
  await api("PATCH", `/specs/${SPEC_ID}/files/${filePath}`, { content });
  console.error(`Updated Postman spec ${SPEC_ID} (root file: ${root.path || root.name}).`);
} else {
  // Create a new spec in the target workspace.
  if (!WORKSPACE_ID) {
    fail("Set POSTMAN_SPEC_ID to update, or POSTMAN_WORKSPACE_ID to create a new spec.");
  }
  const created = await api("POST", `/specs?workspaceId=${WORKSPACE_ID}`, {
    name: SPEC_NAME,
    type: "OPENAPI:3.0",
    files: [{ path: "index.json", content }],
  });
  const id = created.id || created.spec?.id || created.data?.id;
  console.error(`Created Postman spec: ${id}`);
  console.error(
    `>> Save this as the repo variable POSTMAN_SPEC_ID so future runs update it in place.`,
  );
}
