### Essential Proto Annotations for Backend Team

#### 1. Field Descriptions
**Purpose:** Provide clear context for API consumers  
**Implementation:**
```protobuf
string name = 2 [
  (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
    description: "Unique human-readable identifier for the source"
  }
];
```
**Why:** Reduces support requests and clarifies field purpose in Swagger documentation.

#### 2. Enum Values
**Purpose:** Define all possible values  
**Implementation:**
```protobuf
enum SourceType {
  TYPE_UNSPECIFIED = 0;
  CALL = 1;
  CHAT = 2;
  SOCIAL_MEDIA = 3;
  EMAIL = 4;
  API = 5;
  MANUAL = 6;
}
```
**Why:** Critical for frontend validation and documentation - automatically appears in generated clients.

#### 3. Enum Comments
**Purpose:** Clarify non-obvious enum values  
**Implementation:**
```protobuf
enum SourceType {
  TYPE_UNSPECIFIED = 0;
  CALL = 1;  // Phone call source type
  CHAT = 2;  // Real-time chat source
  // ... other enums
}
```
**Why:** Provides context in IDE tooltips and reduces misinterpretation.


#### 4. Required Fields
**Purpose:** Mark mandatory fields  
**Implementation:**
```protobuf
string name = 2 [
  (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
    required: true
  }
];
```
**Why:** Generates automatic validation and prevents API misuse.

#### 5. Field Validation Rules
**Purpose:** Define data constraints  
**Implementation:**
```protobuf
string name = 2 [
  (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
    required: true,
    min_length: 1,  // Prevent empty strings
    max_length: 100,
    pattern: "^[a-zA-Z0-9_\\-\\s]+$"
  }
];
```
**Why:** Enables automatic client-side validation and documents data constraints.

#### 6. Default Values
**Purpose:** Specify omitted field behavior  
**Implementation:**
```protobuf
bool is_active = 25 [
  (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
    default: "true"  // Must be valid JSON string
  }
];
```