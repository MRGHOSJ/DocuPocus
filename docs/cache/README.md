# 📦 Package: `cache`

[← Back to Overview](../README.md)

## 📄 File: `cache.go`

> 📍 `cache\cache.go`

## 📑 Contents

- [🧱 Structs (2)](#-structs)
- [🔧 Functions (7)](#-functions)

## 🧱 Structs

### `CacheKey`

```go
type CacheKey struct {
	Hash SemanticHash 
	Language string 
}
```

_No documentation available._

---

### `Cache`

```go
type Cache struct {
	dir string 
}
```

_No documentation available._

---

## 🔧 Functions

<details>
<summary><b><code>NewCache(dir string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>GenerateSemanticHash(input string)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>filename(key CacheKey)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Get(key CacheKey)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Set(key CacheKey, data []byte)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>GetDoc(c *Cache, key CacheKey, unmarshal func)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>SetDoc(c *Cache, key CacheKey, doc T, marshal func)</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `cache.go`

> 📍 `cache\cache.go`

## 📑 Contents

- [🧱 Structs (2)](#-structs)
- [🔧 Functions (7)](#-functions)

## 🧱 Structs

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
<summary><b><code>GenerateSemanticHash(input string)</code></b></summary>

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

**Summary:** Retrieves cached data by key

**Parameters:**
- `key` (CacheKey): Key to lookup in cache

**Returns:** Cached data as byte slice and existence flag

**Complexity:**
- Time: O(1) average, O(n) worst-case (hash collision)
- Space: O(1)

**Example:**
```go
data, exists := cache.Get("user:123")
```

**Edge Cases:**
- Key not found (returns false)
- Nil or invalid key input
- Expired cache entry


</details>

<details>
<summary><b><code>Set(key CacheKey, data []byte)</code></b></summary>

**Summary:** Stores data in cache with key

**Parameters:**
- `key` (CacheKey): Key for cache storage
- `data` ([]byte): Data to store

**Returns:** Error if operation fails

**Complexity:**
- Time: O(1) average, O(n) worst-case (hash collision)
- Space: O(n) where n is data size

**Example:**
```go
err := cache.Set("user:123", userData)
```

**Edge Cases:**
- Nil data input
- Key collision
- Cache size limits exceeded


</details>

<details>
<summary><b><code>GetDoc(c *Cache, key CacheKey, unmarshal func)</code></b></summary>

**Summary:** Retrieves and unmarshals cached document

**Parameters:**
- `c` (*Cache): Cache instance
- `key` (CacheKey): Document key
- `unmarshal` (func): Unmarshaling function

**Returns:** Unmarshaled document and existence flag

**Complexity:**
- Time: O(1) cache access + unmarshal complexity
- Space: O(n) where n is document size

**Example:**
```go
user, exists := GetDoc(cache, "user:123", json.Unmarshal)
```

**Edge Cases:**
- Unmarshal function failure
- Invalid cached data format
- Key not found


</details>

<details>
<summary><b><code>SetDoc(c *Cache, key CacheKey, doc T, marshal func)</code></b></summary>

**Summary:** Stores a document in cache with serialization

**Parameters:**
- `c` (*Cache): Cache instance
- `key` (CacheKey): Key for cache entry
- `doc` (T): Document to store
- `marshal` (func): Serialization function

**Returns:** Error if operation fails, nil otherwise

**Complexity:**
- Time: O(1) (assuming cache insertion is constant time)
- Space: O(1) (stores one document)

**Example:**
```go
err := SetDoc(cache, "user:123", userDoc, json.Marshal)
```

**Edge Cases:**
- Nil cache instance
- Serialization failure
- Duplicate key overwrite


</details>


---

## 📄 File: `cache.go`

> 📍 `cache\cache.go`

## 📑 Contents

- [🧱 Structs (2)](#-structs)
- [🔧 Functions (7)](#-functions)

## 🧱 Structs

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
<summary><b><code>GenerateSemanticHash(input string)</code></b></summary>

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

**Summary:** Retrieves cached data by key

**Parameters:**
- `key` (CacheKey): Key to lookup in cache

**Returns:** Cached data as byte slice and existence flag

**Complexity:**
- Time: O(1) average, O(n) worst-case (hash collision)
- Space: O(1)

**Example:**
```go
data, exists := cache.Get("user:123")
```

**Edge Cases:**
- Key not found (returns false)
- Nil or invalid key input
- Expired cache entry


</details>

<details>
<summary><b><code>Set(key CacheKey, data []byte)</code></b></summary>

**Summary:** Stores data in cache with key

**Parameters:**
- `key` (CacheKey): Key for cache storage
- `data` ([]byte): Data to store

**Returns:** Error if operation fails

**Complexity:**
- Time: O(1) average, O(n) worst-case (hash collision)
- Space: O(n) where n is data size

**Example:**
```go
err := cache.Set("user:123", userData)
```

**Edge Cases:**
- Nil data input
- Key collision
- Cache size limits exceeded


</details>

<details>
<summary><b><code>GetDoc(c *Cache, key CacheKey, unmarshal func)</code></b></summary>

**Summary:** Retrieves and unmarshals cached document

**Parameters:**
- `c` (*Cache): Cache instance
- `key` (CacheKey): Document key
- `unmarshal` (func): Unmarshaling function

**Returns:** Unmarshaled document and existence flag

**Complexity:**
- Time: O(1) cache access + unmarshal complexity
- Space: O(n) where n is document size

**Example:**
```go
user, exists := GetDoc(cache, "user:123", json.Unmarshal)
```

**Edge Cases:**
- Unmarshal function failure
- Invalid cached data format
- Key not found


</details>

<details>
<summary><b><code>SetDoc(c *Cache, key CacheKey, doc T, marshal func)</code></b></summary>

**Summary:** Stores a document in cache with serialization

**Parameters:**
- `c` (*Cache): Cache instance
- `key` (CacheKey): Key for cache entry
- `doc` (T): Document to store
- `marshal` (func): Serialization function

**Returns:** Error if operation fails, nil otherwise

**Complexity:**
- Time: O(1) (assuming cache insertion is constant time)
- Space: O(1) (stores one document)

**Example:**
```go
err := SetDoc(cache, "user:123", userDoc, json.Marshal)
```

**Edge Cases:**
- Nil cache instance
- Serialization failure
- Duplicate key overwrite


</details>


---

## 📄 File: `cache.go`

> 📍 `cache\cache.go`

## 📑 Contents

- [🧱 Structs (2)](#-structs)
- [🔧 Functions (7)](#-functions)

## 🧱 Structs

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
<summary><b><code>GenerateSemanticHash(input string)</code></b></summary>

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

**Summary:** Retrieves cached data by key

**Parameters:**
- `key` (CacheKey): Key to lookup in cache

**Returns:** Cached data as byte slice and existence flag

**Complexity:**
- Time: O(1) average, O(n) worst-case (hash collision)
- Space: O(1)

**Example:**
```go
data, exists := cache.Get("user:123")
```

**Edge Cases:**
- Key not found (returns false)
- Nil or invalid key input
- Expired cache entry


</details>

<details>
<summary><b><code>Set(key CacheKey, data []byte)</code></b></summary>

**Summary:** Stores data in cache with key

**Parameters:**
- `key` (CacheKey): Key for cache storage
- `data` ([]byte): Data to store

**Returns:** Error if operation fails

**Complexity:**
- Time: O(1) average, O(n) worst-case (hash collision)
- Space: O(n) where n is data size

**Example:**
```go
err := cache.Set("user:123", userData)
```

**Edge Cases:**
- Nil data input
- Key collision
- Cache size limits exceeded


</details>

<details>
<summary><b><code>GetDoc(c *Cache, key CacheKey, unmarshal func)</code></b></summary>

**Summary:** Retrieves and unmarshals cached document

**Parameters:**
- `c` (*Cache): Cache instance
- `key` (CacheKey): Document key
- `unmarshal` (func): Unmarshaling function

**Returns:** Unmarshaled document and existence flag

**Complexity:**
- Time: O(1) cache access + unmarshal complexity
- Space: O(n) where n is document size

**Example:**
```go
user, exists := GetDoc(cache, "user:123", json.Unmarshal)
```

**Edge Cases:**
- Unmarshal function failure
- Invalid cached data format
- Key not found


</details>

<details>
<summary><b><code>SetDoc(c *Cache, key CacheKey, doc T, marshal func)</code></b></summary>

**Summary:** Stores a document in cache with serialization

**Parameters:**
- `c` (*Cache): Cache instance
- `key` (CacheKey): Key for cache entry
- `doc` (T): Document to store
- `marshal` (func): Serialization function

**Returns:** Error if operation fails, nil otherwise

**Complexity:**
- Time: O(1) (assuming cache insertion is constant time)
- Space: O(1) (stores one document)

**Example:**
```go
err := SetDoc(cache, "user:123", userDoc, json.Marshal)
```

**Edge Cases:**
- Nil cache instance
- Serialization failure
- Duplicate key overwrite


</details>


---

## 📄 File: `cache.go`

> 📍 `cache\cache.go`

## 📑 Contents

- [🧱 Structs (2)](#-structs)
- [🔧 Functions (7)](#-functions)

## 🧱 Structs

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
<summary><b><code>GenerateSemanticHash(input string)</code></b></summary>

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

**Summary:** Retrieves cached data by key

**Parameters:**
- `key` (CacheKey): Key to lookup in cache

**Returns:** Cached data as byte slice and existence flag

**Complexity:**
- Time: O(1) average, O(n) worst-case (hash collision)
- Space: O(1)

**Example:**
```go
data, exists := cache.Get("user:123")
```

**Edge Cases:**
- Key not found (returns false)
- Nil or invalid key input
- Expired cache entry


</details>

<details>
<summary><b><code>Set(key CacheKey, data []byte)</code></b></summary>

**Summary:** Stores data in cache with key

**Parameters:**
- `key` (CacheKey): Key for cache storage
- `data` ([]byte): Data to store

**Returns:** Error if operation fails

**Complexity:**
- Time: O(1) average, O(n) worst-case (hash collision)
- Space: O(n) where n is data size

**Example:**
```go
err := cache.Set("user:123", userData)
```

**Edge Cases:**
- Nil data input
- Key collision
- Cache size limits exceeded


</details>

<details>
<summary><b><code>GetDoc(c *Cache, key CacheKey, unmarshal func)</code></b></summary>

**Summary:** Retrieves and unmarshals cached document

**Parameters:**
- `c` (*Cache): Cache instance
- `key` (CacheKey): Document key
- `unmarshal` (func): Unmarshaling function

**Returns:** Unmarshaled document and existence flag

**Complexity:**
- Time: O(1) cache access + unmarshal complexity
- Space: O(n) where n is document size

**Example:**
```go
user, exists := GetDoc(cache, "user:123", json.Unmarshal)
```

**Edge Cases:**
- Unmarshal function failure
- Invalid cached data format
- Key not found


</details>

<details>
<summary><b><code>SetDoc(c *Cache, key CacheKey, doc T, marshal func)</code></b></summary>

**Summary:** Stores a document in cache with serialization

**Parameters:**
- `c` (*Cache): Cache instance
- `key` (CacheKey): Key for cache entry
- `doc` (T): Document to store
- `marshal` (func): Serialization function

**Returns:** Error if operation fails, nil otherwise

**Complexity:**
- Time: O(1) (assuming cache insertion is constant time)
- Space: O(1) (stores one document)

**Example:**
```go
err := SetDoc(cache, "user:123", userDoc, json.Marshal)
```

**Edge Cases:**
- Nil cache instance
- Serialization failure
- Duplicate key overwrite


</details>


---

## 📄 File: `cache.go`

> 📍 `cache\cache.go`

## 📑 Contents

- [🧱 Structs (2)](#-structs)
- [🔧 Functions (7)](#-functions)

## 🧱 Structs

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
<summary><b><code>GenerateSemanticHash(input string)</code></b></summary>

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

**Summary:** Retrieves cached data by key

**Parameters:**
- `key` (CacheKey): Key to lookup in cache

**Returns:** Cached data as byte slice and existence flag

**Complexity:**
- Time: O(1) average, O(n) worst-case (hash collision)
- Space: O(1)

**Example:**
```go
data, exists := cache.Get("user:123")
```

**Edge Cases:**
- Key not found (returns false)
- Nil or invalid key input
- Expired cache entry


</details>

<details>
<summary><b><code>Set(key CacheKey, data []byte)</code></b></summary>

**Summary:** Stores data in cache with key

**Parameters:**
- `key` (CacheKey): Key for cache storage
- `data` ([]byte): Data to store

**Returns:** Error if operation fails

**Complexity:**
- Time: O(1) average, O(n) worst-case (hash collision)
- Space: O(n) where n is data size

**Example:**
```go
err := cache.Set("user:123", userData)
```

**Edge Cases:**
- Nil data input
- Key collision
- Cache size limits exceeded


</details>

<details>
<summary><b><code>GetDoc(c *Cache, key CacheKey, unmarshal func)</code></b></summary>

**Summary:** Retrieves and unmarshals cached document

**Parameters:**
- `c` (*Cache): Cache instance
- `key` (CacheKey): Document key
- `unmarshal` (func): Unmarshaling function

**Returns:** Unmarshaled document and existence flag

**Complexity:**
- Time: O(1) cache access + unmarshal complexity
- Space: O(n) where n is document size

**Example:**
```go
user, exists := GetDoc(cache, "user:123", json.Unmarshal)
```

**Edge Cases:**
- Unmarshal function failure
- Invalid cached data format
- Key not found


</details>

<details>
<summary><b><code>SetDoc(c *Cache, key CacheKey, doc T, marshal func)</code></b></summary>

**Summary:** Stores a document in cache with serialization

**Parameters:**
- `c` (*Cache): Cache instance
- `key` (CacheKey): Key for cache entry
- `doc` (T): Document to store
- `marshal` (func): Serialization function

**Returns:** Error if operation fails, nil otherwise

**Complexity:**
- Time: O(1) (assuming cache insertion is constant time)
- Space: O(1) (stores one document)

**Example:**
```go
err := SetDoc(cache, "user:123", userDoc, json.Marshal)
```

**Edge Cases:**
- Nil cache instance
- Serialization failure
- Duplicate key overwrite


</details>

