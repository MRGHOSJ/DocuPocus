# 📦 Package: `generator`

[← Back to Overview](../README.md)

## 📄 File: `generator.go`

> 📍 `generator\generator.go`

## 📑 Contents

- [🧱 Structs (2)](#-structs)
- [🔧 Functions (12)](#-functions)

## 🧱 Structs

### `GeneratorConfig`

```go
type GeneratorConfig struct {
	AIClient *ai.Client 
	OutputDir string 
	Project  
}
```

**Summary:** Configuration struct for a generator

**Parameters:**
- `AIClient` (*ai.Client): AI client instance
- `OutputDir` (string): Directory for output files
- `Project` (Project): Project configuration

**Returns:** N/A (struct definition)

**Complexity:**
- Time: N/A
- Space: O(1) for struct instance

**Example:**
```go
config := GeneratorConfig{AIClient: client, OutputDir: "./output"}
```

**Edge Cases:**
- Nil AIClient
- Empty OutputDir


---

### `AIRequest`

```go
type AIRequest struct {
	Input string 
	Language string 
	Target *ai.Documentation 
}
```

**Summary:** Struct representing an AI request

**Parameters:**
- `Input` (string): Input text for AI processing
- `Language` (string): Target language for output
- `Target` (*ai.Documentation): Target documentation object

**Returns:** N/A (struct definition)

**Complexity:**
- Time: N/A
- Space: O(1) for struct instance

**Example:**
```go
req := AIRequest{Input: "Explain this code", Language: "en"}
```

**Edge Cases:**
- Empty Input string
- Unsupported Language value
- Nil Target pointer


---

## 🔧 Functions

<details>
<summary><b><code>GeneratePackageDocs(result *analyzer.AnalyzerResult, template string, cfg GeneratorConfig)</code></b></summary>

**Summary:** Generates package documentation using analysis results and template

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to document
- `template` (string): Documentation template string
- `cfg` (GeneratorConfig): Configuration for documentation generation

**Returns:** Error if documentation generation fails

**Complexity:**
- Time: O(n) where n is number of package elements
- Space: O(n) for generated documentation output

**Example:**
```go
err := GeneratePackageDocs(analysisResult, defaultTemplate, config)
```

**Edge Cases:**
- Nil analyzer result pointer
- Invalid template syntax
- Missing required configuration fields


</details>

<details>
<summary><b><code>generateProjectReadme(result *analyzer.AnalyzerResult, cfg GeneratorConfig)</code></b></summary>

**Summary:** Generates project README from analysis results

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Project analysis results
- `cfg` (GeneratorConfig): README generation configuration

**Returns:** Error if README generation fails

**Complexity:**
- Time: O(n) where n is project complexity
- Space: O(1) for in-place file writing

**Example:**
```go
err := generateProjectReadme(projectAnalysis, readmeConfig)
```

**Edge Cases:**
- Empty project structure
- Missing required sections in config
- File system permission issues


</details>

<details>
<summary><b><code>generatePackageDoc(pkg analyzer.Package, filePath string, cfg GeneratorConfig)</code></b></summary>

**Summary:** Creates documentation for a single package

**Parameters:**
- `pkg` (analyzer.Package): Package to document
- `filePath` (string): Output file path
- `cfg` (GeneratorConfig): Documentation configuration

**Returns:** Error if file generation fails

**Complexity:**
- Time: O(m) where m is package size
- Space: O(m) for documentation output

**Example:**
```go
err := generatePackageDoc(pkgData, "docs/pkg.md", docConfig)
```

**Edge Cases:**
- Empty package
- Invalid file path
- Unsupported documentation format


</details>

<details>
<summary><b><code>getDisplayPath(fullPath string)</code></b></summary>

**Summary:** Generates a sidebar from analysis results and configuration

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to process
- `cfg` (GeneratorConfig): Configuration for sidebar generation

**Returns:** Error if generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is size of analysis results
- Space: O(n) for storing generated sidebar structure

**Example:**
```go
err := generateSidebar(analysisResult, config)
```

**Edge Cases:**
- Nil analyzer result pointer
- Invalid configuration options


</details>

<details>
<summary><b><code>generateSidebar(result *analyzer.AnalyzerResult, cfg GeneratorConfig)</code></b></summary>

**Summary:** Generates a sidebar from analysis results and configuration

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to populate sidebar
- `cfg` (GeneratorConfig): Configuration for sidebar generation

**Returns:** error if generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is the size of analysis results
- Space: O(m) where m is the size of generated sidebar

**Example:**
```go
err := generateSidebar(analysisResult, config) // generates sidebar HTML
```

**Edge Cases:**
- Nil analyzer result input
- Invalid generator configuration
- Empty analysis results


</details>

<details>
<summary><b><code>calculateDocCompletion(pkg analyzer.Package)</code></b></summary>

**Summary:** Calculates documentation completion percentage for a package

**Parameters:**
- `pkg` (analyzer.Package): Package to analyze

**Returns:** Completion percentage as integer (0-100)

**Complexity:**
- Time: O(n) where n is number of package elements
- Space: O(1) constant space for calculations

**Example:**
```go
completion := calculateDocCompletion(myPackage) // returns 75
```

**Edge Cases:**
- Empty package with no elements
- Partially documented structures


</details>

<details>
<summary><b><code>processAIRequests(requests []AIRequest, client *ai.Client)</code></b></summary>

**Summary:** Processes batch AI requests using a client

**Parameters:**
- `requests` ([]AIRequest): Batch of requests to process
- `client` (*ai.Client): AI service client

**Returns:** None (handles requests asynchronously)

**Complexity:**
- Time: O(n) where n is number of requests
- Space: O(n) for request processing buffers

**Example:**
```go
processAIRequests(reqBatch, aiClient)
```

**Edge Cases:**
- Empty requests slice
- Nil client pointer
- Network timeouts


</details>

<details>
<summary><b><code>formatParams(params []analyzer.Parameter)</code></b></summary>

**Summary:** Formats a slice of parameters into a string

**Parameters:**
- `params` ([]analyzer.Parameter): Slice of parameters to format

**Returns:** Formatted string representation of parameters

**Complexity:**
- Time: O(n) where n is the number of parameters
- Space: O(n) for the output string

**Example:**
```go
str := formatParams([]Parameter{{Name: "x", Type: "int"}}) // returns "(x int)"
```

**Edge Cases:**
- Empty parameter slice
- Parameters with empty names or types


</details>

<details>
<summary><b><code>formatStruct(s analyzer.Struct)</code></b></summary>

**Summary:** Formats a struct into a string

**Parameters:**
- `s` (analyzer.Struct): Struct to format

**Returns:** Formatted string representation of the struct

**Complexity:**
- Time: O(n) where n is the number of fields
- Space: O(n) for the output string

**Example:**
```go
str := formatStruct(Struct{Name: "Point", Fields: []Field{{Name: "X", Type: "int"}}}) // returns "type Point struct { X int }"
```

**Edge Cases:**
- Empty struct with no fields
- Struct with nested complex types


</details>

<details>
<summary><b><code>formatFunction(f analyzer.Function)</code></b></summary>

**Summary:** Formats a function into a string

**Parameters:**
- `f` (analyzer.Function): Function to format

**Returns:** Formatted string representation of the function

**Complexity:**
- Time: O(n) where n is the complexity of the function signature
- Space: O(n) for the output string

**Example:**
```go
str := formatFunction(Function{Name: "add", Params: []Parameter{{Name: "a", Type: "int"}}, Returns: "int"}) // returns "func add(a int) int"
```

**Edge Cases:**
- Function with no parameters or return values
- Variadic functions
- Functions with complex return types


</details>

<details>
<summary><b><code>formatDocumentation(doc ai.Documentation)</code></b></summary>

**Summary:** Formats AI documentation into a string

**Parameters:**
- `doc` (ai.Documentation): Documentation object to format

**Returns:** Formatted documentation string

**Complexity:**
- Time: O(n) where n is documentation size
- Space: O(n) for output string

**Example:**
```go
formatted := formatDocumentation(doc) // returns markdown string
```

**Edge Cases:**
- Empty documentation input
- Malformed documentation structure


</details>

<details>
<summary><b><code>getLanguage(path string)</code></b></summary>

**Summary:** Detects programming language from file path

**Parameters:**
- `path` (string): File path to analyze

**Returns:** Detected language as string

**Complexity:**
- Time: O(1) for extension lookup
- Space: O(1)

**Example:**
```go
lang := getLanguage('main.go') // returns 'go'
```

**Edge Cases:**
- Unsupported file extensions
- Path without extension


</details>


---

## 📄 File: `generator.go`

> 📍 `generator\generator.go`

## 📑 Contents

- [🧱 Structs (3)](#-structs)
- [🔧 Functions (18)](#-functions)

## 🧱 Structs

### `GeneratorConfig`

```go
type GeneratorConfig struct {
	AIClient *ai.Client 
	OutputDir string 
	Project  
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

## 🔧 Functions

<details>
<summary><b><code>GeneratePackageDocs(result *analyzer.AnalyzerResult, template string, cfg GeneratorConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>generateYAMLDoc(pkg analyzer.Package, filePath string, cfg GeneratorConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>generateYAMLExample(fields []analyzer.Field, indentLevel int)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>getValueOrPlaceholder(f analyzer.Field)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>NormalizeFields(fields []analyzer.Field)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>formatYAMLStruct(s analyzer.Struct)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>formatYAMLFields(fields []analyzer.Field, depth int)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>generateProjectReadme(result *analyzer.AnalyzerResult, cfg GeneratorConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>generatePackageDoc(pkg analyzer.Package, filePath string, cfg GeneratorConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>getDisplayPath(fullPath string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>generateSidebar(result *analyzer.AnalyzerResult, cfg GeneratorConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>calculateDocCompletion(pkg analyzer.Package)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>processAIRequests(codeRequests []AICodeRequest, yamlRequests []AIYAMLRequest, client *ai.Client)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>formatParams(params []analyzer.Parameter)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>formatStruct(s analyzer.Struct)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>formatFunction(f analyzer.Function)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>formatDocumentation(doc aiTypes.Documentation)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>getLanguage(path string)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `generator.go`

> 📍 `generator\generator.go`

## 📑 Contents

- [🧱 Structs (3)](#-structs)
- [🔧 Functions (18)](#-functions)

## 🧱 Structs

### `GeneratorConfig`

```go
type GeneratorConfig struct {
	AIClient *ai.Client 
	OutputDir string 
	Project  
}
```

**Summary:** Configuration struct for code generation

**Parameters:**
- `AIClient` (*ai.Client): AI service client instance
- `OutputDir` (string): Target directory for generated files
- `Project` (Project): Project metadata (type details inferred)

**Returns:** N/A (configuration struct)

**Complexity:**
- Time: N/A
- Space: O(1) for struct allocation

**Example:**
```go
config := GeneratorConfig{AIClient: client, OutputDir: "out"}
```

**Edge Cases:**
- Nil AIClient causing runtime errors
- Invalid OutputDir path
- Incomplete Project struct fields


---

### `AICodeRequest`

```go
type AICodeRequest struct {
	Input string 
	Language string 
	Target *aiTypes.Documentation 
}
```

**Summary:** Struct representing an AI request

**Parameters:**
- `Input` (string): Input text for AI processing
- `Language` (string): Target language for output
- `Target` (*ai.Documentation): Target documentation object

**Returns:** N/A (struct definition)

**Complexity:**
- Time: N/A
- Space: O(1) for struct instance

**Example:**
```go
req := AIRequest{Input: "Explain this code", Language: "en"}
```

**Edge Cases:**
- Empty Input string
- Unsupported Language value
- Nil Target pointer


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

## 🔧 Functions

<details>
<summary><b><code>GeneratePackageDocs(result *analyzer.AnalyzerResult, template string, cfg GeneratorConfig)</code></b></summary>

**Summary:** Generates package documentation from analysis results

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to document
- `template` (string): Documentation template string
- `cfg` (GeneratorConfig): Configuration for documentation generation

**Returns:** Error if generation fails

**Complexity:**
- Time: O(n) where n is size of analysis results
- Space: O(m) where m is output documentation size

**Example:**
```go
err := GeneratePackageDocs(results, "{{.Name}}", config)
```

**Edge Cases:**
- Nil result pointer
- Invalid template syntax
- Empty configuration


</details>

<details>
<summary><b><code>generateYAMLDoc(pkg analyzer.Package, filePath string, cfg GeneratorConfig)</code></b></summary>

**Summary:** Generates project README from analysis results

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Project analysis results
- `cfg` (GeneratorConfig): README generation configuration

**Returns:** Error if README generation fails

**Complexity:**
- Time: O(n) where n is project complexity
- Space: O(1) for in-place file writing

**Example:**
```go
err := generateProjectReadme(projectAnalysis, readmeConfig)
```

**Edge Cases:**
- Empty project structure
- Missing required sections in config
- File system permission issues


</details>

<details>
<summary><b><code>generateYAMLExample(fields []analyzer.Field, indentLevel int)</code></b></summary>

**Summary:** Generates YAML example from field definitions

**Parameters:**
- `fields` ([]analyzer.Field): Fields to include in example
- `indentLevel` (int): Indentation level for YAML output

**Returns:** Generated YAML string

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(m) where m is output string length

**Example:**
```go
yaml := generateYAMLExample(fields, 2)
```

**Edge Cases:**
- Empty fields array
- Negative indentLevel
- Nested complex field types


</details>

<details>
<summary><b><code>getValueOrPlaceholder(f analyzer.Field)</code></b></summary>

**Summary:** Returns field value or a placeholder if empty

**Parameters:**
- `f` (analyzer.Field): Field to check for value

**Returns:** Field value as string or placeholder if empty

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
val := getValueOrPlaceholder(field) // returns "N/A" if empty
```

**Edge Cases:**
- Nil field input
- Field with whitespace-only value


</details>

<details>
<summary><b><code>NormalizeFields(fields []analyzer.Field)</code></b></summary>

**Summary:** Normalizes a slice of fields (e.g., trimming whitespace)

**Parameters:**
- `fields` ([]analyzer.Field): Slice of fields to normalize

**Returns:** Normalized slice of analyzer.Field

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for new slice creation

**Example:**
```go
normFields := NormalizeFields(rawFields) // trims all field values
```

**Edge Cases:**
- Empty input slice
- Fields containing only whitespace


</details>

<details>
<summary><b><code>formatYAMLStruct(s analyzer.Struct)</code></b></summary>

**Summary:** Formats a struct as YAML string

**Parameters:**
- `s` (analyzer.Struct): Struct to format

**Returns:** YAML-formatted string representation

**Complexity:**
- Time: O(n) where n is struct complexity
- Space: O(n) for output string storage

**Example:**
```go
yaml := formatYAMLStruct(myStruct) // returns "field: value\n"
```

**Edge Cases:**
- Empty struct input
- Struct with circular references


</details>

<details>
<summary><b><code>formatYAMLFields(fields []analyzer.Field, depth int)</code></b></summary>

**Summary:** Formats YAML fields with indentation based on depth

**Parameters:**
- `fields` ([]analyzer.Field): List of fields to format
- `depth` (int): Indentation level

**Returns:** Formatted YAML string representation of fields

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for output string storage

**Example:**
```go
yaml := formatYAMLFields([]Field{{Name: "test"}}, 2)
```

**Edge Cases:**
- Empty fields array
- Negative depth values
- Nested field structures


</details>

<details>
<summary><b><code>generateProjectReadme(result *analyzer.AnalyzerResult, cfg GeneratorConfig)</code></b></summary>

**Summary:** Generates project README from analysis results

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to include
- `cfg` (GeneratorConfig): Configuration for generation

**Returns:** Error if generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is size of analysis data
- Space: O(m) for README file content

**Example:**
```go
err := generateProjectReadme(&result, config)
```

**Edge Cases:**
- Nil result pointer
- Invalid config options
- File system permissions


</details>

<details>
<summary><b><code>generatePackageDoc(pkg analyzer.Package, filePath string, cfg GeneratorConfig)</code></b></summary>

**Summary:** Creates documentation for a single package

**Parameters:**
- `pkg` (analyzer.Package): Package to document
- `filePath` (string): Output file path
- `cfg` (GeneratorConfig): Documentation configuration

**Returns:** Error if file generation fails

**Complexity:**
- Time: O(m) where m is package size
- Space: O(m) for documentation output

**Example:**
```go
err := generatePackageDoc(pkgData, "docs/pkg.md", docConfig)
```

**Edge Cases:**
- Empty package
- Invalid file path
- Unsupported documentation format


</details>

<details>
<summary><b><code>getDisplayPath(fullPath string)</code></b></summary>

**Summary:** Generates a sidebar from analysis results and configuration

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to process
- `cfg` (GeneratorConfig): Configuration for sidebar generation

**Returns:** Error if generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is size of analysis results
- Space: O(n) for storing generated sidebar structure

**Example:**
```go
err := generateSidebar(analysisResult, config)
```

**Edge Cases:**
- Nil analyzer result pointer
- Invalid configuration options


</details>

<details>
<summary><b><code>generateSidebar(result *analyzer.AnalyzerResult, cfg GeneratorConfig)</code></b></summary>

**Summary:** Generates a sidebar from analysis results and configuration

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to populate sidebar
- `cfg` (GeneratorConfig): Configuration for sidebar generation

**Returns:** error if generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is the size of analysis results
- Space: O(m) where m is the size of generated sidebar

**Example:**
```go
err := generateSidebar(analysisResult, config) // generates sidebar HTML
```

**Edge Cases:**
- Nil analyzer result input
- Invalid generator configuration
- Empty analysis results


</details>

<details>
<summary><b><code>calculateDocCompletion(pkg analyzer.Package)</code></b></summary>

**Summary:** Calculates documentation completion percentage for a package

**Parameters:**
- `pkg` (analyzer.Package): Package to analyze

**Returns:** Completion percentage as integer (0-100)

**Complexity:**
- Time: O(n) where n is number of package elements
- Space: O(1) constant space for calculations

**Example:**
```go
completion := calculateDocCompletion(myPackage) // returns 75
```

**Edge Cases:**
- Empty package with no elements
- Partially documented structures


</details>

<details>
<summary><b><code>processAIRequests(codeRequests []AICodeRequest, yamlRequests []AIYAMLRequest, client *ai.Client)</code></b></summary>

**Summary:** Processes batch AI requests using a client

**Parameters:**
- `requests` ([]AIRequest): Batch of requests to process
- `client` (*ai.Client): AI service client

**Returns:** None (handles requests asynchronously)

**Complexity:**
- Time: O(n) where n is number of requests
- Space: O(n) for request processing buffers

**Example:**
```go
processAIRequests(reqBatch, aiClient)
```

**Edge Cases:**
- Empty requests slice
- Nil client pointer
- Network timeouts


</details>

<details>
<summary><b><code>formatParams(params []analyzer.Parameter)</code></b></summary>

**Summary:** Formats parameter list into string representation

**Parameters:**
- `params` ([]analyzer.Parameter): Parameters to format

**Returns:** String representation of parameters

**Complexity:**
- Time: O(n) where n is number of parameters
- Space: O(n) for output string storage

**Example:**
```go
paramStr := formatParams([]Parameter{{Name: "id"}})
```

**Edge Cases:**
- Empty parameters array
- Parameters with special characters
- Very long parameter lists


</details>

<details>
<summary><b><code>formatStruct(s analyzer.Struct)</code></b></summary>

**Summary:** Formats a struct into a string

**Parameters:**
- `s` (analyzer.Struct): Struct to format

**Returns:** Formatted string representation of the struct

**Complexity:**
- Time: O(n) where n is the number of fields
- Space: O(n) for the output string

**Example:**
```go
str := formatStruct(Struct{Name: "Point", Fields: []Field{{Name: "X", Type: "int"}}}) // returns "type Point struct { X int }"
```

**Edge Cases:**
- Empty struct with no fields
- Struct with nested complex types


</details>

<details>
<summary><b><code>formatFunction(f analyzer.Function)</code></b></summary>

**Summary:** Formats a function into a string

**Parameters:**
- `f` (analyzer.Function): Function to format

**Returns:** Formatted string representation of the function

**Complexity:**
- Time: O(n) where n is the complexity of the function signature
- Space: O(n) for the output string

**Example:**
```go
str := formatFunction(Function{Name: "add", Params: []Parameter{{Name: "a", Type: "int"}}, Returns: "int"}) // returns "func add(a int) int"
```

**Edge Cases:**
- Function with no parameters or return values
- Variadic functions
- Functions with complex return types


</details>

<details>
<summary><b><code>formatDocumentation(doc aiTypes.Documentation)</code></b></summary>

**Summary:** Formats AI documentation into a string

**Parameters:**
- `doc` (ai.Documentation): Documentation object to format

**Returns:** Formatted documentation string

**Complexity:**
- Time: O(n) where n is documentation size
- Space: O(n) for output string

**Example:**
```go
formatted := formatDocumentation(doc) // returns markdown string
```

**Edge Cases:**
- Empty documentation input
- Malformed documentation structure


</details>

<details>
<summary><b><code>getLanguage(path string)</code></b></summary>

**Summary:** Determines the programming language from a file path

**Parameters:**
- `path` (string): File path to analyze

**Returns:** Detected programming language as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
lang := getLanguage("main.go") // returns "Go"
```

**Edge Cases:**
- Unsupported file extensions return empty string
- Case sensitivity in file extensions
- Non-file paths (e.g., directories) may cause errors


</details>


---

## 📄 File: `generator.go`

> 📍 `generator\generator.go`

## 📑 Contents

- [🧱 Structs (3)](#-structs)
- [🔧 Functions (18)](#-functions)

## 🧱 Structs

### `GeneratorConfig`

```go
type GeneratorConfig struct {
	AIClient *ai.Client 
	OutputDir string 
	Project  
}
```

**Summary:** Configuration struct for code generation

**Parameters:**
- `AIClient` (*ai.Client): AI service client instance
- `OutputDir` (string): Target directory for generated files
- `Project` (Project): Project metadata (type details inferred)

**Returns:** N/A (configuration struct)

**Complexity:**
- Time: N/A
- Space: O(1) for struct allocation

**Example:**
```go
config := GeneratorConfig{AIClient: client, OutputDir: "out"}
```

**Edge Cases:**
- Nil AIClient causing runtime errors
- Invalid OutputDir path
- Incomplete Project struct fields


---

### `AICodeRequest`

```go
type AICodeRequest struct {
	Input string 
	Language string 
	Target *aiTypes.Documentation 
}
```

**Summary:** Struct representing an AI request

**Parameters:**
- `Input` (string): Input text for AI processing
- `Language` (string): Target language for output
- `Target` (*ai.Documentation): Target documentation object

**Returns:** N/A (struct definition)

**Complexity:**
- Time: N/A
- Space: O(1) for struct instance

**Example:**
```go
req := AIRequest{Input: "Explain this code", Language: "en"}
```

**Edge Cases:**
- Empty Input string
- Unsupported Language value
- Nil Target pointer


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

## 🔧 Functions

<details>
<summary><b><code>GeneratePackageDocs(result *analyzer.AnalyzerResult, template string, cfg GeneratorConfig)</code></b></summary>

**Summary:** Generates package documentation from analysis results

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to document
- `template` (string): Documentation template string
- `cfg` (GeneratorConfig): Configuration for documentation generation

**Returns:** Error if generation fails

**Complexity:**
- Time: O(n) where n is size of analysis results
- Space: O(m) where m is output documentation size

**Example:**
```go
err := GeneratePackageDocs(results, "{{.Name}}", config)
```

**Edge Cases:**
- Nil result pointer
- Invalid template syntax
- Empty configuration


</details>

<details>
<summary><b><code>generateYAMLDoc(pkg analyzer.Package, filePath string, cfg GeneratorConfig)</code></b></summary>

**Summary:** Generates project README from analysis results

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Project analysis results
- `cfg` (GeneratorConfig): README generation configuration

**Returns:** Error if README generation fails

**Complexity:**
- Time: O(n) where n is project complexity
- Space: O(1) for in-place file writing

**Example:**
```go
err := generateProjectReadme(projectAnalysis, readmeConfig)
```

**Edge Cases:**
- Empty project structure
- Missing required sections in config
- File system permission issues


</details>

<details>
<summary><b><code>generateYAMLExample(fields []analyzer.Field, indentLevel int)</code></b></summary>

**Summary:** Generates YAML example from field definitions

**Parameters:**
- `fields` ([]analyzer.Field): Fields to include in example
- `indentLevel` (int): Indentation level for YAML output

**Returns:** Generated YAML string

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(m) where m is output string length

**Example:**
```go
yaml := generateYAMLExample(fields, 2)
```

**Edge Cases:**
- Empty fields array
- Negative indentLevel
- Nested complex field types


</details>

<details>
<summary><b><code>getValueOrPlaceholder(f analyzer.Field)</code></b></summary>

**Summary:** Returns field value or a placeholder if empty

**Parameters:**
- `f` (analyzer.Field): Field to check for value

**Returns:** Field value as string or placeholder if empty

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
val := getValueOrPlaceholder(field) // returns "N/A" if empty
```

**Edge Cases:**
- Nil field input
- Field with whitespace-only value


</details>

<details>
<summary><b><code>NormalizeFields(fields []analyzer.Field)</code></b></summary>

**Summary:** Normalizes a slice of fields (e.g., trimming whitespace)

**Parameters:**
- `fields` ([]analyzer.Field): Slice of fields to normalize

**Returns:** Normalized slice of analyzer.Field

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for new slice creation

**Example:**
```go
normFields := NormalizeFields(rawFields) // trims all field values
```

**Edge Cases:**
- Empty input slice
- Fields containing only whitespace


</details>

<details>
<summary><b><code>formatYAMLStruct(s analyzer.Struct)</code></b></summary>

**Summary:** Formats a struct as YAML string

**Parameters:**
- `s` (analyzer.Struct): Struct to format

**Returns:** YAML-formatted string representation

**Complexity:**
- Time: O(n) where n is struct complexity
- Space: O(n) for output string storage

**Example:**
```go
yaml := formatYAMLStruct(myStruct) // returns "field: value\n"
```

**Edge Cases:**
- Empty struct input
- Struct with circular references


</details>

<details>
<summary><b><code>formatYAMLFields(fields []analyzer.Field, depth int)</code></b></summary>

**Summary:** Formats YAML fields with indentation based on depth

**Parameters:**
- `fields` ([]analyzer.Field): List of fields to format
- `depth` (int): Indentation level

**Returns:** Formatted YAML string representation of fields

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for output string storage

**Example:**
```go
yaml := formatYAMLFields([]Field{{Name: "test"}}, 2)
```

**Edge Cases:**
- Empty fields array
- Negative depth values
- Nested field structures


</details>

<details>
<summary><b><code>generateProjectReadme(result *analyzer.AnalyzerResult, cfg GeneratorConfig)</code></b></summary>

**Summary:** Generates project README from analysis results

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to include
- `cfg` (GeneratorConfig): Configuration for generation

**Returns:** Error if generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is size of analysis data
- Space: O(m) for README file content

**Example:**
```go
err := generateProjectReadme(&result, config)
```

**Edge Cases:**
- Nil result pointer
- Invalid config options
- File system permissions


</details>

<details>
<summary><b><code>generatePackageDoc(pkg analyzer.Package, filePath string, cfg GeneratorConfig)</code></b></summary>

**Summary:** Creates documentation for a single package

**Parameters:**
- `pkg` (analyzer.Package): Package to document
- `filePath` (string): Output file path
- `cfg` (GeneratorConfig): Documentation configuration

**Returns:** Error if file generation fails

**Complexity:**
- Time: O(m) where m is package size
- Space: O(m) for documentation output

**Example:**
```go
err := generatePackageDoc(pkgData, "docs/pkg.md", docConfig)
```

**Edge Cases:**
- Empty package
- Invalid file path
- Unsupported documentation format


</details>

<details>
<summary><b><code>getDisplayPath(fullPath string)</code></b></summary>

**Summary:** Generates a sidebar from analysis results and configuration

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to process
- `cfg` (GeneratorConfig): Configuration for sidebar generation

**Returns:** Error if generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is size of analysis results
- Space: O(n) for storing generated sidebar structure

**Example:**
```go
err := generateSidebar(analysisResult, config)
```

**Edge Cases:**
- Nil analyzer result pointer
- Invalid configuration options


</details>

<details>
<summary><b><code>generateSidebar(result *analyzer.AnalyzerResult, cfg GeneratorConfig)</code></b></summary>

**Summary:** Generates a sidebar from analysis results and configuration

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to populate sidebar
- `cfg` (GeneratorConfig): Configuration for sidebar generation

**Returns:** error if generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is the size of analysis results
- Space: O(m) where m is the size of generated sidebar

**Example:**
```go
err := generateSidebar(analysisResult, config) // generates sidebar HTML
```

**Edge Cases:**
- Nil analyzer result input
- Invalid generator configuration
- Empty analysis results


</details>

<details>
<summary><b><code>calculateDocCompletion(pkg analyzer.Package)</code></b></summary>

**Summary:** Calculates documentation completion percentage for a package

**Parameters:**
- `pkg` (analyzer.Package): Package to analyze

**Returns:** Completion percentage as integer (0-100)

**Complexity:**
- Time: O(n) where n is number of package elements
- Space: O(1) constant space for calculations

**Example:**
```go
completion := calculateDocCompletion(myPackage) // returns 75
```

**Edge Cases:**
- Empty package with no elements
- Partially documented structures


</details>

<details>
<summary><b><code>processAIRequests(codeRequests []AICodeRequest, yamlRequests []AIYAMLRequest, client *ai.Client)</code></b></summary>

**Summary:** Processes batch AI requests using a client

**Parameters:**
- `requests` ([]AIRequest): Batch of requests to process
- `client` (*ai.Client): AI service client

**Returns:** None (handles requests asynchronously)

**Complexity:**
- Time: O(n) where n is number of requests
- Space: O(n) for request processing buffers

**Example:**
```go
processAIRequests(reqBatch, aiClient)
```

**Edge Cases:**
- Empty requests slice
- Nil client pointer
- Network timeouts


</details>

<details>
<summary><b><code>formatParams(params []analyzer.Parameter)</code></b></summary>

**Summary:** Formats parameter list into string representation

**Parameters:**
- `params` ([]analyzer.Parameter): Parameters to format

**Returns:** String representation of parameters

**Complexity:**
- Time: O(n) where n is number of parameters
- Space: O(n) for output string storage

**Example:**
```go
paramStr := formatParams([]Parameter{{Name: "id"}})
```

**Edge Cases:**
- Empty parameters array
- Parameters with special characters
- Very long parameter lists


</details>

<details>
<summary><b><code>formatStruct(s analyzer.Struct)</code></b></summary>

**Summary:** Formats a struct into a string

**Parameters:**
- `s` (analyzer.Struct): Struct to format

**Returns:** Formatted string representation of the struct

**Complexity:**
- Time: O(n) where n is the number of fields
- Space: O(n) for the output string

**Example:**
```go
str := formatStruct(Struct{Name: "Point", Fields: []Field{{Name: "X", Type: "int"}}}) // returns "type Point struct { X int }"
```

**Edge Cases:**
- Empty struct with no fields
- Struct with nested complex types


</details>

<details>
<summary><b><code>formatFunction(f analyzer.Function)</code></b></summary>

**Summary:** Formats a function into a string

**Parameters:**
- `f` (analyzer.Function): Function to format

**Returns:** Formatted string representation of the function

**Complexity:**
- Time: O(n) where n is the complexity of the function signature
- Space: O(n) for the output string

**Example:**
```go
str := formatFunction(Function{Name: "add", Params: []Parameter{{Name: "a", Type: "int"}}, Returns: "int"}) // returns "func add(a int) int"
```

**Edge Cases:**
- Function with no parameters or return values
- Variadic functions
- Functions with complex return types


</details>

<details>
<summary><b><code>formatDocumentation(doc aiTypes.Documentation)</code></b></summary>

**Summary:** Formats AI documentation into a string

**Parameters:**
- `doc` (ai.Documentation): Documentation object to format

**Returns:** Formatted documentation string

**Complexity:**
- Time: O(n) where n is documentation size
- Space: O(n) for output string

**Example:**
```go
formatted := formatDocumentation(doc) // returns markdown string
```

**Edge Cases:**
- Empty documentation input
- Malformed documentation structure


</details>

<details>
<summary><b><code>getLanguage(path string)</code></b></summary>

**Summary:** Determines the programming language from a file path

**Parameters:**
- `path` (string): File path to analyze

**Returns:** Detected programming language as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
lang := getLanguage("main.go") // returns "Go"
```

**Edge Cases:**
- Unsupported file extensions return empty string
- Case sensitivity in file extensions
- Non-file paths (e.g., directories) may cause errors


</details>


---

## 📄 File: `ai_requests.go`

> 📍 `generator\ai_requests.go`

## 📑 Contents

- [🔧 Functions (3)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>processAIRequests(codeRequests []cfg.AICodeRequest, yamlRequests []cfg.AIYAMLRequest, client *ai.Client)</code></b></summary>

**Summary:** Configuration struct for code generation

**Parameters:**
- `AIClient` (*ai.Client): AI service client instance
- `OutputDir` (string): Target directory for generated files
- `Project` (Project): Project metadata (type details inferred)

**Returns:** N/A (configuration struct)

**Complexity:**
- Time: N/A
- Space: O(1) for struct allocation

**Example:**
```go
config := GeneratorConfig{AIClient: client, OutputDir: "out"}
```

**Edge Cases:**
- Nil AIClient causing runtime errors
- Invalid OutputDir path
- Incomplete Project struct fields


</details>

<details>
<summary><b><code>formatYAMLStruct(s analyzer.Struct)</code></b></summary>

**Summary:** Formats a YAML structure into a string

**Parameters:**
- `s` (analyzer.Struct): YAML structure to format

**Returns:** Formatted YAML string

**Complexity:**
- Time: O(n) where n is structure size
- Space: O(n) for output string

**Example:**
```go
yamlStr := formatYAMLStruct(myStruct)
```

**Edge Cases:**
- Empty input structure
- Invalid YAML formatting rules
- Nested structure depth limits


</details>

<details>
<summary><b><code>formatYAMLFields(fields []analyzer.Field, depth int)</code></b></summary>

**Summary:** Formats YAML fields with indentation based on depth

**Parameters:**
- `fields` ([]analyzer.Field): List of fields to format
- `depth` (int): Indentation level

**Returns:** Formatted YAML string representation of fields

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for output string storage

**Example:**
```go
yaml := formatYAMLFields([]Field{{Name: "test"}}, 2)
```

**Edge Cases:**
- Empty fields array
- Negative depth values
- Nested field structures


</details>


---

## 📄 File: `generator.go`

> 📍 `generator\generator.go`

## 📑 Contents

- [🔧 Functions (5)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GeneratePackageDocs(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Returns field value or a placeholder if empty

**Parameters:**
- `f` (analyzer.Field): Field to check for value

**Returns:** Field value as string or placeholder if empty

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
val := getValueOrPlaceholder(field) // returns "N/A" if empty
```

**Edge Cases:**
- Nil field input
- Field with whitespace-only value


</details>

<details>
<summary><b><code>prepareOutputStructure(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Prepares output structure from analyzer result using generator config

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis result to process
- `cfg` (docTypes.GeneratorConfig): Configuration for output generation

**Returns:** Error if preparation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is size of result data
- Space: O(n) for output structure allocation

**Example:**
```go
err := prepareOutputStructure(&analysisResult, config)
```

**Edge Cases:**
- Nil result pointer
- Invalid config options
- Large result data causing memory issues


</details>

<details>
<summary><b><code>prepareAIRequests(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Creates AI request objects from analyzer result

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis result to convert
- `cfg` (docTypes.GeneratorConfig): Configuration for request generation

**Returns:** Two slices: AI code requests and AI YAML requests

**Complexity:**
- Time: O(n) where n is number of elements in result
- Space: O(n) for request slices allocation

**Example:**
```go
codeReqs, yamlReqs := prepareAIRequests(&analysisResult, config)
```

**Edge Cases:**
- Empty result data
- Configuration limiting request types
- Memory constraints with large result sets


</details>

<details>
<summary><b><code>enhanceWithAI(codeRequests []docTypes.AICodeRequest, yamlRequests []docTypes.AIYAMLRequest, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Enhances documentation using AI-generated content

**Parameters:**
- `codeRequests` ([]docTypes.AICodeRequest): Code documentation requests
- `yamlRequests` ([]docTypes.AIYAMLRequest): YAML documentation requests
- `cfg` (docTypes.GeneratorConfig): Configuration for enhancement process

**Returns:** Error if enhancement fails, nil otherwise

**Complexity:**
- Time: O(n+m) where n=codeReqs, m=yamlReqs
- Space: O(n+m) for processed results storage

**Example:**
```go
err := enhanceWithAI(codeReqs, yamlReqs, config)
```

**Edge Cases:**
- Empty request slices
- AI service unavailability
- Rate limiting from AI provider
- Large request batches timing out


</details>

<details>
<summary><b><code>generateFinalDocs(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Generates final documentation from analysis results

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to document
- `cfg` (docTypes.GeneratorConfig): Configuration for documentation generation

**Returns:** Error if documentation generation fails

**Complexity:**
- Time: O(n) where n is size of analysis results
- Space: O(n) for generated documentation

**Example:**
```go
err := generateFinalDocs(analysisResults, config)
```

**Edge Cases:**
- Nil analyzer result
- Invalid output directory in config
- Missing required fields in config


</details>


---

## 📄 File: `ai_requests.go`

> 📍 `generator\ai_requests.go`

## 📑 Contents

- [🔧 Functions (3)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>processAIRequests(codeRequests []cfg.AICodeRequest, yamlRequests []cfg.AIYAMLRequest, client *ai.Client)</code></b></summary>

**Summary:** Configuration struct for code generation

**Parameters:**
- `AIClient` (*ai.Client): AI service client instance
- `OutputDir` (string): Target directory for generated files
- `Project` (Project): Project metadata (type details inferred)

**Returns:** N/A (configuration struct)

**Complexity:**
- Time: N/A
- Space: O(1) for struct allocation

**Example:**
```go
config := GeneratorConfig{AIClient: client, OutputDir: "out"}
```

**Edge Cases:**
- Nil AIClient causing runtime errors
- Invalid OutputDir path
- Incomplete Project struct fields


</details>

<details>
<summary><b><code>formatYAMLStruct(s analyzer.Struct)</code></b></summary>

**Summary:** Formats a YAML structure into a string

**Parameters:**
- `s` (analyzer.Struct): YAML structure to format

**Returns:** Formatted YAML string

**Complexity:**
- Time: O(n) where n is structure size
- Space: O(n) for output string

**Example:**
```go
yamlStr := formatYAMLStruct(myStruct)
```

**Edge Cases:**
- Empty input structure
- Invalid YAML formatting rules
- Nested structure depth limits


</details>

<details>
<summary><b><code>formatYAMLFields(fields []analyzer.Field, depth int)</code></b></summary>

**Summary:** Formats YAML fields with indentation based on depth

**Parameters:**
- `fields` ([]analyzer.Field): List of fields to format
- `depth` (int): Indentation level

**Returns:** Formatted YAML string representation of fields

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for output string storage

**Example:**
```go
yaml := formatYAMLFields([]Field{{Name: "test"}}, 2)
```

**Edge Cases:**
- Empty fields array
- Negative depth values
- Nested field structures


</details>


---

## 📄 File: `generator.go`

> 📍 `generator\generator.go`

## 📑 Contents

- [🔧 Functions (5)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GeneratePackageDocs(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Returns field value or a placeholder if empty

**Parameters:**
- `f` (analyzer.Field): Field to check for value

**Returns:** Field value as string or placeholder if empty

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
val := getValueOrPlaceholder(field) // returns "N/A" if empty
```

**Edge Cases:**
- Nil field input
- Field with whitespace-only value


</details>

<details>
<summary><b><code>prepareOutputStructure(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Prepares output structure from analyzer result using generator config

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis result to process
- `cfg` (docTypes.GeneratorConfig): Configuration for output generation

**Returns:** Error if preparation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is size of result data
- Space: O(n) for output structure allocation

**Example:**
```go
err := prepareOutputStructure(&analysisResult, config)
```

**Edge Cases:**
- Nil result pointer
- Invalid config options
- Large result data causing memory issues


</details>

<details>
<summary><b><code>prepareAIRequests(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Creates AI request objects from analyzer result

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis result to convert
- `cfg` (docTypes.GeneratorConfig): Configuration for request generation

**Returns:** Two slices: AI code requests and AI YAML requests

**Complexity:**
- Time: O(n) where n is number of elements in result
- Space: O(n) for request slices allocation

**Example:**
```go
codeReqs, yamlReqs := prepareAIRequests(&analysisResult, config)
```

**Edge Cases:**
- Empty result data
- Configuration limiting request types
- Memory constraints with large result sets


</details>

<details>
<summary><b><code>enhanceWithAI(codeRequests []docTypes.AICodeRequest, yamlRequests []docTypes.AIYAMLRequest, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Enhances documentation using AI-generated content

**Parameters:**
- `codeRequests` ([]docTypes.AICodeRequest): Code documentation requests
- `yamlRequests` ([]docTypes.AIYAMLRequest): YAML documentation requests
- `cfg` (docTypes.GeneratorConfig): Configuration for enhancement process

**Returns:** Error if enhancement fails, nil otherwise

**Complexity:**
- Time: O(n+m) where n=codeReqs, m=yamlReqs
- Space: O(n+m) for processed results storage

**Example:**
```go
err := enhanceWithAI(codeReqs, yamlReqs, config)
```

**Edge Cases:**
- Empty request slices
- AI service unavailability
- Rate limiting from AI provider
- Large request batches timing out


</details>

<details>
<summary><b><code>generateFinalDocs(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Generates final documentation from analysis results

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to document
- `cfg` (docTypes.GeneratorConfig): Configuration for documentation generation

**Returns:** Error if documentation generation fails

**Complexity:**
- Time: O(n) where n is size of analysis results
- Space: O(n) for generated documentation

**Example:**
```go
err := generateFinalDocs(analysisResults, config)
```

**Edge Cases:**
- Nil analyzer result
- Invalid output directory in config
- Missing required fields in config


</details>


---

## 📄 File: `ai_requests.go`

> 📍 `generator\ai_requests.go`

## 📑 Contents

- [🔧 Functions (3)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>processAIRequests(codeRequests []cfg.AICodeRequest, yamlRequests []cfg.AIYAMLRequest, client *ai.Client)</code></b></summary>

**Summary:** Configuration struct for code generation

**Parameters:**
- `AIClient` (*ai.Client): AI service client instance
- `OutputDir` (string): Target directory for generated files
- `Project` (Project): Project metadata (type details inferred)

**Returns:** N/A (configuration struct)

**Complexity:**
- Time: N/A
- Space: O(1) for struct allocation

**Example:**
```go
config := GeneratorConfig{AIClient: client, OutputDir: "out"}
```

**Edge Cases:**
- Nil AIClient causing runtime errors
- Invalid OutputDir path
- Incomplete Project struct fields


</details>

<details>
<summary><b><code>formatYAMLStruct(s analyzer.Struct)</code></b></summary>

**Summary:** Formats a YAML structure into a string

**Parameters:**
- `s` (analyzer.Struct): YAML structure to format

**Returns:** Formatted YAML string

**Complexity:**
- Time: O(n) where n is structure size
- Space: O(n) for output string

**Example:**
```go
yamlStr := formatYAMLStruct(myStruct)
```

**Edge Cases:**
- Empty input structure
- Invalid YAML formatting rules
- Nested structure depth limits


</details>

<details>
<summary><b><code>formatYAMLFields(fields []analyzer.Field, depth int)</code></b></summary>

**Summary:** Formats YAML fields with indentation based on depth

**Parameters:**
- `fields` ([]analyzer.Field): List of fields to format
- `depth` (int): Indentation level

**Returns:** Formatted YAML string representation of fields

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for output string storage

**Example:**
```go
yaml := formatYAMLFields([]Field{{Name: "test"}}, 2)
```

**Edge Cases:**
- Empty fields array
- Negative depth values
- Nested field structures


</details>


---

## 📄 File: `generator.go`

> 📍 `generator\generator.go`

## 📑 Contents

- [🔧 Functions (5)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GeneratePackageDocs(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Returns field value or a placeholder if empty

**Parameters:**
- `f` (analyzer.Field): Field to check for value

**Returns:** Field value as string or placeholder if empty

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
val := getValueOrPlaceholder(field) // returns "N/A" if empty
```

**Edge Cases:**
- Nil field input
- Field with whitespace-only value


</details>

<details>
<summary><b><code>prepareOutputStructure(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Prepares output structure from analyzer result using generator config

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis result to process
- `cfg` (docTypes.GeneratorConfig): Configuration for output generation

**Returns:** Error if preparation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is size of result data
- Space: O(n) for output structure allocation

**Example:**
```go
err := prepareOutputStructure(&analysisResult, config)
```

**Edge Cases:**
- Nil result pointer
- Invalid config options
- Large result data causing memory issues


</details>

<details>
<summary><b><code>prepareAIRequests(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Creates AI request objects from analyzer result

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis result to convert
- `cfg` (docTypes.GeneratorConfig): Configuration for request generation

**Returns:** Two slices: AI code requests and AI YAML requests

**Complexity:**
- Time: O(n) where n is number of elements in result
- Space: O(n) for request slices allocation

**Example:**
```go
codeReqs, yamlReqs := prepareAIRequests(&analysisResult, config)
```

**Edge Cases:**
- Empty result data
- Configuration limiting request types
- Memory constraints with large result sets


</details>

<details>
<summary><b><code>enhanceWithAI(codeRequests []docTypes.AICodeRequest, yamlRequests []docTypes.AIYAMLRequest, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Enhances documentation using AI-generated content

**Parameters:**
- `codeRequests` ([]docTypes.AICodeRequest): Code documentation requests
- `yamlRequests` ([]docTypes.AIYAMLRequest): YAML documentation requests
- `cfg` (docTypes.GeneratorConfig): Configuration for enhancement process

**Returns:** Error if enhancement fails, nil otherwise

**Complexity:**
- Time: O(n+m) where n=codeReqs, m=yamlReqs
- Space: O(n+m) for processed results storage

**Example:**
```go
err := enhanceWithAI(codeReqs, yamlReqs, config)
```

**Edge Cases:**
- Empty request slices
- AI service unavailability
- Rate limiting from AI provider
- Large request batches timing out


</details>

<details>
<summary><b><code>generateFinalDocs(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Generates final documentation from analysis results

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to document
- `cfg` (docTypes.GeneratorConfig): Configuration for documentation generation

**Returns:** Error if documentation generation fails

**Complexity:**
- Time: O(n) where n is size of analysis results
- Space: O(n) for generated documentation

**Example:**
```go
err := generateFinalDocs(analysisResults, config)
```

**Edge Cases:**
- Nil analyzer result
- Invalid output directory in config
- Missing required fields in config


</details>


---

## 📄 File: `ai_requests.go`

> 📍 `generator/ai_requests.go`

## 📑 Contents

- [🔧 Functions (3)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>processAIRequests(codeRequests []cfg.AICodeRequest, yamlRequests []cfg.AIYAMLRequest, client *ai.Client)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>formatYAMLStruct(s analyzer.Struct)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>formatYAMLFields(fields []analyzer.Field, depth int)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `generator.go`

> 📍 `generator/generator.go`

## 📑 Contents

- [🔧 Functions (5)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GeneratePackageDocs(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>prepareOutputStructure(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>prepareAIRequests(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>enhanceWithAI(codeRequests []docTypes.AICodeRequest, yamlRequests []docTypes.AIYAMLRequest, cfg docTypes.GeneratorConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>generateFinalDocs(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

_No documentation available._

</details>

