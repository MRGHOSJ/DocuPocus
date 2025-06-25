# 📦 Package: `backend`

> 📍 `C:\Users\DELL\Documents\GitHub\DocuPocus\internal\ai\backend\openrouter.go`

[← Back to Overview](../README.md)

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

