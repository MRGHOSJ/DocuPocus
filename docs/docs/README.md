# 📦 Package: `docs`

[← Back to Overview](../README.md)

## 📄 File: `code.go`

> 📍 `docs\code.go`

## 📑 Contents

- [🔧 Functions (5)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GeneratePackageDoc(pkg analyzer.Package, filePath string, cfg docTypes.GeneratorConfig)</code></b></summary>

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
<summary><b><code>formatDocumentation(doc aiTypes.Documentation)</code></b></summary>

**Summary:** Formats documentation content into a string

**Parameters:**
- `doc` (aiTypes.Documentation): Documentation object to format

**Returns:** Formatted documentation string

**Complexity:**
- Time: O(n) where n is documentation size
- Space: O(n) for output string

**Example:**
```go
docStr := formatDocumentation(myDoc)
```

**Edge Cases:**
- Empty documentation input
- Unsupported format requirements
- Special character handling


</details>

<details>
<summary><b><code>formatParams(params []analyzer.Parameter)</code></b></summary>

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

<details>
<summary><b><code>FormatFunction(f analyzer.Function)</code></b></summary>

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
<summary><b><code>FormatStruct(s analyzer.Struct)</code></b></summary>

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


---

## 📄 File: `readme.go`

> 📍 `docs\readme.go`

## 📑 Contents

- [🔧 Functions (1)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GenerateProjectReadme(result *analyzer.AnalyzerResult, cfg cfg.GeneratorConfig)</code></b></summary>

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


---

## 📄 File: `sidebar.go`

> 📍 `docs\sidebar.go`

## 📑 Contents

- [🔧 Functions (1)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GenerateSidebar(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Generates a sidebar from analyzer results and config

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to process
- `cfg` (docTypes.GeneratorConfig): Configuration for generation

**Returns:** Error if generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is size of analyzer result
- Space: O(n) for generated sidebar structure

**Example:**
```go
err := GenerateSidebar(analysisResult, config)
```

**Edge Cases:**
- Nil analyzer result input
- Invalid generator configuration


</details>


---

## 📄 File: `yaml.go`

> 📍 `docs\yaml.go`

## 📑 Contents

- [🔧 Functions (2)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GenerateYAMLDoc(pkg analyzer.Package, filePath string, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Generates YAML documentation for a package

**Parameters:**
- `pkg` (analyzer.Package): Package to document
- `filePath` (string): Output file path
- `cfg` (docTypes.GeneratorConfig): Generation configuration

**Returns:** Error if YAML generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is package complexity
- Space: O(n) for YAML output

**Example:**
```go
err := GenerateYAMLDoc(pkg, "docs.yaml", config)
```

**Edge Cases:**
- Empty package input
- Invalid file path permissions


</details>

<details>
<summary><b><code>generateYAMLExample(fields []analyzer.Field, indentLevel int)</code></b></summary>

**Summary:** Generates YAML example from field definitions

**Parameters:**
- `fields` ([]analyzer.Field): Fields to include in example
- `indentLevel` (int): Indentation level for output

**Returns:** Formatted YAML example string

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for generated YAML string

**Example:**
```go
example := generateYAMLExample(fields, 2)
```

**Edge Cases:**
- Empty fields array
- Negative indent level


</details>


---

## 📄 File: `code.go`

> 📍 `docs\code.go`

## 📑 Contents

- [🔧 Functions (5)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GeneratePackageDoc(pkg analyzer.Package, filePath string, cfg docTypes.GeneratorConfig)</code></b></summary>

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
<summary><b><code>formatDocumentation(doc aiTypes.Documentation)</code></b></summary>

**Summary:** Formats documentation content into a string

**Parameters:**
- `doc` (aiTypes.Documentation): Documentation object to format

**Returns:** Formatted documentation string

**Complexity:**
- Time: O(n) where n is documentation size
- Space: O(n) for output string

**Example:**
```go
docStr := formatDocumentation(myDoc)
```

**Edge Cases:**
- Empty documentation input
- Unsupported format requirements
- Special character handling


</details>

<details>
<summary><b><code>formatParams(params []analyzer.Parameter)</code></b></summary>

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

<details>
<summary><b><code>FormatFunction(f analyzer.Function)</code></b></summary>

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
<summary><b><code>FormatStruct(s analyzer.Struct)</code></b></summary>

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


---

## 📄 File: `readme.go`

> 📍 `docs\readme.go`

## 📑 Contents

- [🔧 Functions (1)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GenerateProjectReadme(result *analyzer.AnalyzerResult, cfg cfg.GeneratorConfig)</code></b></summary>

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


---

## 📄 File: `sidebar.go`

> 📍 `docs\sidebar.go`

## 📑 Contents

- [🔧 Functions (1)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GenerateSidebar(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Generates a sidebar from analyzer results and config

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to process
- `cfg` (docTypes.GeneratorConfig): Configuration for generation

**Returns:** Error if generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is size of analyzer result
- Space: O(n) for generated sidebar structure

**Example:**
```go
err := GenerateSidebar(analysisResult, config)
```

**Edge Cases:**
- Nil analyzer result input
- Invalid generator configuration


</details>


---

## 📄 File: `yaml.go`

> 📍 `docs\yaml.go`

## 📑 Contents

- [🔧 Functions (2)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GenerateYAMLDoc(pkg analyzer.Package, filePath string, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Generates YAML documentation for a package

**Parameters:**
- `pkg` (analyzer.Package): Package to document
- `filePath` (string): Output file path
- `cfg` (docTypes.GeneratorConfig): Generation configuration

**Returns:** Error if YAML generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is package complexity
- Space: O(n) for YAML output

**Example:**
```go
err := GenerateYAMLDoc(pkg, "docs.yaml", config)
```

**Edge Cases:**
- Empty package input
- Invalid file path permissions


</details>

<details>
<summary><b><code>generateYAMLExample(fields []analyzer.Field, indentLevel int)</code></b></summary>

**Summary:** Generates YAML example from field definitions

**Parameters:**
- `fields` ([]analyzer.Field): Fields to include in example
- `indentLevel` (int): Indentation level for output

**Returns:** Formatted YAML example string

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for generated YAML string

**Example:**
```go
example := generateYAMLExample(fields, 2)
```

**Edge Cases:**
- Empty fields array
- Negative indent level


</details>


---

## 📄 File: `code.go`

> 📍 `docs\code.go`

## 📑 Contents

- [🔧 Functions (5)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GeneratePackageDoc(pkg analyzer.Package, filePath string, cfg docTypes.GeneratorConfig)</code></b></summary>

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
<summary><b><code>formatDocumentation(doc aiTypes.Documentation)</code></b></summary>

**Summary:** Formats documentation content into a string

**Parameters:**
- `doc` (aiTypes.Documentation): Documentation object to format

**Returns:** Formatted documentation string

**Complexity:**
- Time: O(n) where n is documentation size
- Space: O(n) for output string

**Example:**
```go
docStr := formatDocumentation(myDoc)
```

**Edge Cases:**
- Empty documentation input
- Unsupported format requirements
- Special character handling


</details>

<details>
<summary><b><code>formatParams(params []analyzer.Parameter)</code></b></summary>

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

<details>
<summary><b><code>FormatFunction(f analyzer.Function)</code></b></summary>

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
<summary><b><code>FormatStruct(s analyzer.Struct)</code></b></summary>

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


---

## 📄 File: `readme.go`

> 📍 `docs\readme.go`

## 📑 Contents

- [🔧 Functions (1)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GenerateProjectReadme(result *analyzer.AnalyzerResult, cfg cfg.GeneratorConfig)</code></b></summary>

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


---

## 📄 File: `sidebar.go`

> 📍 `docs\sidebar.go`

## 📑 Contents

- [🔧 Functions (1)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GenerateSidebar(result *analyzer.AnalyzerResult, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Generates a sidebar from analyzer results and config

**Parameters:**
- `result` (*analyzer.AnalyzerResult): Analysis results to process
- `cfg` (docTypes.GeneratorConfig): Configuration for generation

**Returns:** Error if generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is size of analyzer result
- Space: O(n) for generated sidebar structure

**Example:**
```go
err := GenerateSidebar(analysisResult, config)
```

**Edge Cases:**
- Nil analyzer result input
- Invalid generator configuration


</details>


---

## 📄 File: `yaml.go`

> 📍 `docs\yaml.go`

## 📑 Contents

- [🔧 Functions (2)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GenerateYAMLDoc(pkg analyzer.Package, filePath string, cfg docTypes.GeneratorConfig)</code></b></summary>

**Summary:** Generates YAML documentation for a package

**Parameters:**
- `pkg` (analyzer.Package): Package to document
- `filePath` (string): Output file path
- `cfg` (docTypes.GeneratorConfig): Generation configuration

**Returns:** Error if YAML generation fails, nil otherwise

**Complexity:**
- Time: O(n) where n is package complexity
- Space: O(n) for YAML output

**Example:**
```go
err := GenerateYAMLDoc(pkg, "docs.yaml", config)
```

**Edge Cases:**
- Empty package input
- Invalid file path permissions


</details>

<details>
<summary><b><code>generateYAMLExample(fields []analyzer.Field, indentLevel int)</code></b></summary>

**Summary:** Generates YAML example from field definitions

**Parameters:**
- `fields` ([]analyzer.Field): Fields to include in example
- `indentLevel` (int): Indentation level for output

**Returns:** Formatted YAML example string

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for generated YAML string

**Example:**
```go
example := generateYAMLExample(fields, 2)
```

**Edge Cases:**
- Empty fields array
- Negative indent level


</details>

