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

