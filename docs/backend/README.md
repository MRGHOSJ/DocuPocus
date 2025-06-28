# 📦 Package: `backend`

[← Back to Overview](../README.md)

## 📄 File: `backend.go`

> 📍 `backend\backend.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)

## 🧱 Structs

### `BackendConfig`

```go
type BackendConfig struct {
	Model string 
	Endpoint string 
	APIKey string 
	RateLimit int 
	MaxRetries int 
	BatchSize int 
	TokenBudget int 
	Prompt string 
}
```

**Summary:** Configuration struct for AI backend settings

**Parameters:**
- `Model` (string): AI model identifier
- `Endpoint` (string): API endpoint URL
- `APIKey` (string): Authentication key
- `RateLimit` (int): Requests per minute limit
- `MaxRetries` (int): Maximum retry attempts
- `BatchSize` (int): Processing batch size
- `TokenBudget` (int): Maximum token allowance
- `Prompt` (string): Default prompt template

**Returns:** N/A (configuration struct)

**Complexity:**
- Time: N/A
- Space: O(1)

**Example:**
```go
config := BackendConfig{Model: "gpt-4", Endpoint: "https://api.openai.com"}
```

**Edge Cases:**
- Empty required fields
- Invalid rate limit values
- Exceeding token budget


---


---

## 📄 File: `ollama.go`

> 📍 `backend\ollama.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (3)](#-functions)

## 🧱 Structs

### `OllamaBackend`

```go
type OllamaBackend struct {
	client *api.Client 
	config BackendConfig 
}
```

**Summary:** Ollama backend implementation struct

**Parameters:**
- `client` (*api.Client): API client instance
- `config` (BackendConfig): Backend configuration

**Returns:** N/A (implementation struct)

**Complexity:**
- Time: N/A
- Space: O(1)

**Example:**
```go
backend := OllamaBackend{client: apiClient, config: cfg}
```

**Edge Cases:**
- Nil client pointer
- Invalid configuration
- Authentication failures


---

## 🔧 Functions

<details>
<summary><b><code>NewOllamaBackend(cfg BackendConfig)</code></b></summary>

**Summary:** Creates a new OllamaBackend instance with given configuration

**Parameters:**
- `cfg` (BackendConfig): Configuration for the backend

**Returns:** Pointer to OllamaBackend instance and error if initialization fails

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend, err := NewOllamaBackend(config)
```

**Edge Cases:**
- Invalid configuration leading to initialization failure
- Nil configuration input


</details>

<details>
<summary><b><code>Name()</code></b></summary>

**Summary:** Returns the name of the OllamaBackend instance

**Returns:** Name of the backend as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
name := backend.Name()
```

**Edge Cases:**
- Called on nil receiver
- Empty name string returned


</details>

<details>
<summary><b><code>Call(ctx context.Context, prompt string)</code></b></summary>

**Summary:** Executes a prompt call on the OllamaBackend

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `prompt` (string): Input prompt to process

**Returns:** Response string and error if call fails

**Complexity:**
- Time: O(n) where n is processing complexity
- Space: O(m) where m is response size

**Example:**
```go
response, err := backend.Call(ctx, "Hello world")
```

**Edge Cases:**
- Context cancellation during execution
- Empty prompt input
- Network/timeout errors


</details>


---

## 📄 File: `openrouter.go`

> 📍 `backend\openrouter.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `OpenRouterBackend`

```go
type OpenRouterBackend struct {
	httpClient *http.Client 
	config BackendConfig 
	rateLimiter *rate.Limiter 
}
```

**Summary:** Struct defining OpenRouter backend components

**Parameters:**
- `httpClient` (*http.Client): HTTP client for requests
- `config` (BackendConfig): Backend configuration
- `rateLimiter` (*rate.Limiter): Request rate limiter

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend := OpenRouterBackend{httpClient: &http.Client{}, config: cfg, rateLimiter: limiter}
```

**Edge Cases:**
- Nil httpClient causing runtime errors
- Invalid rateLimiter configuration


---

## 🔧 Functions

<details>
<summary><b><code>NewOpenRouterBackend(cfg BackendConfig)</code></b></summary>

**Summary:** Constructor for OpenRouterBackend

**Parameters:**
- `cfg` (BackendConfig): Configuration for the backend

**Returns:** Initialized *OpenRouterBackend instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend := NewOpenRouterBackend(myConfig)
```

**Edge Cases:**
- Nil or invalid config input
- Failure in internal initialization


</details>

<details>
<summary><b><code>Name()</code></b></summary>

**Summary:** Returns backend's identifier name

**Returns:** Backend name as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
name := backend.Name()
```

**Edge Cases:**
- Nil receiver panic
- Empty string return


</details>

<details>
<summary><b><code>Call(ctx context.Context, prompt string)</code></b></summary>

**Summary:** Calls OpenRouter API with a prompt and returns response

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `prompt` (string): Input prompt for the API

**Returns:** API response as string or error if call fails

**Complexity:**
- Time: O(1) (API call dependent)
- Space: O(1) (excluding API response storage)

**Example:**
```go
response, err := backend.Call(ctx, "Hello world")
```

**Edge Cases:**
- Network failures
- Invalid prompt format
- Context cancellation


</details>

<details>
<summary><b><code>callAPI(ctx context.Context, body )</code></b></summary>

**Summary:** Makes HTTP request to OpenRouter API

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `body` (interface{}): Request payload

**Returns:** HTTP response pointer or error if request fails

**Complexity:**
- Time: O(1) (network dependent)
- Space: O(1) (excluding response storage)

**Example:**
```go
resp, err := backend.callAPI(ctx, requestBody)
```

**Edge Cases:**
- Invalid request body
- API rate limiting
- Network timeouts


</details>


---

## 📄 File: `backend.go`

> 📍 `backend\backend.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)

## 🧱 Structs

### `BackendConfig`

```go
type BackendConfig struct {
	Model string 
	Endpoint string 
	APIKey string 
	RateLimit int 
	MaxRetries int 
	BatchSize int 
	TokenBudget int 
	Prompt string 
}
```

_No documentation available._

---


---

## 📄 File: `ollama.go`

> 📍 `backend\ollama.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (3)](#-functions)

## 🧱 Structs

### `OllamaBackend`

```go
type OllamaBackend struct {
	client *api.Client 
	config BackendConfig 
}
```

_No documentation available._

---

## 🔧 Functions

<details>
<summary><b><code>NewOllamaBackend(cfg BackendConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Name()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Call(ctx context.Context, prompt string)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `openrouter.go`

> 📍 `backend\openrouter.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `OpenRouterBackend`

```go
type OpenRouterBackend struct {
	httpClient *http.Client 
	config BackendConfig 
	rateLimiter *rate.Limiter 
}
```

_No documentation available._

---

## 🔧 Functions

<details>
<summary><b><code>NewOpenRouterBackend(cfg BackendConfig)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Name()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Call(ctx context.Context, prompt string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>callAPI(ctx context.Context, body )</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `backend.go`

> 📍 `backend\backend.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)

## 🧱 Structs

### `BackendConfig`

```go
type BackendConfig struct {
	Model string 
	Endpoint string 
	APIKey string 
	RateLimit int 
	MaxRetries int 
	BatchSize int 
	TokenBudget int 
	Prompt string 
}
```

**Summary:** Configuration struct for AI backend settings

**Parameters:**
- `Model` (string): AI model identifier
- `Endpoint` (string): API endpoint URL
- `APIKey` (string): Authentication key
- `RateLimit` (int): Requests per minute limit
- `MaxRetries` (int): Maximum retry attempts
- `BatchSize` (int): Processing batch size
- `TokenBudget` (int): Maximum token allowance
- `Prompt` (string): Default prompt template

**Returns:** N/A (configuration struct)

**Complexity:**
- Time: N/A
- Space: O(1)

**Example:**
```go
config := BackendConfig{Model: "gpt-4", Endpoint: "https://api.openai.com"}
```

**Edge Cases:**
- Empty required fields
- Invalid rate limit values
- Exceeding token budget


---


---

## 📄 File: `ollama.go`

> 📍 `backend\ollama.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (3)](#-functions)

## 🧱 Structs

### `OllamaBackend`

```go
type OllamaBackend struct {
	client *api.Client 
	config BackendConfig 
}
```

**Summary:** Ollama backend implementation struct

**Parameters:**
- `client` (*api.Client): API client instance
- `config` (BackendConfig): Backend configuration

**Returns:** N/A (implementation struct)

**Complexity:**
- Time: N/A
- Space: O(1)

**Example:**
```go
backend := OllamaBackend{client: apiClient, config: cfg}
```

**Edge Cases:**
- Nil client pointer
- Invalid configuration
- Authentication failures


---

## 🔧 Functions

<details>
<summary><b><code>NewOllamaBackend(cfg BackendConfig)</code></b></summary>

**Summary:** Creates a new OllamaBackend instance with given configuration

**Parameters:**
- `cfg` (BackendConfig): Configuration for the backend

**Returns:** Pointer to OllamaBackend instance and error if initialization fails

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend, err := NewOllamaBackend(config)
```

**Edge Cases:**
- Invalid configuration leading to initialization failure
- Nil configuration input


</details>

<details>
<summary><b><code>Name()</code></b></summary>

**Summary:** Returns the name of the OllamaBackend instance

**Returns:** Name of the backend as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
name := backend.Name()
```

**Edge Cases:**
- Called on nil receiver
- Empty name string returned


</details>

<details>
<summary><b><code>Call(ctx context.Context, prompt string)</code></b></summary>

**Summary:** Executes a prompt call on the OllamaBackend

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `prompt` (string): Input prompt to process

**Returns:** Response string and error if call fails

**Complexity:**
- Time: O(n) where n is processing complexity
- Space: O(m) where m is response size

**Example:**
```go
response, err := backend.Call(ctx, "Hello world")
```

**Edge Cases:**
- Context cancellation during execution
- Empty prompt input
- Network/timeout errors


</details>


---

## 📄 File: `openrouter.go`

> 📍 `backend\openrouter.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `OpenRouterBackend`

```go
type OpenRouterBackend struct {
	httpClient *http.Client 
	config BackendConfig 
	rateLimiter *rate.Limiter 
}
```

**Summary:** Struct defining OpenRouter backend components

**Parameters:**
- `httpClient` (*http.Client): HTTP client for requests
- `config` (BackendConfig): Backend configuration
- `rateLimiter` (*rate.Limiter): Request rate limiter

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend := OpenRouterBackend{httpClient: &http.Client{}, config: cfg, rateLimiter: limiter}
```

**Edge Cases:**
- Nil httpClient causing runtime errors
- Invalid rateLimiter configuration


---

## 🔧 Functions

<details>
<summary><b><code>NewOpenRouterBackend(cfg BackendConfig)</code></b></summary>

**Summary:** Constructor for OpenRouterBackend

**Parameters:**
- `cfg` (BackendConfig): Configuration for the backend

**Returns:** Initialized *OpenRouterBackend instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend := NewOpenRouterBackend(myConfig)
```

**Edge Cases:**
- Nil or invalid config input
- Failure in internal initialization


</details>

<details>
<summary><b><code>Name()</code></b></summary>

**Summary:** Returns backend's identifier name

**Returns:** Backend name as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
name := backend.Name()
```

**Edge Cases:**
- Nil receiver panic
- Empty string return


</details>

<details>
<summary><b><code>Call(ctx context.Context, prompt string)</code></b></summary>

**Summary:** Calls OpenRouter API with a prompt and returns response

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `prompt` (string): Input prompt for the API

**Returns:** API response as string or error if call fails

**Complexity:**
- Time: O(1) (API call dependent)
- Space: O(1) (excluding API response storage)

**Example:**
```go
response, err := backend.Call(ctx, "Hello world")
```

**Edge Cases:**
- Network failures
- Invalid prompt format
- Context cancellation


</details>

<details>
<summary><b><code>callAPI(ctx context.Context, body )</code></b></summary>

**Summary:** Makes HTTP request to OpenRouter API

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `body` (interface{}): Request payload

**Returns:** HTTP response pointer or error if request fails

**Complexity:**
- Time: O(1) (network dependent)
- Space: O(1) (excluding response storage)

**Example:**
```go
resp, err := backend.callAPI(ctx, requestBody)
```

**Edge Cases:**
- Invalid request body
- API rate limiting
- Network timeouts


</details>


---

## 📄 File: `backend.go`

> 📍 `backend\backend.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)

## 🧱 Structs

### `BackendConfig`

```go
type BackendConfig struct {
	Model string 
	Endpoint string 
	APIKey string 
	RateLimit int 
	MaxRetries int 
	BatchSize int 
	TokenBudget int 
	Prompt string 
}
```

**Summary:** Configuration struct for AI backend settings

**Parameters:**
- `Model` (string): AI model identifier
- `Endpoint` (string): API endpoint URL
- `APIKey` (string): Authentication key
- `RateLimit` (int): Requests per minute limit
- `MaxRetries` (int): Maximum retry attempts
- `BatchSize` (int): Processing batch size
- `TokenBudget` (int): Maximum token allowance
- `Prompt` (string): Default prompt template

**Returns:** N/A (configuration struct)

**Complexity:**
- Time: N/A
- Space: O(1)

**Example:**
```go
config := BackendConfig{Model: "gpt-4", Endpoint: "https://api.openai.com"}
```

**Edge Cases:**
- Empty required fields
- Invalid rate limit values
- Exceeding token budget


---


---

## 📄 File: `ollama.go`

> 📍 `backend\ollama.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (3)](#-functions)

## 🧱 Structs

### `OllamaBackend`

```go
type OllamaBackend struct {
	client *api.Client 
	config BackendConfig 
}
```

**Summary:** Ollama backend implementation struct

**Parameters:**
- `client` (*api.Client): API client instance
- `config` (BackendConfig): Backend configuration

**Returns:** N/A (implementation struct)

**Complexity:**
- Time: N/A
- Space: O(1)

**Example:**
```go
backend := OllamaBackend{client: apiClient, config: cfg}
```

**Edge Cases:**
- Nil client pointer
- Invalid configuration
- Authentication failures


---

## 🔧 Functions

<details>
<summary><b><code>NewOllamaBackend(cfg BackendConfig)</code></b></summary>

**Summary:** Creates a new OllamaBackend instance with given configuration

**Parameters:**
- `cfg` (BackendConfig): Configuration for the backend

**Returns:** Pointer to OllamaBackend instance and error if initialization fails

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend, err := NewOllamaBackend(config)
```

**Edge Cases:**
- Invalid configuration leading to initialization failure
- Nil configuration input


</details>

<details>
<summary><b><code>Name()</code></b></summary>

**Summary:** Returns the name of the OllamaBackend instance

**Returns:** Name of the backend as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
name := backend.Name()
```

**Edge Cases:**
- Called on nil receiver
- Empty name string returned


</details>

<details>
<summary><b><code>Call(ctx context.Context, prompt string)</code></b></summary>

**Summary:** Executes a prompt call on the OllamaBackend

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `prompt` (string): Input prompt to process

**Returns:** Response string and error if call fails

**Complexity:**
- Time: O(n) where n is processing complexity
- Space: O(m) where m is response size

**Example:**
```go
response, err := backend.Call(ctx, "Hello world")
```

**Edge Cases:**
- Context cancellation during execution
- Empty prompt input
- Network/timeout errors


</details>


---

## 📄 File: `openrouter.go`

> 📍 `backend\openrouter.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `OpenRouterBackend`

```go
type OpenRouterBackend struct {
	httpClient *http.Client 
	config BackendConfig 
	rateLimiter *rate.Limiter 
}
```

**Summary:** Struct defining OpenRouter backend components

**Parameters:**
- `httpClient` (*http.Client): HTTP client for requests
- `config` (BackendConfig): Backend configuration
- `rateLimiter` (*rate.Limiter): Request rate limiter

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend := OpenRouterBackend{httpClient: &http.Client{}, config: cfg, rateLimiter: limiter}
```

**Edge Cases:**
- Nil httpClient causing runtime errors
- Invalid rateLimiter configuration


---

## 🔧 Functions

<details>
<summary><b><code>NewOpenRouterBackend(cfg BackendConfig)</code></b></summary>

**Summary:** Constructor for OpenRouterBackend

**Parameters:**
- `cfg` (BackendConfig): Configuration for the backend

**Returns:** Initialized *OpenRouterBackend instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend := NewOpenRouterBackend(myConfig)
```

**Edge Cases:**
- Nil or invalid config input
- Failure in internal initialization


</details>

<details>
<summary><b><code>Name()</code></b></summary>

**Summary:** Returns backend's identifier name

**Returns:** Backend name as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
name := backend.Name()
```

**Edge Cases:**
- Nil receiver panic
- Empty string return


</details>

<details>
<summary><b><code>Call(ctx context.Context, prompt string)</code></b></summary>

**Summary:** Calls OpenRouter API with a prompt and returns response

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `prompt` (string): Input prompt for the API

**Returns:** API response as string or error if call fails

**Complexity:**
- Time: O(1) (API call dependent)
- Space: O(1) (excluding API response storage)

**Example:**
```go
response, err := backend.Call(ctx, "Hello world")
```

**Edge Cases:**
- Network failures
- Invalid prompt format
- Context cancellation


</details>

<details>
<summary><b><code>callAPI(ctx context.Context, body )</code></b></summary>

**Summary:** Makes HTTP request to OpenRouter API

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `body` (interface{}): Request payload

**Returns:** HTTP response pointer or error if request fails

**Complexity:**
- Time: O(1) (network dependent)
- Space: O(1) (excluding response storage)

**Example:**
```go
resp, err := backend.callAPI(ctx, requestBody)
```

**Edge Cases:**
- Invalid request body
- API rate limiting
- Network timeouts


</details>


---

## 📄 File: `backend.go`

> 📍 `backend\backend.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)

## 🧱 Structs

### `BackendConfig`

```go
type BackendConfig struct {
	Model string 
	Endpoint string 
	APIKey string 
	RateLimit int 
	MaxRetries int 
	BatchSize int 
	TokenBudget int 
	Prompt string 
}
```

**Summary:** Configuration struct for AI backend settings

**Parameters:**
- `Model` (string): AI model identifier
- `Endpoint` (string): API endpoint URL
- `APIKey` (string): Authentication key
- `RateLimit` (int): Requests per minute limit
- `MaxRetries` (int): Maximum retry attempts
- `BatchSize` (int): Processing batch size
- `TokenBudget` (int): Maximum token allowance
- `Prompt` (string): Default prompt template

**Returns:** N/A (configuration struct)

**Complexity:**
- Time: N/A
- Space: O(1)

**Example:**
```go
config := BackendConfig{Model: "gpt-4", Endpoint: "https://api.openai.com"}
```

**Edge Cases:**
- Empty required fields
- Invalid rate limit values
- Exceeding token budget


---


---

## 📄 File: `ollama.go`

> 📍 `backend\ollama.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (3)](#-functions)

## 🧱 Structs

### `OllamaBackend`

```go
type OllamaBackend struct {
	client *api.Client 
	config BackendConfig 
}
```

**Summary:** Ollama backend implementation struct

**Parameters:**
- `client` (*api.Client): API client instance
- `config` (BackendConfig): Backend configuration

**Returns:** N/A (implementation struct)

**Complexity:**
- Time: N/A
- Space: O(1)

**Example:**
```go
backend := OllamaBackend{client: apiClient, config: cfg}
```

**Edge Cases:**
- Nil client pointer
- Invalid configuration
- Authentication failures


---

## 🔧 Functions

<details>
<summary><b><code>NewOllamaBackend(cfg BackendConfig)</code></b></summary>

**Summary:** Creates a new OllamaBackend instance with given configuration

**Parameters:**
- `cfg` (BackendConfig): Configuration for the backend

**Returns:** Pointer to OllamaBackend instance and error if initialization fails

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend, err := NewOllamaBackend(config)
```

**Edge Cases:**
- Invalid configuration leading to initialization failure
- Nil configuration input


</details>

<details>
<summary><b><code>Name()</code></b></summary>

**Summary:** Returns the name of the OllamaBackend instance

**Returns:** Name of the backend as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
name := backend.Name()
```

**Edge Cases:**
- Called on nil receiver
- Empty name string returned


</details>

<details>
<summary><b><code>Call(ctx context.Context, prompt string)</code></b></summary>

**Summary:** Executes a prompt call on the OllamaBackend

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `prompt` (string): Input prompt to process

**Returns:** Response string and error if call fails

**Complexity:**
- Time: O(n) where n is processing complexity
- Space: O(m) where m is response size

**Example:**
```go
response, err := backend.Call(ctx, "Hello world")
```

**Edge Cases:**
- Context cancellation during execution
- Empty prompt input
- Network/timeout errors


</details>


---

## 📄 File: `openrouter.go`

> 📍 `backend\openrouter.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (4)](#-functions)

## 🧱 Structs

### `OpenRouterBackend`

```go
type OpenRouterBackend struct {
	httpClient *http.Client 
	config BackendConfig 
	rateLimiter *rate.Limiter 
}
```

**Summary:** Struct defining OpenRouter backend components

**Parameters:**
- `httpClient` (*http.Client): HTTP client for requests
- `config` (BackendConfig): Backend configuration
- `rateLimiter` (*rate.Limiter): Request rate limiter

**Returns:** None (struct definition)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend := OpenRouterBackend{httpClient: &http.Client{}, config: cfg, rateLimiter: limiter}
```

**Edge Cases:**
- Nil httpClient causing runtime errors
- Invalid rateLimiter configuration


---

## 🔧 Functions

<details>
<summary><b><code>NewOpenRouterBackend(cfg BackendConfig)</code></b></summary>

**Summary:** Constructor for OpenRouterBackend

**Parameters:**
- `cfg` (BackendConfig): Configuration for the backend

**Returns:** Initialized *OpenRouterBackend instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
backend := NewOpenRouterBackend(myConfig)
```

**Edge Cases:**
- Nil or invalid config input
- Failure in internal initialization


</details>

<details>
<summary><b><code>Name()</code></b></summary>

**Summary:** Returns backend's identifier name

**Returns:** Backend name as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
name := backend.Name()
```

**Edge Cases:**
- Nil receiver panic
- Empty string return


</details>

<details>
<summary><b><code>Call(ctx context.Context, prompt string)</code></b></summary>

**Summary:** Calls OpenRouter API with a prompt and returns response

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `prompt` (string): Input prompt for the API

**Returns:** API response as string or error if call fails

**Complexity:**
- Time: O(1) (API call dependent)
- Space: O(1) (excluding API response storage)

**Example:**
```go
response, err := backend.Call(ctx, "Hello world")
```

**Edge Cases:**
- Network failures
- Invalid prompt format
- Context cancellation


</details>

<details>
<summary><b><code>callAPI(ctx context.Context, body )</code></b></summary>

**Summary:** Makes HTTP request to OpenRouter API

**Parameters:**
- `ctx` (context.Context): Context for cancellation/timeout
- `body` (interface{}): Request payload

**Returns:** HTTP response pointer or error if request fails

**Complexity:**
- Time: O(1) (network dependent)
- Space: O(1) (excluding response storage)

**Example:**
```go
resp, err := backend.callAPI(ctx, requestBody)
```

**Edge Cases:**
- Invalid request body
- API rate limiting
- Network timeouts


</details>

