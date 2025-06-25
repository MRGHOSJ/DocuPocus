# 📦 Package: `generator`

> 📍 `C:\Users\DELL\Documents\GitHub\DocuPocus\internal\generator\generator.go`

[← Back to Overview](../README.md)

## 📑 Contents

- [🧱 Structs (2)](#-structs)
- [🔧 Functions (11)](#-functions)

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
<summary><b><code>generateSidebar(result *analyzer.AnalyzerResult, cfg GeneratorConfig)</code></b></summary>

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

