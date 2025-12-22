# Messaging Module Protobuf Contracts

This directory contains protobuf contract definitions for the **Messaging** module.
Each subdirectory represents a separate messaging-related microservice and contains
its own `.proto` files, Buf configuration, and generated artifacts (where applicable).

## Structure

Each service directory usually includes:

- `buf.yaml` / `buf.gen.yaml` – Buf module and generation configs.
- `*.proto` – Protobuf definitions for that microservice's public APIs and shared types.
- `go.mod` / `go.sum` (optional) – Go module configuration when the service has
  language-specific helpers or generated code committed.

## Conventions

To keep contracts consistent and easy to consume:

1. **One microservice per folder**  
   Group all protobuf files that belong to the same microservice under a single
   subdirectory. Shared cross-service types that are not messaging-specific
   should live in more general shared modules (for example, `general/` etc.).

2. **Clear package naming**  
   Use consistent protobuf package names that reflect the domain and service, e.g.:

   - `webitel.messaging.chat.v1` for chat-related APIs
   - `webitel.messaging.notifications.v1` for notifications (example)

   Align the package structure with Buf module configuration and Go package paths.

3. **Versioned APIs**  
   Prefer versioned packages (`...v1`, `...v1alpha1`) for public APIs. Breaking
   changes should be introduced in a new versioned package rather than mutating
   existing contracts.

4. **Buf and linting**  
   - Each service directory should have its own `buf.yaml` referencing the
     relevant module root.
   - Use the root-level `buf.work.yaml` to register new modules.
   - Follow the repository-wide lint and breaking-change policies.

5. **Annotations and documentation**  
   - Use comments on messages, fields, and services to describe behavior and
     semantics. These are used by downstream documentation and client generation
     tools.
   - Reuse common annotations from the `annotations/` module when appropriate
     (for HTTP bindings, OpenAPI, validation hints, etc.).

## Adding a New Messaging Microservice

When you introduce a new microservice under `messaging/`:

1. **Create the directory**

   Example (for a new `notifications` service):

   ```bash
   mkdir -p im/notifications
   ```

2. **Add Buf module configuration**

   Inside the new directory, create a `buf.yaml` (and `buf.gen.yaml` if needed)
   based on an existing messaging service (for example, `messaging/chat`). Adjust
   module name, roots, and generation targets. Name buf repository as messaging-[your-service-name].

3. **Define protobuf contracts**

   - Create one or more `.proto` files describing:
     - Public gRPC / HTTP APIs (`service` definitions)
     - Request/response messages
     - Domain events or stream messages, if applicable
   - Follow naming and versioning conventions from above.

4. **Register in `buf.work.yaml`**

   Add the new module path to the workspace-wide `buf.work.yaml` so it is
   included in linting and generation workflows.

5. **Run Buf checks and generation**

   From the repository root:

   ```bash
   buf lint
   buf generate
   ```

   Fix any lint issues before committing.

## Usage

These protobuf definitions are meant to be:

- **Source of truth** for public contracts between messaging microservices and
  other modules (cases, engine, storage, etc.).
- **Input to code generation** for gRPC/HTTP clients and servers, OpenAPI specs,
  and other derived artifacts.

Downstream services should **not** modify generated code directly. Contract
changes must be made in the corresponding `.proto` files and then regenerated
through the standard tooling (Buf/Makefile).

## Usage

These protobuf definitions are meant to be:

- **Source of truth** for public contracts between messaging microservices and
  other modules (cases, engine, storage, etc.).
- **Input to code generation** for gRPC/HTTP clients and servers, OpenAPI specs,
  and other derived artifacts.

Downstream services should **not** modify generated code directly. Contract
changes must be made in the corresponding `.proto` files and then regenerated
through the standard tooling (Buf/Makefile).


