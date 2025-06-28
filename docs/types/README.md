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


---

## 📄 File: `config.go`

> 📍 `types\config.go`

## 📑 Contents

- [🧱 Structs (7)](#-structs)

## 🧱 Structs

### `GeneratorConfig`

```go
type GeneratorConfig struct {
	AIClient *ai.Client 
	OutputDir string 
	Project ProjectMeta 
}
```

**Summary:** Configuration for documentation generator

**Parameters:**
- `AIClient` (*ai.Client): AI client for enhanced generation
- `OutputDir` (string): Directory to save documentation
- `Project` (ProjectMeta): Project metadata

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(1) for struct allocation

**Example:**
```go
config := GeneratorConfig{client, "./docs", projectData}
```

**Edge Cases:**
- Nil AI client
- Non-writable output directory
- Incomplete project metadata


---

### `ProjectMeta`

```go
type ProjectMeta struct {
	Name string 
	Description string 
	RepoURL string 
	Features []Feature 
	TechStack []string 
	QuickStart []QuickStartBlock 
	BestPractices BestPractices 
}
```

**Summary:** Metadata about a software project

**Parameters:**
- `Name` (string): Project name
- `Description` (string): Brief project description
- `RepoURL` (string): Source repository URL
- `Features` ([]Feature): List of project features
- `TechStack` ([]string): Technologies used
- `QuickStart` ([]QuickStartBlock): Getting started guide
- `BestPractices` (BestPractices): Recommended practices

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(n) where n is size of slices

**Example:**
```go
meta := ProjectMeta{Name: "MyProject", ...}
```

**Edge Cases:**
- Empty required fields (name, description)
- Invalid repository URL format
- Empty slices vs nil slices


---

### `Feature`

```go
type Feature struct {
	Title string 
	Description string 
}
```

**Summary:** Defines a Feature struct with title and description fields.

**Parameters:**
- `Title` (string): Title of the feature
- `Description` (string): Description of the feature

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
feature := Feature{Title: "AI", Description: "Advanced AI capabilities"}
```

**Edge Cases:**
- Empty strings for Title or Description
- Non-string values assigned


---

### `QuickStartBlock`

```go
type QuickStartBlock struct {
	Title string 
	Shell string 
	Command string 
}
```

**Summary:** Defines a QuickStartBlock struct for shell command instructions.

**Parameters:**
- `Title` (string): Title of the quick start block
- `Shell` (string): Shell type (e.g., bash, zsh)
- `Command` (string): Command to execute

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
block := QuickStartBlock{Title: "Install", Shell: "bash", Command: "npm install"}
```

**Edge Cases:**
- Empty Command field
- Invalid shell type specified


---

### `BestPractices`

```go
type BestPractices struct {
	Do []string 
	Dont []string 
}
```

**Summary:** Defines a BestPractices struct with Do and Dont lists.

**Parameters:**
- `Do` ([]string): List of recommended practices
- `Dont` ([]string): List of practices to avoid

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(n)

**Example:**
```go
bp := BestPractices{Do: ["use interfaces"], Dont: ["ignore errors"]}
```

**Edge Cases:**
- Empty Do/Dont lists
- Duplicate entries in lists


---

### `AICodeRequest`

```go
type AICodeRequest struct {
	Input string 
	Language string 
	Target *aiTypes.Documentation 
}
```

**Summary:** Defines a struct for AI code request with input, language, and target documentation.

**Parameters:**
- `Input` (string): Input code or text
- `Language` (string): Programming language of the input
- `Target` (*aiTypes.Documentation): Pointer to target documentation type

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
req := AICodeRequest{Input: "func() {}", Language: "go", Target: &doc}
```

**Edge Cases:**
- Nil Target pointer
- Empty Input or Language strings


---

### `AIYAMLRequest`

```go
type AIYAMLRequest struct {
	Input string 
	Language string 
	Target *aiTypes.YAMLDocumentation 
}
```

**Summary:** Struct for AI YAML documentation request

**Parameters:**
- `Input` (string): Input text for documentation
- `Language` (string): Programming language of the input
- `Target` (*aiTypes.YAMLDocumentation): Pointer to target YAML documentation object

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
req := AIYAMLRequest{Input: "func()", Language: "go", Target: &doc}
```

**Edge Cases:**
- Nil Target pointer
- Empty Input string


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


---

## 📄 File: `config.go`

> 📍 `types\config.go`

## 📑 Contents

- [🧱 Structs (7)](#-structs)

## 🧱 Structs

### `GeneratorConfig`

```go
type GeneratorConfig struct {
	AIClient *ai.Client 
	OutputDir string 
	Project ProjectMeta 
}
```

**Summary:** Configuration for documentation generator

**Parameters:**
- `AIClient` (*ai.Client): AI client for enhanced generation
- `OutputDir` (string): Directory to save documentation
- `Project` (ProjectMeta): Project metadata

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(1) for struct allocation

**Example:**
```go
config := GeneratorConfig{client, "./docs", projectData}
```

**Edge Cases:**
- Nil AI client
- Non-writable output directory
- Incomplete project metadata


---

### `ProjectMeta`

```go
type ProjectMeta struct {
	Name string 
	Description string 
	RepoURL string 
	Features []Feature 
	TechStack []string 
	QuickStart []QuickStartBlock 
	BestPractices BestPractices 
}
```

**Summary:** Metadata about a software project

**Parameters:**
- `Name` (string): Project name
- `Description` (string): Brief project description
- `RepoURL` (string): Source repository URL
- `Features` ([]Feature): List of project features
- `TechStack` ([]string): Technologies used
- `QuickStart` ([]QuickStartBlock): Getting started guide
- `BestPractices` (BestPractices): Recommended practices

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(n) where n is size of slices

**Example:**
```go
meta := ProjectMeta{Name: "MyProject", ...}
```

**Edge Cases:**
- Empty required fields (name, description)
- Invalid repository URL format
- Empty slices vs nil slices


---

### `Feature`

```go
type Feature struct {
	Title string 
	Description string 
}
```

**Summary:** Defines a Feature struct with title and description fields.

**Parameters:**
- `Title` (string): Title of the feature
- `Description` (string): Description of the feature

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
feature := Feature{Title: "AI", Description: "Advanced AI capabilities"}
```

**Edge Cases:**
- Empty strings for Title or Description
- Non-string values assigned


---

### `QuickStartBlock`

```go
type QuickStartBlock struct {
	Title string 
	Shell string 
	Command string 
}
```

**Summary:** Defines a QuickStartBlock struct for shell command instructions.

**Parameters:**
- `Title` (string): Title of the quick start block
- `Shell` (string): Shell type (e.g., bash, zsh)
- `Command` (string): Command to execute

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
block := QuickStartBlock{Title: "Install", Shell: "bash", Command: "npm install"}
```

**Edge Cases:**
- Empty Command field
- Invalid shell type specified


---

### `BestPractices`

```go
type BestPractices struct {
	Do []string 
	Dont []string 
}
```

**Summary:** Defines a BestPractices struct with Do and Dont lists.

**Parameters:**
- `Do` ([]string): List of recommended practices
- `Dont` ([]string): List of practices to avoid

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(n)

**Example:**
```go
bp := BestPractices{Do: ["use interfaces"], Dont: ["ignore errors"]}
```

**Edge Cases:**
- Empty Do/Dont lists
- Duplicate entries in lists


---

### `AICodeRequest`

```go
type AICodeRequest struct {
	Input string 
	Language string 
	Target *aiTypes.Documentation 
}
```

**Summary:** Defines a struct for AI code request with input, language, and target documentation.

**Parameters:**
- `Input` (string): Input code or text
- `Language` (string): Programming language of the input
- `Target` (*aiTypes.Documentation): Pointer to target documentation type

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
req := AICodeRequest{Input: "func() {}", Language: "go", Target: &doc}
```

**Edge Cases:**
- Nil Target pointer
- Empty Input or Language strings


---

### `AIYAMLRequest`

```go
type AIYAMLRequest struct {
	Input string 
	Language string 
	Target *aiTypes.YAMLDocumentation 
}
```

**Summary:** Struct for AI YAML documentation request

**Parameters:**
- `Input` (string): Input text for documentation
- `Language` (string): Programming language of the input
- `Target` (*aiTypes.YAMLDocumentation): Pointer to target YAML documentation object

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
req := AIYAMLRequest{Input: "func()", Language: "go", Target: &doc}
```

**Edge Cases:**
- Nil Target pointer
- Empty Input string


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


---

## 📄 File: `config.go`

> 📍 `types\config.go`

## 📑 Contents

- [🧱 Structs (7)](#-structs)

## 🧱 Structs

### `GeneratorConfig`

```go
type GeneratorConfig struct {
	AIClient *ai.Client 
	OutputDir string 
	Project ProjectMeta 
}
```

**Summary:** Configuration for documentation generator

**Parameters:**
- `AIClient` (*ai.Client): AI client for enhanced generation
- `OutputDir` (string): Directory to save documentation
- `Project` (ProjectMeta): Project metadata

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(1) for struct allocation

**Example:**
```go
config := GeneratorConfig{client, "./docs", projectData}
```

**Edge Cases:**
- Nil AI client
- Non-writable output directory
- Incomplete project metadata


---

### `ProjectMeta`

```go
type ProjectMeta struct {
	Name string 
	Description string 
	RepoURL string 
	Features []Feature 
	TechStack []string 
	QuickStart []QuickStartBlock 
	BestPractices BestPractices 
}
```

**Summary:** Metadata about a software project

**Parameters:**
- `Name` (string): Project name
- `Description` (string): Brief project description
- `RepoURL` (string): Source repository URL
- `Features` ([]Feature): List of project features
- `TechStack` ([]string): Technologies used
- `QuickStart` ([]QuickStartBlock): Getting started guide
- `BestPractices` (BestPractices): Recommended practices

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(n) where n is size of slices

**Example:**
```go
meta := ProjectMeta{Name: "MyProject", ...}
```

**Edge Cases:**
- Empty required fields (name, description)
- Invalid repository URL format
- Empty slices vs nil slices


---

### `Feature`

```go
type Feature struct {
	Title string 
	Description string 
}
```

**Summary:** Defines a Feature struct with title and description fields.

**Parameters:**
- `Title` (string): Title of the feature
- `Description` (string): Description of the feature

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
feature := Feature{Title: "AI", Description: "Advanced AI capabilities"}
```

**Edge Cases:**
- Empty strings for Title or Description
- Non-string values assigned


---

### `QuickStartBlock`

```go
type QuickStartBlock struct {
	Title string 
	Shell string 
	Command string 
}
```

**Summary:** Defines a QuickStartBlock struct for shell command instructions.

**Parameters:**
- `Title` (string): Title of the quick start block
- `Shell` (string): Shell type (e.g., bash, zsh)
- `Command` (string): Command to execute

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
block := QuickStartBlock{Title: "Install", Shell: "bash", Command: "npm install"}
```

**Edge Cases:**
- Empty Command field
- Invalid shell type specified


---

### `BestPractices`

```go
type BestPractices struct {
	Do []string 
	Dont []string 
}
```

**Summary:** Defines a BestPractices struct with Do and Dont lists.

**Parameters:**
- `Do` ([]string): List of recommended practices
- `Dont` ([]string): List of practices to avoid

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(n)

**Example:**
```go
bp := BestPractices{Do: ["use interfaces"], Dont: ["ignore errors"]}
```

**Edge Cases:**
- Empty Do/Dont lists
- Duplicate entries in lists


---

### `AICodeRequest`

```go
type AICodeRequest struct {
	Input string 
	Language string 
	Target *aiTypes.Documentation 
}
```

**Summary:** Defines a struct for AI code request with input, language, and target documentation.

**Parameters:**
- `Input` (string): Input code or text
- `Language` (string): Programming language of the input
- `Target` (*aiTypes.Documentation): Pointer to target documentation type

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
req := AICodeRequest{Input: "func() {}", Language: "go", Target: &doc}
```

**Edge Cases:**
- Nil Target pointer
- Empty Input or Language strings


---

### `AIYAMLRequest`

```go
type AIYAMLRequest struct {
	Input string 
	Language string 
	Target *aiTypes.YAMLDocumentation 
}
```

**Summary:** Struct for AI YAML documentation request

**Parameters:**
- `Input` (string): Input text for documentation
- `Language` (string): Programming language of the input
- `Target` (*aiTypes.YAMLDocumentation): Pointer to target YAML documentation object

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
req := AIYAMLRequest{Input: "func()", Language: "go", Target: &doc}
```

**Edge Cases:**
- Nil Target pointer
- Empty Input string


---


---

## 📄 File: `documentation.go`

> 📍 `types/documentation.go`

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

## 📄 File: `config.go`

> 📍 `types/config.go`

## 📑 Contents

- [🧱 Structs (7)](#-structs)

## 🧱 Structs

### `GeneratorConfig`

```go
type GeneratorConfig struct {
	AIClient *ai.Client 
	OutputDir string 
	Project ProjectMeta 
}
```

_No documentation available._

---

### `ProjectMeta`

```go
type ProjectMeta struct {
	Name string 
	Description string 
	RepoURL string 
	Features []Feature 
	TechStack []string 
	QuickStart []QuickStartBlock 
	BestPractices BestPractices 
}
```

_No documentation available._

---

### `Feature`

```go
type Feature struct {
	Title string 
	Description string 
}
```

_No documentation available._

---

### `QuickStartBlock`

```go
type QuickStartBlock struct {
	Title string 
	Shell string 
	Command string 
}
```

_No documentation available._

---

### `BestPractices`

```go
type BestPractices struct {
	Do []string 
	Dont []string 
}
```

_No documentation available._

---

### `AICodeRequest`

```go
type AICodeRequest struct {
	Input string 
	Language string 
	Target *aiTypes.Documentation 
}
```

_No documentation available._

---

### `AIYAMLRequest`

```go
type AIYAMLRequest struct {
	Input string 
	Language string 
	Target *aiTypes.YAMLDocumentation 
}
```

_No documentation available._

---

