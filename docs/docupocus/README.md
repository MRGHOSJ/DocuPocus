# 📦 Package: `docupocus`

[← Back to Overview](../README.md)

## 📄 File: `main.go`

> 📍 `docupocus\main.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>main()</code></b></summary>

**Summary:** Main function entry point in Go

**Returns:** None (void function)

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
func main() { fmt.Println("Hello, world!") }
```

**Edge Cases:**
- None (empty function does nothing)
- May panic if internal logic fails


</details>

<details>
<summary><b><code>run()</code></b></summary>

**Summary:** Function that executes logic and may return error

**Returns:** error if execution fails, nil otherwise

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
if err := run(); err != nil { log.Fatal(err) }
```

**Edge Cases:**
- May return various error types
- Could deadlock if using concurrency improperly


</details>

<details>
<summary><b><code>setupAIClient(backend string, model string, endpoint string, apiKey string, verbose bool)</code></b></summary>

**Summary:** Creates AI client with configuration parameters

**Parameters:**
- `backend` (string): Backend service name
- `model` (string): AI model identifier
- `endpoint` (string): Service API endpoint
- `apiKey` (string): Authentication key
- `verbose` (bool): Enable debug logging

**Returns:** Initialized *ai.Client or error if creation fails

**Complexity:**
- Time: O(1) (constant time initialization)
- Space: O(1) (fixed memory allocation)

**Example:**
```go
client, err := setupAIClient("openai", "gpt-4", "https://api.openai.com", "sk-...", true)
```

**Edge Cases:**
- Empty/invalid API key
- Unreachable endpoint
- Unsupported backend/model combination


</details>

<details>
<summary><b><code>generateDocs(projectDir string, template string, outputFile string, aiClient *ai.Client, verbose bool)</code></b></summary>

**Summary:** Generates documentation for a project using AI and templates

**Parameters:**
- `projectDir` (string): Path to project directory
- `template` (string): Template file path
- `outputFile` (string): Output file path
- `aiClient` (*ai.Client): AI client for processing
- `verbose` (bool): Enable verbose logging

**Returns:** Error if documentation generation fails

**Complexity:**
- Time: O(n) where n is project complexity
- Space: O(n) for output storage

**Example:**
```go
err := generateDocs("./myproject", "template.md", "docs.md", client, true)
```

**Edge Cases:**
- Invalid project directory
- Missing template file
- AI client connection failure


</details>


---

## 📄 File: `main.go`

> 📍 `docupocus\main.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>main()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>run()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>setupAIClient(backend string, model string, endpoint string, apiKey string, verbose bool)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>generateDocs(projectDir string, template string, outputFile string, aiClient *ai.Client, verbose bool)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `main.go`

> 📍 `docupocus\main.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>main()</code></b></summary>

**Summary:** Main function entry point in Go

**Returns:** None (void function)

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
func main() { fmt.Println("Hello, world!") }
```

**Edge Cases:**
- None (empty function does nothing)
- May panic if internal logic fails


</details>

<details>
<summary><b><code>run()</code></b></summary>

**Summary:** Function that executes logic and may return error

**Returns:** error if execution fails, nil otherwise

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
if err := run(); err != nil { log.Fatal(err) }
```

**Edge Cases:**
- May return various error types
- Could deadlock if using concurrency improperly


</details>

<details>
<summary><b><code>setupAIClient(backend string, model string, endpoint string, apiKey string, verbose bool)</code></b></summary>

**Summary:** Creates AI client with configuration parameters

**Parameters:**
- `backend` (string): Backend service name
- `model` (string): AI model identifier
- `endpoint` (string): Service API endpoint
- `apiKey` (string): Authentication key
- `verbose` (bool): Enable debug logging

**Returns:** Initialized *ai.Client or error if creation fails

**Complexity:**
- Time: O(1) (constant time initialization)
- Space: O(1) (fixed memory allocation)

**Example:**
```go
client, err := setupAIClient("openai", "gpt-4", "https://api.openai.com", "sk-...", true)
```

**Edge Cases:**
- Empty/invalid API key
- Unreachable endpoint
- Unsupported backend/model combination


</details>

<details>
<summary><b><code>generateDocs(projectDir string, template string, outputFile string, aiClient *ai.Client, verbose bool)</code></b></summary>

**Summary:** Generates documentation for a project using AI and templates

**Parameters:**
- `projectDir` (string): Path to project directory
- `template` (string): Template file path
- `outputFile` (string): Output file path
- `aiClient` (*ai.Client): AI client for processing
- `verbose` (bool): Enable verbose logging

**Returns:** Error if documentation generation fails

**Complexity:**
- Time: O(n) where n is project complexity
- Space: O(n) for output storage

**Example:**
```go
err := generateDocs("./myproject", "template.md", "docs.md", client, true)
```

**Edge Cases:**
- Invalid project directory
- Missing template file
- AI client connection failure


</details>


---

## 📄 File: `docupocus.go`

> 📍 `docupocus\docupocus.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>main()</code></b></summary>

**Summary:** Main function entry point in Go

**Returns:** None (void function)

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
func main() { fmt.Println("Hello, world!") }
```

**Edge Cases:**
- None (empty function does nothing)
- May panic if internal logic fails


</details>

<details>
<summary><b><code>run()</code></b></summary>

**Summary:** Function that executes logic and may return error

**Returns:** error if execution fails, nil otherwise

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
if err := run(); err != nil { log.Fatal(err) }
```

**Edge Cases:**
- May return various error types
- Could deadlock if using concurrency improperly


</details>

<details>
<summary><b><code>setupAIClient(backend string, model string, endpoint string, apiKey string, verbose bool)</code></b></summary>

**Summary:** Creates AI client with configuration parameters

**Parameters:**
- `backend` (string): Backend service name
- `model` (string): AI model identifier
- `endpoint` (string): Service API endpoint
- `apiKey` (string): Authentication key
- `verbose` (bool): Enable debug logging

**Returns:** Initialized *ai.Client or error if creation fails

**Complexity:**
- Time: O(1) (constant time initialization)
- Space: O(1) (fixed memory allocation)

**Example:**
```go
client, err := setupAIClient("openai", "gpt-4", "https://api.openai.com", "sk-...", true)
```

**Edge Cases:**
- Empty/invalid API key
- Unreachable endpoint
- Unsupported backend/model combination


</details>

<details>
<summary><b><code>generateDocs(projectDir string, template string, outputFile string, aiClient *ai.Client, verbose bool)</code></b></summary>

**Summary:** Generates documentation for a project using AI and templates

**Parameters:**
- `projectDir` (string): Path to project directory
- `template` (string): Template file path
- `outputFile` (string): Output file path
- `aiClient` (*ai.Client): AI client for processing
- `verbose` (bool): Enable verbose logging

**Returns:** Error if documentation generation fails

**Complexity:**
- Time: O(n) where n is project complexity
- Space: O(n) for output storage

**Example:**
```go
err := generateDocs("./myproject", "template.md", "docs.md", client, true)
```

**Edge Cases:**
- Invalid project directory
- Missing template file
- AI client connection failure


</details>


---

## 📄 File: `docupocus.go`

> 📍 `docupocus\docupocus.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>main()</code></b></summary>

**Summary:** Main function entry point in Go

**Returns:** None (void function)

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
func main() { fmt.Println("Hello, world!") }
```

**Edge Cases:**
- None (empty function does nothing)
- May panic if internal logic fails


</details>

<details>
<summary><b><code>run()</code></b></summary>

**Summary:** Function that executes logic and may return error

**Returns:** error if execution fails, nil otherwise

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
if err := run(); err != nil { log.Fatal(err) }
```

**Edge Cases:**
- May return various error types
- Could deadlock if using concurrency improperly


</details>

<details>
<summary><b><code>setupAIClient(backend string, Model string, endpoint string, apiKey string, verbose bool)</code></b></summary>

**Summary:** Creates AI client with configuration parameters

**Parameters:**
- `backend` (string): Backend service name
- `model` (string): AI model identifier
- `endpoint` (string): Service API endpoint
- `apiKey` (string): Authentication key
- `verbose` (bool): Enable debug logging

**Returns:** Initialized *ai.Client or error if creation fails

**Complexity:**
- Time: O(1) (constant time initialization)
- Space: O(1) (fixed memory allocation)

**Example:**
```go
client, err := setupAIClient("openai", "gpt-4", "https://api.openai.com", "sk-...", true)
```

**Edge Cases:**
- Empty/invalid API key
- Unreachable endpoint
- Unsupported backend/model combination


</details>

<details>
<summary><b><code>generateDocs(projectDir string, outputFolder string, aiClient *ai.Client, verbose bool)</code></b></summary>

**Summary:** Generates documentation for a project using AI assistance

**Parameters:**
- `projectDir` (string): Path to project directory
- `outputFolder` (string): Path to output documentation
- `aiClient` (*ai.Client): AI client for processing
- `verbose` (bool): Enable verbose logging

**Returns:** Error if documentation generation fails

**Complexity:**
- Time: O(n) where n is project complexity
- Space: O(n) for output storage

**Example:**
```go
err := generateDocs("./myproject", "./docs", client, true)
```

**Edge Cases:**
- Invalid project directory path
- AI client connection failure
- Insufficient permissions for output folder


</details>


---

## 📄 File: `docupocus.go`

> 📍 `docupocus\docupocus.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>main()</code></b></summary>

**Summary:** Main function entry point in Go

**Returns:** None (void function)

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
func main() { fmt.Println("Hello, world!") }
```

**Edge Cases:**
- None (empty function does nothing)
- May panic if internal logic fails


</details>

<details>
<summary><b><code>run()</code></b></summary>

**Summary:** Function that executes logic and may return error

**Returns:** error if execution fails, nil otherwise

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
if err := run(); err != nil { log.Fatal(err) }
```

**Edge Cases:**
- May return various error types
- Could deadlock if using concurrency improperly


</details>

<details>
<summary><b><code>setupAIClient(backend string, Model string, endpoint string, apiKey string, verbose bool)</code></b></summary>

**Summary:** Creates AI client with configuration parameters

**Parameters:**
- `backend` (string): Backend service name
- `model` (string): AI model identifier
- `endpoint` (string): Service API endpoint
- `apiKey` (string): Authentication key
- `verbose` (bool): Enable debug logging

**Returns:** Initialized *ai.Client or error if creation fails

**Complexity:**
- Time: O(1) (constant time initialization)
- Space: O(1) (fixed memory allocation)

**Example:**
```go
client, err := setupAIClient("openai", "gpt-4", "https://api.openai.com", "sk-...", true)
```

**Edge Cases:**
- Empty/invalid API key
- Unreachable endpoint
- Unsupported backend/model combination


</details>

<details>
<summary><b><code>generateDocs(projectDir string, outputFolder string, aiClient *ai.Client, verbose bool)</code></b></summary>

**Summary:** Generates documentation for a project using AI assistance

**Parameters:**
- `projectDir` (string): Path to project directory
- `outputFolder` (string): Path to output documentation
- `aiClient` (*ai.Client): AI client for processing
- `verbose` (bool): Enable verbose logging

**Returns:** Error if documentation generation fails

**Complexity:**
- Time: O(n) where n is project complexity
- Space: O(n) for output storage

**Example:**
```go
err := generateDocs("./myproject", "./docs", client, true)
```

**Edge Cases:**
- Invalid project directory path
- AI client connection failure
- Insufficient permissions for output folder


</details>


---

## 📄 File: `docupocus.go`

> 📍 `docupocus\docupocus.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>main()</code></b></summary>

**Summary:** Main function entry point in Go

**Returns:** None (void function)

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
func main() { fmt.Println("Hello, world!") }
```

**Edge Cases:**
- None (empty function does nothing)
- May panic if internal logic fails


</details>

<details>
<summary><b><code>run()</code></b></summary>

**Summary:** Function that executes logic and may return error

**Returns:** error if execution fails, nil otherwise

**Complexity:**
- Time: Depends on implementation
- Space: Depends on implementation

**Example:**
```go
if err := run(); err != nil { log.Fatal(err) }
```

**Edge Cases:**
- May return various error types
- Could deadlock if using concurrency improperly


</details>

<details>
<summary><b><code>setupAIClient(backend string, Model string, endpoint string, apiKey string, verbose bool)</code></b></summary>

**Summary:** Creates AI client with configuration parameters

**Parameters:**
- `backend` (string): Backend service name
- `model` (string): AI model identifier
- `endpoint` (string): Service API endpoint
- `apiKey` (string): Authentication key
- `verbose` (bool): Enable debug logging

**Returns:** Initialized *ai.Client or error if creation fails

**Complexity:**
- Time: O(1) (constant time initialization)
- Space: O(1) (fixed memory allocation)

**Example:**
```go
client, err := setupAIClient("openai", "gpt-4", "https://api.openai.com", "sk-...", true)
```

**Edge Cases:**
- Empty/invalid API key
- Unreachable endpoint
- Unsupported backend/model combination


</details>

<details>
<summary><b><code>generateDocs(projectDir string, outputFolder string, aiClient *ai.Client, verbose bool)</code></b></summary>

**Summary:** Generates documentation for a project using AI assistance

**Parameters:**
- `projectDir` (string): Path to project directory
- `outputFolder` (string): Path to output documentation
- `aiClient` (*ai.Client): AI client for processing
- `verbose` (bool): Enable verbose logging

**Returns:** Error if documentation generation fails

**Complexity:**
- Time: O(n) where n is project complexity
- Space: O(n) for output storage

**Example:**
```go
err := generateDocs("./myproject", "./docs", client, true)
```

**Edge Cases:**
- Invalid project directory path
- AI client connection failure
- Insufficient permissions for output folder


</details>


---

## 📄 File: `docupocus.go`

> 📍 `docupocus/docupocus.go`

## 📑 Contents

- [🔧 Functions (6)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>main()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>run()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>setupAIClient(backend string, Model string, endpoint string, apiKey string, verbose bool)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>generateDocs(projectDir string, outputFolder string, aiClient *ai.Client, verbose bool)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>getAllDiff(projectDir string, baseBranch string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>generatePRSummary(projectDir string, baseBranch string, aiClient *ai.Client)</code></b></summary>

_No documentation available._

</details>

