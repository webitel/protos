#!/usr/bin/env node
// Generate a Postman collection from the OpenAPI 3.0 spec and push it to Postman,
// grouped into one folder per protos module (folderStrategy: Tags — the spec is
// already re-tagged by module in build-public-spec.mjs).
//
// Create-or-update:
//   - POSTMAN_COLLECTION_UID set  -> PUT (replace) that collection
//   - unset                       -> POST create in POSTMAN_WORKSPACE_ID,
//                                     then print the uid to save as a repo var
//
// NOTE: the collection is fully regenerated each run — manual edits made in
// Postman (saved examples, tweaks) are overwritten. It is a derived artifact.
//
// Env (required): POSTMAN_API_KEY, SPEC_FILE
// Env (one of):   POSTMAN_COLLECTION_UID (update) | POSTMAN_WORKSPACE_ID (create)
// Env (optional): POSTMAN_COLLECTION_NAME (default "Webitel API")

import { readFileSync } from "node:fs";
import Converter from "openapi-to-postmanv2";

const BASE = "https://api.getpostman.com";
const KEY = req("POSTMAN_API_KEY");
const SPEC_FILE = req("SPEC_FILE");
const COLLECTION_UID = process.env.POSTMAN_COLLECTION_UID || "";
const WORKSPACE_ID = process.env.POSTMAN_WORKSPACE_ID || "";
const NAME = process.env.POSTMAN_COLLECTION_NAME || "Webitel API";

function req(name) {
  const v = process.env[name];
  if (!v) fail(`Missing required env ${name}`);
  return v;
}
function fail(msg) {
  console.error(`postman-collection: ${msg}`);
  process.exit(1);
}

const openapi = JSON.parse(readFileSync(SPEC_FILE, "utf8"));

const collection = await new Promise((resolve, reject) => {
  Converter.convert(
    { type: "json", data: openapi },
    {
      folderStrategy: "Tags", // one folder per module tag
      requestParametersResolution: "Example",
      exampleParametersResolution: "Example",
      collapseFolders: false,
    },
    (err, res) => {
      if (err) return reject(err);
      if (!res || !res.result) return reject(new Error(res && res.reason));
      resolve(res.output[0].data);
    },
  );
}).catch((e) => fail(`OpenAPI → collection conversion failed: ${e.message}`));

collection.info = collection.info || {};
collection.info.name = NAME;

async function api(method, path, body) {
  const res = await fetch(`${BASE}${path}`, {
    method,
    headers: { "X-Api-Key": KEY, "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });
  const text = await res.text();
  let json;
  try {
    json = text ? JSON.parse(text) : {};
  } catch {
    json = { raw: text };
  }
  if (!res.ok) fail(`${method} ${path} -> ${res.status}\n${JSON.stringify(json, null, 2)}`);
  return json;
}

if (COLLECTION_UID) {
  await api("PUT", `/collections/${COLLECTION_UID}`, { collection });
  console.error(`Updated Postman collection ${COLLECTION_UID}.`);
} else {
  if (!WORKSPACE_ID) {
    fail("Set POSTMAN_COLLECTION_UID to update, or POSTMAN_WORKSPACE_ID to create.");
  }
  const created = await api("POST", `/collections?workspace=${WORKSPACE_ID}`, { collection });
  const uid = created.collection?.uid || created.collection?.id;
  console.error(`Created Postman collection: ${uid}`);
  console.error(`>> Save this as the repo variable POSTMAN_COLLECTION_UID for in-place updates.`);
}
