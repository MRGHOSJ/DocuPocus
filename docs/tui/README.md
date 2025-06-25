# 📦 Package: `tui`

> 📍 `C:\Users\DELL\Documents\GitHub\DocuPocus\internal\tui\wizard.go`

[← Back to Overview](../README.md)

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (5)](#-functions)

## 🧱 Structs

### `model`

```go
type model struct {
	step int 
	projectDir string 
	template string 
	aiEnabled bool 
	inputs []string 
	cursor int 
}
```

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


---

## 🔧 Functions

<details>
<summary><b><code>initialModel()</code></b></summary>

**Summary:** Creates and returns an initial model instance

**Returns:** Initialized model instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
m := initialModel() // returns a new model
```

**Edge Cases:**
- None (pure initialization function)


</details>

<details>
<summary><b><code>Init()</code></b></summary>

**Summary:** Initializes the model and returns a command

**Returns:** tea.Cmd representing initial command

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
cmd := model.Init() // returns initial command
```

**Edge Cases:**
- None (depends on implementation)


</details>

<details>
<summary><b><code>Update(msg tea.Msg)</code></b></summary>

**Summary:** Updates model state based on message

**Parameters:**
- `msg` (tea.Msg): Message triggering update

**Returns:** Updated model and optional command

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
newModel, cmd := model.Update(msg) // processes message
```

**Edge Cases:**
- Unhandled message types
- Nil message input


</details>

<details>
<summary><b><code>View()</code></b></summary>

**Summary:** Returns a string representation of the model's view

**Returns:** String representing the model's view

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
view := model.View() // returns view as string
```

**Edge Cases:**
- Empty or nil model may return an empty string


</details>

<details>
<summary><b><code>StartWizard()</code></b></summary>

**Summary:** Initializes and starts a wizard process

**Returns:** Error if wizard fails to start, nil otherwise

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
err := StartWizard() // starts interactive wizard
```

**Edge Cases:**
- Dependencies not installed may cause errors
- User interrupts may return specific errors


</details>

