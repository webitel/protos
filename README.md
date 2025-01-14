# Webitel Protos

Use the [Buf CLI](https://buf.build/product/cli) to generate code and swagger schemes. The project currently supports Buf CLI v1, but it is planned to upgrade to v2. So far, the following modules support generation:
- engine
- chat

### Example:

- **Step 1**: `cd chat` - Navigating to the **`chat`** directory where your **Protobuf** files and **`buf.work.yaml`** configuration file are located.
- **Step 2**: `buf generate` -  Running **`buf generate`** will generate code based on the configurations defined in **`buf.gen.yaml`**.
- **Step 3**: After generation, the output files (e.g., Go code and Swagger definitions) are stored in specified directories.
