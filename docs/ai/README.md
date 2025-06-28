# 📦 Package: `ai`

[← Back to Overview](../README.md)

## 📄 File: `cache.go`

> 📍 `ai\cache.go`

## 📑 Contents

- [🧱 Structs (2)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `Cache`

```go
type Cache struct {
	dir string 
}
```

**Summary:** Directory-based cache structure

**Returns:** None (type definition)

**Complexity:**
- Time: N/A (type definition)
- Space: O(1) (single string storage)

**Example:**
```go
cache := Cache{dir: "/tmp/cache"}
```

**Edge Cases:**
- Invalid directory path
- Permission issues
- Non-existent directory


---

### `CacheKey`

```go
type CacheKey struct {
	Hash SemanticHash 
	Language string 
}
```

**Summary:** Defines a cache key with hash and language fields

**Parameters:**
- `Hash` (SemanticHash): Unique identifier for cached content
- `Language` (string): Language identifier for the cached content

**Returns:** None (type definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
key := CacheKey{Hash: "abc123", Language: "go"}
```

**Edge Cases:**
- Empty or invalid SemanticHash
- Empty language string


---

## 🔧 Functions

<details>
<summary><b><code>NewCache(dir string)</code></b></summary>

**Summary:** Creates a new Cache instance with given directory

**Parameters:**
- `dir` (string): Directory path for cache storage

**Returns:** Pointer to initialized Cache instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
cache := NewCache("/tmp/mycache")
```

**Edge Cases:**
- Invalid or non-writable directory path
- Empty directory string


</details>

<details>
<summary><b><code>filename(key CacheKey)</code></b></summary>

**Summary:** Generates a filename from cache key

**Parameters:**
- `key` (CacheKey): Cache key containing hash and language

**Returns:** Generated filename as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
fname := cache.filename(CacheKey{Hash: "xyz", Language: "python"})
```

**Edge Cases:**
- Empty CacheKey fields
- Special characters in hash or language


</details>

<details>
<summary><b><code>Get(key CacheKey)</code></b></summary>

**Summary:** Retrieves documentation from cache using a key

**Parameters:**
- `key` (CacheKey): Key to lookup in cache

**Returns:** Documentation (if found) and boolean indicating success

**Complexity:**
- Time: O(1) average (hash map lookup)
- Space: O(1) (no additional allocation)

**Example:**
```go
doc, ok := cache.Get("user:123")
```

**Edge Cases:**
- Key not found in cache
- Nil key input
- Expired cache entry


</details>

<details>
<summary><b><code>Set(key CacheKey, doc Documentation)</code></b></summary>

**Summary:** Stores documentation in cache with a key

**Parameters:**
- `key` (CacheKey): Key for storage
- `doc` (Documentation): Value to cache

**Returns:** Error if operation fails

**Complexity:**
- Time: O(1) average (hash map insertion)
- Space: O(n) where n is cache size

**Example:**
```go
err := cache.Set("user:123", userDoc)
```

**Edge Cases:**
- Duplicate key overwrite
- Cache eviction during full capacity
- Nil value storage


</details>


---

## 📄 File: `client.go`

> 📍 `ai\client.go`

## 📑 Contents

- [🧱 Structs (3)](#-structs)
- [🔧 Functions (10)](#-functions)

## 🧱 Structs

### `Client`

```go
type Client struct {
	backend ai.Backend 
	cache *Cache 
	logger Logger 
	config ai.BackendConfig 
}
```

**Summary:** Client struct with AI backend and caching

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(1) for struct allocation

**Example:**
```go
client := Client{backend: myBackend, cache: NewCache()}
```

**Edge Cases:**
- Nil backend causing runtime errors
- Uninitialized logger
- Thread safety with shared cache


---

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

## 🔧 Functions

<details>
<summary><b><code>NewClient(backend ai.Backend, cfg ai.BackendConfig)</code></b></summary>

**Summary:** Constructor for AI client with backend configuration

**Parameters:**
- `backend` (ai.Backend): Implementation interface
- `cfg` (ai.BackendConfig): Connection settings

**Returns:** Initialized *Client instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
client := NewClient(&GPTBackend{}, config)
```

**Edge Cases:**
- Nil backend implementation
- Invalid configuration


</details>

<details>
<summary><b><code>EnhanceDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Enhances documentation for multiple inputs in specified languages

**Parameters:**
- `ctx` (context.Context): Context for cancellation and timeouts
- `inputs` ([]string): List of input texts to document
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of enhanced Documentation objects or error

**Complexity:**
- Time: O(n * m) where n=inputs, m=avg processing time per item
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.EnhanceDocumentationBatch(ctx, []string{"func foo()"}, []string{"go"})
```

**Edge Cases:**
- Empty input list
- Mismatched input/language lengths
- Unsupported languages


</details>

<details>
<summary><b><code>processBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes batch with automatic retry logic on failures

**Parameters:**
- `ctx` (context.Context): Context for cancellation and timeouts
- `inputs` ([]string): Input texts to process
- `languages` ([]string): Target languages for processing

**Returns:** Processed Documentation slice or error after retries

**Complexity:**
- Time: O(r * n) where r=retries, n=base processing time
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processBatchWithRetry(ctx, []string{"code"}, []string{"python"})
```

**Edge Cases:**
- Permanent failures after max retries
- Context cancellation during retries
- Partial batch failures


</details>

<details>
<summary><b><code>processBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Core batch processing without retry mechanism

**Parameters:**
- `ctx` (context.Context): Context for cancellation
- `inputs` ([]string): Input texts to process
- `languages` ([]string): Target processing languages

**Returns:** Processed Documentation slice or immediate error

**Complexity:**
- Time: O(n) where n=input items
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processBatch(ctx, []string{"SELECT *"}, []string{"sql"})
```

**Edge Cases:**
- First-failure behavior
- Partial processing before error
- Context deadline exceeded


</details>

<details>
<summary><b><code>callBatchAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

**Summary:** Calls batch API to generate documentation for multiple inputs

**Parameters:**
- `ctx` (context.Context): Context for cancellation and timeouts
- `inputs` ([]string): List of input strings to process
- `languages` ([]string): List of target languages for each input
- `indices` ([]int): List of indices mapping inputs to languages

**Returns:** Slice of Documentation objects and error if any

**Complexity:**
- Time: O(n) where n is the number of inputs
- Space: O(n) for storing results

**Example:**
```go
docs, err := client.callBatchAPI(ctx, []string{"code1", "code2"}, []string{"go", "python"}, []int{0, 1})
```

**Edge Cases:**
- Empty input slices
- Mismatched slice lengths
- Context cancellation


</details>

<details>
<summary><b><code>buildPrompt(input string, language string)</code></b></summary>

**Summary:** Builds a prompt string from input and language

**Parameters:**
- `input` (string): Code or text input
- `language` (string): Target programming language

**Returns:** Formatted prompt string

**Complexity:**
- Time: O(1)
- Space: O(n) where n is input length

**Example:**
```go
prompt := client.buildPrompt("func add() {}", "go")
```

**Edge Cases:**
- Empty input string
- Unsupported language


</details>

<details>
<summary><b><code>buildBatchPrompt(prompts []string)</code></b></summary>

**Summary:** Combines multiple prompts into a batch prompt

**Parameters:**
- `prompts` ([]string): List of individual prompts

**Returns:** Single concatenated prompt string

**Complexity:**
- Time: O(n) where n is total characters
- Space: O(n) for combined output

**Example:**
```go
batchPrompt := client.buildBatchPrompt([]string{"prompt1", "prompt2"})
```

**Edge Cases:**
- Empty prompts slice
- Very large combined prompt exceeding limits


</details>

<details>
<summary><b><code>parseBatchResponse(response string, expectedCount int)</code></b></summary>

**Summary:** Parses batch API response into Documentation objects

**Parameters:**
- `response` (string): Raw API response string
- `expectedCount` (int): Expected number of documents

**Returns:** Slice of Documentation objects or error if parsing fails

**Complexity:**
- Time: O(n) where n is response length
- Space: O(n) for parsed documents storage

**Example:**
```go
docs, err := client.parseBatchResponse(jsonResponse, 10)
```

**Edge Cases:**
- Malformed JSON input
- Response contains fewer items than expectedCount


</details>

<details>
<summary><b><code>generateSemanticHash(input string)</code></b></summary>

**Summary:** Generates semantic hash from input string

**Parameters:**
- `input` (string): Text to be hashed

**Returns:** SemanticHash representing input's meaning

**Complexity:**
- Time: O(n) where n is input length
- Space: O(1) for fixed-size hash output

**Example:**
```go
hash := generateSemanticHash("hello world")
```

**Edge Cases:**
- Empty input string
- Unicode/non-ASCII characters


</details>

<details>
<summary><b><code>groupByTokenCounts(inputs []string, tokenCounts []int, budget int)</code></b></summary>

**Summary:** Groups strings by token counts within budget

**Parameters:**
- `inputs` ([]string): Strings to group
- `tokenCounts` ([]int): Corresponding token counts
- `budget` (int): Max tokens per group

**Returns:** Slice of grouped input indices

**Complexity:**
- Time: O(n log n) for sorting + O(n) grouping
- Space: O(n) for output groups

**Example:**
```go
groups := groupByTokenCounts(texts, counts, 4000)
```

**Edge Cases:**
- Single item exceeds budget
- Empty inputs array
- Mismatched input/tokenCounts lengths


</details>


---

## 📄 File: `logger.go`

> 📍 `ai\logger.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `StdLogger`

```go
type StdLogger struct {
	logger *log.Logger 
}
```

**Summary:** Struct wrapping a standard logger instance

**Parameters:**
- `logger` (*log.Logger): Underlying logger instance

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
sl := StdLogger{logger: log.New(os.Stdout, "", 0)}
```

**Edge Cases:**
- Nil logger pointer initialization
- Thread safety when shared across goroutines


---

## 🔧 Functions

<details>
<summary><b><code>NewStdLogger()</code></b></summary>

**Summary:** Constructor for StdLogger with default settings

**Returns:** Initialized *StdLogger instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
logger := NewStdLogger()
```

**Edge Cases:**
- Resource exhaustion during logger creation
- Default output destination may need configuration


</details>

<details>
<summary><b><code>Info(msg string, args )</code></b></summary>

**Summary:** Logs an informational message with variadic arguments

**Parameters:**
- `msg` (string): Base log message
- `args` (...interface{}): Optional format arguments

**Returns:** None (side-effect operation)

**Complexity:**
- Time: O(n) where n is message length
- Space: O(1)

**Example:**
```go
logger.Info("User logged in", "userID", 123)
```

**Edge Cases:**
- Nil receiver panic
- Format string/argument mismatch
- Concurrent write operations


</details>

<details>
<summary><b><code>Warn(msg string, args )</code></b></summary>

**Summary:** Logs a warning message with variable arguments

**Parameters:**
- `msg` (string): Warning message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Warn("Low disk space: %dGB remaining", 5)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>

<details>
<summary><b><code>Error(msg string, args )</code></b></summary>

**Summary:** Logs an error message with variable arguments

**Parameters:**
- `msg` (string): Error message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Error("Failed to open file: %v", err)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>


---

## 📄 File: `token.go`

> 📍 `ai\token.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

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
<summary><b><code>CountTokens(text string)</code></b></summary>

**Summary:** Counts tokens in a given text string

**Parameters:**
- `text` (string): Input text to tokenize and count

**Returns:** Number of tokens as int

**Complexity:**
- Time: O(n) where n is text length
- Space: O(1)

**Example:**
```go
count := CountTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string returns 0
- Unicode/whitespace handling may vary


</details>

<details>
<summary><b><code>cheapSkipFilter(input string)</code></b></summary>

**Summary:** Cheaply filters input strings based on skip conditions

**Parameters:**
- `input` (string): String to evaluate

**Returns:** Boolean indicating whether to skip

**Complexity:**
- Time: O(1) or O(n) depending on implementation
- Space: O(1)

**Example:**
```go
skip := cheapSkipFilter("example") // returns false
```

**Edge Cases:**
- Empty string handling
- Special characters in input


</details>

<details>
<summary><b><code>groupByTokenBudget(snippets []string, languages []string, tokenBudget int)</code></b></summary>

**Summary:** Groups code snippets by language within token budget

**Parameters:**
- `snippets` ([]string): Code snippets to group
- `languages` ([]string): Programming languages for snippets
- `tokenBudget` (int): Max tokens per group

**Returns:** 2D array of snippet indices grouped by constraints

**Complexity:**
- Time: O(n^2) in worst case
- Space: O(n) for output storage

**Example:**
```go
groups := groupByTokenBudget(snippets, langs, 1000)
```

**Edge Cases:**
- Empty input arrays
- Token budget smaller than any snippet
- Mismatched snippet/language array lengths


</details>


---

## 📄 File: `utils.go`

> 📍 `ai\utils.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>ExtractJSONArray(raw string)</code></b></summary>

**Summary:** Extracts JSON array from raw string input

**Parameters:**
- `raw` (string): Raw string containing JSON array

**Returns:** Extracted JSON array as string or error if parsing fails

**Complexity:**
- Time: O(n) where n is string length (JSON parsing)
- Space: O(n) for parsed JSON storage

**Example:**
```go
arr, err := ExtractJSONArray(`{"data": [1,2,3]}`) // returns `[1,2,3]`
```

**Edge Cases:**
- Malformed JSON input
- Non-array JSON structures
- Empty input string


</details>

<details>
<summary><b><code>ParseObjectKeysAsArray(raw string, expectedCount int)</code></b></summary>

**Summary:** Parses object keys into string array with size validation

**Parameters:**
- `raw` (string): JSON object string
- `expectedCount` (int): Required number of keys

**Returns:** Array of keys or error if count mismatches

**Complexity:**
- Time: O(n) for JSON parsing + O(k) for key extraction
- Space: O(k) for key storage (k = key count)

**Example:**
```go
keys, err := ParseObjectKeysAsArray(`{"a":1,"b":2}`, 2) // returns ["a","b"]
```

**Edge Cases:**
- Non-object JSON input
- Key count mismatch
- Duplicate keys in JSON


</details>

<details>
<summary><b><code>MapResults(cachedResults []string, reverseMap []int, inputLen int)</code></b></summary>

**Summary:** Maps cached results using reverse index mapping

**Parameters:**
- `cachedResults` ([]string): Precomputed result array
- `reverseMap` ([]int): Index mapping array
- `inputLen` (int): Expected output length

**Returns:** Remapped string array according to indices

**Complexity:**
- Time: O(m) where m is inputLen
- Space: O(m) for output array

**Example:**
```go
mapped := MapResults(["a","b"], [1,0], 2) // returns ["b","a"]
```

**Edge Cases:**
- Index out of bounds in reverseMap
- Mismatch between inputLen and reverseMap length
- Empty input arrays


</details>

<details>
<summary><b><code>EstimateTokens(input string)</code></b></summary>

**Summary:** Estimates token count from input string

**Parameters:**
- `input` (string): Text to analyze for token count

**Returns:** Estimated token count as integer

**Complexity:**
- Time: O(n) where n is input length
- Space: O(1)

**Example:**
```go
tokens := EstimateTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string input
- Unicode/emoji characters
- Very long strings (memory considerations)


</details>


---

## 📄 File: `batching.go`

> 📍 `ai\batching.go`

## 📑 Contents

- [🔧 Functions (6)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>processBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>processYAMLBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>processBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>processYAMLBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>callBatchAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>callBatchYamlAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `client.go`

> 📍 `ai\client.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (9)](#-functions)

## 🧱 Structs

### `Client`

```go
type Client struct {
	backend ai.Backend 
	cache *aiCache.Cache 
	logger Logger 
	config ai.BackendConfig 
}
```

_No documentation available._

---

## 🔧 Functions

<details>
<summary><b><code>NewClient(backend ai.Backend, cfg ai.BackendConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>ApplyDefaults()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>EnhanceDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>EnhanceYAMLDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>groupByTokenCounts(inputs []string, tokenCounts []int, budget int)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>deduplicateInputs(inputs []string, languages []string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>EnhanceGenericBatch(ctx context.Context, c *Client, inputs []string, languages []string, getCache func, setCache func, callBatch func)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>jsonUnmarshalAdapter(data []byte, v *T)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>jsonMarshalIndentAdapter(v T)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `logger.go`

> 📍 `ai\logger.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `StdLogger`

```go
type StdLogger struct {
	logger *log.Logger 
}
```

_No documentation available._

---

## 🔧 Functions

<details>
<summary><b><code>NewStdLogger()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Info(msg string, args )</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Warn(msg string, args )</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Error(msg string, args )</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `parser.go`

> 📍 `ai\parser.go`

## 📑 Contents

- [🔧 Functions (2)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>parseBatchResponse(response string, expectedCount int)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>parseYAMLBatchResponse(response string, expectedCount int)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `prompt.go`

> 📍 `ai\prompt.go`

## 📑 Contents

- [🔧 Functions (3)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>buildPrompt(input string, language string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>buildBatchPromptYamlDocumentation(prompts []string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>buildBatchPromptCodeAssistant(prompts []string)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `token.go`

> 📍 `ai\token.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>init()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>CountTokens(text string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>cheapSkipFilter(input string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>groupByTokenBudget(snippets []string, languages []string, tokenBudget int)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `utils.go`

> 📍 `ai\utils.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>ExtractJSONArray(raw string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>ParseObjectKeysAsArray(raw string, expectedCount int)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>MapResults(cachedResults []string, reverseMap []int, inputLen int)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>EstimateTokens(input string)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `batching.go`

> 📍 `ai\batching.go`

## 📑 Contents

- [🔧 Functions (6)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>processBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes batch of inputs with retry logic, returns documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): Input strings to process
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of Documentation objects or error if processing fails

**Complexity:**
- Time: O(n*m) where n=inputs, m=retries
- Space: O(n) for output docs storage

**Example:**
```go
docs, err := client.processBatchWithRetry(ctx, []string{"code1", "code2"}, []string{"go"})
```

**Edge Cases:**
- Empty inputs/languages slices
- Context cancellation mid-execution
- Partial failures in batch


</details>

<details>
<summary><b><code>processYAMLBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes YAML inputs with retries, returns YAML documentation

**Parameters:**
- `ctx` (context.Context): Execution context
- `inputs` ([]string): YAML content strings
- `languages` ([]string): Target documentation languages

**Returns:** Slice of YAMLDocumentation objects or processing error

**Complexity:**
- Time: O(n*m) where n=inputs, m=retries
- Space: O(n) for YAML output storage

**Example:**
```go
yamlDocs, err := client.processYAMLBatchWithRetry(ctx, []string{"yaml1", "yaml2"}, []string{"en"})
```

**Edge Cases:**
- Invalid YAML inputs
- Language-specific YAML parsing issues
- Network failures during retries


</details>

<details>
<summary><b><code>processBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes input batch without retries, returns documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation
- `inputs` ([]string): Input strings to document
- `languages` ([]string): Target output languages

**Returns:** Documentation slice or immediate error on failure

**Complexity:**
- Time: O(n) where n=inputs count
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processBatch(ctx, []string{"func() {}"}, []string{"python"})
```

**Edge Cases:**
- First-attempt failures terminate entire batch
- Mixed-language input/output mismatches
- Memory pressure with large batches


</details>

<details>
<summary><b><code>processYAMLBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes YAML batch inputs into documentation objects

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): List of input strings to process
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of YAMLDocumentation objects or error if processing fails

**Complexity:**
- Time: O(n) where n is total input size
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processYAMLBatch(ctx, []string{"config.yml"}, []string{"en"})
```

**Edge Cases:**
- Empty inputs array
- Unsupported language codes
- Malformed YAML content


</details>

<details>
<summary><b><code>callBatchAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

**Summary:** Calls batch API to generate documentation for inputs

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): Input strings to document
- `languages` ([]string): Target documentation languages
- `indices` ([]int): Position tracking for input mapping

**Returns:** Slice of Documentation objects or error on API failure

**Complexity:**
- Time: O(n) for input processing + API call overhead
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.callBatchAPI(ctx, []string{"func() {}"}, []string{"go"}, []int{0})
```

**Edge Cases:**
- Mismatched input/language lengths
- Invalid indices
- API rate limiting


</details>

<details>
<summary><b><code>callBatchYamlAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

**Summary:** Batch processes YAML inputs via API into documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): YAML content strings
- `languages` ([]string): Target output languages
- `indices` ([]int): Input position tracking indices

**Returns:** YAMLDocumentation slice or API error

**Complexity:**
- Time: O(n) processing + O(k) API calls (k=batches)
- Space: O(n) for output + O(k) batch buffers

**Example:**
```go
yamlDocs, err := client.callBatchYamlAPI(ctx, []string{"key: value"}, []string{"en"}, []int{0})
```

**Edge Cases:**
- Invalid YAML syntax
- Partial batch failures
- Index out-of-bounds errors


</details>


---

## 📄 File: `client.go`

> 📍 `ai\client.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (9)](#-functions)

## 🧱 Structs

### `Client`

```go
type Client struct {
	backend ai.Backend 
	cache *aiCache.Cache 
	logger Logger 
	config ai.BackendConfig 
}
```

**Summary:** Client struct with backend, cache, logger, and config

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(1) (fixed-size struct)

**Example:**
```go
client := Client{backend: myBackend, cache: myCache}
```

**Edge Cases:**
- Nil dependencies
- Uninitialized config causing runtime errors


---

## 🔧 Functions

<details>
<summary><b><code>NewClient(backend ai.Backend, cfg ai.BackendConfig)</code></b></summary>

**Summary:** Constructor for AI client with backend configuration

**Parameters:**
- `backend` (ai.Backend): Implementation interface
- `cfg` (ai.BackendConfig): Connection settings

**Returns:** Initialized *Client instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
client := NewClient(&GPTBackend{}, config)
```

**Edge Cases:**
- Nil backend implementation
- Invalid configuration


</details>

<details>
<summary><b><code>ApplyDefaults()</code></b></summary>

**Summary:** Applies default settings to Client

**Returns:** None (method modifies receiver)

**Complexity:**
- Time: O(1) (assuming fixed config operations)
- Space: O(1)

**Example:**
```go
client.ApplyDefaults()
```

**Edge Cases:**
- Overwriting existing non-default values
- Nil receiver panic


</details>

<details>
<summary><b><code>EnhanceDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Generates enhanced documentation for multiple inputs in specified languages

**Parameters:**
- `ctx` (context.Context): Context for request cancellation
- `inputs` ([]string): List of input strings to document
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of enhanced documentation objects and potential error

**Complexity:**
- Time: O(n) where n is number of inputs
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.EnhanceDocumentationBatch(ctx, []string{"func foo()"}, []string{"go"})
```

**Edge Cases:**
- Empty input list
- Mismatched input/language lengths
- Unsupported languages


</details>

<details>
<summary><b><code>EnhanceYAMLDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Generates enhanced YAML documentation for multiple inputs in specified languages

**Parameters:**
- `ctx` (context.Context): Context for request cancellation
- `inputs` ([]string): List of input strings to document
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of YAML documentation objects and potential error

**Complexity:**
- Time: O(n) where n is number of inputs
- Space: O(n) for output storage

**Example:**
```go
yamlDocs, err := client.EnhanceYAMLDocumentationBatch(ctx, []string{"key: value"}, []string{"yaml"})
```

**Edge Cases:**
- Invalid YAML inputs
- Empty input list
- Mismatched input/language lengths


</details>

<details>
<summary><b><code>groupByTokenCounts(inputs []string, tokenCounts []int, budget int)</code></b></summary>

**Summary:** Groups strings by token counts within budget

**Parameters:**
- `inputs` ([]string): Strings to group
- `tokenCounts` ([]int): Corresponding token counts
- `budget` (int): Max tokens per group

**Returns:** Slice of grouped input indices

**Complexity:**
- Time: O(n log n) for sorting + O(n) grouping
- Space: O(n) for output groups

**Example:**
```go
groups := groupByTokenCounts(texts, counts, 4000)
```

**Edge Cases:**
- Single item exceeds budget
- Empty inputs array
- Mismatched input/tokenCounts lengths


</details>

<details>
<summary><b><code>deduplicateInputs(inputs []string, languages []string)</code></b></summary>

**Summary:** Deduplicates inputs while preserving mapping to original indices

**Parameters:**
- `inputs` ([]string): Input strings to deduplicate
- `languages` ([]string): Corresponding languages for each input

**Returns:** Tuple containing semantic hashes, unique inputs/languages, and original index mapping

**Complexity:**
- Time: O(n) for n inputs (assuming hash computation is O(1))
- Space: O(n) for storing unique elements and mappings

**Example:**
```go
hashes, uniqIn, uniqLang, map := client.deduplicateInputs([]string{"a","a"}, []string{"go","go"})
```

**Edge Cases:**
- Empty input lists
- Mismatched input/language lengths
- Hash collisions


</details>

<details>
<summary><b><code>EnhanceGenericBatch(ctx context.Context, c *Client, inputs []string, languages []string, getCache func, setCache func, callBatch func)</code></b></summary>

**Summary:** Processes a batch of inputs with language-specific enhancements using caching and batch calls.

**Parameters:**
- `ctx` (context.Context): Context for request cancellation and timeouts
- `c` (*Client): Client instance for making requests
- `inputs` ([]string): List of input strings to process
- `languages` ([]string): List of target languages for enhancement
- `getCache` (func): Function to retrieve cached results
- `setCache` (func): Function to store results in cache
- `callBatch` (func): Function to make batch API calls

**Returns:** Slice of processed results (type T) and an error if any

**Complexity:**
- Time: O(n * m) where n is inputs count and m is languages count
- Space: O(n * m) for storing results

**Example:**
```go
results, err := EnhanceGenericBatch(ctx, client, []string{"text1"}, []string{"en"}, getCache, setCache, callBatch)
```

**Edge Cases:**
- Empty inputs or languages lists
- Cache functions returning errors
- Context cancellation during processing


</details>

<details>
<summary><b><code>jsonUnmarshalAdapter(data []byte, v *T)</code></b></summary>

**Summary:** Unmarshals JSON data into a target variable of generic type T.

**Parameters:**
- `data` ([]byte): JSON-encoded data to unmarshal
- `v` (*T): Pointer to target variable for unmarshaling

**Returns:** Error if unmarshaling fails

**Complexity:**
- Time: O(n) where n is length of input data
- Space: O(1) (excluding space for output structure)

**Example:**
```go
err := jsonUnmarshalAdapter([]byte('{"key":"value"}'), &result)
```

**Edge Cases:**
- Malformed JSON input
- Type mismatches between JSON and target structure
- Nil pointer passed as target


</details>

<details>
<summary><b><code>jsonMarshalIndentAdapter(v T)</code></b></summary>

**Summary:** Marshals a value into indented JSON format.

**Parameters:**
- `v` (T): Value to marshal (generic type)

**Returns:** Indented JSON as byte slice and error if marshaling fails

**Complexity:**
- Time: O(n) where n is size of input structure
- Space: O(n) for JSON output

**Example:**
```go
jsonData, err := jsonMarshalIndentAdapter(myStruct)
```

**Edge Cases:**
- Cyclic references in input structure
- Unsupported types in input
- Very large structures causing memory issues


</details>


---

## 📄 File: `logger.go`

> 📍 `ai\logger.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `StdLogger`

```go
type StdLogger struct {
	logger *log.Logger 
}
```

**Summary:** Struct wrapping a standard logger instance

**Parameters:**
- `logger` (*log.Logger): Underlying logger instance

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
sl := StdLogger{logger: log.New(os.Stdout, "", 0)}
```

**Edge Cases:**
- Nil logger pointer initialization
- Thread safety when shared across goroutines


---

## 🔧 Functions

<details>
<summary><b><code>NewStdLogger()</code></b></summary>

**Summary:** Constructor for StdLogger with default settings

**Returns:** Initialized *StdLogger instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
logger := NewStdLogger()
```

**Edge Cases:**
- Resource exhaustion during logger creation
- Default output destination may need configuration


</details>

<details>
<summary><b><code>Info(msg string, args )</code></b></summary>

**Summary:** Logs an informational message with variadic arguments

**Parameters:**
- `msg` (string): Base log message
- `args` (...interface{}): Optional format arguments

**Returns:** None (side-effect operation)

**Complexity:**
- Time: O(n) where n is message length
- Space: O(1)

**Example:**
```go
logger.Info("User logged in", "userID", 123)
```

**Edge Cases:**
- Nil receiver panic
- Format string/argument mismatch
- Concurrent write operations


</details>

<details>
<summary><b><code>Warn(msg string, args )</code></b></summary>

**Summary:** Logs a warning message with variable arguments

**Parameters:**
- `msg` (string): Warning message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Warn("Low disk space: %dGB remaining", 5)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>

<details>
<summary><b><code>Error(msg string, args )</code></b></summary>

**Summary:** Logs an error message with variable arguments

**Parameters:**
- `msg` (string): Error message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Error("Failed to open file: %v", err)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>


---

## 📄 File: `parser.go`

> 📍 `ai\parser.go`

## 📑 Contents

- [🔧 Functions (2)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>parseBatchResponse(response string, expectedCount int)</code></b></summary>

**Summary:** Parses batch response string into Documentation objects

**Parameters:**
- `response` (string): Raw batch response string
- `expectedCount` (int): Expected number of documents

**Returns:** Slice of parsed Documentation objects or error

**Complexity:**
- Time: O(n) where n is response length
- Space: O(n) for output slice

**Example:**
```go
docs, err := client.parseBatchResponse(jsonResponse, 5)
```

**Edge Cases:**
- Empty response string
- Malformed response data
- Actual count ≠ expectedCount


</details>

<details>
<summary><b><code>parseYAMLBatchResponse(response string, expectedCount int)</code></b></summary>

**Summary:** Parses YAML batch response into YAMLDocumentation objects

**Parameters:**
- `response` (string): Raw YAML response string
- `expectedCount` (int): Expected document count

**Returns:** Slice of YAMLDocumentation objects or error

**Complexity:**
- Time: O(n) where n is YAML size
- Space: O(n) for output slice

**Example:**
```go
yamlDocs, err := client.parseYAMLBatchResponse(yamlResponse, 3)
```

**Edge Cases:**
- Invalid YAML syntax
- Mismatched document count
- Empty YAML content


</details>


---

## 📄 File: `prompt.go`

> 📍 `ai\prompt.go`

## 📑 Contents

- [🔧 Functions (3)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>buildPrompt(input string, language string)</code></b></summary>

**Summary:** Builds a prompt string from input and language

**Parameters:**
- `input` (string): Code or text input
- `language` (string): Target programming language

**Returns:** Formatted prompt string

**Complexity:**
- Time: O(1)
- Space: O(n) where n is input length

**Example:**
```go
prompt := client.buildPrompt("func add() {}", "go")
```

**Edge Cases:**
- Empty input string
- Unsupported language


</details>

<details>
<summary><b><code>buildBatchPromptYamlDocumentation(prompts []string)</code></b></summary>

**Summary:** Builds batch prompt for YAML documentation generation

**Parameters:**
- `prompts` ([]string): Input prompt strings

**Returns:** Formatted batch prompt string

**Complexity:**
- Time: O(n) where n is total prompt length
- Space: O(n) for output string

**Example:**
```go
batchPrompt := client.buildBatchPromptYamlDocumentation([]string{"func1", "func2"})
```

**Edge Cases:**
- Empty prompts slice
- Very large prompt strings
- Special characters in prompts


</details>

<details>
<summary><b><code>buildBatchPromptCodeAssistant(prompts []string)</code></b></summary>

**Summary:** Builds a batch prompt string from an array of prompts

**Parameters:**
- `prompts` ([]string): Array of input prompts

**Returns:** Concatenated string of all prompts

**Complexity:**
- Time: O(n), where n is the total length of all prompts
- Space: O(n), where n is the total length of the output string

**Example:**
```go
batchPrompt := client.buildBatchPromptCodeAssistant([]string{"hello", "world"})
```

**Edge Cases:**
- Empty input array returns an empty string
- Large input array may cause memory issues


</details>


---

## 📄 File: `token.go`

> 📍 `ai\token.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

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
<summary><b><code>CountTokens(text string)</code></b></summary>

**Summary:** Counts tokens in a given text string

**Parameters:**
- `text` (string): Input text to tokenize and count

**Returns:** Number of tokens as int

**Complexity:**
- Time: O(n) where n is text length
- Space: O(1)

**Example:**
```go
count := CountTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string returns 0
- Unicode/whitespace handling may vary


</details>

<details>
<summary><b><code>cheapSkipFilter(input string)</code></b></summary>

**Summary:** Cheaply filters input strings based on skip conditions

**Parameters:**
- `input` (string): String to evaluate

**Returns:** Boolean indicating whether to skip

**Complexity:**
- Time: O(1) or O(n) depending on implementation
- Space: O(1)

**Example:**
```go
skip := cheapSkipFilter("example") // returns false
```

**Edge Cases:**
- Empty string handling
- Special characters in input


</details>

<details>
<summary><b><code>groupByTokenBudget(snippets []string, languages []string, tokenBudget int)</code></b></summary>

**Summary:** Groups code snippets by language within token budget

**Parameters:**
- `snippets` ([]string): Code snippets to group
- `languages` ([]string): Programming languages for snippets
- `tokenBudget` (int): Max tokens per group

**Returns:** 2D array of snippet indices grouped by constraints

**Complexity:**
- Time: O(n^2) in worst case
- Space: O(n) for output storage

**Example:**
```go
groups := groupByTokenBudget(snippets, langs, 1000)
```

**Edge Cases:**
- Empty input arrays
- Token budget smaller than any snippet
- Mismatched snippet/language array lengths


</details>


---

## 📄 File: `utils.go`

> 📍 `ai\utils.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>ExtractJSONArray(raw string)</code></b></summary>

**Summary:** Extracts JSON array from raw string input

**Parameters:**
- `raw` (string): Raw string containing JSON array

**Returns:** Extracted JSON array as string or error if parsing fails

**Complexity:**
- Time: O(n) where n is string length (JSON parsing)
- Space: O(n) for parsed JSON storage

**Example:**
```go
arr, err := ExtractJSONArray(`{"data": [1,2,3]}`) // returns `[1,2,3]`
```

**Edge Cases:**
- Malformed JSON input
- Non-array JSON structures
- Empty input string


</details>

<details>
<summary><b><code>ParseObjectKeysAsArray(raw string, expectedCount int)</code></b></summary>

**Summary:** Parses object keys into string array with size validation

**Parameters:**
- `raw` (string): JSON object string
- `expectedCount` (int): Required number of keys

**Returns:** Array of keys or error if count mismatches

**Complexity:**
- Time: O(n) for JSON parsing + O(k) for key extraction
- Space: O(k) for key storage (k = key count)

**Example:**
```go
keys, err := ParseObjectKeysAsArray(`{"a":1,"b":2}`, 2) // returns ["a","b"]
```

**Edge Cases:**
- Non-object JSON input
- Key count mismatch
- Duplicate keys in JSON


</details>

<details>
<summary><b><code>MapResults(cachedResults []string, reverseMap []int, inputLen int)</code></b></summary>

**Summary:** Maps cached results using reverse index mapping

**Parameters:**
- `cachedResults` ([]string): Precomputed result array
- `reverseMap` ([]int): Index mapping array
- `inputLen` (int): Expected output length

**Returns:** Remapped string array according to indices

**Complexity:**
- Time: O(m) where m is inputLen
- Space: O(m) for output array

**Example:**
```go
mapped := MapResults(["a","b"], [1,0], 2) // returns ["b","a"]
```

**Edge Cases:**
- Index out of bounds in reverseMap
- Mismatch between inputLen and reverseMap length
- Empty input arrays


</details>

<details>
<summary><b><code>EstimateTokens(input string)</code></b></summary>

**Summary:** Estimates token count from input string

**Parameters:**
- `input` (string): Text to analyze for token count

**Returns:** Estimated token count as integer

**Complexity:**
- Time: O(n) where n is input length
- Space: O(1)

**Example:**
```go
tokens := EstimateTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string input
- Unicode/emoji characters
- Very long strings (memory considerations)


</details>


---

## 📄 File: `batching.go`

> 📍 `ai\batching.go`

## 📑 Contents

- [🔧 Functions (6)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>processBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes batch of inputs with retry logic, returns documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): Input strings to process
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of Documentation objects or error if processing fails

**Complexity:**
- Time: O(n*m) where n=inputs, m=retries
- Space: O(n) for output docs storage

**Example:**
```go
docs, err := client.processBatchWithRetry(ctx, []string{"code1", "code2"}, []string{"go"})
```

**Edge Cases:**
- Empty inputs/languages slices
- Context cancellation mid-execution
- Partial failures in batch


</details>

<details>
<summary><b><code>processYAMLBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes YAML inputs with retries, returns YAML documentation

**Parameters:**
- `ctx` (context.Context): Execution context
- `inputs` ([]string): YAML content strings
- `languages` ([]string): Target documentation languages

**Returns:** Slice of YAMLDocumentation objects or processing error

**Complexity:**
- Time: O(n*m) where n=inputs, m=retries
- Space: O(n) for YAML output storage

**Example:**
```go
yamlDocs, err := client.processYAMLBatchWithRetry(ctx, []string{"yaml1", "yaml2"}, []string{"en"})
```

**Edge Cases:**
- Invalid YAML inputs
- Language-specific YAML parsing issues
- Network failures during retries


</details>

<details>
<summary><b><code>processBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes input batch without retries, returns documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation
- `inputs` ([]string): Input strings to document
- `languages` ([]string): Target output languages

**Returns:** Documentation slice or immediate error on failure

**Complexity:**
- Time: O(n) where n=inputs count
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processBatch(ctx, []string{"func() {}"}, []string{"python"})
```

**Edge Cases:**
- First-attempt failures terminate entire batch
- Mixed-language input/output mismatches
- Memory pressure with large batches


</details>

<details>
<summary><b><code>processYAMLBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes YAML batch inputs into documentation objects

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): List of input strings to process
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of YAMLDocumentation objects or error if processing fails

**Complexity:**
- Time: O(n) where n is total input size
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processYAMLBatch(ctx, []string{"config.yml"}, []string{"en"})
```

**Edge Cases:**
- Empty inputs array
- Unsupported language codes
- Malformed YAML content


</details>

<details>
<summary><b><code>callBatchAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

**Summary:** Calls batch API to generate documentation for inputs

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): Input strings to document
- `languages` ([]string): Target documentation languages
- `indices` ([]int): Position tracking for input mapping

**Returns:** Slice of Documentation objects or error on API failure

**Complexity:**
- Time: O(n) for input processing + API call overhead
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.callBatchAPI(ctx, []string{"func() {}"}, []string{"go"}, []int{0})
```

**Edge Cases:**
- Mismatched input/language lengths
- Invalid indices
- API rate limiting


</details>

<details>
<summary><b><code>callBatchYamlAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

**Summary:** Batch processes YAML inputs via API into documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): YAML content strings
- `languages` ([]string): Target output languages
- `indices` ([]int): Input position tracking indices

**Returns:** YAMLDocumentation slice or API error

**Complexity:**
- Time: O(n) processing + O(k) API calls (k=batches)
- Space: O(n) for output + O(k) batch buffers

**Example:**
```go
yamlDocs, err := client.callBatchYamlAPI(ctx, []string{"key: value"}, []string{"en"}, []int{0})
```

**Edge Cases:**
- Invalid YAML syntax
- Partial batch failures
- Index out-of-bounds errors


</details>


---

## 📄 File: `client.go`

> 📍 `ai\client.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (9)](#-functions)

## 🧱 Structs

### `Client`

```go
type Client struct {
	backend ai.Backend 
	cache *aiCache.Cache 
	logger Logger 
	config ai.BackendConfig 
}
```

**Summary:** Client struct with backend, cache, logger, and config

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(1) (fixed-size struct)

**Example:**
```go
client := Client{backend: myBackend, cache: myCache}
```

**Edge Cases:**
- Nil dependencies
- Uninitialized config causing runtime errors


---

## 🔧 Functions

<details>
<summary><b><code>NewClient(backend ai.Backend, cfg ai.BackendConfig)</code></b></summary>

**Summary:** Constructor for AI client with backend configuration

**Parameters:**
- `backend` (ai.Backend): Implementation interface
- `cfg` (ai.BackendConfig): Connection settings

**Returns:** Initialized *Client instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
client := NewClient(&GPTBackend{}, config)
```

**Edge Cases:**
- Nil backend implementation
- Invalid configuration


</details>

<details>
<summary><b><code>ApplyDefaults()</code></b></summary>

**Summary:** Applies default settings to Client

**Returns:** None (method modifies receiver)

**Complexity:**
- Time: O(1) (assuming fixed config operations)
- Space: O(1)

**Example:**
```go
client.ApplyDefaults()
```

**Edge Cases:**
- Overwriting existing non-default values
- Nil receiver panic


</details>

<details>
<summary><b><code>EnhanceDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Generates enhanced documentation for multiple inputs in specified languages

**Parameters:**
- `ctx` (context.Context): Context for request cancellation
- `inputs` ([]string): List of input strings to document
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of enhanced documentation objects and potential error

**Complexity:**
- Time: O(n) where n is number of inputs
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.EnhanceDocumentationBatch(ctx, []string{"func foo()"}, []string{"go"})
```

**Edge Cases:**
- Empty input list
- Mismatched input/language lengths
- Unsupported languages


</details>

<details>
<summary><b><code>EnhanceYAMLDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Generates enhanced YAML documentation for multiple inputs in specified languages

**Parameters:**
- `ctx` (context.Context): Context for request cancellation
- `inputs` ([]string): List of input strings to document
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of YAML documentation objects and potential error

**Complexity:**
- Time: O(n) where n is number of inputs
- Space: O(n) for output storage

**Example:**
```go
yamlDocs, err := client.EnhanceYAMLDocumentationBatch(ctx, []string{"key: value"}, []string{"yaml"})
```

**Edge Cases:**
- Invalid YAML inputs
- Empty input list
- Mismatched input/language lengths


</details>

<details>
<summary><b><code>groupByTokenCounts(inputs []string, tokenCounts []int, budget int)</code></b></summary>

**Summary:** Groups strings by token counts within budget

**Parameters:**
- `inputs` ([]string): Strings to group
- `tokenCounts` ([]int): Corresponding token counts
- `budget` (int): Max tokens per group

**Returns:** Slice of grouped input indices

**Complexity:**
- Time: O(n log n) for sorting + O(n) grouping
- Space: O(n) for output groups

**Example:**
```go
groups := groupByTokenCounts(texts, counts, 4000)
```

**Edge Cases:**
- Single item exceeds budget
- Empty inputs array
- Mismatched input/tokenCounts lengths


</details>

<details>
<summary><b><code>deduplicateInputs(inputs []string, languages []string)</code></b></summary>

**Summary:** Deduplicates inputs while preserving mapping to original indices

**Parameters:**
- `inputs` ([]string): Input strings to deduplicate
- `languages` ([]string): Corresponding languages for each input

**Returns:** Tuple containing semantic hashes, unique inputs/languages, and original index mapping

**Complexity:**
- Time: O(n) for n inputs (assuming hash computation is O(1))
- Space: O(n) for storing unique elements and mappings

**Example:**
```go
hashes, uniqIn, uniqLang, map := client.deduplicateInputs([]string{"a","a"}, []string{"go","go"})
```

**Edge Cases:**
- Empty input lists
- Mismatched input/language lengths
- Hash collisions


</details>

<details>
<summary><b><code>EnhanceGenericBatch(ctx context.Context, c *Client, inputs []string, languages []string, getCache func, setCache func, callBatch func)</code></b></summary>

**Summary:** Processes a batch of inputs with language-specific enhancements using caching and batch calls.

**Parameters:**
- `ctx` (context.Context): Context for request cancellation and timeouts
- `c` (*Client): Client instance for making requests
- `inputs` ([]string): List of input strings to process
- `languages` ([]string): List of target languages for enhancement
- `getCache` (func): Function to retrieve cached results
- `setCache` (func): Function to store results in cache
- `callBatch` (func): Function to make batch API calls

**Returns:** Slice of processed results (type T) and an error if any

**Complexity:**
- Time: O(n * m) where n is inputs count and m is languages count
- Space: O(n * m) for storing results

**Example:**
```go
results, err := EnhanceGenericBatch(ctx, client, []string{"text1"}, []string{"en"}, getCache, setCache, callBatch)
```

**Edge Cases:**
- Empty inputs or languages lists
- Cache functions returning errors
- Context cancellation during processing


</details>

<details>
<summary><b><code>jsonUnmarshalAdapter(data []byte, v *T)</code></b></summary>

**Summary:** Unmarshals JSON data into a target variable of generic type T.

**Parameters:**
- `data` ([]byte): JSON-encoded data to unmarshal
- `v` (*T): Pointer to target variable for unmarshaling

**Returns:** Error if unmarshaling fails

**Complexity:**
- Time: O(n) where n is length of input data
- Space: O(1) (excluding space for output structure)

**Example:**
```go
err := jsonUnmarshalAdapter([]byte('{"key":"value"}'), &result)
```

**Edge Cases:**
- Malformed JSON input
- Type mismatches between JSON and target structure
- Nil pointer passed as target


</details>

<details>
<summary><b><code>jsonMarshalIndentAdapter(v T)</code></b></summary>

**Summary:** Marshals a value into indented JSON format.

**Parameters:**
- `v` (T): Value to marshal (generic type)

**Returns:** Indented JSON as byte slice and error if marshaling fails

**Complexity:**
- Time: O(n) where n is size of input structure
- Space: O(n) for JSON output

**Example:**
```go
jsonData, err := jsonMarshalIndentAdapter(myStruct)
```

**Edge Cases:**
- Cyclic references in input structure
- Unsupported types in input
- Very large structures causing memory issues


</details>


---

## 📄 File: `logger.go`

> 📍 `ai\logger.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `StdLogger`

```go
type StdLogger struct {
	logger *log.Logger 
}
```

**Summary:** Struct wrapping a standard logger instance

**Parameters:**
- `logger` (*log.Logger): Underlying logger instance

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
sl := StdLogger{logger: log.New(os.Stdout, "", 0)}
```

**Edge Cases:**
- Nil logger pointer initialization
- Thread safety when shared across goroutines


---

## 🔧 Functions

<details>
<summary><b><code>NewStdLogger()</code></b></summary>

**Summary:** Constructor for StdLogger with default settings

**Returns:** Initialized *StdLogger instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
logger := NewStdLogger()
```

**Edge Cases:**
- Resource exhaustion during logger creation
- Default output destination may need configuration


</details>

<details>
<summary><b><code>Info(msg string, args )</code></b></summary>

**Summary:** Logs an informational message with variadic arguments

**Parameters:**
- `msg` (string): Base log message
- `args` (...interface{}): Optional format arguments

**Returns:** None (side-effect operation)

**Complexity:**
- Time: O(n) where n is message length
- Space: O(1)

**Example:**
```go
logger.Info("User logged in", "userID", 123)
```

**Edge Cases:**
- Nil receiver panic
- Format string/argument mismatch
- Concurrent write operations


</details>

<details>
<summary><b><code>Warn(msg string, args )</code></b></summary>

**Summary:** Logs a warning message with variable arguments

**Parameters:**
- `msg` (string): Warning message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Warn("Low disk space: %dGB remaining", 5)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>

<details>
<summary><b><code>Error(msg string, args )</code></b></summary>

**Summary:** Logs an error message with variable arguments

**Parameters:**
- `msg` (string): Error message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Error("Failed to open file: %v", err)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>


---

## 📄 File: `parser.go`

> 📍 `ai\parser.go`

## 📑 Contents

- [🔧 Functions (2)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>parseBatchResponse(response string, expectedCount int)</code></b></summary>

**Summary:** Parses batch response string into Documentation objects

**Parameters:**
- `response` (string): Raw batch response string
- `expectedCount` (int): Expected number of documents

**Returns:** Slice of parsed Documentation objects or error

**Complexity:**
- Time: O(n) where n is response length
- Space: O(n) for output slice

**Example:**
```go
docs, err := client.parseBatchResponse(jsonResponse, 5)
```

**Edge Cases:**
- Empty response string
- Malformed response data
- Actual count ≠ expectedCount


</details>

<details>
<summary><b><code>parseYAMLBatchResponse(response string, expectedCount int)</code></b></summary>

**Summary:** Parses YAML batch response into YAMLDocumentation objects

**Parameters:**
- `response` (string): Raw YAML response string
- `expectedCount` (int): Expected document count

**Returns:** Slice of YAMLDocumentation objects or error

**Complexity:**
- Time: O(n) where n is YAML size
- Space: O(n) for output slice

**Example:**
```go
yamlDocs, err := client.parseYAMLBatchResponse(yamlResponse, 3)
```

**Edge Cases:**
- Invalid YAML syntax
- Mismatched document count
- Empty YAML content


</details>


---

## 📄 File: `prompt.go`

> 📍 `ai\prompt.go`

## 📑 Contents

- [🔧 Functions (3)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>buildPrompt(input string, language string)</code></b></summary>

**Summary:** Builds a prompt string from input and language

**Parameters:**
- `input` (string): Code or text input
- `language` (string): Target programming language

**Returns:** Formatted prompt string

**Complexity:**
- Time: O(1)
- Space: O(n) where n is input length

**Example:**
```go
prompt := client.buildPrompt("func add() {}", "go")
```

**Edge Cases:**
- Empty input string
- Unsupported language


</details>

<details>
<summary><b><code>buildBatchPromptYamlDocumentation(prompts []string)</code></b></summary>

**Summary:** Builds batch prompt for YAML documentation generation

**Parameters:**
- `prompts` ([]string): Input prompt strings

**Returns:** Formatted batch prompt string

**Complexity:**
- Time: O(n) where n is total prompt length
- Space: O(n) for output string

**Example:**
```go
batchPrompt := client.buildBatchPromptYamlDocumentation([]string{"func1", "func2"})
```

**Edge Cases:**
- Empty prompts slice
- Very large prompt strings
- Special characters in prompts


</details>

<details>
<summary><b><code>buildBatchPromptCodeAssistant(prompts []string)</code></b></summary>

**Summary:** Builds a batch prompt string from an array of prompts

**Parameters:**
- `prompts` ([]string): Array of input prompts

**Returns:** Concatenated string of all prompts

**Complexity:**
- Time: O(n), where n is the total length of all prompts
- Space: O(n), where n is the total length of the output string

**Example:**
```go
batchPrompt := client.buildBatchPromptCodeAssistant([]string{"hello", "world"})
```

**Edge Cases:**
- Empty input array returns an empty string
- Large input array may cause memory issues


</details>


---

## 📄 File: `token.go`

> 📍 `ai\token.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

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
<summary><b><code>CountTokens(text string)</code></b></summary>

**Summary:** Counts tokens in a given text string

**Parameters:**
- `text` (string): Input text to tokenize and count

**Returns:** Number of tokens as int

**Complexity:**
- Time: O(n) where n is text length
- Space: O(1)

**Example:**
```go
count := CountTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string returns 0
- Unicode/whitespace handling may vary


</details>

<details>
<summary><b><code>cheapSkipFilter(input string)</code></b></summary>

**Summary:** Cheaply filters input strings based on skip conditions

**Parameters:**
- `input` (string): String to evaluate

**Returns:** Boolean indicating whether to skip

**Complexity:**
- Time: O(1) or O(n) depending on implementation
- Space: O(1)

**Example:**
```go
skip := cheapSkipFilter("example") // returns false
```

**Edge Cases:**
- Empty string handling
- Special characters in input


</details>

<details>
<summary><b><code>groupByTokenBudget(snippets []string, languages []string, tokenBudget int)</code></b></summary>

**Summary:** Groups code snippets by language within token budget

**Parameters:**
- `snippets` ([]string): Code snippets to group
- `languages` ([]string): Programming languages for snippets
- `tokenBudget` (int): Max tokens per group

**Returns:** 2D array of snippet indices grouped by constraints

**Complexity:**
- Time: O(n^2) in worst case
- Space: O(n) for output storage

**Example:**
```go
groups := groupByTokenBudget(snippets, langs, 1000)
```

**Edge Cases:**
- Empty input arrays
- Token budget smaller than any snippet
- Mismatched snippet/language array lengths


</details>


---

## 📄 File: `utils.go`

> 📍 `ai\utils.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>ExtractJSONArray(raw string)</code></b></summary>

**Summary:** Extracts JSON array from raw string input

**Parameters:**
- `raw` (string): Raw string containing JSON array

**Returns:** Extracted JSON array as string or error if parsing fails

**Complexity:**
- Time: O(n) where n is string length (JSON parsing)
- Space: O(n) for parsed JSON storage

**Example:**
```go
arr, err := ExtractJSONArray(`{"data": [1,2,3]}`) // returns `[1,2,3]`
```

**Edge Cases:**
- Malformed JSON input
- Non-array JSON structures
- Empty input string


</details>

<details>
<summary><b><code>ParseObjectKeysAsArray(raw string, expectedCount int)</code></b></summary>

**Summary:** Parses object keys into string array with size validation

**Parameters:**
- `raw` (string): JSON object string
- `expectedCount` (int): Required number of keys

**Returns:** Array of keys or error if count mismatches

**Complexity:**
- Time: O(n) for JSON parsing + O(k) for key extraction
- Space: O(k) for key storage (k = key count)

**Example:**
```go
keys, err := ParseObjectKeysAsArray(`{"a":1,"b":2}`, 2) // returns ["a","b"]
```

**Edge Cases:**
- Non-object JSON input
- Key count mismatch
- Duplicate keys in JSON


</details>

<details>
<summary><b><code>MapResults(cachedResults []string, reverseMap []int, inputLen int)</code></b></summary>

**Summary:** Maps cached results using reverse index mapping

**Parameters:**
- `cachedResults` ([]string): Precomputed result array
- `reverseMap` ([]int): Index mapping array
- `inputLen` (int): Expected output length

**Returns:** Remapped string array according to indices

**Complexity:**
- Time: O(m) where m is inputLen
- Space: O(m) for output array

**Example:**
```go
mapped := MapResults(["a","b"], [1,0], 2) // returns ["b","a"]
```

**Edge Cases:**
- Index out of bounds in reverseMap
- Mismatch between inputLen and reverseMap length
- Empty input arrays


</details>

<details>
<summary><b><code>EstimateTokens(input string)</code></b></summary>

**Summary:** Estimates token count from input string

**Parameters:**
- `input` (string): Text to analyze for token count

**Returns:** Estimated token count as integer

**Complexity:**
- Time: O(n) where n is input length
- Space: O(1)

**Example:**
```go
tokens := EstimateTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string input
- Unicode/emoji characters
- Very long strings (memory considerations)


</details>


---

## 📄 File: `batching.go`

> 📍 `ai\batching.go`

## 📑 Contents

- [🔧 Functions (6)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>processBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes batch of inputs with retry logic, returns documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): Input strings to process
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of Documentation objects or error if processing fails

**Complexity:**
- Time: O(n*m) where n=inputs, m=retries
- Space: O(n) for output docs storage

**Example:**
```go
docs, err := client.processBatchWithRetry(ctx, []string{"code1", "code2"}, []string{"go"})
```

**Edge Cases:**
- Empty inputs/languages slices
- Context cancellation mid-execution
- Partial failures in batch


</details>

<details>
<summary><b><code>processYAMLBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes YAML inputs with retries, returns YAML documentation

**Parameters:**
- `ctx` (context.Context): Execution context
- `inputs` ([]string): YAML content strings
- `languages` ([]string): Target documentation languages

**Returns:** Slice of YAMLDocumentation objects or processing error

**Complexity:**
- Time: O(n*m) where n=inputs, m=retries
- Space: O(n) for YAML output storage

**Example:**
```go
yamlDocs, err := client.processYAMLBatchWithRetry(ctx, []string{"yaml1", "yaml2"}, []string{"en"})
```

**Edge Cases:**
- Invalid YAML inputs
- Language-specific YAML parsing issues
- Network failures during retries


</details>

<details>
<summary><b><code>processBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes input batch without retries, returns documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation
- `inputs` ([]string): Input strings to document
- `languages` ([]string): Target output languages

**Returns:** Documentation slice or immediate error on failure

**Complexity:**
- Time: O(n) where n=inputs count
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processBatch(ctx, []string{"func() {}"}, []string{"python"})
```

**Edge Cases:**
- First-attempt failures terminate entire batch
- Mixed-language input/output mismatches
- Memory pressure with large batches


</details>

<details>
<summary><b><code>processYAMLBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes YAML batch inputs into documentation objects

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): List of input strings to process
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of YAMLDocumentation objects or error if processing fails

**Complexity:**
- Time: O(n) where n is total input size
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processYAMLBatch(ctx, []string{"config.yml"}, []string{"en"})
```

**Edge Cases:**
- Empty inputs array
- Unsupported language codes
- Malformed YAML content


</details>

<details>
<summary><b><code>callBatchAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

**Summary:** Calls batch API to generate documentation for inputs

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): Input strings to document
- `languages` ([]string): Target documentation languages
- `indices` ([]int): Position tracking for input mapping

**Returns:** Slice of Documentation objects or error on API failure

**Complexity:**
- Time: O(n) for input processing + API call overhead
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.callBatchAPI(ctx, []string{"func() {}"}, []string{"go"}, []int{0})
```

**Edge Cases:**
- Mismatched input/language lengths
- Invalid indices
- API rate limiting


</details>

<details>
<summary><b><code>callBatchYamlAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

**Summary:** Batch processes YAML inputs via API into documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): YAML content strings
- `languages` ([]string): Target output languages
- `indices` ([]int): Input position tracking indices

**Returns:** YAMLDocumentation slice or API error

**Complexity:**
- Time: O(n) processing + O(k) API calls (k=batches)
- Space: O(n) for output + O(k) batch buffers

**Example:**
```go
yamlDocs, err := client.callBatchYamlAPI(ctx, []string{"key: value"}, []string{"en"}, []int{0})
```

**Edge Cases:**
- Invalid YAML syntax
- Partial batch failures
- Index out-of-bounds errors


</details>


---

## 📄 File: `client.go`

> 📍 `ai\client.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (9)](#-functions)

## 🧱 Structs

### `Client`

```go
type Client struct {
	backend ai.Backend 
	cache *aiCache.Cache 
	logger Logger 
	config ai.BackendConfig 
}
```

**Summary:** Client struct with backend, cache, logger, and config

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(1) (fixed-size struct)

**Example:**
```go
client := Client{backend: myBackend, cache: myCache}
```

**Edge Cases:**
- Nil dependencies
- Uninitialized config causing runtime errors


---

## 🔧 Functions

<details>
<summary><b><code>NewClient(backend ai.Backend, cfg ai.BackendConfig)</code></b></summary>

**Summary:** Constructor for AI client with backend configuration

**Parameters:**
- `backend` (ai.Backend): Implementation interface
- `cfg` (ai.BackendConfig): Connection settings

**Returns:** Initialized *Client instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
client := NewClient(&GPTBackend{}, config)
```

**Edge Cases:**
- Nil backend implementation
- Invalid configuration


</details>

<details>
<summary><b><code>ApplyDefaults()</code></b></summary>

**Summary:** Applies default settings to Client

**Returns:** None (method modifies receiver)

**Complexity:**
- Time: O(1) (assuming fixed config operations)
- Space: O(1)

**Example:**
```go
client.ApplyDefaults()
```

**Edge Cases:**
- Overwriting existing non-default values
- Nil receiver panic


</details>

<details>
<summary><b><code>EnhanceDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Generates enhanced documentation for multiple inputs in specified languages

**Parameters:**
- `ctx` (context.Context): Context for request cancellation
- `inputs` ([]string): List of input strings to document
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of enhanced documentation objects and potential error

**Complexity:**
- Time: O(n) where n is number of inputs
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.EnhanceDocumentationBatch(ctx, []string{"func foo()"}, []string{"go"})
```

**Edge Cases:**
- Empty input list
- Mismatched input/language lengths
- Unsupported languages


</details>

<details>
<summary><b><code>EnhanceYAMLDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Generates enhanced YAML documentation for multiple inputs in specified languages

**Parameters:**
- `ctx` (context.Context): Context for request cancellation
- `inputs` ([]string): List of input strings to document
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of YAML documentation objects and potential error

**Complexity:**
- Time: O(n) where n is number of inputs
- Space: O(n) for output storage

**Example:**
```go
yamlDocs, err := client.EnhanceYAMLDocumentationBatch(ctx, []string{"key: value"}, []string{"yaml"})
```

**Edge Cases:**
- Invalid YAML inputs
- Empty input list
- Mismatched input/language lengths


</details>

<details>
<summary><b><code>groupByTokenCounts(inputs []string, tokenCounts []int, budget int)</code></b></summary>

**Summary:** Groups strings by token counts within budget

**Parameters:**
- `inputs` ([]string): Strings to group
- `tokenCounts` ([]int): Corresponding token counts
- `budget` (int): Max tokens per group

**Returns:** Slice of grouped input indices

**Complexity:**
- Time: O(n log n) for sorting + O(n) grouping
- Space: O(n) for output groups

**Example:**
```go
groups := groupByTokenCounts(texts, counts, 4000)
```

**Edge Cases:**
- Single item exceeds budget
- Empty inputs array
- Mismatched input/tokenCounts lengths


</details>

<details>
<summary><b><code>deduplicateInputs(inputs []string, languages []string)</code></b></summary>

**Summary:** Deduplicates inputs while preserving mapping to original indices

**Parameters:**
- `inputs` ([]string): Input strings to deduplicate
- `languages` ([]string): Corresponding languages for each input

**Returns:** Tuple containing semantic hashes, unique inputs/languages, and original index mapping

**Complexity:**
- Time: O(n) for n inputs (assuming hash computation is O(1))
- Space: O(n) for storing unique elements and mappings

**Example:**
```go
hashes, uniqIn, uniqLang, map := client.deduplicateInputs([]string{"a","a"}, []string{"go","go"})
```

**Edge Cases:**
- Empty input lists
- Mismatched input/language lengths
- Hash collisions


</details>

<details>
<summary><b><code>EnhanceGenericBatch(ctx context.Context, c *Client, inputs []string, languages []string, getCache func, setCache func, callBatch func)</code></b></summary>

**Summary:** Processes a batch of inputs with language-specific enhancements using caching and batch calls.

**Parameters:**
- `ctx` (context.Context): Context for request cancellation and timeouts
- `c` (*Client): Client instance for making requests
- `inputs` ([]string): List of input strings to process
- `languages` ([]string): List of target languages for enhancement
- `getCache` (func): Function to retrieve cached results
- `setCache` (func): Function to store results in cache
- `callBatch` (func): Function to make batch API calls

**Returns:** Slice of processed results (type T) and an error if any

**Complexity:**
- Time: O(n * m) where n is inputs count and m is languages count
- Space: O(n * m) for storing results

**Example:**
```go
results, err := EnhanceGenericBatch(ctx, client, []string{"text1"}, []string{"en"}, getCache, setCache, callBatch)
```

**Edge Cases:**
- Empty inputs or languages lists
- Cache functions returning errors
- Context cancellation during processing


</details>

<details>
<summary><b><code>jsonUnmarshalAdapter(data []byte, v *T)</code></b></summary>

**Summary:** Unmarshals JSON data into a target variable of generic type T.

**Parameters:**
- `data` ([]byte): JSON-encoded data to unmarshal
- `v` (*T): Pointer to target variable for unmarshaling

**Returns:** Error if unmarshaling fails

**Complexity:**
- Time: O(n) where n is length of input data
- Space: O(1) (excluding space for output structure)

**Example:**
```go
err := jsonUnmarshalAdapter([]byte('{"key":"value"}'), &result)
```

**Edge Cases:**
- Malformed JSON input
- Type mismatches between JSON and target structure
- Nil pointer passed as target


</details>

<details>
<summary><b><code>jsonMarshalIndentAdapter(v T)</code></b></summary>

**Summary:** Marshals a value into indented JSON format.

**Parameters:**
- `v` (T): Value to marshal (generic type)

**Returns:** Indented JSON as byte slice and error if marshaling fails

**Complexity:**
- Time: O(n) where n is size of input structure
- Space: O(n) for JSON output

**Example:**
```go
jsonData, err := jsonMarshalIndentAdapter(myStruct)
```

**Edge Cases:**
- Cyclic references in input structure
- Unsupported types in input
- Very large structures causing memory issues


</details>


---

## 📄 File: `logger.go`

> 📍 `ai\logger.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `StdLogger`

```go
type StdLogger struct {
	logger *log.Logger 
}
```

**Summary:** Struct wrapping a standard logger instance

**Parameters:**
- `logger` (*log.Logger): Underlying logger instance

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
sl := StdLogger{logger: log.New(os.Stdout, "", 0)}
```

**Edge Cases:**
- Nil logger pointer initialization
- Thread safety when shared across goroutines


---

## 🔧 Functions

<details>
<summary><b><code>NewStdLogger()</code></b></summary>

**Summary:** Constructor for StdLogger with default settings

**Returns:** Initialized *StdLogger instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
logger := NewStdLogger()
```

**Edge Cases:**
- Resource exhaustion during logger creation
- Default output destination may need configuration


</details>

<details>
<summary><b><code>Info(msg string, args )</code></b></summary>

**Summary:** Logs an informational message with variadic arguments

**Parameters:**
- `msg` (string): Base log message
- `args` (...interface{}): Optional format arguments

**Returns:** None (side-effect operation)

**Complexity:**
- Time: O(n) where n is message length
- Space: O(1)

**Example:**
```go
logger.Info("User logged in", "userID", 123)
```

**Edge Cases:**
- Nil receiver panic
- Format string/argument mismatch
- Concurrent write operations


</details>

<details>
<summary><b><code>Warn(msg string, args )</code></b></summary>

**Summary:** Logs a warning message with variable arguments

**Parameters:**
- `msg` (string): Warning message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Warn("Low disk space: %dGB remaining", 5)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>

<details>
<summary><b><code>Error(msg string, args )</code></b></summary>

**Summary:** Logs an error message with variable arguments

**Parameters:**
- `msg` (string): Error message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Error("Failed to open file: %v", err)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>


---

## 📄 File: `parser.go`

> 📍 `ai\parser.go`

## 📑 Contents

- [🔧 Functions (2)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>parseBatchResponse(response string, expectedCount int)</code></b></summary>

**Summary:** Parses batch response string into Documentation objects

**Parameters:**
- `response` (string): Raw batch response string
- `expectedCount` (int): Expected number of documents

**Returns:** Slice of parsed Documentation objects or error

**Complexity:**
- Time: O(n) where n is response length
- Space: O(n) for output slice

**Example:**
```go
docs, err := client.parseBatchResponse(jsonResponse, 5)
```

**Edge Cases:**
- Empty response string
- Malformed response data
- Actual count ≠ expectedCount


</details>

<details>
<summary><b><code>parseYAMLBatchResponse(response string, expectedCount int)</code></b></summary>

**Summary:** Parses YAML batch response into YAMLDocumentation objects

**Parameters:**
- `response` (string): Raw YAML response string
- `expectedCount` (int): Expected document count

**Returns:** Slice of YAMLDocumentation objects or error

**Complexity:**
- Time: O(n) where n is YAML size
- Space: O(n) for output slice

**Example:**
```go
yamlDocs, err := client.parseYAMLBatchResponse(yamlResponse, 3)
```

**Edge Cases:**
- Invalid YAML syntax
- Mismatched document count
- Empty YAML content


</details>


---

## 📄 File: `prompt.go`

> 📍 `ai\prompt.go`

## 📑 Contents

- [🔧 Functions (3)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>buildPrompt(input string, language string)</code></b></summary>

**Summary:** Builds a prompt string from input and language

**Parameters:**
- `input` (string): Code or text input
- `language` (string): Target programming language

**Returns:** Formatted prompt string

**Complexity:**
- Time: O(1)
- Space: O(n) where n is input length

**Example:**
```go
prompt := client.buildPrompt("func add() {}", "go")
```

**Edge Cases:**
- Empty input string
- Unsupported language


</details>

<details>
<summary><b><code>buildBatchPromptYamlDocumentation(prompts []string)</code></b></summary>

**Summary:** Builds batch prompt for YAML documentation generation

**Parameters:**
- `prompts` ([]string): Input prompt strings

**Returns:** Formatted batch prompt string

**Complexity:**
- Time: O(n) where n is total prompt length
- Space: O(n) for output string

**Example:**
```go
batchPrompt := client.buildBatchPromptYamlDocumentation([]string{"func1", "func2"})
```

**Edge Cases:**
- Empty prompts slice
- Very large prompt strings
- Special characters in prompts


</details>

<details>
<summary><b><code>buildBatchPromptCodeAssistant(prompts []string)</code></b></summary>

**Summary:** Builds a batch prompt string from an array of prompts

**Parameters:**
- `prompts` ([]string): Array of input prompts

**Returns:** Concatenated string of all prompts

**Complexity:**
- Time: O(n), where n is the total length of all prompts
- Space: O(n), where n is the total length of the output string

**Example:**
```go
batchPrompt := client.buildBatchPromptCodeAssistant([]string{"hello", "world"})
```

**Edge Cases:**
- Empty input array returns an empty string
- Large input array may cause memory issues


</details>


---

## 📄 File: `token.go`

> 📍 `ai\token.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

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
<summary><b><code>CountTokens(text string)</code></b></summary>

**Summary:** Counts tokens in a given text string

**Parameters:**
- `text` (string): Input text to tokenize and count

**Returns:** Number of tokens as int

**Complexity:**
- Time: O(n) where n is text length
- Space: O(1)

**Example:**
```go
count := CountTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string returns 0
- Unicode/whitespace handling may vary


</details>

<details>
<summary><b><code>cheapSkipFilter(input string)</code></b></summary>

**Summary:** Cheaply filters input strings based on skip conditions

**Parameters:**
- `input` (string): String to evaluate

**Returns:** Boolean indicating whether to skip

**Complexity:**
- Time: O(1) or O(n) depending on implementation
- Space: O(1)

**Example:**
```go
skip := cheapSkipFilter("example") // returns false
```

**Edge Cases:**
- Empty string handling
- Special characters in input


</details>

<details>
<summary><b><code>groupByTokenBudget(snippets []string, languages []string, tokenBudget int)</code></b></summary>

**Summary:** Groups code snippets by language within token budget

**Parameters:**
- `snippets` ([]string): Code snippets to group
- `languages` ([]string): Programming languages for snippets
- `tokenBudget` (int): Max tokens per group

**Returns:** 2D array of snippet indices grouped by constraints

**Complexity:**
- Time: O(n^2) in worst case
- Space: O(n) for output storage

**Example:**
```go
groups := groupByTokenBudget(snippets, langs, 1000)
```

**Edge Cases:**
- Empty input arrays
- Token budget smaller than any snippet
- Mismatched snippet/language array lengths


</details>


---

## 📄 File: `utils.go`

> 📍 `ai\utils.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>ExtractJSONArray(raw string)</code></b></summary>

**Summary:** Extracts JSON array from raw string input

**Parameters:**
- `raw` (string): Raw string containing JSON array

**Returns:** Extracted JSON array as string or error if parsing fails

**Complexity:**
- Time: O(n) where n is string length (JSON parsing)
- Space: O(n) for parsed JSON storage

**Example:**
```go
arr, err := ExtractJSONArray(`{"data": [1,2,3]}`) // returns `[1,2,3]`
```

**Edge Cases:**
- Malformed JSON input
- Non-array JSON structures
- Empty input string


</details>

<details>
<summary><b><code>ParseObjectKeysAsArray(raw string, expectedCount int)</code></b></summary>

**Summary:** Parses object keys into string array with size validation

**Parameters:**
- `raw` (string): JSON object string
- `expectedCount` (int): Required number of keys

**Returns:** Array of keys or error if count mismatches

**Complexity:**
- Time: O(n) for JSON parsing + O(k) for key extraction
- Space: O(k) for key storage (k = key count)

**Example:**
```go
keys, err := ParseObjectKeysAsArray(`{"a":1,"b":2}`, 2) // returns ["a","b"]
```

**Edge Cases:**
- Non-object JSON input
- Key count mismatch
- Duplicate keys in JSON


</details>

<details>
<summary><b><code>MapResults(cachedResults []string, reverseMap []int, inputLen int)</code></b></summary>

**Summary:** Maps cached results using reverse index mapping

**Parameters:**
- `cachedResults` ([]string): Precomputed result array
- `reverseMap` ([]int): Index mapping array
- `inputLen` (int): Expected output length

**Returns:** Remapped string array according to indices

**Complexity:**
- Time: O(m) where m is inputLen
- Space: O(m) for output array

**Example:**
```go
mapped := MapResults(["a","b"], [1,0], 2) // returns ["b","a"]
```

**Edge Cases:**
- Index out of bounds in reverseMap
- Mismatch between inputLen and reverseMap length
- Empty input arrays


</details>

<details>
<summary><b><code>EstimateTokens(input string)</code></b></summary>

**Summary:** Estimates token count from input string

**Parameters:**
- `input` (string): Text to analyze for token count

**Returns:** Estimated token count as integer

**Complexity:**
- Time: O(n) where n is input length
- Space: O(1)

**Example:**
```go
tokens := EstimateTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string input
- Unicode/emoji characters
- Very long strings (memory considerations)


</details>


---

## 📄 File: `batching.go`

> 📍 `ai\batching.go`

## 📑 Contents

- [🔧 Functions (6)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>processBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes batch of inputs with retry logic, returns documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): Input strings to process
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of Documentation objects or error if processing fails

**Complexity:**
- Time: O(n*m) where n=inputs, m=retries
- Space: O(n) for output docs storage

**Example:**
```go
docs, err := client.processBatchWithRetry(ctx, []string{"code1", "code2"}, []string{"go"})
```

**Edge Cases:**
- Empty inputs/languages slices
- Context cancellation mid-execution
- Partial failures in batch


</details>

<details>
<summary><b><code>processYAMLBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes YAML inputs with retries, returns YAML documentation

**Parameters:**
- `ctx` (context.Context): Execution context
- `inputs` ([]string): YAML content strings
- `languages` ([]string): Target documentation languages

**Returns:** Slice of YAMLDocumentation objects or processing error

**Complexity:**
- Time: O(n*m) where n=inputs, m=retries
- Space: O(n) for YAML output storage

**Example:**
```go
yamlDocs, err := client.processYAMLBatchWithRetry(ctx, []string{"yaml1", "yaml2"}, []string{"en"})
```

**Edge Cases:**
- Invalid YAML inputs
- Language-specific YAML parsing issues
- Network failures during retries


</details>

<details>
<summary><b><code>processBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes input batch without retries, returns documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation
- `inputs` ([]string): Input strings to document
- `languages` ([]string): Target output languages

**Returns:** Documentation slice or immediate error on failure

**Complexity:**
- Time: O(n) where n=inputs count
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processBatch(ctx, []string{"func() {}"}, []string{"python"})
```

**Edge Cases:**
- First-attempt failures terminate entire batch
- Mixed-language input/output mismatches
- Memory pressure with large batches


</details>

<details>
<summary><b><code>processYAMLBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes YAML batch inputs into documentation objects

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): List of input strings to process
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of YAMLDocumentation objects or error if processing fails

**Complexity:**
- Time: O(n) where n is total input size
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processYAMLBatch(ctx, []string{"config.yml"}, []string{"en"})
```

**Edge Cases:**
- Empty inputs array
- Unsupported language codes
- Malformed YAML content


</details>

<details>
<summary><b><code>callBatchAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

**Summary:** Calls batch API to generate documentation for inputs

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): Input strings to document
- `languages` ([]string): Target documentation languages
- `indices` ([]int): Position tracking for input mapping

**Returns:** Slice of Documentation objects or error on API failure

**Complexity:**
- Time: O(n) for input processing + API call overhead
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.callBatchAPI(ctx, []string{"func() {}"}, []string{"go"}, []int{0})
```

**Edge Cases:**
- Mismatched input/language lengths
- Invalid indices
- API rate limiting


</details>

<details>
<summary><b><code>callBatchYamlAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

**Summary:** Batch processes YAML inputs via API into documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): YAML content strings
- `languages` ([]string): Target output languages
- `indices` ([]int): Input position tracking indices

**Returns:** YAMLDocumentation slice or API error

**Complexity:**
- Time: O(n) processing + O(k) API calls (k=batches)
- Space: O(n) for output + O(k) batch buffers

**Example:**
```go
yamlDocs, err := client.callBatchYamlAPI(ctx, []string{"key: value"}, []string{"en"}, []int{0})
```

**Edge Cases:**
- Invalid YAML syntax
- Partial batch failures
- Index out-of-bounds errors


</details>


---

## 📄 File: `client.go`

> 📍 `ai\client.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (9)](#-functions)

## 🧱 Structs

### `Client`

```go
type Client struct {
	backend ai.Backend 
	cache *aiCache.Cache 
	logger Logger 
	config ai.BackendConfig 
}
```

**Summary:** Client struct with backend, cache, logger, and config

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(1) (fixed-size struct)

**Example:**
```go
client := Client{backend: myBackend, cache: myCache}
```

**Edge Cases:**
- Nil dependencies
- Uninitialized config causing runtime errors


---

## 🔧 Functions

<details>
<summary><b><code>NewClient(backend ai.Backend, cfg ai.BackendConfig)</code></b></summary>

**Summary:** Constructor for AI client with backend configuration

**Parameters:**
- `backend` (ai.Backend): Implementation interface
- `cfg` (ai.BackendConfig): Connection settings

**Returns:** Initialized *Client instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
client := NewClient(&GPTBackend{}, config)
```

**Edge Cases:**
- Nil backend implementation
- Invalid configuration


</details>

<details>
<summary><b><code>ApplyDefaults()</code></b></summary>

**Summary:** Applies default settings to Client

**Returns:** None (method modifies receiver)

**Complexity:**
- Time: O(1) (assuming fixed config operations)
- Space: O(1)

**Example:**
```go
client.ApplyDefaults()
```

**Edge Cases:**
- Overwriting existing non-default values
- Nil receiver panic


</details>

<details>
<summary><b><code>EnhanceDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Generates enhanced documentation for multiple inputs in specified languages

**Parameters:**
- `ctx` (context.Context): Context for request cancellation
- `inputs` ([]string): List of input strings to document
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of enhanced documentation objects and potential error

**Complexity:**
- Time: O(n) where n is number of inputs
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.EnhanceDocumentationBatch(ctx, []string{"func foo()"}, []string{"go"})
```

**Edge Cases:**
- Empty input list
- Mismatched input/language lengths
- Unsupported languages


</details>

<details>
<summary><b><code>EnhanceYAMLDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Generates enhanced YAML documentation for multiple inputs in specified languages

**Parameters:**
- `ctx` (context.Context): Context for request cancellation
- `inputs` ([]string): List of input strings to document
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of YAML documentation objects and potential error

**Complexity:**
- Time: O(n) where n is number of inputs
- Space: O(n) for output storage

**Example:**
```go
yamlDocs, err := client.EnhanceYAMLDocumentationBatch(ctx, []string{"key: value"}, []string{"yaml"})
```

**Edge Cases:**
- Invalid YAML inputs
- Empty input list
- Mismatched input/language lengths


</details>

<details>
<summary><b><code>groupByTokenCounts(inputs []string, tokenCounts []int, budget int)</code></b></summary>

**Summary:** Groups strings by token counts within budget

**Parameters:**
- `inputs` ([]string): Strings to group
- `tokenCounts` ([]int): Corresponding token counts
- `budget` (int): Max tokens per group

**Returns:** Slice of grouped input indices

**Complexity:**
- Time: O(n log n) for sorting + O(n) grouping
- Space: O(n) for output groups

**Example:**
```go
groups := groupByTokenCounts(texts, counts, 4000)
```

**Edge Cases:**
- Single item exceeds budget
- Empty inputs array
- Mismatched input/tokenCounts lengths


</details>

<details>
<summary><b><code>deduplicateInputs(inputs []string, languages []string)</code></b></summary>

**Summary:** Deduplicates inputs while preserving mapping to original indices

**Parameters:**
- `inputs` ([]string): Input strings to deduplicate
- `languages` ([]string): Corresponding languages for each input

**Returns:** Tuple containing semantic hashes, unique inputs/languages, and original index mapping

**Complexity:**
- Time: O(n) for n inputs (assuming hash computation is O(1))
- Space: O(n) for storing unique elements and mappings

**Example:**
```go
hashes, uniqIn, uniqLang, map := client.deduplicateInputs([]string{"a","a"}, []string{"go","go"})
```

**Edge Cases:**
- Empty input lists
- Mismatched input/language lengths
- Hash collisions


</details>

<details>
<summary><b><code>EnhanceGenericBatch(ctx context.Context, c *Client, inputs []string, languages []string, getCache func, setCache func, callBatch func)</code></b></summary>

**Summary:** Processes a batch of inputs with language-specific enhancements using caching and batch calls.

**Parameters:**
- `ctx` (context.Context): Context for request cancellation and timeouts
- `c` (*Client): Client instance for making requests
- `inputs` ([]string): List of input strings to process
- `languages` ([]string): List of target languages for enhancement
- `getCache` (func): Function to retrieve cached results
- `setCache` (func): Function to store results in cache
- `callBatch` (func): Function to make batch API calls

**Returns:** Slice of processed results (type T) and an error if any

**Complexity:**
- Time: O(n * m) where n is inputs count and m is languages count
- Space: O(n * m) for storing results

**Example:**
```go
results, err := EnhanceGenericBatch(ctx, client, []string{"text1"}, []string{"en"}, getCache, setCache, callBatch)
```

**Edge Cases:**
- Empty inputs or languages lists
- Cache functions returning errors
- Context cancellation during processing


</details>

<details>
<summary><b><code>jsonUnmarshalAdapter(data []byte, v *T)</code></b></summary>

**Summary:** Unmarshals JSON data into a target variable of generic type T.

**Parameters:**
- `data` ([]byte): JSON-encoded data to unmarshal
- `v` (*T): Pointer to target variable for unmarshaling

**Returns:** Error if unmarshaling fails

**Complexity:**
- Time: O(n) where n is length of input data
- Space: O(1) (excluding space for output structure)

**Example:**
```go
err := jsonUnmarshalAdapter([]byte('{"key":"value"}'), &result)
```

**Edge Cases:**
- Malformed JSON input
- Type mismatches between JSON and target structure
- Nil pointer passed as target


</details>

<details>
<summary><b><code>jsonMarshalIndentAdapter(v T)</code></b></summary>

**Summary:** Marshals a value into indented JSON format.

**Parameters:**
- `v` (T): Value to marshal (generic type)

**Returns:** Indented JSON as byte slice and error if marshaling fails

**Complexity:**
- Time: O(n) where n is size of input structure
- Space: O(n) for JSON output

**Example:**
```go
jsonData, err := jsonMarshalIndentAdapter(myStruct)
```

**Edge Cases:**
- Cyclic references in input structure
- Unsupported types in input
- Very large structures causing memory issues


</details>


---

## 📄 File: `logger.go`

> 📍 `ai\logger.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `StdLogger`

```go
type StdLogger struct {
	logger *log.Logger 
}
```

**Summary:** Struct wrapping a standard logger instance

**Parameters:**
- `logger` (*log.Logger): Underlying logger instance

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
sl := StdLogger{logger: log.New(os.Stdout, "", 0)}
```

**Edge Cases:**
- Nil logger pointer initialization
- Thread safety when shared across goroutines


---

## 🔧 Functions

<details>
<summary><b><code>NewStdLogger()</code></b></summary>

**Summary:** Constructor for StdLogger with default settings

**Returns:** Initialized *StdLogger instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
logger := NewStdLogger()
```

**Edge Cases:**
- Resource exhaustion during logger creation
- Default output destination may need configuration


</details>

<details>
<summary><b><code>Info(msg string, args )</code></b></summary>

**Summary:** Logs an informational message with variadic arguments

**Parameters:**
- `msg` (string): Base log message
- `args` (...interface{}): Optional format arguments

**Returns:** None (side-effect operation)

**Complexity:**
- Time: O(n) where n is message length
- Space: O(1)

**Example:**
```go
logger.Info("User logged in", "userID", 123)
```

**Edge Cases:**
- Nil receiver panic
- Format string/argument mismatch
- Concurrent write operations


</details>

<details>
<summary><b><code>Warn(msg string, args )</code></b></summary>

**Summary:** Logs a warning message with variable arguments

**Parameters:**
- `msg` (string): Warning message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Warn("Low disk space: %dGB remaining", 5)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>

<details>
<summary><b><code>Error(msg string, args )</code></b></summary>

**Summary:** Logs an error message with variable arguments

**Parameters:**
- `msg` (string): Error message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Error("Failed to open file: %v", err)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>


---

## 📄 File: `parser.go`

> 📍 `ai\parser.go`

## 📑 Contents

- [🔧 Functions (2)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>parseBatchResponse(response string, expectedCount int)</code></b></summary>

**Summary:** Parses batch response string into Documentation objects

**Parameters:**
- `response` (string): Raw batch response string
- `expectedCount` (int): Expected number of documents

**Returns:** Slice of parsed Documentation objects or error

**Complexity:**
- Time: O(n) where n is response length
- Space: O(n) for output slice

**Example:**
```go
docs, err := client.parseBatchResponse(jsonResponse, 5)
```

**Edge Cases:**
- Empty response string
- Malformed response data
- Actual count ≠ expectedCount


</details>

<details>
<summary><b><code>parseYAMLBatchResponse(response string, expectedCount int)</code></b></summary>

**Summary:** Parses YAML batch response into YAMLDocumentation objects

**Parameters:**
- `response` (string): Raw YAML response string
- `expectedCount` (int): Expected document count

**Returns:** Slice of YAMLDocumentation objects or error

**Complexity:**
- Time: O(n) where n is YAML size
- Space: O(n) for output slice

**Example:**
```go
yamlDocs, err := client.parseYAMLBatchResponse(yamlResponse, 3)
```

**Edge Cases:**
- Invalid YAML syntax
- Mismatched document count
- Empty YAML content


</details>


---

## 📄 File: `prompt.go`

> 📍 `ai\prompt.go`

## 📑 Contents

- [🔧 Functions (3)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>buildPrompt(input string, language string)</code></b></summary>

**Summary:** Builds a prompt string from input and language

**Parameters:**
- `input` (string): Code or text input
- `language` (string): Target programming language

**Returns:** Formatted prompt string

**Complexity:**
- Time: O(1)
- Space: O(n) where n is input length

**Example:**
```go
prompt := client.buildPrompt("func add() {}", "go")
```

**Edge Cases:**
- Empty input string
- Unsupported language


</details>

<details>
<summary><b><code>buildBatchPromptYamlDocumentation(prompts []string)</code></b></summary>

**Summary:** Builds batch prompt for YAML documentation generation

**Parameters:**
- `prompts` ([]string): Input prompt strings

**Returns:** Formatted batch prompt string

**Complexity:**
- Time: O(n) where n is total prompt length
- Space: O(n) for output string

**Example:**
```go
batchPrompt := client.buildBatchPromptYamlDocumentation([]string{"func1", "func2"})
```

**Edge Cases:**
- Empty prompts slice
- Very large prompt strings
- Special characters in prompts


</details>

<details>
<summary><b><code>buildBatchPromptCodeAssistant(prompts []string)</code></b></summary>

**Summary:** Builds a batch prompt string from an array of prompts

**Parameters:**
- `prompts` ([]string): Array of input prompts

**Returns:** Concatenated string of all prompts

**Complexity:**
- Time: O(n), where n is the total length of all prompts
- Space: O(n), where n is the total length of the output string

**Example:**
```go
batchPrompt := client.buildBatchPromptCodeAssistant([]string{"hello", "world"})
```

**Edge Cases:**
- Empty input array returns an empty string
- Large input array may cause memory issues


</details>


---

## 📄 File: `token.go`

> 📍 `ai\token.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

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
<summary><b><code>CountTokens(text string)</code></b></summary>

**Summary:** Counts tokens in a given text string

**Parameters:**
- `text` (string): Input text to tokenize and count

**Returns:** Number of tokens as int

**Complexity:**
- Time: O(n) where n is text length
- Space: O(1)

**Example:**
```go
count := CountTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string returns 0
- Unicode/whitespace handling may vary


</details>

<details>
<summary><b><code>cheapSkipFilter(input string)</code></b></summary>

**Summary:** Cheaply filters input strings based on skip conditions

**Parameters:**
- `input` (string): String to evaluate

**Returns:** Boolean indicating whether to skip

**Complexity:**
- Time: O(1) or O(n) depending on implementation
- Space: O(1)

**Example:**
```go
skip := cheapSkipFilter("example") // returns false
```

**Edge Cases:**
- Empty string handling
- Special characters in input


</details>

<details>
<summary><b><code>groupByTokenBudget(snippets []string, languages []string, tokenBudget int)</code></b></summary>

**Summary:** Groups code snippets by language within token budget

**Parameters:**
- `snippets` ([]string): Code snippets to group
- `languages` ([]string): Programming languages for snippets
- `tokenBudget` (int): Max tokens per group

**Returns:** 2D array of snippet indices grouped by constraints

**Complexity:**
- Time: O(n^2) in worst case
- Space: O(n) for output storage

**Example:**
```go
groups := groupByTokenBudget(snippets, langs, 1000)
```

**Edge Cases:**
- Empty input arrays
- Token budget smaller than any snippet
- Mismatched snippet/language array lengths


</details>


---

## 📄 File: `utils.go`

> 📍 `ai\utils.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>ExtractJSONArray(raw string)</code></b></summary>

**Summary:** Extracts JSON array from raw string input

**Parameters:**
- `raw` (string): Raw string containing JSON array

**Returns:** Extracted JSON array as string or error if parsing fails

**Complexity:**
- Time: O(n) where n is string length (JSON parsing)
- Space: O(n) for parsed JSON storage

**Example:**
```go
arr, err := ExtractJSONArray(`{"data": [1,2,3]}`) // returns `[1,2,3]`
```

**Edge Cases:**
- Malformed JSON input
- Non-array JSON structures
- Empty input string


</details>

<details>
<summary><b><code>ParseObjectKeysAsArray(raw string, expectedCount int)</code></b></summary>

**Summary:** Parses object keys into string array with size validation

**Parameters:**
- `raw` (string): JSON object string
- `expectedCount` (int): Required number of keys

**Returns:** Array of keys or error if count mismatches

**Complexity:**
- Time: O(n) for JSON parsing + O(k) for key extraction
- Space: O(k) for key storage (k = key count)

**Example:**
```go
keys, err := ParseObjectKeysAsArray(`{"a":1,"b":2}`, 2) // returns ["a","b"]
```

**Edge Cases:**
- Non-object JSON input
- Key count mismatch
- Duplicate keys in JSON


</details>

<details>
<summary><b><code>MapResults(cachedResults []string, reverseMap []int, inputLen int)</code></b></summary>

**Summary:** Maps cached results using reverse index mapping

**Parameters:**
- `cachedResults` ([]string): Precomputed result array
- `reverseMap` ([]int): Index mapping array
- `inputLen` (int): Expected output length

**Returns:** Remapped string array according to indices

**Complexity:**
- Time: O(m) where m is inputLen
- Space: O(m) for output array

**Example:**
```go
mapped := MapResults(["a","b"], [1,0], 2) // returns ["b","a"]
```

**Edge Cases:**
- Index out of bounds in reverseMap
- Mismatch between inputLen and reverseMap length
- Empty input arrays


</details>

<details>
<summary><b><code>EstimateTokens(input string)</code></b></summary>

**Summary:** Estimates token count from input string

**Parameters:**
- `input` (string): Text to analyze for token count

**Returns:** Estimated token count as integer

**Complexity:**
- Time: O(n) where n is input length
- Space: O(1)

**Example:**
```go
tokens := EstimateTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string input
- Unicode/emoji characters
- Very long strings (memory considerations)


</details>


---

## 📄 File: `batching.go`

> 📍 `ai\batching.go`

## 📑 Contents

- [🔧 Functions (6)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>processBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes batch of inputs with retry logic, returns documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): Input strings to process
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of Documentation objects or error if processing fails

**Complexity:**
- Time: O(n*m) where n=inputs, m=retries
- Space: O(n) for output docs storage

**Example:**
```go
docs, err := client.processBatchWithRetry(ctx, []string{"code1", "code2"}, []string{"go"})
```

**Edge Cases:**
- Empty inputs/languages slices
- Context cancellation mid-execution
- Partial failures in batch


</details>

<details>
<summary><b><code>processYAMLBatchWithRetry(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes YAML inputs with retries, returns YAML documentation

**Parameters:**
- `ctx` (context.Context): Execution context
- `inputs` ([]string): YAML content strings
- `languages` ([]string): Target documentation languages

**Returns:** Slice of YAMLDocumentation objects or processing error

**Complexity:**
- Time: O(n*m) where n=inputs, m=retries
- Space: O(n) for YAML output storage

**Example:**
```go
yamlDocs, err := client.processYAMLBatchWithRetry(ctx, []string{"yaml1", "yaml2"}, []string{"en"})
```

**Edge Cases:**
- Invalid YAML inputs
- Language-specific YAML parsing issues
- Network failures during retries


</details>

<details>
<summary><b><code>processBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes input batch without retries, returns documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation
- `inputs` ([]string): Input strings to document
- `languages` ([]string): Target output languages

**Returns:** Documentation slice or immediate error on failure

**Complexity:**
- Time: O(n) where n=inputs count
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processBatch(ctx, []string{"func() {}"}, []string{"python"})
```

**Edge Cases:**
- First-attempt failures terminate entire batch
- Mixed-language input/output mismatches
- Memory pressure with large batches


</details>

<details>
<summary><b><code>processYAMLBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Processes YAML batch inputs into documentation objects

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): List of input strings to process
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of YAMLDocumentation objects or error if processing fails

**Complexity:**
- Time: O(n) where n is total input size
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.processYAMLBatch(ctx, []string{"config.yml"}, []string{"en"})
```

**Edge Cases:**
- Empty inputs array
- Unsupported language codes
- Malformed YAML content


</details>

<details>
<summary><b><code>callBatchAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

**Summary:** Calls batch API to generate documentation for inputs

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): Input strings to document
- `languages` ([]string): Target documentation languages
- `indices` ([]int): Position tracking for input mapping

**Returns:** Slice of Documentation objects or error on API failure

**Complexity:**
- Time: O(n) for input processing + API call overhead
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.callBatchAPI(ctx, []string{"func() {}"}, []string{"go"}, []int{0})
```

**Edge Cases:**
- Mismatched input/language lengths
- Invalid indices
- API rate limiting


</details>

<details>
<summary><b><code>callBatchYamlAPI(ctx context.Context, inputs []string, languages []string, indices []int)</code></b></summary>

**Summary:** Batch processes YAML inputs via API into documentation

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `inputs` ([]string): YAML content strings
- `languages` ([]string): Target output languages
- `indices` ([]int): Input position tracking indices

**Returns:** YAMLDocumentation slice or API error

**Complexity:**
- Time: O(n) processing + O(k) API calls (k=batches)
- Space: O(n) for output + O(k) batch buffers

**Example:**
```go
yamlDocs, err := client.callBatchYamlAPI(ctx, []string{"key: value"}, []string{"en"}, []int{0})
```

**Edge Cases:**
- Invalid YAML syntax
- Partial batch failures
- Index out-of-bounds errors


</details>


---

## 📄 File: `client.go`

> 📍 `ai\client.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (9)](#-functions)

## 🧱 Structs

### `Client`

```go
type Client struct {
	backend ai.Backend 
	cache *aiCache.Cache 
	logger Logger 
	config ai.BackendConfig 
}
```

**Summary:** Client struct with backend, cache, logger, and config

**Returns:** N/A (type definition)

**Complexity:**
- Time: N/A
- Space: O(1) (fixed-size struct)

**Example:**
```go
client := Client{backend: myBackend, cache: myCache}
```

**Edge Cases:**
- Nil dependencies
- Uninitialized config causing runtime errors


---

## 🔧 Functions

<details>
<summary><b><code>NewClient(backend ai.Backend, cfg ai.BackendConfig)</code></b></summary>

**Summary:** Constructor for AI client with backend configuration

**Parameters:**
- `backend` (ai.Backend): Implementation interface
- `cfg` (ai.BackendConfig): Connection settings

**Returns:** Initialized *Client instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
client := NewClient(&GPTBackend{}, config)
```

**Edge Cases:**
- Nil backend implementation
- Invalid configuration


</details>

<details>
<summary><b><code>ApplyDefaults()</code></b></summary>

**Summary:** Applies default settings to Client

**Returns:** None (method modifies receiver)

**Complexity:**
- Time: O(1) (assuming fixed config operations)
- Space: O(1)

**Example:**
```go
client.ApplyDefaults()
```

**Edge Cases:**
- Overwriting existing non-default values
- Nil receiver panic


</details>

<details>
<summary><b><code>EnhanceDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Generates enhanced documentation for multiple inputs in specified languages

**Parameters:**
- `ctx` (context.Context): Context for request cancellation
- `inputs` ([]string): List of input strings to document
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of enhanced documentation objects and potential error

**Complexity:**
- Time: O(n) where n is number of inputs
- Space: O(n) for output storage

**Example:**
```go
docs, err := client.EnhanceDocumentationBatch(ctx, []string{"func foo()"}, []string{"go"})
```

**Edge Cases:**
- Empty input list
- Mismatched input/language lengths
- Unsupported languages


</details>

<details>
<summary><b><code>EnhanceYAMLDocumentationBatch(ctx context.Context, inputs []string, languages []string)</code></b></summary>

**Summary:** Generates enhanced YAML documentation for multiple inputs in specified languages

**Parameters:**
- `ctx` (context.Context): Context for request cancellation
- `inputs` ([]string): List of input strings to document
- `languages` ([]string): Target languages for documentation

**Returns:** Slice of YAML documentation objects and potential error

**Complexity:**
- Time: O(n) where n is number of inputs
- Space: O(n) for output storage

**Example:**
```go
yamlDocs, err := client.EnhanceYAMLDocumentationBatch(ctx, []string{"key: value"}, []string{"yaml"})
```

**Edge Cases:**
- Invalid YAML inputs
- Empty input list
- Mismatched input/language lengths


</details>

<details>
<summary><b><code>groupByTokenCounts(inputs []string, tokenCounts []int, budget int)</code></b></summary>

**Summary:** Groups strings by token counts within budget

**Parameters:**
- `inputs` ([]string): Strings to group
- `tokenCounts` ([]int): Corresponding token counts
- `budget` (int): Max tokens per group

**Returns:** Slice of grouped input indices

**Complexity:**
- Time: O(n log n) for sorting + O(n) grouping
- Space: O(n) for output groups

**Example:**
```go
groups := groupByTokenCounts(texts, counts, 4000)
```

**Edge Cases:**
- Single item exceeds budget
- Empty inputs array
- Mismatched input/tokenCounts lengths


</details>

<details>
<summary><b><code>deduplicateInputs(inputs []string, languages []string)</code></b></summary>

**Summary:** Deduplicates inputs while preserving mapping to original indices

**Parameters:**
- `inputs` ([]string): Input strings to deduplicate
- `languages` ([]string): Corresponding languages for each input

**Returns:** Tuple containing semantic hashes, unique inputs/languages, and original index mapping

**Complexity:**
- Time: O(n) for n inputs (assuming hash computation is O(1))
- Space: O(n) for storing unique elements and mappings

**Example:**
```go
hashes, uniqIn, uniqLang, map := client.deduplicateInputs([]string{"a","a"}, []string{"go","go"})
```

**Edge Cases:**
- Empty input lists
- Mismatched input/language lengths
- Hash collisions


</details>

<details>
<summary><b><code>EnhanceGenericBatch(ctx context.Context, c *Client, inputs []string, languages []string, getCache func, setCache func, callBatch func)</code></b></summary>

**Summary:** Processes a batch of inputs with language-specific enhancements using caching and batch calls.

**Parameters:**
- `ctx` (context.Context): Context for request cancellation and timeouts
- `c` (*Client): Client instance for making requests
- `inputs` ([]string): List of input strings to process
- `languages` ([]string): List of target languages for enhancement
- `getCache` (func): Function to retrieve cached results
- `setCache` (func): Function to store results in cache
- `callBatch` (func): Function to make batch API calls

**Returns:** Slice of processed results (type T) and an error if any

**Complexity:**
- Time: O(n * m) where n is inputs count and m is languages count
- Space: O(n * m) for storing results

**Example:**
```go
results, err := EnhanceGenericBatch(ctx, client, []string{"text1"}, []string{"en"}, getCache, setCache, callBatch)
```

**Edge Cases:**
- Empty inputs or languages lists
- Cache functions returning errors
- Context cancellation during processing


</details>

<details>
<summary><b><code>jsonUnmarshalAdapter(data []byte, v *T)</code></b></summary>

**Summary:** Unmarshals JSON data into a target variable of generic type T.

**Parameters:**
- `data` ([]byte): JSON-encoded data to unmarshal
- `v` (*T): Pointer to target variable for unmarshaling

**Returns:** Error if unmarshaling fails

**Complexity:**
- Time: O(n) where n is length of input data
- Space: O(1) (excluding space for output structure)

**Example:**
```go
err := jsonUnmarshalAdapter([]byte('{"key":"value"}'), &result)
```

**Edge Cases:**
- Malformed JSON input
- Type mismatches between JSON and target structure
- Nil pointer passed as target


</details>

<details>
<summary><b><code>jsonMarshalIndentAdapter(v T)</code></b></summary>

**Summary:** Marshals a value into indented JSON format.

**Parameters:**
- `v` (T): Value to marshal (generic type)

**Returns:** Indented JSON as byte slice and error if marshaling fails

**Complexity:**
- Time: O(n) where n is size of input structure
- Space: O(n) for JSON output

**Example:**
```go
jsonData, err := jsonMarshalIndentAdapter(myStruct)
```

**Edge Cases:**
- Cyclic references in input structure
- Unsupported types in input
- Very large structures causing memory issues


</details>


---

## 📄 File: `logger.go`

> 📍 `ai\logger.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `StdLogger`

```go
type StdLogger struct {
	logger *log.Logger 
}
```

**Summary:** Struct wrapping a standard logger instance

**Parameters:**
- `logger` (*log.Logger): Underlying logger instance

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
sl := StdLogger{logger: log.New(os.Stdout, "", 0)}
```

**Edge Cases:**
- Nil logger pointer initialization
- Thread safety when shared across goroutines


---

## 🔧 Functions

<details>
<summary><b><code>NewStdLogger()</code></b></summary>

**Summary:** Constructor for StdLogger with default settings

**Returns:** Initialized *StdLogger instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
logger := NewStdLogger()
```

**Edge Cases:**
- Resource exhaustion during logger creation
- Default output destination may need configuration


</details>

<details>
<summary><b><code>Info(msg string, args )</code></b></summary>

**Summary:** Logs an informational message with variadic arguments

**Parameters:**
- `msg` (string): Base log message
- `args` (...interface{}): Optional format arguments

**Returns:** None (side-effect operation)

**Complexity:**
- Time: O(n) where n is message length
- Space: O(1)

**Example:**
```go
logger.Info("User logged in", "userID", 123)
```

**Edge Cases:**
- Nil receiver panic
- Format string/argument mismatch
- Concurrent write operations


</details>

<details>
<summary><b><code>Warn(msg string, args )</code></b></summary>

**Summary:** Logs a warning message with variable arguments

**Parameters:**
- `msg` (string): Warning message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Warn("Low disk space: %dGB remaining", 5)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>

<details>
<summary><b><code>Error(msg string, args )</code></b></summary>

**Summary:** Logs an error message with variable arguments

**Parameters:**
- `msg` (string): Error message to log
- `args` (...interface{}): Variable arguments for message formatting

**Returns:** None (void function)

**Complexity:**
- Time: O(n) where n is the length of args (due to formatting)
- Space: O(1) (in-place logging, no additional storage)

**Example:**
```go
logger.Error("Failed to open file: %v", err)
```

**Edge Cases:**
- Empty msg string
- Nil args causing formatting issues


</details>


---

## 📄 File: `parser.go`

> 📍 `ai\parser.go`

## 📑 Contents

- [🔧 Functions (2)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>parseBatchResponse(response string, expectedCount int)</code></b></summary>

**Summary:** Parses batch response string into Documentation objects

**Parameters:**
- `response` (string): Raw batch response string
- `expectedCount` (int): Expected number of documents

**Returns:** Slice of parsed Documentation objects or error

**Complexity:**
- Time: O(n) where n is response length
- Space: O(n) for output slice

**Example:**
```go
docs, err := client.parseBatchResponse(jsonResponse, 5)
```

**Edge Cases:**
- Empty response string
- Malformed response data
- Actual count ≠ expectedCount


</details>

<details>
<summary><b><code>parseYAMLBatchResponse(response string, expectedCount int)</code></b></summary>

**Summary:** Parses YAML batch response into YAMLDocumentation objects

**Parameters:**
- `response` (string): Raw YAML response string
- `expectedCount` (int): Expected document count

**Returns:** Slice of YAMLDocumentation objects or error

**Complexity:**
- Time: O(n) where n is YAML size
- Space: O(n) for output slice

**Example:**
```go
yamlDocs, err := client.parseYAMLBatchResponse(yamlResponse, 3)
```

**Edge Cases:**
- Invalid YAML syntax
- Mismatched document count
- Empty YAML content


</details>


---

## 📄 File: `prompt.go`

> 📍 `ai\prompt.go`

## 📑 Contents

- [🔧 Functions (3)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>buildPrompt(input string, language string)</code></b></summary>

**Summary:** Builds a prompt string from input and language

**Parameters:**
- `input` (string): Code or text input
- `language` (string): Target programming language

**Returns:** Formatted prompt string

**Complexity:**
- Time: O(1)
- Space: O(n) where n is input length

**Example:**
```go
prompt := client.buildPrompt("func add() {}", "go")
```

**Edge Cases:**
- Empty input string
- Unsupported language


</details>

<details>
<summary><b><code>buildBatchPromptYamlDocumentation(prompts []string)</code></b></summary>

**Summary:** Builds batch prompt for YAML documentation generation

**Parameters:**
- `prompts` ([]string): Input prompt strings

**Returns:** Formatted batch prompt string

**Complexity:**
- Time: O(n) where n is total prompt length
- Space: O(n) for output string

**Example:**
```go
batchPrompt := client.buildBatchPromptYamlDocumentation([]string{"func1", "func2"})
```

**Edge Cases:**
- Empty prompts slice
- Very large prompt strings
- Special characters in prompts


</details>

<details>
<summary><b><code>buildBatchPromptCodeAssistant(prompts []string)</code></b></summary>

**Summary:** Builds a batch prompt string from an array of prompts

**Parameters:**
- `prompts` ([]string): Array of input prompts

**Returns:** Concatenated string of all prompts

**Complexity:**
- Time: O(n), where n is the total length of all prompts
- Space: O(n), where n is the total length of the output string

**Example:**
```go
batchPrompt := client.buildBatchPromptCodeAssistant([]string{"hello", "world"})
```

**Edge Cases:**
- Empty input array returns an empty string
- Large input array may cause memory issues


</details>


---

## 📄 File: `token.go`

> 📍 `ai\token.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

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
<summary><b><code>CountTokens(text string)</code></b></summary>

**Summary:** Counts tokens in a given text string

**Parameters:**
- `text` (string): Input text to tokenize and count

**Returns:** Number of tokens as int

**Complexity:**
- Time: O(n) where n is text length
- Space: O(1)

**Example:**
```go
count := CountTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string returns 0
- Unicode/whitespace handling may vary


</details>

<details>
<summary><b><code>cheapSkipFilter(input string)</code></b></summary>

**Summary:** Cheaply filters input strings based on skip conditions

**Parameters:**
- `input` (string): String to evaluate

**Returns:** Boolean indicating whether to skip

**Complexity:**
- Time: O(1) or O(n) depending on implementation
- Space: O(1)

**Example:**
```go
skip := cheapSkipFilter("example") // returns false
```

**Edge Cases:**
- Empty string handling
- Special characters in input


</details>

<details>
<summary><b><code>groupByTokenBudget(snippets []string, languages []string, tokenBudget int)</code></b></summary>

**Summary:** Groups code snippets by language within token budget

**Parameters:**
- `snippets` ([]string): Code snippets to group
- `languages` ([]string): Programming languages for snippets
- `tokenBudget` (int): Max tokens per group

**Returns:** 2D array of snippet indices grouped by constraints

**Complexity:**
- Time: O(n^2) in worst case
- Space: O(n) for output storage

**Example:**
```go
groups := groupByTokenBudget(snippets, langs, 1000)
```

**Edge Cases:**
- Empty input arrays
- Token budget smaller than any snippet
- Mismatched snippet/language array lengths


</details>


---

## 📄 File: `utils.go`

> 📍 `ai\utils.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>ExtractJSONArray(raw string)</code></b></summary>

**Summary:** Extracts JSON array from raw string input

**Parameters:**
- `raw` (string): Raw string containing JSON array

**Returns:** Extracted JSON array as string or error if parsing fails

**Complexity:**
- Time: O(n) where n is string length (JSON parsing)
- Space: O(n) for parsed JSON storage

**Example:**
```go
arr, err := ExtractJSONArray(`{"data": [1,2,3]}`) // returns `[1,2,3]`
```

**Edge Cases:**
- Malformed JSON input
- Non-array JSON structures
- Empty input string


</details>

<details>
<summary><b><code>ParseObjectKeysAsArray(raw string, expectedCount int)</code></b></summary>

**Summary:** Parses object keys into string array with size validation

**Parameters:**
- `raw` (string): JSON object string
- `expectedCount` (int): Required number of keys

**Returns:** Array of keys or error if count mismatches

**Complexity:**
- Time: O(n) for JSON parsing + O(k) for key extraction
- Space: O(k) for key storage (k = key count)

**Example:**
```go
keys, err := ParseObjectKeysAsArray(`{"a":1,"b":2}`, 2) // returns ["a","b"]
```

**Edge Cases:**
- Non-object JSON input
- Key count mismatch
- Duplicate keys in JSON


</details>

<details>
<summary><b><code>MapResults(cachedResults []string, reverseMap []int, inputLen int)</code></b></summary>

**Summary:** Maps cached results using reverse index mapping

**Parameters:**
- `cachedResults` ([]string): Precomputed result array
- `reverseMap` ([]int): Index mapping array
- `inputLen` (int): Expected output length

**Returns:** Remapped string array according to indices

**Complexity:**
- Time: O(m) where m is inputLen
- Space: O(m) for output array

**Example:**
```go
mapped := MapResults(["a","b"], [1,0], 2) // returns ["b","a"]
```

**Edge Cases:**
- Index out of bounds in reverseMap
- Mismatch between inputLen and reverseMap length
- Empty input arrays


</details>

<details>
<summary><b><code>EstimateTokens(input string)</code></b></summary>

**Summary:** Estimates token count from input string

**Parameters:**
- `input` (string): Text to analyze for token count

**Returns:** Estimated token count as integer

**Complexity:**
- Time: O(n) where n is input length
- Space: O(1)

**Example:**
```go
tokens := EstimateTokens("Hello world") // returns 2
```

**Edge Cases:**
- Empty string input
- Unicode/emoji characters
- Very long strings (memory considerations)


</details>

