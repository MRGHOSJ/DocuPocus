# 📦 Package: `ai`

> 📍 `C:\Users\DELL\Documents\GitHub\DocuPocus\internal\ai\utils.go`

[← Back to Overview](../README.md)

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

