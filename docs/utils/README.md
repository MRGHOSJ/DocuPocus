# 📦 Package: `utils`

> 📍 `C:\Users\DELL\Documents\GitHub\DocuPocus\internal\utils\utils.go`

[← Back to Overview](../README.md)

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

