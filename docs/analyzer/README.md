# 📦 Package: `analyzer`

[← Back to Overview](../README.md)

## 📄 File: `analyzer.go`

> 📍 `analyzer\analyzer.go`

## 📑 Contents

- [🧱 Structs (7)](#-structs)
- [🔧 Functions (2)](#-functions)

## 🧱 Structs

### `AnalyzerResult`

```go
type AnalyzerResult struct {
	Files []*AnalyzedFile 
}
```

**Summary:** Container for multiple file analysis results

**Returns:** None (type definition)

**Complexity:**
- Time: N/A (data structure)
- Space: O(n) where n is number of files

**Example:**
```go
result := AnalyzerResult{Files: analyzedFilesSlice}
```

**Edge Cases:**
- Empty Files slice
- Nil elements in Files array
- Concurrent access/modification


---

### `AnalyzedFile`

```go
type AnalyzedFile struct {
	Path string 
	Packages []Package 
}
```

**Summary:** Stores analysis data for a single file

**Returns:** None (type definition)

**Complexity:**
- Time: N/A (data structure)
- Space: O(m) where m is number of packages

**Example:**
```go
file := AnalyzedFile{Path: "/src/main.go", Packages: pkgs}
```

**Edge Cases:**
- Empty Path string
- Nil Packages slice
- Duplicate package entries


---

### `Package`

```go
type Package struct {
	Name string 
	Path string 
	Imports []string 
	Structs []Struct 
	Funcs []Function 
	Files []string 
}
```

**Summary:** Defines a Package structure with metadata and components

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
pkg := Package{Name: "example", Path: "github.com/example"}
```

**Edge Cases:**
- Empty slices for Imports, Structs, Funcs, or Files
- Nil values in any field


---

### `Struct`

```go
type Struct struct {
	Name string 
	Fields []Field 
	Methods []Function 
	Doc ai.Documentation 
}
```

**Summary:** Defines a Struct with fields, methods, and documentation

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
s := Struct{Name: "User", Fields: []Field{...}}
```

**Edge Cases:**
- Empty Fields or Methods slices
- Nil Doc field


---

### `Field`

```go
type Field struct {
	Name string 
	Type string 
	Tag string 
	Doc ai.Documentation 
}
```

**Summary:** Defines a Field with name, type, tag, and documentation

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
f := Field{Name: "ID", Type: "int", Tag: "json:\"id\""}
```

**Edge Cases:**
- Empty Type or Tag strings
- Nil Doc field


---

### `Function`

```go
type Function struct {
	Name string 
	Receiver string 
	Parameters []Parameter 
	Results []Parameter 
	Doc ai.Documentation 
	Calls []string 
}
```

**Summary:** Struct representing a function with metadata and calls

**Parameters:**
- `Name` (string): Function name
- `Receiver` (string): Method receiver type
- `Parameters` ([]Parameter): Function parameters
- `Results` ([]Parameter): Return types
- `Doc` (ai.Documentation): Associated documentation
- `Calls` ([]string): Names of called functions

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1) for instantiation
- Space: O(n) where n is total parameters/results/calls

**Example:**
```go
funcObj := Function{Name: "Calculate", Parameters: []Parameter{{Name: "x", Type: "int"}}}
```

**Edge Cases:**
- Empty function name
- Nil slices for Parameters/Results/Calls


---

### `Parameter`

```go
type Parameter struct {
	Name string 
	Type string 
}
```

**Summary:** Struct representing a function parameter or return type

**Parameters:**
- `Name` (string): Parameter name
- `Type` (string): Parameter type

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1) for instantiation
- Space: O(1)

**Example:**
```go
param := Parameter{Name: "userId", Type: "string"}
```

**Edge Cases:**
- Empty type string
- Unsanitized special characters in name


---

## 🔧 Functions

<details>
<summary><b><code>init()</code></b></summary>

**Summary:** Analyzes Go project directory and returns results

**Parameters:**
- `projectDir` (string): Path to project root directory

**Returns:** *AnalyzerResult with analysis data or error if failed

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is total code elements analyzed

**Example:**
```go
result, err := AnalyzeProject("./myproject")
```

**Edge Cases:**
- Non-existent directory path
- Insufficient permissions
- Malformed Go files


</details>

<details>
<summary><b><code>AnalyzeProject(projectDir string)</code></b></summary>

**Summary:** Analyzes Go project directory and returns results

**Parameters:**
- `projectDir` (string): Path to project root directory

**Returns:** *AnalyzerResult with analysis data or error if failed

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is total code elements analyzed

**Example:**
```go
result, err := AnalyzeProject("./myproject")
```

**Edge Cases:**
- Non-existent directory path
- Insufficient permissions
- Malformed Go files


</details>


---

## 📄 File: `go_analyzer.go`

> 📍 `analyzer\go_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (6)](#-functions)

## 🧱 Structs

### `GoAnalyzer`

```go
type GoAnalyzer struct {
}
```

**Summary:** Empty struct defining a Go analyzer type

**Returns:** None (type definition only)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := GoAnalyzer{} // creates an instance
```

**Edge Cases:**
- No functionality (empty struct)


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if analyzer supports a project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) (assuming simple checks)
- Space: O(1)

**Example:**
```go
supported := analyzer.Supports("./myproject")
```

**Edge Cases:**
- Nonexistent directory path
- Empty/malformed path string


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes a Go project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** AnalyzerResult pointer and error (nil on success)

**Complexity:**
- Time: O(n) (depends on project size)
- Space: O(n) (depends on analysis output)

**Example:**
```go
result, err := analyzer.Analyze("./myproject")
```

**Edge Cases:**
- Nonexistent directory
- Permission errors
- Malformed Go code


</details>

<details>
<summary><b><code>Merge(other *AnalyzerResult)</code></b></summary>

**Summary:** Merges another AnalyzerResult into the current one

**Parameters:**
- `other` (*AnalyzerResult): Result to merge into the current instance

**Returns:** None (modifies the receiver in-place)

**Complexity:**
- Time: O(n) where n is size of other result
- Space: O(1) (modifies existing structure)

**Example:**
```go
result1.Merge(result2) // combines analysis results
```

**Edge Cases:**
- Nil input parameter
- Conflicting data between results


</details>

<details>
<summary><b><code>exists(path string)</code></b></summary>

**Summary:** Checks if a file/path exists

**Parameters:**
- `path` (string): Filesystem path to check

**Returns:** Boolean indicating existence

**Complexity:**
- Time: O(1) (filesystem operation)
- Space: O(1)

**Example:**
```go
if exists("/tmp/file") { ... }
```

**Edge Cases:**
- Empty path string
- Permission issues
- Symbolic links


</details>

<details>
<summary><b><code>AnalyzeFile(path string)</code></b></summary>

**Summary:** Converts an AST basic literal tag to its string representation

**Parameters:**
- `tag` (*ast.BasicLit): AST node representing a basic literal

**Returns:** String representation of the tag

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
str := TagToString(astNode) // returns "foo" for tag `foo`
```

**Edge Cases:**
- Nil input node
- Empty or malformed tag values


</details>

<details>
<summary><b><code>extractParams(fl *ast.FieldList)</code></b></summary>

**Summary:** Extracts parameters from an AST field list

**Parameters:**
- `fl` (*ast.FieldList): AST field list to extract parameters from

**Returns:** Slice of Parameter structs representing extracted parameters

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for storing parameters

**Example:**
```go
params := extractParams(funcDecl.Type.Params)
```

**Edge Cases:**
- Empty field list returns empty slice
- Nil input may cause panic


</details>


---

## 📄 File: `js_analyzer.go`

> 📍 `analyzer\js_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (8)](#-functions)

## 🧱 Structs

### `JSAnalyzer`

```go
type JSAnalyzer struct {
}
```

**Summary:** Empty struct for JavaScript code analysis

**Returns:** N/A (type declaration)

**Complexity:**
- Time: N/A
- Space: O(1) for instance creation

**Example:**
```go
analyzer := JSAnalyzer{}
```

**Edge Cases:**
- No immediate edge cases for empty struct


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if analyzer supports given project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) assuming simple checks
- Space: O(1)

**Example:**
```go
if analyzer.Supports("./my-project") { ... }
```

**Edge Cases:**
- Empty path string
- Non-existent directory
- Permission issues


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes JavaScript project directory and returns results

**Parameters:**
- `projectDir` (string): Path to the project directory

**Returns:** AnalyzerResult pointer and error, containing analysis results

**Complexity:**
- Time: O(n) where n is number of files
- Space: O(m) where m is size of analysis data

**Example:**
```go
result, err := analyzer.Analyze("./my-js-project")
```

**Edge Cases:**
- Non-existent directory path
- Permission issues accessing files
- Malformed JavaScript files


</details>

<details>
<summary><b><code>extractJSFunctions(src string)</code></b></summary>

**Summary:** Extracts all function declarations from JS source

**Parameters:**
- `src` (string): JavaScript source code

**Returns:** Slice of Function objects representing found functions

**Complexity:**
- Time: O(n) where n is source length
- Space: O(k) where k is number of functions

**Example:**
```go
funcs := extractJSFunctions("function foo(){}")
```

**Edge Cases:**
- Empty source string
- Nested functions
- Arrow functions
- Invalid syntax


</details>

<details>
<summary><b><code>extractJSFunctionBody(src string, funcDecl string)</code></b></summary>

**Summary:** Extracts body of specified function from JS source

**Parameters:**
- `src` (string): JavaScript source code
- `funcDecl` (string): Target function declaration

**Returns:** String containing the function body

**Complexity:**
- Time: O(n) where n is source length
- Space: O(m) where m is body size

**Example:**
```go
body := extractJSFunctionBody("function foo(){ return 1; }", "foo")
```

**Edge Cases:**
- Function not found
- Multiple matching functions
- Nested function with same name
- Invalid function declaration format


</details>

<details>
<summary><b><code>extractJSClasses(src string)</code></b></summary>

**Summary:** Extracts JavaScript class structures from source code

**Parameters:**
- `src` (string): JavaScript source code to parse

**Returns:** Array of Struct objects representing class definitions

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(m) where m is number of classes found

**Example:**
```go
classes := extractJSClasses('class A {} class B {}')
```

**Edge Cases:**
- Malformed JavaScript syntax
- Nested class declarations
- Source containing non-class code


</details>

<details>
<summary><b><code>extractJsClassBody(src string, classDecl string)</code></b></summary>

**Summary:** Extracts body of a specific JavaScript class

**Parameters:**
- `src` (string): JavaScript source code
- `classDecl` (string): Class declaration to match

**Returns:** String containing the matched class body content

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(k) where k is body length

**Example:**
```go
body := extractJsClassBody('class X { method() {} }', 'class X')
```

**Edge Cases:**
- Class declaration not found
- Partial/incomplete class declarations
- Multiple matching declarations


</details>

<details>
<summary><b><code>parseParamList(_ string, raw string)</code></b></summary>

**Summary:** Parses parameter list string into structured objects

**Parameters:**
- `_` (string): Unused parameter (likely placeholder)
- `raw` (string): Raw parameter list string

**Returns:** Array of Parameter objects with parsed details

**Complexity:**
- Time: O(n) where n is parameter list length
- Space: O(m) where m is number of parameters

**Example:**
```go
params := parseParamList('', 'a, b=5, ...rest')
```

**Edge Cases:**
- Empty parameter list
- Complex default values
- Rest parameters
- Malformed syntax


</details>

<details>
<summary><b><code>findDocBefore(context string, src string, docRegex *regexp.Regexp)</code></b></summary>

**Summary:** Finds documentation before a context string using regex

**Parameters:**
- `context` (string): Context string to search before
- `src` (string): Source text to search in
- `docRegex` (*regexp.Regexp): Regex pattern for documentation

**Returns:** Documentation string found before the context

**Complexity:**
- Time: O(n), where n is the length of src (regex dependent)
- Space: O(1) (excluding regex storage)

**Example:**
```go
doc := findDocBefore("func example()", "/* doc */ func example()", regexp.MustCompile(`/\*.*?\*/`))
```

**Edge Cases:**
- Empty src string
- No match for docRegex
- Context string not found in src


</details>


---

## 📄 File: `python_analyzer.go`

> 📍 `analyzer\python_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (7)](#-functions)

## 🧱 Structs

### `PythonAnalyzer`

```go
type PythonAnalyzer struct {
}
```

**Summary:** Empty struct for Python code analysis

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := PythonAnalyzer{}
```

**Edge Cases:**
- None (empty struct)


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if PythonAnalyzer supports a project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) (implementation dependent)
- Space: O(1)

**Example:**
```go
supported := analyzer.Supports("/path/to/project")
```

**Edge Cases:**
- Non-existent directory
- Empty directory
- Permission issues


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes Python project directory and returns results

**Parameters:**
- `projectDir` (string): Path to Python project directory

**Returns:** AnalyzerResult pointer and error (nil if successful)

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is analysis data size

**Example:**
```go
result, err := analyzer.Analyze("./my_project")
```

**Edge Cases:**
- Non-existent directory path
- Invalid Python project structure
- Permission issues


</details>

<details>
<summary><b><code>extractPythonClasses(src string)</code></b></summary>

**Summary:** Extracts Python class structures from source code

**Parameters:**
- `src` (string): Python source code text

**Returns:** Slice of Struct objects representing classes

**Complexity:**
- Time: O(n) where n is source length
- Space: O(k) where k is class count

**Example:**
```go
classes := extractPythonClasses("class Foo: pass")
```

**Edge Cases:**
- Malformed class syntax
- Nested class declarations
- Source with no classes


</details>

<details>
<summary><b><code>extractPythonClassBody(src string, classDecl string)</code></b></summary>

**Summary:** Extracts body of specified Python class

**Parameters:**
- `src` (string): Python source code text
- `classDecl` (string): Target class declaration

**Returns:** Class body as string (excluding declaration)

**Complexity:**
- Time: O(n) where n is source length
- Space: O(m) where m is body size

**Example:**
```go
body := extractPythonClassBody(src, "class MyClass")
```

**Edge Cases:**
- Class not found in source
- Malformed class syntax
- Multiple matching declarations


</details>

<details>
<summary><b><code>extractPythonFunctions(src string)</code></b></summary>

**Summary:** Extracts Python function declarations from source code

**Parameters:**
- `src` (string): Python source code to parse

**Returns:** Slice of Function structs containing function metadata

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(m) where m is number of functions found

**Example:**
```go
funcs := extractPythonFunctions('def foo(): pass\ndef bar(): pass')
```

**Edge Cases:**
- Malformed function declarations
- Nested functions
- Source with syntax errors


</details>

<details>
<summary><b><code>extractFunctionBody(src string, funcDecl string)</code></b></summary>

**Summary:** Extracts body of a specific function from source

**Parameters:**
- `src` (string): Source code containing the function
- `funcDecl` (string): Function declaration to match

**Returns:** String containing the matched function's body

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(k) where k is body length

**Example:**
```go
body := extractFunctionBody('def foo():\n    return 42', 'def foo():')
```

**Edge Cases:**
- Multiple matching declarations
- Function not found
- Incorrect indentation in source


</details>

<details>
<summary><b><code>findDocstringAfter(context string, src string, docRegex *regexp.Regexp)</code></b></summary>

**Summary:** Finds docstring following specific context in source

**Parameters:**
- `context` (string): Preceding code pattern
- `src` (string): Source code to search
- `docRegex` (*regexp.Regexp): Pattern matching docstring format

**Returns:** Matched docstring or empty string

**Complexity:**
- Time: O(n) where n is source length
- Space: O(1) for search, O(m) for result where m is docstring length

**Example:**
```go
doc := findDocstringAfter('def foo():', source, regexp.MustCompile(`"""(.+?)"""`))
```

**Edge Cases:**
- Context appears multiple times
- Malformed docstrings
- Context not found


</details>


---

## 📄 File: `yaml_analyzer.go`

> 📍 `analyzer\yaml_analyzer.go`

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


---

## 📄 File: `analyzer.go`

> 📍 `analyzer\analyzer.go`

## 📑 Contents

- [🧱 Structs (7)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `AnalyzerResult`

```go
type AnalyzerResult struct {
	Files []*AnalyzedFile 
}
```

_No documentation available._

---

### `AnalyzedFile`

```go
type AnalyzedFile struct {
	Path string 
	Packages []Package 
}
```

_No documentation available._

---

### `Package`

```go
type Package struct {
	Name string 
	Path string 
	Imports []string 
	Structs []Struct 
	Funcs []Function 
	Files []string 
}
```

_No documentation available._

---

### `Struct`

```go
type Struct struct {
	Name string 
	Fields []Field 
	Methods []Function 
	Doc ai.Documentation 
	DocYAML ai.YAMLDocumentation 
}
```

_No documentation available._

---

### `Field`

```go
type Field struct {
	Name string 
	Type string 
	Tag string 
	Doc ai.Documentation 
	DocYAML ai.YAMLDocumentation 
	Value string 
	Fields []Field 
}
```

_No documentation available._

---

### `Function`

```go
type Function struct {
	Name string 
	Receiver string 
	Parameters []Parameter 
	Results []Parameter 
	Doc ai.Documentation 
	Calls []string 
}
```

_No documentation available._

---

### `Parameter`

```go
type Parameter struct {
	Name string 
	Type string 
}
```

_No documentation available._

---

## 🔧 Functions

<details>
<summary><b><code>init()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>AnalyzeProject(projectDir string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>yamlPrintFileAnalysis(file *AnalyzedFile)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>yamlPrintField(f Field, indent int)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `go_analyzer.go`

> 📍 `analyzer\go_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (6)](#-functions)

## 🧱 Structs

### `GoAnalyzer`

```go
type GoAnalyzer struct {
}
```

_No documentation available._

---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Merge(other *AnalyzerResult)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>exists(path string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>AnalyzeFile(path string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractParams(fl *ast.FieldList)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `js_analyzer.go`

> 📍 `analyzer\js_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (8)](#-functions)

## 🧱 Structs

### `JSAnalyzer`

```go
type JSAnalyzer struct {
}
```

_No documentation available._

---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractJSFunctions(src string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractJSFunctionBody(src string, funcDecl string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractJSClasses(src string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractJsClassBody(src string, classDecl string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>parseParamList(_ string, raw string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>findDocBefore(context string, src string, docRegex *regexp.Regexp)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `python_analyzer.go`

> 📍 `analyzer\python_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (7)](#-functions)

## 🧱 Structs

### `PythonAnalyzer`

```go
type PythonAnalyzer struct {
}
```

_No documentation available._

---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractPythonClasses(src string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractPythonClassBody(src string, classDecl string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractPythonFunctions(src string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractFunctionBody(src string, funcDecl string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>findDocstringAfter(context string, src string, docRegex *regexp.Regexp)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `yaml_analyzer.go`

> 📍 `analyzer\yaml_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (25)](#-functions)

## 🧱 Structs

### `YAMLAnalyzer`

```go
type YAMLAnalyzer struct {
}
```

_No documentation available._

---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>analyzeYAMLNode(node *yaml.Node, baseName string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>isKubernetesResource(node *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>isHelmChart(node *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>isAnsiblePlaybook(node *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>isDockerCompose(node *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>analyzeKubernetesResource(node *yaml.Node, baseName string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractKubernetesSpec(node *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>findSpecNode(node *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>analyzeHelmChart(node *yaml.Node, baseName string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>analyzeAnsiblePlaybook(node *yaml.Node, baseName string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractAnsiblePlayFields(playNode *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractAnsibleTasks(tasksNode *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>analyzeDockerCompose(node *yaml.Node, baseName string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractDockerServices(servicesNode *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>getValueNode(node *yaml.Node, key string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractArrayItems(node *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractYAMLFields(node *yaml.Node, depth int)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>extractYAMLDescription(node *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>hasKey(node *yaml.Node, key string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>getValue(node *yaml.Node, key string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>getMappingKeys(node *yaml.Node)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>nodeKindToString(kind yaml.Kind)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>sequenceItemName(index int)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `analyzer.go`

> 📍 `analyzer\analyzer.go`

## 📑 Contents

- [🧱 Structs (7)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `AnalyzerResult`

```go
type AnalyzerResult struct {
	Files []*AnalyzedFile 
}
```

**Summary:** Container for multiple file analysis results

**Returns:** None (type definition)

**Complexity:**
- Time: N/A (data structure)
- Space: O(n) where n is number of files

**Example:**
```go
result := AnalyzerResult{Files: analyzedFilesSlice}
```

**Edge Cases:**
- Empty Files slice
- Nil elements in Files array
- Concurrent access/modification


---

### `AnalyzedFile`

```go
type AnalyzedFile struct {
	Path string 
	Packages []Package 
}
```

**Summary:** Stores analysis data for a single file

**Returns:** None (type definition)

**Complexity:**
- Time: N/A (data structure)
- Space: O(m) where m is number of packages

**Example:**
```go
file := AnalyzedFile{Path: "/src/main.go", Packages: pkgs}
```

**Edge Cases:**
- Empty Path string
- Nil Packages slice
- Duplicate package entries


---

### `Package`

```go
type Package struct {
	Name string 
	Path string 
	Imports []string 
	Structs []Struct 
	Funcs []Function 
	Files []string 
}
```

**Summary:** Defines a Package structure with metadata and components

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
pkg := Package{Name: "example", Path: "github.com/example"}
```

**Edge Cases:**
- Empty slices for Imports, Structs, Funcs, or Files
- Nil values in any field


---

### `Struct`

```go
type Struct struct {
	Name string 
	Fields []Field 
	Methods []Function 
	Doc ai.Documentation 
	DocYAML ai.YAMLDocumentation 
}
```

**Summary:** Defines a struct type with name, fields, methods, and documentation.

**Parameters:**
- `Name` (string): Name of the struct
- `Fields` ([]Field): List of fields in the struct
- `Methods` ([]Function): List of methods associated with the struct
- `Doc` (ai.Documentation): Documentation for the struct
- `DocYAML` (ai.YAMLDocumentation): YAML-formatted documentation for the struct

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
type MyStruct Struct
```

**Edge Cases:**
- Empty or nil fields/methods
- Missing documentation


---

### `Field`

```go
type Field struct {
	Name string 
	Type string 
	Tag string 
	Doc ai.Documentation 
	DocYAML ai.YAMLDocumentation 
	Value string 
	Fields []Field 
}
```

**Summary:** Defines a field type with name, type, tag, and nested fields.

**Parameters:**
- `Name` (string): Name of the field
- `Type` (string): Type of the field
- `Tag` (string): Tag associated with the field
- `Doc` (ai.Documentation): Documentation for the field
- `DocYAML` (ai.YAMLDocumentation): YAML-formatted documentation for the field
- `Value` (string): Default value of the field
- `Fields` ([]Field): Nested fields (for complex types)

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
field := Field{Name: "Age", Type: "int"}
```

**Edge Cases:**
- Recursive field definitions
- Empty or invalid type/tag


---

### `Function`

```go
type Function struct {
	Name string 
	Receiver string 
	Parameters []Parameter 
	Results []Parameter 
	Doc ai.Documentation 
	Calls []string 
}
```

**Summary:** Struct representing a function with metadata and calls

**Parameters:**
- `Name` (string): Function name
- `Receiver` (string): Method receiver type
- `Parameters` ([]Parameter): Function parameters
- `Results` ([]Parameter): Return types
- `Doc` (ai.Documentation): Associated documentation
- `Calls` ([]string): Names of called functions

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1) for instantiation
- Space: O(n) where n is total parameters/results/calls

**Example:**
```go
funcObj := Function{Name: "Calculate", Parameters: []Parameter{{Name: "x", Type: "int"}}}
```

**Edge Cases:**
- Empty function name
- Nil slices for Parameters/Results/Calls


---

### `Parameter`

```go
type Parameter struct {
	Name string 
	Type string 
}
```

**Summary:** Struct representing a function parameter or return type

**Parameters:**
- `Name` (string): Parameter name
- `Type` (string): Parameter type

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1) for instantiation
- Space: O(1)

**Example:**
```go
param := Parameter{Name: "userId", Type: "string"}
```

**Edge Cases:**
- Empty type string
- Unsanitized special characters in name


---

## 🔧 Functions

<details>
<summary><b><code>init()</code></b></summary>

**Summary:** Analyzes Go project directory and returns results

**Parameters:**
- `projectDir` (string): Path to project root directory

**Returns:** *AnalyzerResult with analysis data or error if failed

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is total code elements analyzed

**Example:**
```go
result, err := AnalyzeProject("./myproject")
```

**Edge Cases:**
- Non-existent directory path
- Insufficient permissions
- Malformed Go files


</details>

<details>
<summary><b><code>AnalyzeProject(projectDir string)</code></b></summary>

**Summary:** Analyzes Go project directory and returns results

**Parameters:**
- `projectDir` (string): Path to project root directory

**Returns:** *AnalyzerResult with analysis data or error if failed

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is total code elements analyzed

**Example:**
```go
result, err := AnalyzeProject("./myproject")
```

**Edge Cases:**
- Non-existent directory path
- Insufficient permissions
- Malformed Go files


</details>

<details>
<summary><b><code>yamlPrintFileAnalysis(file *AnalyzedFile)</code></b></summary>

**Summary:** Empty struct defining a Go analyzer type

**Returns:** None (type definition only)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := GoAnalyzer{} // creates an instance
```

**Edge Cases:**
- No functionality (empty struct)


</details>

<details>
<summary><b><code>yamlPrintField(f Field, indent int)</code></b></summary>

**Summary:** Prints a field's details in YAML format with indentation.

**Parameters:**
- `f` (Field): Field to print
- `indent` (int): Indentation level

**Returns:** None (prints to stdout)

**Complexity:**
- Time: O(n) where n is the number of nested fields
- Space: O(1) (excluding recursion stack)

**Example:**
```go
yamlPrintField(Field{Name: "ID", Type: "string"}, 2)
```

**Edge Cases:**
- Deeply nested fields causing stack overflow
- Large indentation values


</details>


---

## 📄 File: `go_analyzer.go`

> 📍 `analyzer\go_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (6)](#-functions)

## 🧱 Structs

### `GoAnalyzer`

```go
type GoAnalyzer struct {
}
```

**Summary:** Empty struct defining a Go code analyzer type

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := GoAnalyzer{}
```

**Edge Cases:**
- No functionality implemented yet
- Struct may be extended in future


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if analyzer supports a project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) (assuming simple checks)
- Space: O(1)

**Example:**
```go
supported := analyzer.Supports("./myproject")
```

**Edge Cases:**
- Nonexistent directory path
- Empty/malformed path string


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes a Go project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** AnalyzerResult pointer and error (nil on success)

**Complexity:**
- Time: O(n) (depends on project size)
- Space: O(n) (depends on analysis output)

**Example:**
```go
result, err := analyzer.Analyze("./myproject")
```

**Edge Cases:**
- Nonexistent directory
- Permission errors
- Malformed Go code


</details>

<details>
<summary><b><code>Merge(other *AnalyzerResult)</code></b></summary>

**Summary:** Merges another AnalyzerResult into the current one

**Parameters:**
- `other` (*AnalyzerResult): Result to merge into the current instance

**Returns:** None (modifies the receiver in-place)

**Complexity:**
- Time: O(n) where n is size of other result
- Space: O(1) (modifies existing structure)

**Example:**
```go
result1.Merge(result2) // combines analysis results
```

**Edge Cases:**
- Nil input parameter
- Conflicting data between results


</details>

<details>
<summary><b><code>exists(path string)</code></b></summary>

**Summary:** Checks if a file/path exists

**Parameters:**
- `path` (string): Filesystem path to check

**Returns:** Boolean indicating existence

**Complexity:**
- Time: O(1) (filesystem operation)
- Space: O(1)

**Example:**
```go
if exists("/tmp/file") { ... }
```

**Edge Cases:**
- Empty path string
- Permission issues
- Symbolic links


</details>

<details>
<summary><b><code>AnalyzeFile(path string)</code></b></summary>

**Summary:** Converts an AST basic literal tag to its string representation

**Parameters:**
- `tag` (*ast.BasicLit): AST node representing a basic literal

**Returns:** String representation of the tag

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
str := TagToString(astNode) // returns "foo" for tag `foo`
```

**Edge Cases:**
- Nil input node
- Empty or malformed tag values


</details>

<details>
<summary><b><code>extractParams(fl *ast.FieldList)</code></b></summary>

**Summary:** Extracts parameters from an AST field list

**Parameters:**
- `fl` (*ast.FieldList): AST field list to extract parameters from

**Returns:** Slice of Parameter structs representing extracted parameters

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for storing parameters

**Example:**
```go
params := extractParams(funcDecl.Type.Params)
```

**Edge Cases:**
- Empty field list returns empty slice
- Nil input may cause panic


</details>


---

## 📄 File: `js_analyzer.go`

> 📍 `analyzer\js_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (8)](#-functions)

## 🧱 Structs

### `JSAnalyzer`

```go
type JSAnalyzer struct {
}
```

**Summary:** Empty struct for JavaScript code analysis

**Returns:** N/A (type declaration)

**Complexity:**
- Time: N/A
- Space: O(1) for instance creation

**Example:**
```go
analyzer := JSAnalyzer{}
```

**Edge Cases:**
- No immediate edge cases for empty struct


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if analyzer supports given project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) assuming simple checks
- Space: O(1)

**Example:**
```go
if analyzer.Supports("./my-project") { ... }
```

**Edge Cases:**
- Empty path string
- Non-existent directory
- Permission issues


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes JavaScript project directory and returns results

**Parameters:**
- `projectDir` (string): Path to the project directory

**Returns:** AnalyzerResult pointer and error, containing analysis results

**Complexity:**
- Time: O(n) where n is number of files
- Space: O(m) where m is size of analysis data

**Example:**
```go
result, err := analyzer.Analyze("./my-js-project")
```

**Edge Cases:**
- Non-existent directory path
- Permission issues accessing files
- Malformed JavaScript files


</details>

<details>
<summary><b><code>extractJSFunctions(src string)</code></b></summary>

**Summary:** Extracts all function declarations from JS source

**Parameters:**
- `src` (string): JavaScript source code

**Returns:** Slice of Function objects representing found functions

**Complexity:**
- Time: O(n) where n is source length
- Space: O(k) where k is number of functions

**Example:**
```go
funcs := extractJSFunctions("function foo(){}")
```

**Edge Cases:**
- Empty source string
- Nested functions
- Arrow functions
- Invalid syntax


</details>

<details>
<summary><b><code>extractJSFunctionBody(src string, funcDecl string)</code></b></summary>

**Summary:** Extracts body of specified function from JS source

**Parameters:**
- `src` (string): JavaScript source code
- `funcDecl` (string): Target function declaration

**Returns:** String containing the function body

**Complexity:**
- Time: O(n) where n is source length
- Space: O(m) where m is body size

**Example:**
```go
body := extractJSFunctionBody("function foo(){ return 1; }", "foo")
```

**Edge Cases:**
- Function not found
- Multiple matching functions
- Nested function with same name
- Invalid function declaration format


</details>

<details>
<summary><b><code>extractJSClasses(src string)</code></b></summary>

**Summary:** Extracts JavaScript class structures from source code

**Parameters:**
- `src` (string): JavaScript source code to parse

**Returns:** Array of Struct objects representing class definitions

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(m) where m is number of classes found

**Example:**
```go
classes := extractJSClasses('class A {} class B {}')
```

**Edge Cases:**
- Malformed JavaScript syntax
- Nested class declarations
- Source containing non-class code


</details>

<details>
<summary><b><code>extractJsClassBody(src string, classDecl string)</code></b></summary>

**Summary:** Extracts body of a specific JavaScript class

**Parameters:**
- `src` (string): JavaScript source code
- `classDecl` (string): Class declaration to match

**Returns:** String containing the matched class body content

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(k) where k is body length

**Example:**
```go
body := extractJsClassBody('class X { method() {} }', 'class X')
```

**Edge Cases:**
- Class declaration not found
- Partial/incomplete class declarations
- Multiple matching declarations


</details>

<details>
<summary><b><code>parseParamList(_ string, raw string)</code></b></summary>

**Summary:** Parses parameter list string into structured objects

**Parameters:**
- `_` (string): Unused parameter (likely placeholder)
- `raw` (string): Raw parameter list string

**Returns:** Array of Parameter objects with parsed details

**Complexity:**
- Time: O(n) where n is parameter list length
- Space: O(m) where m is number of parameters

**Example:**
```go
params := parseParamList('', 'a, b=5, ...rest')
```

**Edge Cases:**
- Empty parameter list
- Complex default values
- Rest parameters
- Malformed syntax


</details>

<details>
<summary><b><code>findDocBefore(context string, src string, docRegex *regexp.Regexp)</code></b></summary>

**Summary:** Finds documentation before a context string using regex

**Parameters:**
- `context` (string): Context string to search before
- `src` (string): Source text to search in
- `docRegex` (*regexp.Regexp): Regex pattern for documentation

**Returns:** Documentation string found before the context

**Complexity:**
- Time: O(n), where n is the length of src (regex dependent)
- Space: O(1) (excluding regex storage)

**Example:**
```go
doc := findDocBefore("func example()", "/* doc */ func example()", regexp.MustCompile(`/\*.*?\*/`))
```

**Edge Cases:**
- Empty src string
- No match for docRegex
- Context string not found in src


</details>


---

## 📄 File: `python_analyzer.go`

> 📍 `analyzer\python_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (7)](#-functions)

## 🧱 Structs

### `PythonAnalyzer`

```go
type PythonAnalyzer struct {
}
```

**Summary:** Empty struct for Python code analysis

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := PythonAnalyzer{}
```

**Edge Cases:**
- None (empty struct)


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if PythonAnalyzer supports a project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) (implementation dependent)
- Space: O(1)

**Example:**
```go
supported := analyzer.Supports("/path/to/project")
```

**Edge Cases:**
- Non-existent directory
- Empty directory
- Permission issues


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes Python project directory and returns results

**Parameters:**
- `projectDir` (string): Path to Python project directory

**Returns:** AnalyzerResult pointer and error (nil if successful)

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is analysis data size

**Example:**
```go
result, err := analyzer.Analyze("./my_project")
```

**Edge Cases:**
- Non-existent directory path
- Invalid Python project structure
- Permission issues


</details>

<details>
<summary><b><code>extractPythonClasses(src string)</code></b></summary>

**Summary:** Extracts Python class structures from source code

**Parameters:**
- `src` (string): Python source code text

**Returns:** Slice of Struct objects representing classes

**Complexity:**
- Time: O(n) where n is source length
- Space: O(k) where k is class count

**Example:**
```go
classes := extractPythonClasses("class Foo: pass")
```

**Edge Cases:**
- Malformed class syntax
- Nested class declarations
- Source with no classes


</details>

<details>
<summary><b><code>extractPythonClassBody(src string, classDecl string)</code></b></summary>

**Summary:** Extracts body of specified Python class

**Parameters:**
- `src` (string): Python source code text
- `classDecl` (string): Target class declaration

**Returns:** Class body as string (excluding declaration)

**Complexity:**
- Time: O(n) where n is source length
- Space: O(m) where m is body size

**Example:**
```go
body := extractPythonClassBody(src, "class MyClass")
```

**Edge Cases:**
- Class not found in source
- Malformed class syntax
- Multiple matching declarations


</details>

<details>
<summary><b><code>extractPythonFunctions(src string)</code></b></summary>

**Summary:** Extracts Python function declarations from source code

**Parameters:**
- `src` (string): Python source code to parse

**Returns:** Slice of Function structs containing function metadata

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(m) where m is number of functions found

**Example:**
```go
funcs := extractPythonFunctions('def foo(): pass\ndef bar(): pass')
```

**Edge Cases:**
- Malformed function declarations
- Nested functions
- Source with syntax errors


</details>

<details>
<summary><b><code>extractFunctionBody(src string, funcDecl string)</code></b></summary>

**Summary:** Extracts body of a specific function from source

**Parameters:**
- `src` (string): Source code containing the function
- `funcDecl` (string): Function declaration to match

**Returns:** String containing the matched function's body

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(k) where k is body length

**Example:**
```go
body := extractFunctionBody('def foo():\n    return 42', 'def foo():')
```

**Edge Cases:**
- Multiple matching declarations
- Function not found
- Incorrect indentation in source


</details>

<details>
<summary><b><code>findDocstringAfter(context string, src string, docRegex *regexp.Regexp)</code></b></summary>

**Summary:** Finds docstring following specific context in source

**Parameters:**
- `context` (string): Preceding code pattern
- `src` (string): Source code to search
- `docRegex` (*regexp.Regexp): Pattern matching docstring format

**Returns:** Matched docstring or empty string

**Complexity:**
- Time: O(n) where n is source length
- Space: O(1) for search, O(m) for result where m is docstring length

**Example:**
```go
doc := findDocstringAfter('def foo():', source, regexp.MustCompile(`"""(.+?)"""`))
```

**Edge Cases:**
- Context appears multiple times
- Malformed docstrings
- Context not found


</details>


---

## 📄 File: `yaml_analyzer.go`

> 📍 `analyzer\yaml_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (25)](#-functions)

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
<summary><b><code>analyzeYAMLNode(node *yaml.Node, baseName string)</code></b></summary>

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

<details>
<summary><b><code>isKubernetesResource(node *yaml.Node)</code></b></summary>

**Summary:** Checks if YAML node represents a Kubernetes resource

**Parameters:**
- `node` (*yaml.Node): YAML node to inspect

**Returns:** Boolean indicating Kubernetes resource status

**Complexity:**
- Time: O(n) where n is node depth/complexity
- Space: O(1)

**Example:**
```go
isKube := analyzer.isKubernetesResource(yamlNode)
```

**Edge Cases:**
- Nil node input
- Malformed YAML structure
- Non-Kubernetes YAML with similar fields


</details>

<details>
<summary><b><code>isHelmChart(node *yaml.Node)</code></b></summary>

**Summary:** Determines if YAML node is a Helm chart definition

**Parameters:**
- `node` (*yaml.Node): YAML node to evaluate

**Returns:** Boolean indicating Helm chart status

**Complexity:**
- Time: O(n) where n is node depth/complexity
- Space: O(1)

**Example:**
```go
isHelm := analyzer.isHelmChart(yamlNode)
```

**Edge Cases:**
- Nil node input
- Partial Helm chart structures
- Version compatibility issues


</details>

<details>
<summary><b><code>isAnsiblePlaybook(node *yaml.Node)</code></b></summary>

**Summary:** Checks if YAML node represents an Ansible playbook

**Parameters:**
- `node` (*yaml.Node): YAML node to analyze

**Returns:** Boolean indicating Ansible playbook status

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
isPlaybook := analyzer.isAnsiblePlaybook(yamlNode)
```

**Edge Cases:**
- Empty YAML node
- Malformed YAML structure


</details>

<details>
<summary><b><code>isDockerCompose(node *yaml.Node)</code></b></summary>

**Summary:** Determines if YAML node is a Docker Compose file

**Parameters:**
- `node` (*yaml.Node): YAML node to check

**Returns:** Boolean indicating Docker Compose format

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
isCompose := analyzer.isDockerCompose(yamlNode)
```

**Edge Cases:**
- Partial Docker Compose syntax
- Version mismatch in compose spec


</details>

<details>
<summary><b><code>analyzeKubernetesResource(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes YAML node for Kubernetes resource metadata

**Parameters:**
- `node` (*yaml.Node): YAML node to parse
- `baseName` (string): Base filename for context

**Returns:** Structured analysis result as custom type

**Complexity:**
- Time: O(n) where n is node children
- Space: O(n) for recursive parsing

**Example:**
```go
resource := analyzer.analyzeKubernetesResource(node, 'deployment.yaml')
```

**Edge Cases:**
- Custom resource definitions
- Multi-document YAML files
- Missing required Kubernetes fields


</details>

<details>
<summary><b><code>extractKubernetesSpec(node *yaml.Node)</code></b></summary>

**Summary:** Extracts Kubernetes spec fields from a YAML node

**Parameters:**
- `node` (*yaml.Node): YAML node to analyze

**Returns:** List of extracted fields ([]Field)

**Complexity:**
- Time: O(n) where n is the number of nodes in the YAML tree
- Space: O(m) where m is the number of fields extracted

**Example:**
```go
fields := analyzer.extractKubernetesSpec(yamlNode)
```

**Edge Cases:**
- Nil input node
- Malformed YAML structure
- Missing Kubernetes spec section


</details>

<details>
<summary><b><code>findSpecNode(node *yaml.Node)</code></b></summary>

**Summary:** Finds the spec node in a YAML document

**Parameters:**
- `node` (*yaml.Node): Root YAML node to search

**Returns:** Pointer to the found spec node (*yaml.Node) or nil

**Complexity:**
- Time: O(n) where n is the number of nodes in the YAML tree
- Space: O(1)

**Example:**
```go
specNode := analyzer.findSpecNode(rootNode)
```

**Edge Cases:**
- Nil input node
- Spec node not present
- Multiple spec nodes in document


</details>

<details>
<summary><b><code>analyzeHelmChart(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes a Helm chart's YAML structure

**Parameters:**
- `node` (*yaml.Node): Root YAML node of the chart
- `baseName` (string): Base name for the chart

**Returns:** Structured analysis result (Struct)

**Complexity:**
- Time: O(n) where n is the number of nodes in the YAML tree
- Space: O(m) where m is the size of the returned structure

**Example:**
```go
chartStruct := analyzer.analyzeHelmChart(chartNode, "my-chart")
```

**Edge Cases:**
- Nil input node
- Invalid Helm chart structure
- Empty baseName parameter


</details>

<details>
<summary><b><code>analyzeAnsiblePlaybook(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes an Ansible playbook YAML node and returns a structured representation.

**Parameters:**
- `node` (*yaml.Node): YAML node representing the playbook
- `baseName` (string): Base name for the playbook

**Returns:** Struct containing analyzed playbook data

**Complexity:**
- Time: O(n), where n is the number of nodes in the YAML
- Space: O(n), for storing the structured output

**Example:**
```go
result := analyzer.analyzeAnsiblePlaybook(yamlNode, "deploy.yml")
```

**Edge Cases:**
- Empty YAML node
- Malformed YAML structure
- Missing required playbook fields


</details>

<details>
<summary><b><code>extractAnsiblePlayFields(playNode *yaml.Node)</code></b></summary>

**Summary:** Extracts fields from an Ansible play YAML node.

**Parameters:**
- `playNode` (*yaml.Node): YAML node representing a play

**Returns:** Slice of Field objects representing play attributes

**Complexity:**
- Time: O(n), where n is the number of fields in the play
- Space: O(n), for storing the extracted fields

**Example:**
```go
fields := analyzer.extractAnsiblePlayFields(playNode)
```

**Edge Cases:**
- Empty play node
- Nested field structures
- Unsupported field types


</details>

<details>
<summary><b><code>extractAnsibleTasks(tasksNode *yaml.Node)</code></b></summary>

**Summary:** Extracts task fields from an Ansible tasks YAML node.

**Parameters:**
- `tasksNode` (*yaml.Node): YAML node containing tasks

**Returns:** Slice of Field objects representing task definitions

**Complexity:**
- Time: O(n), where n is the number of tasks
- Space: O(n), for storing the task fields

**Example:**
```go
tasks := analyzer.extractAnsibleTasks(tasksNode)
```

**Edge Cases:**
- Empty tasks section
- Tasks with complex nested structures
- Invalid task syntax


</details>

<details>
<summary><b><code>analyzeDockerCompose(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes Docker Compose YAML node and returns structured data

**Parameters:**
- `node` (*yaml.Node): YAML node to analyze
- `baseName` (string): Base name for the analysis

**Returns:** Structured analysis result as Struct

**Complexity:**
- Time: O(n) where n is YAML node complexity
- Space: O(n) for output structure

**Example:**
```go
result := analyzer.analyzeDockerCompose(doc.Root, "app")
```

**Edge Cases:**
- Nil node input
- Malformed YAML structure
- Empty baseName


</details>

<details>
<summary><b><code>extractDockerServices(servicesNode *yaml.Node)</code></b></summary>

**Summary:** Extracts Docker services from YAML services node

**Parameters:**
- `servicesNode` (*yaml.Node): YAML node containing services

**Returns:** Array of extracted service fields as []Field

**Complexity:**
- Time: O(n) where n is number of services
- Space: O(n) for output array

**Example:**
```go
services := analyzer.extractDockerServices(composeNode.Content[0])
```

**Edge Cases:**
- Nil servicesNode
- Non-service YAML structure
- Empty services section


</details>

<details>
<summary><b><code>getValueNode(node *yaml.Node, key string)</code></b></summary>

**Summary:** Finds YAML value node by key

**Parameters:**
- `node` (*yaml.Node): Parent YAML node
- `key` (string): Key to search for

**Returns:** Found value node or nil as *yaml.Node

**Complexity:**
- Time: O(n) where n is child nodes count
- Space: O(1)

**Example:**
```go
portNode := analyzer.getValueNode(serviceNode, "ports")
```

**Edge Cases:**
- Nil parent node
- Non-existent key
- Case-sensitive key matching


</details>

<details>
<summary><b><code>extractArrayItems(node *yaml.Node)</code></b></summary>

**Summary:** Extracts array items from a YAML node into Field objects

**Parameters:**
- `node` (*yaml.Node): YAML node containing array items

**Returns:** Slice of Field objects representing array items

**Complexity:**
- Time: O(n) where n is the number of array items
- Space: O(n) for storing extracted Field objects

**Example:**
```go
fields := analyzer.extractArrayItems(yamlArrayNode)
```

**Edge Cases:**
- Empty array node
- Nested array structures
- Non-array node input


</details>

<details>
<summary><b><code>extractYAMLFields(node *yaml.Node, depth int)</code></b></summary>

**Summary:** Recursively extracts YAML fields up to specified depth

**Parameters:**
- `node` (*yaml.Node): Root YAML node to analyze
- `depth` (int): Maximum recursion depth

**Returns:** Slice of Field objects representing YAML structure

**Complexity:**
- Time: O(n) where n is total nodes visited
- Space: O(d) where d is maximum depth (call stack)

**Example:**
```go
fields := analyzer.extractYAMLFields(rootNode, 3)
```

**Edge Cases:**
- Depth limit reached before complete traversal
- Cyclic references in YAML
- Very large depth values causing stack overflow


</details>

<details>
<summary><b><code>extractYAMLDescription(node *yaml.Node)</code></b></summary>

**Summary:** Extracts description text from a YAML node

**Parameters:**
- `node` (*yaml.Node): YAML node containing description

**Returns:** Extracted description as string

**Complexity:**
- Time: O(1) for simple nodes
- Space: O(1) for returned string

**Example:**
```go
desc := analyzer.extractYAMLDescription(commentNode)
```

**Edge Cases:**
- Node without description content
- Multi-line descriptions
- Special characters in description


</details>

<details>
<summary><b><code>hasKey(node *yaml.Node, key string)</code></b></summary>

**Summary:** Checks if a YAML node contains a specific key

**Parameters:**
- `node` (*yaml.Node): YAML node to inspect
- `key` (string): Key to search for

**Returns:** Boolean indicating if the key exists

**Complexity:**
- Time: O(n) where n is the number of keys in the node
- Space: O(1)

**Example:**
```go
exists := analyzer.hasKey(node, "version")
```

**Edge Cases:**
- Empty node
- Non-mapping node type
- Case-sensitive key matching


</details>

<details>
<summary><b><code>getValue(node *yaml.Node, key string)</code></b></summary>

**Summary:** Retrieves the value associated with a key in a YAML node

**Parameters:**
- `node` (*yaml.Node): YAML node to inspect
- `key` (string): Key whose value to retrieve

**Returns:** String value of the key, or empty string if not found

**Complexity:**
- Time: O(n) where n is the number of keys in the node
- Space: O(1)

**Example:**
```go
value := analyzer.getValue(node, "name")
```

**Edge Cases:**
- Key doesn't exist
- Non-string value type
- Nested mapping values


</details>

<details>
<summary><b><code>getMappingKeys(node *yaml.Node)</code></b></summary>

**Summary:** Extracts all keys from a YAML mapping node

**Parameters:**
- `node` (*yaml.Node): YAML mapping node

**Returns:** Slice of strings containing all keys

**Complexity:**
- Time: O(n) where n is the number of keys
- Space: O(n) for storing the keys

**Example:**
```go
keys := analyzer.getMappingKeys(node)
```

**Edge Cases:**
- Empty mapping node
- Non-mapping node type
- Duplicate keys in YAML


</details>

<details>
<summary><b><code>nodeKindToString(kind yaml.Kind)</code></b></summary>

**Summary:** Converts YAML node kind to its string representation

**Parameters:**
- `kind` (yaml.Kind): YAML node kind enum value

**Returns:** String representation of the YAML node kind

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
str := analyzer.nodeKindToString(yaml.SequenceNode) // returns "sequence"
```

**Edge Cases:**
- Unknown/unsupported yaml.Kind values
- Nil receiver (though method is on pointer receiver)


</details>

<details>
<summary><b><code>sequenceItemName(index int)</code></b></summary>

**Summary:** Generates a name for sequence items by index

**Parameters:**
- `index` (int): Zero-based sequence position

**Returns:** Formatted string name for the sequence item

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
name := analyzer.sequenceItemName(2) // returns "item_2"
```

**Edge Cases:**
- Negative index values
- Extremely large indices causing formatting issues


</details>


---

## 📄 File: `analyzer.go`

> 📍 `analyzer\analyzer.go`

## 📑 Contents

- [🧱 Structs (7)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `AnalyzerResult`

```go
type AnalyzerResult struct {
	Files []*AnalyzedFile 
}
```

**Summary:** Container for multiple file analysis results

**Returns:** None (type definition)

**Complexity:**
- Time: N/A (data structure)
- Space: O(n) where n is number of files

**Example:**
```go
result := AnalyzerResult{Files: analyzedFilesSlice}
```

**Edge Cases:**
- Empty Files slice
- Nil elements in Files array
- Concurrent access/modification


---

### `AnalyzedFile`

```go
type AnalyzedFile struct {
	Path string 
	Packages []Package 
}
```

**Summary:** Stores analysis data for a single file

**Returns:** None (type definition)

**Complexity:**
- Time: N/A (data structure)
- Space: O(m) where m is number of packages

**Example:**
```go
file := AnalyzedFile{Path: "/src/main.go", Packages: pkgs}
```

**Edge Cases:**
- Empty Path string
- Nil Packages slice
- Duplicate package entries


---

### `Package`

```go
type Package struct {
	Name string 
	Path string 
	Imports []string 
	Structs []Struct 
	Funcs []Function 
	Files []string 
}
```

**Summary:** Defines a Package structure with metadata and components

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
pkg := Package{Name: "example", Path: "github.com/example"}
```

**Edge Cases:**
- Empty slices for Imports, Structs, Funcs, or Files
- Nil values in any field


---

### `Struct`

```go
type Struct struct {
	Name string 
	Fields []Field 
	Methods []Function 
	Doc ai.Documentation 
	DocYAML ai.YAMLDocumentation 
}
```

**Summary:** Defines a struct type with name, fields, methods, and documentation.

**Parameters:**
- `Name` (string): Name of the struct
- `Fields` ([]Field): List of fields in the struct
- `Methods` ([]Function): List of methods associated with the struct
- `Doc` (ai.Documentation): Documentation for the struct
- `DocYAML` (ai.YAMLDocumentation): YAML-formatted documentation for the struct

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
type MyStruct Struct
```

**Edge Cases:**
- Empty or nil fields/methods
- Missing documentation


---

### `Field`

```go
type Field struct {
	Name string 
	Type string 
	Tag string 
	Doc ai.Documentation 
	DocYAML ai.YAMLDocumentation 
	Value string 
	Fields []Field 
}
```

**Summary:** Defines a field type with name, type, tag, and nested fields.

**Parameters:**
- `Name` (string): Name of the field
- `Type` (string): Type of the field
- `Tag` (string): Tag associated with the field
- `Doc` (ai.Documentation): Documentation for the field
- `DocYAML` (ai.YAMLDocumentation): YAML-formatted documentation for the field
- `Value` (string): Default value of the field
- `Fields` ([]Field): Nested fields (for complex types)

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
field := Field{Name: "Age", Type: "int"}
```

**Edge Cases:**
- Recursive field definitions
- Empty or invalid type/tag


---

### `Function`

```go
type Function struct {
	Name string 
	Receiver string 
	Parameters []Parameter 
	Results []Parameter 
	Doc ai.Documentation 
	Calls []string 
}
```

**Summary:** Struct representing a function with metadata and calls

**Parameters:**
- `Name` (string): Function name
- `Receiver` (string): Method receiver type
- `Parameters` ([]Parameter): Function parameters
- `Results` ([]Parameter): Return types
- `Doc` (ai.Documentation): Associated documentation
- `Calls` ([]string): Names of called functions

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1) for instantiation
- Space: O(n) where n is total parameters/results/calls

**Example:**
```go
funcObj := Function{Name: "Calculate", Parameters: []Parameter{{Name: "x", Type: "int"}}}
```

**Edge Cases:**
- Empty function name
- Nil slices for Parameters/Results/Calls


---

### `Parameter`

```go
type Parameter struct {
	Name string 
	Type string 
}
```

**Summary:** Struct representing a function parameter or return type

**Parameters:**
- `Name` (string): Parameter name
- `Type` (string): Parameter type

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1) for instantiation
- Space: O(1)

**Example:**
```go
param := Parameter{Name: "userId", Type: "string"}
```

**Edge Cases:**
- Empty type string
- Unsanitized special characters in name


---

## 🔧 Functions

<details>
<summary><b><code>init()</code></b></summary>

**Summary:** Analyzes Go project directory and returns results

**Parameters:**
- `projectDir` (string): Path to project root directory

**Returns:** *AnalyzerResult with analysis data or error if failed

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is total code elements analyzed

**Example:**
```go
result, err := AnalyzeProject("./myproject")
```

**Edge Cases:**
- Non-existent directory path
- Insufficient permissions
- Malformed Go files


</details>

<details>
<summary><b><code>AnalyzeProject(projectDir string)</code></b></summary>

**Summary:** Analyzes Go project directory and returns results

**Parameters:**
- `projectDir` (string): Path to project root directory

**Returns:** *AnalyzerResult with analysis data or error if failed

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is total code elements analyzed

**Example:**
```go
result, err := AnalyzeProject("./myproject")
```

**Edge Cases:**
- Non-existent directory path
- Insufficient permissions
- Malformed Go files


</details>

<details>
<summary><b><code>yamlPrintFileAnalysis(file *AnalyzedFile)</code></b></summary>

**Summary:** Empty struct defining a Go analyzer type

**Returns:** None (type definition only)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := GoAnalyzer{} // creates an instance
```

**Edge Cases:**
- No functionality (empty struct)


</details>

<details>
<summary><b><code>yamlPrintField(f Field, indent int)</code></b></summary>

**Summary:** Prints a field's details in YAML format with indentation.

**Parameters:**
- `f` (Field): Field to print
- `indent` (int): Indentation level

**Returns:** None (prints to stdout)

**Complexity:**
- Time: O(n) where n is the number of nested fields
- Space: O(1) (excluding recursion stack)

**Example:**
```go
yamlPrintField(Field{Name: "ID", Type: "string"}, 2)
```

**Edge Cases:**
- Deeply nested fields causing stack overflow
- Large indentation values


</details>


---

## 📄 File: `go_analyzer.go`

> 📍 `analyzer\go_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (6)](#-functions)

## 🧱 Structs

### `GoAnalyzer`

```go
type GoAnalyzer struct {
}
```

**Summary:** Empty struct defining a Go code analyzer type

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := GoAnalyzer{}
```

**Edge Cases:**
- No functionality implemented yet
- Struct may be extended in future


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if analyzer supports a project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) (assuming simple checks)
- Space: O(1)

**Example:**
```go
supported := analyzer.Supports("./myproject")
```

**Edge Cases:**
- Nonexistent directory path
- Empty/malformed path string


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes a Go project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** AnalyzerResult pointer and error (nil on success)

**Complexity:**
- Time: O(n) (depends on project size)
- Space: O(n) (depends on analysis output)

**Example:**
```go
result, err := analyzer.Analyze("./myproject")
```

**Edge Cases:**
- Nonexistent directory
- Permission errors
- Malformed Go code


</details>

<details>
<summary><b><code>Merge(other *AnalyzerResult)</code></b></summary>

**Summary:** Merges another AnalyzerResult into the current one

**Parameters:**
- `other` (*AnalyzerResult): Result to merge into the current instance

**Returns:** None (modifies the receiver in-place)

**Complexity:**
- Time: O(n) where n is size of other result
- Space: O(1) (modifies existing structure)

**Example:**
```go
result1.Merge(result2) // combines analysis results
```

**Edge Cases:**
- Nil input parameter
- Conflicting data between results


</details>

<details>
<summary><b><code>exists(path string)</code></b></summary>

**Summary:** Checks if a file/path exists

**Parameters:**
- `path` (string): Filesystem path to check

**Returns:** Boolean indicating existence

**Complexity:**
- Time: O(1) (filesystem operation)
- Space: O(1)

**Example:**
```go
if exists("/tmp/file") { ... }
```

**Edge Cases:**
- Empty path string
- Permission issues
- Symbolic links


</details>

<details>
<summary><b><code>AnalyzeFile(path string)</code></b></summary>

**Summary:** Converts an AST basic literal tag to its string representation

**Parameters:**
- `tag` (*ast.BasicLit): AST node representing a basic literal

**Returns:** String representation of the tag

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
str := TagToString(astNode) // returns "foo" for tag `foo`
```

**Edge Cases:**
- Nil input node
- Empty or malformed tag values


</details>

<details>
<summary><b><code>extractParams(fl *ast.FieldList)</code></b></summary>

**Summary:** Extracts parameters from an AST field list

**Parameters:**
- `fl` (*ast.FieldList): AST field list to extract parameters from

**Returns:** Slice of Parameter structs representing extracted parameters

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for storing parameters

**Example:**
```go
params := extractParams(funcDecl.Type.Params)
```

**Edge Cases:**
- Empty field list returns empty slice
- Nil input may cause panic


</details>


---

## 📄 File: `js_analyzer.go`

> 📍 `analyzer\js_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (8)](#-functions)

## 🧱 Structs

### `JSAnalyzer`

```go
type JSAnalyzer struct {
}
```

**Summary:** Empty struct for JavaScript code analysis

**Returns:** N/A (type declaration)

**Complexity:**
- Time: N/A
- Space: O(1) for instance creation

**Example:**
```go
analyzer := JSAnalyzer{}
```

**Edge Cases:**
- No immediate edge cases for empty struct


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if analyzer supports given project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) assuming simple checks
- Space: O(1)

**Example:**
```go
if analyzer.Supports("./my-project") { ... }
```

**Edge Cases:**
- Empty path string
- Non-existent directory
- Permission issues


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes JavaScript project directory and returns results

**Parameters:**
- `projectDir` (string): Path to the project directory

**Returns:** AnalyzerResult pointer and error, containing analysis results

**Complexity:**
- Time: O(n) where n is number of files
- Space: O(m) where m is size of analysis data

**Example:**
```go
result, err := analyzer.Analyze("./my-js-project")
```

**Edge Cases:**
- Non-existent directory path
- Permission issues accessing files
- Malformed JavaScript files


</details>

<details>
<summary><b><code>extractJSFunctions(src string)</code></b></summary>

**Summary:** Extracts all function declarations from JS source

**Parameters:**
- `src` (string): JavaScript source code

**Returns:** Slice of Function objects representing found functions

**Complexity:**
- Time: O(n) where n is source length
- Space: O(k) where k is number of functions

**Example:**
```go
funcs := extractJSFunctions("function foo(){}")
```

**Edge Cases:**
- Empty source string
- Nested functions
- Arrow functions
- Invalid syntax


</details>

<details>
<summary><b><code>extractJSFunctionBody(src string, funcDecl string)</code></b></summary>

**Summary:** Extracts body of specified function from JS source

**Parameters:**
- `src` (string): JavaScript source code
- `funcDecl` (string): Target function declaration

**Returns:** String containing the function body

**Complexity:**
- Time: O(n) where n is source length
- Space: O(m) where m is body size

**Example:**
```go
body := extractJSFunctionBody("function foo(){ return 1; }", "foo")
```

**Edge Cases:**
- Function not found
- Multiple matching functions
- Nested function with same name
- Invalid function declaration format


</details>

<details>
<summary><b><code>extractJSClasses(src string)</code></b></summary>

**Summary:** Extracts JavaScript class structures from source code

**Parameters:**
- `src` (string): JavaScript source code to parse

**Returns:** Array of Struct objects representing class definitions

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(m) where m is number of classes found

**Example:**
```go
classes := extractJSClasses('class A {} class B {}')
```

**Edge Cases:**
- Malformed JavaScript syntax
- Nested class declarations
- Source containing non-class code


</details>

<details>
<summary><b><code>extractJsClassBody(src string, classDecl string)</code></b></summary>

**Summary:** Extracts body of a specific JavaScript class

**Parameters:**
- `src` (string): JavaScript source code
- `classDecl` (string): Class declaration to match

**Returns:** String containing the matched class body content

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(k) where k is body length

**Example:**
```go
body := extractJsClassBody('class X { method() {} }', 'class X')
```

**Edge Cases:**
- Class declaration not found
- Partial/incomplete class declarations
- Multiple matching declarations


</details>

<details>
<summary><b><code>parseParamList(_ string, raw string)</code></b></summary>

**Summary:** Parses parameter list string into structured objects

**Parameters:**
- `_` (string): Unused parameter (likely placeholder)
- `raw` (string): Raw parameter list string

**Returns:** Array of Parameter objects with parsed details

**Complexity:**
- Time: O(n) where n is parameter list length
- Space: O(m) where m is number of parameters

**Example:**
```go
params := parseParamList('', 'a, b=5, ...rest')
```

**Edge Cases:**
- Empty parameter list
- Complex default values
- Rest parameters
- Malformed syntax


</details>

<details>
<summary><b><code>findDocBefore(context string, src string, docRegex *regexp.Regexp)</code></b></summary>

**Summary:** Finds documentation before a context string using regex

**Parameters:**
- `context` (string): Context string to search before
- `src` (string): Source text to search in
- `docRegex` (*regexp.Regexp): Regex pattern for documentation

**Returns:** Documentation string found before the context

**Complexity:**
- Time: O(n), where n is the length of src (regex dependent)
- Space: O(1) (excluding regex storage)

**Example:**
```go
doc := findDocBefore("func example()", "/* doc */ func example()", regexp.MustCompile(`/\*.*?\*/`))
```

**Edge Cases:**
- Empty src string
- No match for docRegex
- Context string not found in src


</details>


---

## 📄 File: `python_analyzer.go`

> 📍 `analyzer\python_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (7)](#-functions)

## 🧱 Structs

### `PythonAnalyzer`

```go
type PythonAnalyzer struct {
}
```

**Summary:** Empty struct for Python code analysis

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := PythonAnalyzer{}
```

**Edge Cases:**
- None (empty struct)


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if PythonAnalyzer supports a project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) (implementation dependent)
- Space: O(1)

**Example:**
```go
supported := analyzer.Supports("/path/to/project")
```

**Edge Cases:**
- Non-existent directory
- Empty directory
- Permission issues


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes Python project directory and returns results

**Parameters:**
- `projectDir` (string): Path to Python project directory

**Returns:** AnalyzerResult pointer and error (nil if successful)

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is analysis data size

**Example:**
```go
result, err := analyzer.Analyze("./my_project")
```

**Edge Cases:**
- Non-existent directory path
- Invalid Python project structure
- Permission issues


</details>

<details>
<summary><b><code>extractPythonClasses(src string)</code></b></summary>

**Summary:** Extracts Python class structures from source code

**Parameters:**
- `src` (string): Python source code text

**Returns:** Slice of Struct objects representing classes

**Complexity:**
- Time: O(n) where n is source length
- Space: O(k) where k is class count

**Example:**
```go
classes := extractPythonClasses("class Foo: pass")
```

**Edge Cases:**
- Malformed class syntax
- Nested class declarations
- Source with no classes


</details>

<details>
<summary><b><code>extractPythonClassBody(src string, classDecl string)</code></b></summary>

**Summary:** Extracts body of specified Python class

**Parameters:**
- `src` (string): Python source code text
- `classDecl` (string): Target class declaration

**Returns:** Class body as string (excluding declaration)

**Complexity:**
- Time: O(n) where n is source length
- Space: O(m) where m is body size

**Example:**
```go
body := extractPythonClassBody(src, "class MyClass")
```

**Edge Cases:**
- Class not found in source
- Malformed class syntax
- Multiple matching declarations


</details>

<details>
<summary><b><code>extractPythonFunctions(src string)</code></b></summary>

**Summary:** Extracts Python function declarations from source code

**Parameters:**
- `src` (string): Python source code to parse

**Returns:** Slice of Function structs containing function metadata

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(m) where m is number of functions found

**Example:**
```go
funcs := extractPythonFunctions('def foo(): pass\ndef bar(): pass')
```

**Edge Cases:**
- Malformed function declarations
- Nested functions
- Source with syntax errors


</details>

<details>
<summary><b><code>extractFunctionBody(src string, funcDecl string)</code></b></summary>

**Summary:** Extracts body of a specific function from source

**Parameters:**
- `src` (string): Source code containing the function
- `funcDecl` (string): Function declaration to match

**Returns:** String containing the matched function's body

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(k) where k is body length

**Example:**
```go
body := extractFunctionBody('def foo():\n    return 42', 'def foo():')
```

**Edge Cases:**
- Multiple matching declarations
- Function not found
- Incorrect indentation in source


</details>

<details>
<summary><b><code>findDocstringAfter(context string, src string, docRegex *regexp.Regexp)</code></b></summary>

**Summary:** Finds docstring following specific context in source

**Parameters:**
- `context` (string): Preceding code pattern
- `src` (string): Source code to search
- `docRegex` (*regexp.Regexp): Pattern matching docstring format

**Returns:** Matched docstring or empty string

**Complexity:**
- Time: O(n) where n is source length
- Space: O(1) for search, O(m) for result where m is docstring length

**Example:**
```go
doc := findDocstringAfter('def foo():', source, regexp.MustCompile(`"""(.+?)"""`))
```

**Edge Cases:**
- Context appears multiple times
- Malformed docstrings
- Context not found


</details>


---

## 📄 File: `yaml_analyzer.go`

> 📍 `analyzer\yaml_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (25)](#-functions)

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
<summary><b><code>analyzeYAMLNode(node *yaml.Node, baseName string)</code></b></summary>

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

<details>
<summary><b><code>isKubernetesResource(node *yaml.Node)</code></b></summary>

**Summary:** Checks if YAML node represents a Kubernetes resource

**Parameters:**
- `node` (*yaml.Node): YAML node to inspect

**Returns:** Boolean indicating Kubernetes resource status

**Complexity:**
- Time: O(n) where n is node depth/complexity
- Space: O(1)

**Example:**
```go
isKube := analyzer.isKubernetesResource(yamlNode)
```

**Edge Cases:**
- Nil node input
- Malformed YAML structure
- Non-Kubernetes YAML with similar fields


</details>

<details>
<summary><b><code>isHelmChart(node *yaml.Node)</code></b></summary>

**Summary:** Determines if YAML node is a Helm chart definition

**Parameters:**
- `node` (*yaml.Node): YAML node to evaluate

**Returns:** Boolean indicating Helm chart status

**Complexity:**
- Time: O(n) where n is node depth/complexity
- Space: O(1)

**Example:**
```go
isHelm := analyzer.isHelmChart(yamlNode)
```

**Edge Cases:**
- Nil node input
- Partial Helm chart structures
- Version compatibility issues


</details>

<details>
<summary><b><code>isAnsiblePlaybook(node *yaml.Node)</code></b></summary>

**Summary:** Checks if YAML node represents an Ansible playbook

**Parameters:**
- `node` (*yaml.Node): YAML node to analyze

**Returns:** Boolean indicating Ansible playbook status

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
isPlaybook := analyzer.isAnsiblePlaybook(yamlNode)
```

**Edge Cases:**
- Empty YAML node
- Malformed YAML structure


</details>

<details>
<summary><b><code>isDockerCompose(node *yaml.Node)</code></b></summary>

**Summary:** Determines if YAML node is a Docker Compose file

**Parameters:**
- `node` (*yaml.Node): YAML node to check

**Returns:** Boolean indicating Docker Compose format

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
isCompose := analyzer.isDockerCompose(yamlNode)
```

**Edge Cases:**
- Partial Docker Compose syntax
- Version mismatch in compose spec


</details>

<details>
<summary><b><code>analyzeKubernetesResource(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes YAML node for Kubernetes resource metadata

**Parameters:**
- `node` (*yaml.Node): YAML node to parse
- `baseName` (string): Base filename for context

**Returns:** Structured analysis result as custom type

**Complexity:**
- Time: O(n) where n is node children
- Space: O(n) for recursive parsing

**Example:**
```go
resource := analyzer.analyzeKubernetesResource(node, 'deployment.yaml')
```

**Edge Cases:**
- Custom resource definitions
- Multi-document YAML files
- Missing required Kubernetes fields


</details>

<details>
<summary><b><code>extractKubernetesSpec(node *yaml.Node)</code></b></summary>

**Summary:** Extracts Kubernetes spec fields from a YAML node

**Parameters:**
- `node` (*yaml.Node): YAML node to analyze

**Returns:** List of extracted fields ([]Field)

**Complexity:**
- Time: O(n) where n is the number of nodes in the YAML tree
- Space: O(m) where m is the number of fields extracted

**Example:**
```go
fields := analyzer.extractKubernetesSpec(yamlNode)
```

**Edge Cases:**
- Nil input node
- Malformed YAML structure
- Missing Kubernetes spec section


</details>

<details>
<summary><b><code>findSpecNode(node *yaml.Node)</code></b></summary>

**Summary:** Finds the spec node in a YAML document

**Parameters:**
- `node` (*yaml.Node): Root YAML node to search

**Returns:** Pointer to the found spec node (*yaml.Node) or nil

**Complexity:**
- Time: O(n) where n is the number of nodes in the YAML tree
- Space: O(1)

**Example:**
```go
specNode := analyzer.findSpecNode(rootNode)
```

**Edge Cases:**
- Nil input node
- Spec node not present
- Multiple spec nodes in document


</details>

<details>
<summary><b><code>analyzeHelmChart(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes a Helm chart's YAML structure

**Parameters:**
- `node` (*yaml.Node): Root YAML node of the chart
- `baseName` (string): Base name for the chart

**Returns:** Structured analysis result (Struct)

**Complexity:**
- Time: O(n) where n is the number of nodes in the YAML tree
- Space: O(m) where m is the size of the returned structure

**Example:**
```go
chartStruct := analyzer.analyzeHelmChart(chartNode, "my-chart")
```

**Edge Cases:**
- Nil input node
- Invalid Helm chart structure
- Empty baseName parameter


</details>

<details>
<summary><b><code>analyzeAnsiblePlaybook(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes an Ansible playbook YAML node and returns a structured representation.

**Parameters:**
- `node` (*yaml.Node): YAML node representing the playbook
- `baseName` (string): Base name for the playbook

**Returns:** Struct containing analyzed playbook data

**Complexity:**
- Time: O(n), where n is the number of nodes in the YAML
- Space: O(n), for storing the structured output

**Example:**
```go
result := analyzer.analyzeAnsiblePlaybook(yamlNode, "deploy.yml")
```

**Edge Cases:**
- Empty YAML node
- Malformed YAML structure
- Missing required playbook fields


</details>

<details>
<summary><b><code>extractAnsiblePlayFields(playNode *yaml.Node)</code></b></summary>

**Summary:** Extracts fields from an Ansible play YAML node.

**Parameters:**
- `playNode` (*yaml.Node): YAML node representing a play

**Returns:** Slice of Field objects representing play attributes

**Complexity:**
- Time: O(n), where n is the number of fields in the play
- Space: O(n), for storing the extracted fields

**Example:**
```go
fields := analyzer.extractAnsiblePlayFields(playNode)
```

**Edge Cases:**
- Empty play node
- Nested field structures
- Unsupported field types


</details>

<details>
<summary><b><code>extractAnsibleTasks(tasksNode *yaml.Node)</code></b></summary>

**Summary:** Extracts task fields from an Ansible tasks YAML node.

**Parameters:**
- `tasksNode` (*yaml.Node): YAML node containing tasks

**Returns:** Slice of Field objects representing task definitions

**Complexity:**
- Time: O(n), where n is the number of tasks
- Space: O(n), for storing the task fields

**Example:**
```go
tasks := analyzer.extractAnsibleTasks(tasksNode)
```

**Edge Cases:**
- Empty tasks section
- Tasks with complex nested structures
- Invalid task syntax


</details>

<details>
<summary><b><code>analyzeDockerCompose(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes Docker Compose YAML node and returns structured data

**Parameters:**
- `node` (*yaml.Node): YAML node to analyze
- `baseName` (string): Base name for the analysis

**Returns:** Structured analysis result as Struct

**Complexity:**
- Time: O(n) where n is YAML node complexity
- Space: O(n) for output structure

**Example:**
```go
result := analyzer.analyzeDockerCompose(doc.Root, "app")
```

**Edge Cases:**
- Nil node input
- Malformed YAML structure
- Empty baseName


</details>

<details>
<summary><b><code>extractDockerServices(servicesNode *yaml.Node)</code></b></summary>

**Summary:** Extracts Docker services from YAML services node

**Parameters:**
- `servicesNode` (*yaml.Node): YAML node containing services

**Returns:** Array of extracted service fields as []Field

**Complexity:**
- Time: O(n) where n is number of services
- Space: O(n) for output array

**Example:**
```go
services := analyzer.extractDockerServices(composeNode.Content[0])
```

**Edge Cases:**
- Nil servicesNode
- Non-service YAML structure
- Empty services section


</details>

<details>
<summary><b><code>getValueNode(node *yaml.Node, key string)</code></b></summary>

**Summary:** Finds YAML value node by key

**Parameters:**
- `node` (*yaml.Node): Parent YAML node
- `key` (string): Key to search for

**Returns:** Found value node or nil as *yaml.Node

**Complexity:**
- Time: O(n) where n is child nodes count
- Space: O(1)

**Example:**
```go
portNode := analyzer.getValueNode(serviceNode, "ports")
```

**Edge Cases:**
- Nil parent node
- Non-existent key
- Case-sensitive key matching


</details>

<details>
<summary><b><code>extractArrayItems(node *yaml.Node)</code></b></summary>

**Summary:** Extracts array items from a YAML node into Field objects

**Parameters:**
- `node` (*yaml.Node): YAML node containing array items

**Returns:** Slice of Field objects representing array items

**Complexity:**
- Time: O(n) where n is the number of array items
- Space: O(n) for storing extracted Field objects

**Example:**
```go
fields := analyzer.extractArrayItems(yamlArrayNode)
```

**Edge Cases:**
- Empty array node
- Nested array structures
- Non-array node input


</details>

<details>
<summary><b><code>extractYAMLFields(node *yaml.Node, depth int)</code></b></summary>

**Summary:** Recursively extracts YAML fields up to specified depth

**Parameters:**
- `node` (*yaml.Node): Root YAML node to analyze
- `depth` (int): Maximum recursion depth

**Returns:** Slice of Field objects representing YAML structure

**Complexity:**
- Time: O(n) where n is total nodes visited
- Space: O(d) where d is maximum depth (call stack)

**Example:**
```go
fields := analyzer.extractYAMLFields(rootNode, 3)
```

**Edge Cases:**
- Depth limit reached before complete traversal
- Cyclic references in YAML
- Very large depth values causing stack overflow


</details>

<details>
<summary><b><code>extractYAMLDescription(node *yaml.Node)</code></b></summary>

**Summary:** Extracts description text from a YAML node

**Parameters:**
- `node` (*yaml.Node): YAML node containing description

**Returns:** Extracted description as string

**Complexity:**
- Time: O(1) for simple nodes
- Space: O(1) for returned string

**Example:**
```go
desc := analyzer.extractYAMLDescription(commentNode)
```

**Edge Cases:**
- Node without description content
- Multi-line descriptions
- Special characters in description


</details>

<details>
<summary><b><code>hasKey(node *yaml.Node, key string)</code></b></summary>

**Summary:** Checks if a YAML node contains a specific key

**Parameters:**
- `node` (*yaml.Node): YAML node to inspect
- `key` (string): Key to search for

**Returns:** Boolean indicating if the key exists

**Complexity:**
- Time: O(n) where n is the number of keys in the node
- Space: O(1)

**Example:**
```go
exists := analyzer.hasKey(node, "version")
```

**Edge Cases:**
- Empty node
- Non-mapping node type
- Case-sensitive key matching


</details>

<details>
<summary><b><code>getValue(node *yaml.Node, key string)</code></b></summary>

**Summary:** Retrieves the value associated with a key in a YAML node

**Parameters:**
- `node` (*yaml.Node): YAML node to inspect
- `key` (string): Key whose value to retrieve

**Returns:** String value of the key, or empty string if not found

**Complexity:**
- Time: O(n) where n is the number of keys in the node
- Space: O(1)

**Example:**
```go
value := analyzer.getValue(node, "name")
```

**Edge Cases:**
- Key doesn't exist
- Non-string value type
- Nested mapping values


</details>

<details>
<summary><b><code>getMappingKeys(node *yaml.Node)</code></b></summary>

**Summary:** Extracts all keys from a YAML mapping node

**Parameters:**
- `node` (*yaml.Node): YAML mapping node

**Returns:** Slice of strings containing all keys

**Complexity:**
- Time: O(n) where n is the number of keys
- Space: O(n) for storing the keys

**Example:**
```go
keys := analyzer.getMappingKeys(node)
```

**Edge Cases:**
- Empty mapping node
- Non-mapping node type
- Duplicate keys in YAML


</details>

<details>
<summary><b><code>nodeKindToString(kind yaml.Kind)</code></b></summary>

**Summary:** Converts YAML node kind to its string representation

**Parameters:**
- `kind` (yaml.Kind): YAML node kind enum value

**Returns:** String representation of the YAML node kind

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
str := analyzer.nodeKindToString(yaml.SequenceNode) // returns "sequence"
```

**Edge Cases:**
- Unknown/unsupported yaml.Kind values
- Nil receiver (though method is on pointer receiver)


</details>

<details>
<summary><b><code>sequenceItemName(index int)</code></b></summary>

**Summary:** Generates a name for sequence items by index

**Parameters:**
- `index` (int): Zero-based sequence position

**Returns:** Formatted string name for the sequence item

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
name := analyzer.sequenceItemName(2) // returns "item_2"
```

**Edge Cases:**
- Negative index values
- Extremely large indices causing formatting issues


</details>


---

## 📄 File: `analyzer.go`

> 📍 `analyzer\analyzer.go`

## 📑 Contents

- [🧱 Structs (7)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `AnalyzerResult`

```go
type AnalyzerResult struct {
	Files []*AnalyzedFile 
}
```

**Summary:** Container for multiple file analysis results

**Returns:** None (type definition)

**Complexity:**
- Time: N/A (data structure)
- Space: O(n) where n is number of files

**Example:**
```go
result := AnalyzerResult{Files: analyzedFilesSlice}
```

**Edge Cases:**
- Empty Files slice
- Nil elements in Files array
- Concurrent access/modification


---

### `AnalyzedFile`

```go
type AnalyzedFile struct {
	Path string 
	Packages []Package 
}
```

**Summary:** Stores analysis data for a single file

**Returns:** None (type definition)

**Complexity:**
- Time: N/A (data structure)
- Space: O(m) where m is number of packages

**Example:**
```go
file := AnalyzedFile{Path: "/src/main.go", Packages: pkgs}
```

**Edge Cases:**
- Empty Path string
- Nil Packages slice
- Duplicate package entries


---

### `Package`

```go
type Package struct {
	Name string 
	Path string 
	Imports []string 
	Structs []Struct 
	Funcs []Function 
	Files []string 
}
```

**Summary:** Defines a Package structure with metadata and components

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
pkg := Package{Name: "example", Path: "github.com/example"}
```

**Edge Cases:**
- Empty slices for Imports, Structs, Funcs, or Files
- Nil values in any field


---

### `Struct`

```go
type Struct struct {
	Name string 
	Fields []Field 
	Methods []Function 
	Doc ai.Documentation 
	DocYAML ai.YAMLDocumentation 
}
```

**Summary:** Defines a struct type with name, fields, methods, and documentation.

**Parameters:**
- `Name` (string): Name of the struct
- `Fields` ([]Field): List of fields in the struct
- `Methods` ([]Function): List of methods associated with the struct
- `Doc` (ai.Documentation): Documentation for the struct
- `DocYAML` (ai.YAMLDocumentation): YAML-formatted documentation for the struct

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
type MyStruct Struct
```

**Edge Cases:**
- Empty or nil fields/methods
- Missing documentation


---

### `Field`

```go
type Field struct {
	Name string 
	Type string 
	Tag string 
	Doc ai.Documentation 
	DocYAML ai.YAMLDocumentation 
	Value string 
	Fields []Field 
}
```

**Summary:** Defines a field type with name, type, tag, and nested fields.

**Parameters:**
- `Name` (string): Name of the field
- `Type` (string): Type of the field
- `Tag` (string): Tag associated with the field
- `Doc` (ai.Documentation): Documentation for the field
- `DocYAML` (ai.YAMLDocumentation): YAML-formatted documentation for the field
- `Value` (string): Default value of the field
- `Fields` ([]Field): Nested fields (for complex types)

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
field := Field{Name: "Age", Type: "int"}
```

**Edge Cases:**
- Recursive field definitions
- Empty or invalid type/tag


---

### `Function`

```go
type Function struct {
	Name string 
	Receiver string 
	Parameters []Parameter 
	Results []Parameter 
	Doc ai.Documentation 
	Calls []string 
}
```

**Summary:** Struct representing a function with metadata and calls

**Parameters:**
- `Name` (string): Function name
- `Receiver` (string): Method receiver type
- `Parameters` ([]Parameter): Function parameters
- `Results` ([]Parameter): Return types
- `Doc` (ai.Documentation): Associated documentation
- `Calls` ([]string): Names of called functions

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1) for instantiation
- Space: O(n) where n is total parameters/results/calls

**Example:**
```go
funcObj := Function{Name: "Calculate", Parameters: []Parameter{{Name: "x", Type: "int"}}}
```

**Edge Cases:**
- Empty function name
- Nil slices for Parameters/Results/Calls


---

### `Parameter`

```go
type Parameter struct {
	Name string 
	Type string 
}
```

**Summary:** Struct representing a function parameter or return type

**Parameters:**
- `Name` (string): Parameter name
- `Type` (string): Parameter type

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1) for instantiation
- Space: O(1)

**Example:**
```go
param := Parameter{Name: "userId", Type: "string"}
```

**Edge Cases:**
- Empty type string
- Unsanitized special characters in name


---

## 🔧 Functions

<details>
<summary><b><code>init()</code></b></summary>

**Summary:** Analyzes Go project directory and returns results

**Parameters:**
- `projectDir` (string): Path to project root directory

**Returns:** *AnalyzerResult with analysis data or error if failed

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is total code elements analyzed

**Example:**
```go
result, err := AnalyzeProject("./myproject")
```

**Edge Cases:**
- Non-existent directory path
- Insufficient permissions
- Malformed Go files


</details>

<details>
<summary><b><code>AnalyzeProject(projectDir string)</code></b></summary>

**Summary:** Analyzes Go project directory and returns results

**Parameters:**
- `projectDir` (string): Path to project root directory

**Returns:** *AnalyzerResult with analysis data or error if failed

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is total code elements analyzed

**Example:**
```go
result, err := AnalyzeProject("./myproject")
```

**Edge Cases:**
- Non-existent directory path
- Insufficient permissions
- Malformed Go files


</details>

<details>
<summary><b><code>yamlPrintFileAnalysis(file *AnalyzedFile)</code></b></summary>

**Summary:** Empty struct defining a Go analyzer type

**Returns:** None (type definition only)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := GoAnalyzer{} // creates an instance
```

**Edge Cases:**
- No functionality (empty struct)


</details>

<details>
<summary><b><code>yamlPrintField(f Field, indent int)</code></b></summary>

**Summary:** Prints a field's details in YAML format with indentation.

**Parameters:**
- `f` (Field): Field to print
- `indent` (int): Indentation level

**Returns:** None (prints to stdout)

**Complexity:**
- Time: O(n) where n is the number of nested fields
- Space: O(1) (excluding recursion stack)

**Example:**
```go
yamlPrintField(Field{Name: "ID", Type: "string"}, 2)
```

**Edge Cases:**
- Deeply nested fields causing stack overflow
- Large indentation values


</details>


---

## 📄 File: `go_analyzer.go`

> 📍 `analyzer\go_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (6)](#-functions)

## 🧱 Structs

### `GoAnalyzer`

```go
type GoAnalyzer struct {
}
```

**Summary:** Empty struct defining a Go code analyzer type

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := GoAnalyzer{}
```

**Edge Cases:**
- No functionality implemented yet
- Struct may be extended in future


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if analyzer supports a project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) (assuming simple checks)
- Space: O(1)

**Example:**
```go
supported := analyzer.Supports("./myproject")
```

**Edge Cases:**
- Nonexistent directory path
- Empty/malformed path string


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes a Go project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** AnalyzerResult pointer and error (nil on success)

**Complexity:**
- Time: O(n) (depends on project size)
- Space: O(n) (depends on analysis output)

**Example:**
```go
result, err := analyzer.Analyze("./myproject")
```

**Edge Cases:**
- Nonexistent directory
- Permission errors
- Malformed Go code


</details>

<details>
<summary><b><code>Merge(other *AnalyzerResult)</code></b></summary>

**Summary:** Merges another AnalyzerResult into the current one

**Parameters:**
- `other` (*AnalyzerResult): Result to merge into the current instance

**Returns:** None (modifies the receiver in-place)

**Complexity:**
- Time: O(n) where n is size of other result
- Space: O(1) (modifies existing structure)

**Example:**
```go
result1.Merge(result2) // combines analysis results
```

**Edge Cases:**
- Nil input parameter
- Conflicting data between results


</details>

<details>
<summary><b><code>exists(path string)</code></b></summary>

**Summary:** Checks if a file/path exists

**Parameters:**
- `path` (string): Filesystem path to check

**Returns:** Boolean indicating existence

**Complexity:**
- Time: O(1) (filesystem operation)
- Space: O(1)

**Example:**
```go
if exists("/tmp/file") { ... }
```

**Edge Cases:**
- Empty path string
- Permission issues
- Symbolic links


</details>

<details>
<summary><b><code>AnalyzeFile(path string)</code></b></summary>

**Summary:** Converts an AST basic literal tag to its string representation

**Parameters:**
- `tag` (*ast.BasicLit): AST node representing a basic literal

**Returns:** String representation of the tag

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
str := TagToString(astNode) // returns "foo" for tag `foo`
```

**Edge Cases:**
- Nil input node
- Empty or malformed tag values


</details>

<details>
<summary><b><code>extractParams(fl *ast.FieldList)</code></b></summary>

**Summary:** Extracts parameters from an AST field list

**Parameters:**
- `fl` (*ast.FieldList): AST field list to extract parameters from

**Returns:** Slice of Parameter structs representing extracted parameters

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for storing parameters

**Example:**
```go
params := extractParams(funcDecl.Type.Params)
```

**Edge Cases:**
- Empty field list returns empty slice
- Nil input may cause panic


</details>


---

## 📄 File: `js_analyzer.go`

> 📍 `analyzer\js_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (8)](#-functions)

## 🧱 Structs

### `JSAnalyzer`

```go
type JSAnalyzer struct {
}
```

**Summary:** Empty struct for JavaScript code analysis

**Returns:** N/A (type declaration)

**Complexity:**
- Time: N/A
- Space: O(1) for instance creation

**Example:**
```go
analyzer := JSAnalyzer{}
```

**Edge Cases:**
- No immediate edge cases for empty struct


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if analyzer supports given project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) assuming simple checks
- Space: O(1)

**Example:**
```go
if analyzer.Supports("./my-project") { ... }
```

**Edge Cases:**
- Empty path string
- Non-existent directory
- Permission issues


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes JavaScript project directory and returns results

**Parameters:**
- `projectDir` (string): Path to the project directory

**Returns:** AnalyzerResult pointer and error, containing analysis results

**Complexity:**
- Time: O(n) where n is number of files
- Space: O(m) where m is size of analysis data

**Example:**
```go
result, err := analyzer.Analyze("./my-js-project")
```

**Edge Cases:**
- Non-existent directory path
- Permission issues accessing files
- Malformed JavaScript files


</details>

<details>
<summary><b><code>extractJSFunctions(src string)</code></b></summary>

**Summary:** Extracts all function declarations from JS source

**Parameters:**
- `src` (string): JavaScript source code

**Returns:** Slice of Function objects representing found functions

**Complexity:**
- Time: O(n) where n is source length
- Space: O(k) where k is number of functions

**Example:**
```go
funcs := extractJSFunctions("function foo(){}")
```

**Edge Cases:**
- Empty source string
- Nested functions
- Arrow functions
- Invalid syntax


</details>

<details>
<summary><b><code>extractJSFunctionBody(src string, funcDecl string)</code></b></summary>

**Summary:** Extracts body of specified function from JS source

**Parameters:**
- `src` (string): JavaScript source code
- `funcDecl` (string): Target function declaration

**Returns:** String containing the function body

**Complexity:**
- Time: O(n) where n is source length
- Space: O(m) where m is body size

**Example:**
```go
body := extractJSFunctionBody("function foo(){ return 1; }", "foo")
```

**Edge Cases:**
- Function not found
- Multiple matching functions
- Nested function with same name
- Invalid function declaration format


</details>

<details>
<summary><b><code>extractJSClasses(src string)</code></b></summary>

**Summary:** Extracts JavaScript class structures from source code

**Parameters:**
- `src` (string): JavaScript source code to parse

**Returns:** Array of Struct objects representing class definitions

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(m) where m is number of classes found

**Example:**
```go
classes := extractJSClasses('class A {} class B {}')
```

**Edge Cases:**
- Malformed JavaScript syntax
- Nested class declarations
- Source containing non-class code


</details>

<details>
<summary><b><code>extractJsClassBody(src string, classDecl string)</code></b></summary>

**Summary:** Extracts body of a specific JavaScript class

**Parameters:**
- `src` (string): JavaScript source code
- `classDecl` (string): Class declaration to match

**Returns:** String containing the matched class body content

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(k) where k is body length

**Example:**
```go
body := extractJsClassBody('class X { method() {} }', 'class X')
```

**Edge Cases:**
- Class declaration not found
- Partial/incomplete class declarations
- Multiple matching declarations


</details>

<details>
<summary><b><code>parseParamList(_ string, raw string)</code></b></summary>

**Summary:** Parses parameter list string into structured objects

**Parameters:**
- `_` (string): Unused parameter (likely placeholder)
- `raw` (string): Raw parameter list string

**Returns:** Array of Parameter objects with parsed details

**Complexity:**
- Time: O(n) where n is parameter list length
- Space: O(m) where m is number of parameters

**Example:**
```go
params := parseParamList('', 'a, b=5, ...rest')
```

**Edge Cases:**
- Empty parameter list
- Complex default values
- Rest parameters
- Malformed syntax


</details>

<details>
<summary><b><code>findDocBefore(context string, src string, docRegex *regexp.Regexp)</code></b></summary>

**Summary:** Finds documentation before a context string using regex

**Parameters:**
- `context` (string): Context string to search before
- `src` (string): Source text to search in
- `docRegex` (*regexp.Regexp): Regex pattern for documentation

**Returns:** Documentation string found before the context

**Complexity:**
- Time: O(n), where n is the length of src (regex dependent)
- Space: O(1) (excluding regex storage)

**Example:**
```go
doc := findDocBefore("func example()", "/* doc */ func example()", regexp.MustCompile(`/\*.*?\*/`))
```

**Edge Cases:**
- Empty src string
- No match for docRegex
- Context string not found in src


</details>


---

## 📄 File: `python_analyzer.go`

> 📍 `analyzer\python_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (7)](#-functions)

## 🧱 Structs

### `PythonAnalyzer`

```go
type PythonAnalyzer struct {
}
```

**Summary:** Empty struct for Python code analysis

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
analyzer := PythonAnalyzer{}
```

**Edge Cases:**
- None (empty struct)


---

## 🔧 Functions

<details>
<summary><b><code>Supports(projectDir string)</code></b></summary>

**Summary:** Checks if PythonAnalyzer supports a project directory

**Parameters:**
- `projectDir` (string): Path to project directory

**Returns:** Boolean indicating support status

**Complexity:**
- Time: O(1) (implementation dependent)
- Space: O(1)

**Example:**
```go
supported := analyzer.Supports("/path/to/project")
```

**Edge Cases:**
- Non-existent directory
- Empty directory
- Permission issues


</details>

<details>
<summary><b><code>Analyze(projectDir string)</code></b></summary>

**Summary:** Analyzes Python project directory and returns results

**Parameters:**
- `projectDir` (string): Path to Python project directory

**Returns:** AnalyzerResult pointer and error (nil if successful)

**Complexity:**
- Time: O(n) where n is project file count
- Space: O(m) where m is analysis data size

**Example:**
```go
result, err := analyzer.Analyze("./my_project")
```

**Edge Cases:**
- Non-existent directory path
- Invalid Python project structure
- Permission issues


</details>

<details>
<summary><b><code>extractPythonClasses(src string)</code></b></summary>

**Summary:** Extracts Python class structures from source code

**Parameters:**
- `src` (string): Python source code text

**Returns:** Slice of Struct objects representing classes

**Complexity:**
- Time: O(n) where n is source length
- Space: O(k) where k is class count

**Example:**
```go
classes := extractPythonClasses("class Foo: pass")
```

**Edge Cases:**
- Malformed class syntax
- Nested class declarations
- Source with no classes


</details>

<details>
<summary><b><code>extractPythonClassBody(src string, classDecl string)</code></b></summary>

**Summary:** Extracts body of specified Python class

**Parameters:**
- `src` (string): Python source code text
- `classDecl` (string): Target class declaration

**Returns:** Class body as string (excluding declaration)

**Complexity:**
- Time: O(n) where n is source length
- Space: O(m) where m is body size

**Example:**
```go
body := extractPythonClassBody(src, "class MyClass")
```

**Edge Cases:**
- Class not found in source
- Malformed class syntax
- Multiple matching declarations


</details>

<details>
<summary><b><code>extractPythonFunctions(src string)</code></b></summary>

**Summary:** Extracts Python function declarations from source code

**Parameters:**
- `src` (string): Python source code to parse

**Returns:** Slice of Function structs containing function metadata

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(m) where m is number of functions found

**Example:**
```go
funcs := extractPythonFunctions('def foo(): pass\ndef bar(): pass')
```

**Edge Cases:**
- Malformed function declarations
- Nested functions
- Source with syntax errors


</details>

<details>
<summary><b><code>extractFunctionBody(src string, funcDecl string)</code></b></summary>

**Summary:** Extracts body of a specific function from source

**Parameters:**
- `src` (string): Source code containing the function
- `funcDecl` (string): Function declaration to match

**Returns:** String containing the matched function's body

**Complexity:**
- Time: O(n) where n is length of source
- Space: O(k) where k is body length

**Example:**
```go
body := extractFunctionBody('def foo():\n    return 42', 'def foo():')
```

**Edge Cases:**
- Multiple matching declarations
- Function not found
- Incorrect indentation in source


</details>

<details>
<summary><b><code>findDocstringAfter(context string, src string, docRegex *regexp.Regexp)</code></b></summary>

**Summary:** Finds docstring following specific context in source

**Parameters:**
- `context` (string): Preceding code pattern
- `src` (string): Source code to search
- `docRegex` (*regexp.Regexp): Pattern matching docstring format

**Returns:** Matched docstring or empty string

**Complexity:**
- Time: O(n) where n is source length
- Space: O(1) for search, O(m) for result where m is docstring length

**Example:**
```go
doc := findDocstringAfter('def foo():', source, regexp.MustCompile(`"""(.+?)"""`))
```

**Edge Cases:**
- Context appears multiple times
- Malformed docstrings
- Context not found


</details>


---

## 📄 File: `yaml_analyzer.go`

> 📍 `analyzer\yaml_analyzer.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (25)](#-functions)

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
<summary><b><code>analyzeYAMLNode(node *yaml.Node, baseName string)</code></b></summary>

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

<details>
<summary><b><code>isKubernetesResource(node *yaml.Node)</code></b></summary>

**Summary:** Checks if YAML node represents a Kubernetes resource

**Parameters:**
- `node` (*yaml.Node): YAML node to inspect

**Returns:** Boolean indicating Kubernetes resource status

**Complexity:**
- Time: O(n) where n is node depth/complexity
- Space: O(1)

**Example:**
```go
isKube := analyzer.isKubernetesResource(yamlNode)
```

**Edge Cases:**
- Nil node input
- Malformed YAML structure
- Non-Kubernetes YAML with similar fields


</details>

<details>
<summary><b><code>isHelmChart(node *yaml.Node)</code></b></summary>

**Summary:** Determines if YAML node is a Helm chart definition

**Parameters:**
- `node` (*yaml.Node): YAML node to evaluate

**Returns:** Boolean indicating Helm chart status

**Complexity:**
- Time: O(n) where n is node depth/complexity
- Space: O(1)

**Example:**
```go
isHelm := analyzer.isHelmChart(yamlNode)
```

**Edge Cases:**
- Nil node input
- Partial Helm chart structures
- Version compatibility issues


</details>

<details>
<summary><b><code>isAnsiblePlaybook(node *yaml.Node)</code></b></summary>

**Summary:** Checks if YAML node represents an Ansible playbook

**Parameters:**
- `node` (*yaml.Node): YAML node to analyze

**Returns:** Boolean indicating Ansible playbook status

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
isPlaybook := analyzer.isAnsiblePlaybook(yamlNode)
```

**Edge Cases:**
- Empty YAML node
- Malformed YAML structure


</details>

<details>
<summary><b><code>isDockerCompose(node *yaml.Node)</code></b></summary>

**Summary:** Determines if YAML node is a Docker Compose file

**Parameters:**
- `node` (*yaml.Node): YAML node to check

**Returns:** Boolean indicating Docker Compose format

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
isCompose := analyzer.isDockerCompose(yamlNode)
```

**Edge Cases:**
- Partial Docker Compose syntax
- Version mismatch in compose spec


</details>

<details>
<summary><b><code>analyzeKubernetesResource(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes YAML node for Kubernetes resource metadata

**Parameters:**
- `node` (*yaml.Node): YAML node to parse
- `baseName` (string): Base filename for context

**Returns:** Structured analysis result as custom type

**Complexity:**
- Time: O(n) where n is node children
- Space: O(n) for recursive parsing

**Example:**
```go
resource := analyzer.analyzeKubernetesResource(node, 'deployment.yaml')
```

**Edge Cases:**
- Custom resource definitions
- Multi-document YAML files
- Missing required Kubernetes fields


</details>

<details>
<summary><b><code>extractKubernetesSpec(node *yaml.Node)</code></b></summary>

**Summary:** Extracts Kubernetes spec fields from a YAML node

**Parameters:**
- `node` (*yaml.Node): YAML node to analyze

**Returns:** List of extracted fields ([]Field)

**Complexity:**
- Time: O(n) where n is the number of nodes in the YAML tree
- Space: O(m) where m is the number of fields extracted

**Example:**
```go
fields := analyzer.extractKubernetesSpec(yamlNode)
```

**Edge Cases:**
- Nil input node
- Malformed YAML structure
- Missing Kubernetes spec section


</details>

<details>
<summary><b><code>findSpecNode(node *yaml.Node)</code></b></summary>

**Summary:** Finds the spec node in a YAML document

**Parameters:**
- `node` (*yaml.Node): Root YAML node to search

**Returns:** Pointer to the found spec node (*yaml.Node) or nil

**Complexity:**
- Time: O(n) where n is the number of nodes in the YAML tree
- Space: O(1)

**Example:**
```go
specNode := analyzer.findSpecNode(rootNode)
```

**Edge Cases:**
- Nil input node
- Spec node not present
- Multiple spec nodes in document


</details>

<details>
<summary><b><code>analyzeHelmChart(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes a Helm chart's YAML structure

**Parameters:**
- `node` (*yaml.Node): Root YAML node of the chart
- `baseName` (string): Base name for the chart

**Returns:** Structured analysis result (Struct)

**Complexity:**
- Time: O(n) where n is the number of nodes in the YAML tree
- Space: O(m) where m is the size of the returned structure

**Example:**
```go
chartStruct := analyzer.analyzeHelmChart(chartNode, "my-chart")
```

**Edge Cases:**
- Nil input node
- Invalid Helm chart structure
- Empty baseName parameter


</details>

<details>
<summary><b><code>analyzeAnsiblePlaybook(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes an Ansible playbook YAML node and returns a structured representation.

**Parameters:**
- `node` (*yaml.Node): YAML node representing the playbook
- `baseName` (string): Base name for the playbook

**Returns:** Struct containing analyzed playbook data

**Complexity:**
- Time: O(n), where n is the number of nodes in the YAML
- Space: O(n), for storing the structured output

**Example:**
```go
result := analyzer.analyzeAnsiblePlaybook(yamlNode, "deploy.yml")
```

**Edge Cases:**
- Empty YAML node
- Malformed YAML structure
- Missing required playbook fields


</details>

<details>
<summary><b><code>extractAnsiblePlayFields(playNode *yaml.Node)</code></b></summary>

**Summary:** Extracts fields from an Ansible play YAML node.

**Parameters:**
- `playNode` (*yaml.Node): YAML node representing a play

**Returns:** Slice of Field objects representing play attributes

**Complexity:**
- Time: O(n), where n is the number of fields in the play
- Space: O(n), for storing the extracted fields

**Example:**
```go
fields := analyzer.extractAnsiblePlayFields(playNode)
```

**Edge Cases:**
- Empty play node
- Nested field structures
- Unsupported field types


</details>

<details>
<summary><b><code>extractAnsibleTasks(tasksNode *yaml.Node)</code></b></summary>

**Summary:** Extracts task fields from an Ansible tasks YAML node.

**Parameters:**
- `tasksNode` (*yaml.Node): YAML node containing tasks

**Returns:** Slice of Field objects representing task definitions

**Complexity:**
- Time: O(n), where n is the number of tasks
- Space: O(n), for storing the task fields

**Example:**
```go
tasks := analyzer.extractAnsibleTasks(tasksNode)
```

**Edge Cases:**
- Empty tasks section
- Tasks with complex nested structures
- Invalid task syntax


</details>

<details>
<summary><b><code>analyzeDockerCompose(node *yaml.Node, baseName string)</code></b></summary>

**Summary:** Analyzes Docker Compose YAML node and returns structured data

**Parameters:**
- `node` (*yaml.Node): YAML node to analyze
- `baseName` (string): Base name for the analysis

**Returns:** Structured analysis result as Struct

**Complexity:**
- Time: O(n) where n is YAML node complexity
- Space: O(n) for output structure

**Example:**
```go
result := analyzer.analyzeDockerCompose(doc.Root, "app")
```

**Edge Cases:**
- Nil node input
- Malformed YAML structure
- Empty baseName


</details>

<details>
<summary><b><code>extractDockerServices(servicesNode *yaml.Node)</code></b></summary>

**Summary:** Extracts Docker services from YAML services node

**Parameters:**
- `servicesNode` (*yaml.Node): YAML node containing services

**Returns:** Array of extracted service fields as []Field

**Complexity:**
- Time: O(n) where n is number of services
- Space: O(n) for output array

**Example:**
```go
services := analyzer.extractDockerServices(composeNode.Content[0])
```

**Edge Cases:**
- Nil servicesNode
- Non-service YAML structure
- Empty services section


</details>

<details>
<summary><b><code>getValueNode(node *yaml.Node, key string)</code></b></summary>

**Summary:** Finds YAML value node by key

**Parameters:**
- `node` (*yaml.Node): Parent YAML node
- `key` (string): Key to search for

**Returns:** Found value node or nil as *yaml.Node

**Complexity:**
- Time: O(n) where n is child nodes count
- Space: O(1)

**Example:**
```go
portNode := analyzer.getValueNode(serviceNode, "ports")
```

**Edge Cases:**
- Nil parent node
- Non-existent key
- Case-sensitive key matching


</details>

<details>
<summary><b><code>extractArrayItems(node *yaml.Node)</code></b></summary>

**Summary:** Extracts array items from a YAML node into Field objects

**Parameters:**
- `node` (*yaml.Node): YAML node containing array items

**Returns:** Slice of Field objects representing array items

**Complexity:**
- Time: O(n) where n is the number of array items
- Space: O(n) for storing extracted Field objects

**Example:**
```go
fields := analyzer.extractArrayItems(yamlArrayNode)
```

**Edge Cases:**
- Empty array node
- Nested array structures
- Non-array node input


</details>

<details>
<summary><b><code>extractYAMLFields(node *yaml.Node, depth int)</code></b></summary>

**Summary:** Recursively extracts YAML fields up to specified depth

**Parameters:**
- `node` (*yaml.Node): Root YAML node to analyze
- `depth` (int): Maximum recursion depth

**Returns:** Slice of Field objects representing YAML structure

**Complexity:**
- Time: O(n) where n is total nodes visited
- Space: O(d) where d is maximum depth (call stack)

**Example:**
```go
fields := analyzer.extractYAMLFields(rootNode, 3)
```

**Edge Cases:**
- Depth limit reached before complete traversal
- Cyclic references in YAML
- Very large depth values causing stack overflow


</details>

<details>
<summary><b><code>extractYAMLDescription(node *yaml.Node)</code></b></summary>

**Summary:** Extracts description text from a YAML node

**Parameters:**
- `node` (*yaml.Node): YAML node containing description

**Returns:** Extracted description as string

**Complexity:**
- Time: O(1) for simple nodes
- Space: O(1) for returned string

**Example:**
```go
desc := analyzer.extractYAMLDescription(commentNode)
```

**Edge Cases:**
- Node without description content
- Multi-line descriptions
- Special characters in description


</details>

<details>
<summary><b><code>hasKey(node *yaml.Node, key string)</code></b></summary>

**Summary:** Checks if a YAML node contains a specific key

**Parameters:**
- `node` (*yaml.Node): YAML node to inspect
- `key` (string): Key to search for

**Returns:** Boolean indicating if the key exists

**Complexity:**
- Time: O(n) where n is the number of keys in the node
- Space: O(1)

**Example:**
```go
exists := analyzer.hasKey(node, "version")
```

**Edge Cases:**
- Empty node
- Non-mapping node type
- Case-sensitive key matching


</details>

<details>
<summary><b><code>getValue(node *yaml.Node, key string)</code></b></summary>

**Summary:** Retrieves the value associated with a key in a YAML node

**Parameters:**
- `node` (*yaml.Node): YAML node to inspect
- `key` (string): Key whose value to retrieve

**Returns:** String value of the key, or empty string if not found

**Complexity:**
- Time: O(n) where n is the number of keys in the node
- Space: O(1)

**Example:**
```go
value := analyzer.getValue(node, "name")
```

**Edge Cases:**
- Key doesn't exist
- Non-string value type
- Nested mapping values


</details>

<details>
<summary><b><code>getMappingKeys(node *yaml.Node)</code></b></summary>

**Summary:** Extracts all keys from a YAML mapping node

**Parameters:**
- `node` (*yaml.Node): YAML mapping node

**Returns:** Slice of strings containing all keys

**Complexity:**
- Time: O(n) where n is the number of keys
- Space: O(n) for storing the keys

**Example:**
```go
keys := analyzer.getMappingKeys(node)
```

**Edge Cases:**
- Empty mapping node
- Non-mapping node type
- Duplicate keys in YAML


</details>

<details>
<summary><b><code>nodeKindToString(kind yaml.Kind)</code></b></summary>

**Summary:** Converts YAML node kind to its string representation

**Parameters:**
- `kind` (yaml.Kind): YAML node kind enum value

**Returns:** String representation of the YAML node kind

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
str := analyzer.nodeKindToString(yaml.SequenceNode) // returns "sequence"
```

**Edge Cases:**
- Unknown/unsupported yaml.Kind values
- Nil receiver (though method is on pointer receiver)


</details>

<details>
<summary><b><code>sequenceItemName(index int)</code></b></summary>

**Summary:** Generates a name for sequence items by index

**Parameters:**
- `index` (int): Zero-based sequence position

**Returns:** Formatted string name for the sequence item

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
name := analyzer.sequenceItemName(2) // returns "item_2"
```

**Edge Cases:**
- Negative index values
- Extremely large indices causing formatting issues


</details>

