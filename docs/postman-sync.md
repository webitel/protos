# Webitel API in Postman

This workspace hosts the **Webitel HTTP API** as an OpenAPI 3.0 specification.
It is **generated automatically** from the gRPC/protobuf definitions in the
`protos` repository — do not edit it by hand in Postman, changes are overwritten
on the next sync.

## Getting started

1. Select an **Environment** (top-right): `dev` or `test`.
2. Make sure the environment defines:
   - `baseUrl` — e.g. `https://dev.webitel.com/api`
   - `apiKey` — your Webitel access token
3. Authorization is inherited from the collection: header
   `X-Webitel-Access: {{apiKey}}`.
4. Open any request and press **Send**.

| Environment | baseUrl |
|-------------|---------|
| dev  | `https://dev.webitel.com/api` |
| test | `https://test.webitel.me/api` |

## How the spec is produced

```
*.proto  ──►  make (buf)  ──►  swagger/*.swagger.json   (per service, Swagger 2.0)
             │
             ▼  swagger-mix.yml (CI)
        swagger/api.json           (merged, Swagger 2.0, 800+ ops)
             │
             ▼  postman-sync.yml (CI)  →  scripts/build-public-spec.mjs
        swagger/api.public.json    (OpenAPI 3.0, servers set from PUBLIC_API_HOST)
             │
             ▼  scripts/postman-sync.mjs (Spec Hub API, X-Api-Key)
        Postman spec  ← updated in place on every merge to main
```

Every merge to `main` that changes a swagger file regenerates the merged bundle,
converts it to OpenAPI 3.0, and pushes it to this Postman spec. No manual steps.

## Notes

- **Swagger 2.0 → OpenAPI 3.0** conversion (`swagger2openapi`) is required because
  Postman Spec Hub only accepts OpenAPI 3.x.
- **Colliding path templates** (e.g. `/im/gates/facebook/{id}` declared under two
  proto packages) are merged into one path — OpenAPI 3.0 cannot hold two paths that
  differ only by parameter name. No endpoint is lost functionally; duplicates are
  collapsed. The permanent fix is consistent `google.api.http` path params in the
  `.proto` files.
- The generated `swagger/api.public.json` is not committed (derived artifact).

## Configuration (GitHub repo `protos`)

| Kind | Name | Purpose |
|------|------|---------|
| Secret   | `POSTMAN_API_KEY`      | Postman API key (write access to the spec) |
| Variable | `POSTMAN_WORKSPACE_ID` | target workspace |
| Variable | `POSTMAN_SPEC_ID`      | spec updated in place (unset ⇒ a new spec is created) |
| Variable | `PUBLIC_API_HOST`      | host written into `servers` (e.g. `api.webitel.com`) |

To publish a new environment or change the host, update `PUBLIC_API_HOST` and
re-run the **Postman sync** workflow.
