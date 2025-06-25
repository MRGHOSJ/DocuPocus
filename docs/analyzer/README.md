# 📦 Package: `analyzer`

> 📍 `C:\Users\DELL\Documents\GitHub\DocuPocus\internal\analyzer\yaml_analyzer.go`

[← Back to Overview](../README.md)

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (3)](#-functions)

## 🧱 Structs

### `YAMLAnalyzer`

```go
type YAMLAnalyzer struct {
}
```

**Summary:** Empty struct for YAML analysis functionality

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := YAMLAnalyzer{}
```

**Edge Cases:**
- No functionality implemented yet
- Struct may need method implementations


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if YAML analyzer supports given project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
supported := analyzer.Supports("./my-project")
```

**Edge Cases:**
- Empty/non-existent directory path
- Insufficient permissions to check directory


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes project directory for YAML content

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Analysis result or error

**Complexity:**
- Time: O(n) where n is number of YAML files
- Space: O(m) where m is size of YAML content

**Example:**
```go
result, err := analyzer.Analyze("./configs")
```

**Edge Cases:**
- Malformed YAML files
- Large YAML files causing memory issues
- Permission errors on directory access


</details>

<details>
<summary><b><code>extractYAMLDescription(node *yaml.Node)</code></b></summary>

**Summary:** Extracts YAML description from a YAML node

**Parameters:**
- `node` (*yaml.Node): YAML node to extract description from

**Returns:** Extracted description as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
desc := extractYAMLDescription(yamlNode) // returns 'Example YAML description'
```

**Edge Cases:**
- Nil node input
- Node without a description field


</details>

