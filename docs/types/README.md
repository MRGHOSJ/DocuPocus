# 📦 Package: `types`

[← Back to Overview](../README.md)

## 📄 File: `documentation.go`

> 📍 `types\documentation.go`

## 📑 Contents

- [🧱 Structs (4)](#-structs)

## 🧱 Structs

### `Documentation`

```go
type Documentation struct {
	Summary string json:"summary
	Parameters []Param json:"parameters
	Returns string json:"returns
	TimeComplexity string json:"time_complexity
	SpaceComplexity string json:"space_complexity
	UsageExample string json:"usage_example
	EdgeCases []string json:"edge_cases
}
```

_No documentation available._

---

### `Param`

```go
type Param struct {
	Name string json:"name
	Type string json:"type
	Description string json:"description
}
```

_No documentation available._

---

### `YAMLField`

```go
type YAMLField struct {
	Name string json:"name
	Type string json:"type
	Description string json:"description
}
```

_No documentation available._

---

### `YAMLDocumentation`

```go
type YAMLDocumentation struct {
	Summary string json:"summary
	Fields []YAMLField json:"fields
	Examples map[string] json:"examples,omitempty
	Defaults map[string] json:"defaults,omitempty
	Usage string json:"usage,omitempty
	BestPractices []string json:"best_practices,omitempty
}
```

_No documentation available._

---


---

## 📄 File: `documentation.go`

> 📍 `types\documentation.go`

## 📑 Contents

- [🧱 Structs (4)](#-structs)

## 🧱 Structs

### `Documentation`

```go
type Documentation struct {
	Summary string json:"summary
	Parameters []Param json:"parameters
	Returns string json:"returns
	TimeComplexity string json:"time_complexity
	SpaceComplexity string json:"space_complexity
	UsageExample string json:"usage_example
	EdgeCases []string json:"edge_cases
}
```

**Summary:** Struct for documenting code snippets in JSON format

**Parameters:**
- `Summary` (string): Brief functional description
- `Parameters` ([]Param): List of input parameters
- `Returns` (string): Output description
- `TimeComplexity` (string): Big-O time complexity
- `SpaceComplexity` (string): Memory usage analysis
- `UsageExample` (string): Example usage scenario
- `EdgeCases` ([]string): Potential edge cases

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
Used as a template for code documentation generation
```

**Edge Cases:**
- Missing required fields in JSON
- Invalid complexity notation


---

### `Param`

```go
type Param struct {
	Name string json:"name
	Type string json:"type
	Description string json:"description
}
```

**Summary:** Struct for parameter metadata in documentation

**Parameters:**
- `Name` (string): Parameter identifier
- `Type` (string): Data type
- `Description` (string): Purpose/usage

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
Embedded in Documentation struct for parameter details
```

**Edge Cases:**
- Type/name mismatch
- Empty description fields


---

### `YAMLField`

```go
type YAMLField struct {
	Name string json:"name
	Type string json:"type
	Description string json:"description
}
```

**Summary:** Defines a YAML field structure with name, type, and description

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
field := YAMLField{Name: "port", Type: "int", Description: "Server port"}
```

**Edge Cases:**
- Empty name or type fields may cause validation issues
- JSON tags must match struct fields exactly


---

### `YAMLDocumentation`

```go
type YAMLDocumentation struct {
	Summary string json:"summary
	Fields []YAMLField json:"fields
	Examples map[string] json:"examples,omitempty
	Defaults map[string] json:"defaults,omitempty
	Usage string json:"usage,omitempty
	BestPractices []string json:"best_practices,omitempty
}
```

**Summary:** Defines comprehensive YAML documentation structure

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
doc := YAMLDocumentation{Summary: "API config", Fields: []YAMLField{...}}
```

**Edge Cases:**
- Optional fields may be nil if omitempty is used
- Large examples/defaults maps may impact memory


---

