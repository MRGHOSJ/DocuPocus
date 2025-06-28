# 📦 Package: `utils`

[← Back to Overview](../README.md)

## 📄 File: `utils.go`

> 📍 `utils\utils.go`

## 📑 Contents

- [🔧 Functions (8)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>Exists(path string)</code></b></summary>

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
<summary><b><code>DocToString(doc *ast.CommentGroup)</code></b></summary>

**Summary:** Converts an AST comment group to string

**Parameters:**
- `doc` (*ast.CommentGroup): AST comment group to convert

**Returns:** String representation of the comment group

**Complexity:**
- Time: O(n) where n is comment length
- Space: O(n) where n is comment length

**Example:**
```go
str := DocToString(commentGroup) // converts to text
```

**Edge Cases:**
- Nil input may return empty string
- Multi-line comments may need special handling


</details>

<details>
<summary><b><code>TagToString(tag *ast.BasicLit)</code></b></summary>

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
<summary><b><code>ExprToString(expr ast.Expr)</code></b></summary>

**Summary:** Converts an AST expression to its string representation

**Parameters:**
- `expr` (ast.Expr): AST node representing an expression

**Returns:** String representation of the expression

**Complexity:**
- Time: O(n) where n is expression complexity
- Space: O(n) for recursive expressions

**Example:**
```go
str := ExprToString(astNode) // returns "x + 1" for expression x+1
```

**Edge Cases:**
- Nil input expression
- Deeply nested expressions causing stack overflow


</details>

<details>
<summary><b><code>ExprToRaw(expr ast.Expr)</code></b></summary>

**Summary:** Extracts raw string value from an AST expression

**Parameters:**
- `expr` (ast.Expr): AST node representing an expression

**Returns:** Raw string value of the expression

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
raw := ExprToRaw(astNode) // returns "42" for literal 42
```

**Edge Cases:**
- Nil input expression
- Non-literal expressions returning empty string


</details>

<details>
<summary><b><code>RecvToString(recv *ast.FieldList)</code></b></summary>

**Summary:** Converts AST field list to string representation

**Parameters:**
- `recv` (*ast.FieldList): AST field list to convert

**Returns:** String representation of the field list

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for output string storage

**Example:**
```go
str := RecvToString(fieldList) // returns "(x int, y string)"
```

**Edge Cases:**
- Nil input returns empty string
- Empty field list returns empty parentheses


</details>

<details>
<summary><b><code>HasFilesWithExtension(dir string, ext string)</code></b></summary>

**Summary:** Checks if directory contains files with extension

**Parameters:**
- `dir` (string): Directory path to search
- `ext` (string): File extension to match (e.g. ".go")

**Returns:** Boolean indicating if matching files exist

**Complexity:**
- Time: O(n) where n is number of files in directory
- Space: O(1) (constant space for directory traversal)

**Example:**
```go
found := HasFilesWithExtension("./", ".go") // returns true if .go files exist
```

**Edge Cases:**
- Non-existent directory returns false
- Empty extension matches all files
- Permission issues may cause errors


</details>

<details>
<summary><b><code>FieldTagToString(tag *ast.BasicLit)</code></b></summary>

**Summary:** Extracts string value from AST basic literal tag

**Parameters:**
- `tag` (*ast.BasicLit): AST basic literal node containing tag

**Returns:** String value of the tag without quotes

**Complexity:**
- Time: O(1)
- Space: O(n) where n is tag length

**Example:**
```go
tagStr := FieldTagToString(lit) // returns "json:\"name\"" from source
```

**Edge Cases:**
- Nil input returns empty string
- Untagged field returns empty string
- Malformed tags may return partial results


</details>


---

## 📄 File: `utils.go`

> 📍 `utils\utils.go`

## 📑 Contents

- [🔧 Functions (8)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>Exists(path string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>DocToString(doc *ast.CommentGroup)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>TagToString(tag *ast.BasicLit)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>ExprToString(expr ast.Expr)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>ExprToRaw(expr ast.Expr)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>RecvToString(recv *ast.FieldList)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>HasFilesWithExtension(dir string, ext string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>FieldTagToString(tag *ast.BasicLit)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `utils.go`

> 📍 `utils\utils.go`

## 📑 Contents

- [🔧 Functions (8)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>Exists(path string)</code></b></summary>

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
<summary><b><code>DocToString(doc *ast.CommentGroup)</code></b></summary>

**Summary:** Converts an AST comment group to string

**Parameters:**
- `doc` (*ast.CommentGroup): AST comment group to convert

**Returns:** String representation of the comment group

**Complexity:**
- Time: O(n) where n is comment length
- Space: O(n) where n is comment length

**Example:**
```go
str := DocToString(commentGroup) // converts to text
```

**Edge Cases:**
- Nil input may return empty string
- Multi-line comments may need special handling


</details>

<details>
<summary><b><code>TagToString(tag *ast.BasicLit)</code></b></summary>

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
<summary><b><code>ExprToString(expr ast.Expr)</code></b></summary>

**Summary:** Converts an AST expression to its string representation

**Parameters:**
- `expr` (ast.Expr): AST node representing an expression

**Returns:** String representation of the expression

**Complexity:**
- Time: O(n) where n is expression complexity
- Space: O(n) for recursive expressions

**Example:**
```go
str := ExprToString(astNode) // returns "x + 1" for expression x+1
```

**Edge Cases:**
- Nil input expression
- Deeply nested expressions causing stack overflow


</details>

<details>
<summary><b><code>ExprToRaw(expr ast.Expr)</code></b></summary>

**Summary:** Extracts raw string value from an AST expression

**Parameters:**
- `expr` (ast.Expr): AST node representing an expression

**Returns:** Raw string value of the expression

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
raw := ExprToRaw(astNode) // returns "42" for literal 42
```

**Edge Cases:**
- Nil input expression
- Non-literal expressions returning empty string


</details>

<details>
<summary><b><code>RecvToString(recv *ast.FieldList)</code></b></summary>

**Summary:** Converts AST field list to string representation

**Parameters:**
- `recv` (*ast.FieldList): AST field list to convert

**Returns:** String representation of the field list

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for output string storage

**Example:**
```go
str := RecvToString(fieldList) // returns "(x int, y string)"
```

**Edge Cases:**
- Nil input returns empty string
- Empty field list returns empty parentheses


</details>

<details>
<summary><b><code>HasFilesWithExtension(dir string, ext string)</code></b></summary>

**Summary:** Checks if directory contains files with extension

**Parameters:**
- `dir` (string): Directory path to search
- `ext` (string): File extension to match (e.g. ".go")

**Returns:** Boolean indicating if matching files exist

**Complexity:**
- Time: O(n) where n is number of files in directory
- Space: O(1) (constant space for directory traversal)

**Example:**
```go
found := HasFilesWithExtension("./", ".go") // returns true if .go files exist
```

**Edge Cases:**
- Non-existent directory returns false
- Empty extension matches all files
- Permission issues may cause errors


</details>

<details>
<summary><b><code>FieldTagToString(tag *ast.BasicLit)</code></b></summary>

**Summary:** Extracts string value from AST basic literal tag

**Parameters:**
- `tag` (*ast.BasicLit): AST basic literal node containing tag

**Returns:** String value of the tag without quotes

**Complexity:**
- Time: O(1)
- Space: O(n) where n is tag length

**Example:**
```go
tagStr := FieldTagToString(lit) // returns "json:\"name\"" from source
```

**Edge Cases:**
- Nil input returns empty string
- Untagged field returns empty string
- Malformed tags may return partial results


</details>


---

## 📄 File: `utils.go`

> 📍 `utils\utils.go`

## 📑 Contents

- [🔧 Functions (8)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>Exists(path string)</code></b></summary>

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
<summary><b><code>DocToString(doc *ast.CommentGroup)</code></b></summary>

**Summary:** Converts an AST comment group to string

**Parameters:**
- `doc` (*ast.CommentGroup): AST comment group to convert

**Returns:** String representation of the comment group

**Complexity:**
- Time: O(n) where n is comment length
- Space: O(n) where n is comment length

**Example:**
```go
str := DocToString(commentGroup) // converts to text
```

**Edge Cases:**
- Nil input may return empty string
- Multi-line comments may need special handling


</details>

<details>
<summary><b><code>TagToString(tag *ast.BasicLit)</code></b></summary>

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
<summary><b><code>ExprToString(expr ast.Expr)</code></b></summary>

**Summary:** Converts an AST expression to its string representation

**Parameters:**
- `expr` (ast.Expr): AST node representing an expression

**Returns:** String representation of the expression

**Complexity:**
- Time: O(n) where n is expression complexity
- Space: O(n) for recursive expressions

**Example:**
```go
str := ExprToString(astNode) // returns "x + 1" for expression x+1
```

**Edge Cases:**
- Nil input expression
- Deeply nested expressions causing stack overflow


</details>

<details>
<summary><b><code>ExprToRaw(expr ast.Expr)</code></b></summary>

**Summary:** Extracts raw string value from an AST expression

**Parameters:**
- `expr` (ast.Expr): AST node representing an expression

**Returns:** Raw string value of the expression

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
raw := ExprToRaw(astNode) // returns "42" for literal 42
```

**Edge Cases:**
- Nil input expression
- Non-literal expressions returning empty string


</details>

<details>
<summary><b><code>RecvToString(recv *ast.FieldList)</code></b></summary>

**Summary:** Converts AST field list to string representation

**Parameters:**
- `recv` (*ast.FieldList): AST field list to convert

**Returns:** String representation of the field list

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for output string storage

**Example:**
```go
str := RecvToString(fieldList) // returns "(x int, y string)"
```

**Edge Cases:**
- Nil input returns empty string
- Empty field list returns empty parentheses


</details>

<details>
<summary><b><code>HasFilesWithExtension(dir string, ext string)</code></b></summary>

**Summary:** Checks if directory contains files with extension

**Parameters:**
- `dir` (string): Directory path to search
- `ext` (string): File extension to match (e.g. ".go")

**Returns:** Boolean indicating if matching files exist

**Complexity:**
- Time: O(n) where n is number of files in directory
- Space: O(1) (constant space for directory traversal)

**Example:**
```go
found := HasFilesWithExtension("./", ".go") // returns true if .go files exist
```

**Edge Cases:**
- Non-existent directory returns false
- Empty extension matches all files
- Permission issues may cause errors


</details>

<details>
<summary><b><code>FieldTagToString(tag *ast.BasicLit)</code></b></summary>

**Summary:** Extracts string value from AST basic literal tag

**Parameters:**
- `tag` (*ast.BasicLit): AST basic literal node containing tag

**Returns:** String value of the tag without quotes

**Complexity:**
- Time: O(1)
- Space: O(n) where n is tag length

**Example:**
```go
tagStr := FieldTagToString(lit) // returns "json:\"name\"" from source
```

**Edge Cases:**
- Nil input returns empty string
- Untagged field returns empty string
- Malformed tags may return partial results


</details>


---

## 📄 File: `utils.go`

> 📍 `utils\utils.go`

## 📑 Contents

- [🔧 Functions (5)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GetDisplayPath(fullPath string)</code></b></summary>

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
<summary><b><code>GetLanguage(path string)</code></b></summary>

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
<summary><b><code>NormalizeFields(fields []analyzer.Field)</code></b></summary>

**Summary:** Model configuration for documentation generation

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(n) for inputs slice

**Example:**
```go
m := model{step: 1, projectDir: './'}
```

**Edge Cases:**
- Empty project directory
- Invalid template format
- Cursor out of inputs bounds


</details>

<details>
<summary><b><code>GetValueOrPlaceholder(f analyzer.Field)</code></b></summary>

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
<summary><b><code>CalculateDocCompletion(pkg analyzer.Package)</code></b></summary>

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


---

## 📄 File: `Globalutils.go`

> 📍 `utils\Globalutils.go`

## 📑 Contents

- [🔧 Functions (8)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>Exists(path string)</code></b></summary>

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
<summary><b><code>DocToString(doc *ast.CommentGroup)</code></b></summary>

**Summary:** Converts an AST comment group to a string

**Parameters:**
- `doc` (*ast.CommentGroup): AST comment group to convert

**Returns:** String representation of the comment group

**Complexity:**
- Time: O(n)
- Space: O(n)

**Example:**
```go
str := DocToString(commentGroup) // returns "// Example comment"
```

**Edge Cases:**
- Nil input may return empty string
- Multiline comments may require special handling


</details>

<details>
<summary><b><code>TagToString(tag *ast.BasicLit)</code></b></summary>

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
<summary><b><code>ExprToString(expr ast.Expr)</code></b></summary>

**Summary:** Converts an AST expression to its string representation

**Parameters:**
- `expr` (ast.Expr): AST node representing an expression

**Returns:** String representation of the expression

**Complexity:**
- Time: O(n) where n is expression complexity
- Space: O(n) for recursive expressions

**Example:**
```go
str := ExprToString(astNode) // returns "x + 1" for expression x+1
```

**Edge Cases:**
- Nil input expression
- Deeply nested expressions causing stack overflow


</details>

<details>
<summary><b><code>ExprToRaw(expr ast.Expr)</code></b></summary>

**Summary:** Extracts raw string value from an AST expression

**Parameters:**
- `expr` (ast.Expr): AST node representing an expression

**Returns:** Raw string value of the expression

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
raw := ExprToRaw(astNode) // returns "42" for literal 42
```

**Edge Cases:**
- Nil input expression
- Non-literal expressions returning empty string


</details>

<details>
<summary><b><code>RecvToString(recv *ast.FieldList)</code></b></summary>

**Summary:** Converts AST field list to string representation

**Parameters:**
- `recv` (*ast.FieldList): AST field list to convert

**Returns:** String representation of the field list

**Complexity:**
- Time: O(n) where n is number of fields
- Space: O(n) for output string storage

**Example:**
```go
str := RecvToString(fieldList) // returns "(x int, y string)"
```

**Edge Cases:**
- Nil input returns empty string
- Empty field list returns empty parentheses


</details>

<details>
<summary><b><code>HasFilesWithExtension(dir string, ext string)</code></b></summary>

**Summary:** Checks if directory contains files with extension

**Parameters:**
- `dir` (string): Directory path to search
- `ext` (string): File extension to match (e.g. ".go")

**Returns:** Boolean indicating if matching files exist

**Complexity:**
- Time: O(n) where n is number of files in directory
- Space: O(1) (constant space for directory traversal)

**Example:**
```go
found := HasFilesWithExtension("./", ".go") // returns true if .go files exist
```

**Edge Cases:**
- Non-existent directory returns false
- Empty extension matches all files
- Permission issues may cause errors


</details>

<details>
<summary><b><code>FieldTagToString(tag *ast.BasicLit)</code></b></summary>

**Summary:** Extracts string value from AST basic literal tag

**Parameters:**
- `tag` (*ast.BasicLit): AST basic literal node containing tag

**Returns:** String value of the tag without quotes

**Complexity:**
- Time: O(1)
- Space: O(n) where n is tag length

**Example:**
```go
tagStr := FieldTagToString(lit) // returns "json:\"name\"" from source
```

**Edge Cases:**
- Nil input returns empty string
- Untagged field returns empty string
- Malformed tags may return partial results


</details>


---

## 📄 File: `utils.go`

> 📍 `utils\utils.go`

## 📑 Contents

- [🔧 Functions (7)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>GetGitRemoteURL(projectDir string)</code></b></summary>

**Summary:** Retrieves the Git remote URL for a project directory

**Parameters:**
- `projectDir` (string): Path to the project directory

**Returns:** Git remote URL as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
url := GetGitRemoteURL("./my-project") // returns "https://github.com/user/repo.git"
```

**Edge Cases:**
- Non-Git directory may return empty string
- Missing remote configuration may cause errors


</details>

<details>
<summary><b><code>DetectLanguages(projectDir string)</code></b></summary>

**Summary:** Detects programming languages used in a project

**Parameters:**
- `projectDir` (string): Path to the project directory

**Returns:** Array of detected language names

**Complexity:**
- Time: O(n)
- Space: O(n)

**Example:**
```go
langs := DetectLanguages("./my-project") // returns ["Go", "JavaScript"]
```

**Edge Cases:**
- Empty directory may return empty array
- Unrecognized file extensions may be omitted


</details>

<details>
<summary><b><code>ParseGoMod(projectDir string)</code></b></summary>

**Summary:** Parses go.mod file in a project directory

**Parameters:**
- `projectDir` (string): Path to the project directory containing go.mod

**Returns:** Parsed content of go.mod file as string

**Complexity:**
- Time: O(n)
- Space: O(n)

**Example:**
```go
content := ParseGoMod("./myproject") // returns go.mod content
```

**Edge Cases:**
- Nonexistent project directory
- Missing or invalid go.mod file


</details>

<details>
<summary><b><code>ParsePackageJSON(projectDir string)</code></b></summary>

**Summary:** Parses package.json file in a project directory

**Parameters:**
- `projectDir` (string): Path to the project directory containing package.json

**Returns:** Parsed content and file path of package.json as two strings

**Complexity:**
- Time: O(n)
- Space: O(n)

**Example:**
```go
content, path := ParsePackageJSON("./myproject") // returns package.json content and path
```

**Edge Cases:**
- Nonexistent project directory
- Missing or invalid package.json file


</details>

<details>
<summary><b><code>ExtractDescriptionFromReadme(projectDir string)</code></b></summary>

**Summary:** Extracts description from README file in project directory

**Parameters:**
- `projectDir` (string): Path to the project directory containing README

**Returns:** Extracted description from README as string

**Complexity:**
- Time: O(n)
- Space: O(n)

**Example:**
```go
desc := ExtractDescriptionFromReadme("./myproject") // returns README description
```

**Edge Cases:**
- Nonexistent project directory
- Missing or unreadable README file
- No description found in README


</details>

<details>
<summary><b><code>DetectTechStack(projectDir string, langs []string)</code></b></summary>

**Summary:** Detects technology stack from project directory and language list

**Parameters:**
- `projectDir` (string): Path to project directory
- `langs` ([]string): List of programming languages to check

**Returns:** List of detected technologies as strings

**Complexity:**
- Time: O(n*m) where n=files, m=langs (linear scan through files for each language)
- Space: O(k) where k=detected techs (output storage)

**Example:**
```go
techs := DetectTechStack("./myapp", []string{"Go", "Python"})
```

**Edge Cases:**
- Empty project directory
- Unsupported languages in input list
- Nested node_modules directories bloating scan


</details>

<details>
<summary><b><code>DetectFeaturesAndQuickstart(projectDir string, langs []string)</code></b></summary>

**Summary:** Detects project features and generates quickstart documentation blocks

**Parameters:**
- `projectDir` (string): Path to project directory
- `langs` ([]string): List of programming languages to analyze

**Returns:** Tuple of detected features and quickstart documentation blocks

**Complexity:**
- Time: O(n*m) where n=files, m=langs (file analysis for each language)
- Space: O(f + q) where f=features, q=quickstart blocks (output storage)

**Example:**
```go
features, quickstarts := DetectFeaturesAndQuickstart("./myapp", []string{"Go", "Markdown"})
```

**Edge Cases:**
- Projects with mixed language configurations
- Missing configuration files
- Unparseable documentation comments


</details>

