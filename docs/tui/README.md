# 📦 Package: `tui`

[← Back to Overview](../README.md)

## 📄 File: `wizard.go`

> 📍 `tui\wizard.go`

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


---

## 📄 File: `wizard.go`

> 📍 `tui\wizard.go`

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

_No documentation available._

---

## 🔧 Functions

<details>
<summary><b><code>initialModel()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Init()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>Update(msg tea.Msg)</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>View()</code></b></summary>

_No documentation available._

</details>

<details>
<summary><b><code>StartWizard()</code></b></summary>

_No documentation available._

</details>


---

## 📄 File: `wizard.go`

> 📍 `tui\wizard.go`

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


---

## 📄 File: `wizard.go`

> 📍 `tui\wizard.go`

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


---

## 📄 File: `model.go`

> 📍 `tui\model.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (1)](#-functions)

## 🧱 Structs

### `Model`

```go
type Model struct {
	step int 
	ProjectDir string 
	AiBackend string 
	AiModel string 
	AiEndpoint string 
	AiAPIKey string 
	OutputFolder string 
	verbose bool 
	inputValue string 
	choices []string 
	cursor int 
}
```

**Summary:** Processes batch AI requests using a client

**Parameters:**
- `requests` ([]AIRequest): Batch of requests to process
- `client` (*ai.Client): AI service client

**Returns:** None (handles requests asynchronously)

**Complexity:**
- Time: O(n) where n is number of requests
- Space: O(n) for request processing buffers

**Example:**
```go
processAIRequests(reqBatch, aiClient)
```

**Edge Cases:**
- Empty requests slice
- Nil client pointer
- Network timeouts


---

## 🔧 Functions

<details>
<summary><b><code>initialModel()</code></b></summary>

**Summary:** Initializes and returns a default Model instance.

**Returns:** Initialized Model instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model := initialModel()
```

**Edge Cases:**
- Model initialization failure
- Uninitialized fields in returned Model


</details>


---

## 📄 File: `steps.go`

> 📍 `tui\steps.go`

## 📑 Contents

- [🔧 Functions (13)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>handleTextInput(m Model, msg tea.KeyMsg)</code></b></summary>

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
<summary><b><code>stepProjectDir(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Updates Model state based on key message and returns updated Model with command.

**Parameters:**
- `m` (Model): Current model state
- `msg` (tea.KeyMsg): Key press message

**Returns:** Updated Model and optional tea.Cmd

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
newModel, cmd := stepProjectDir(currentModel, keyMsg)
```

**Edge Cases:**
- Nil or invalid key message
- Model state corruption during update


</details>

<details>
<summary><b><code>stepAIBackend(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Updates AI backend model state based on keyboard input

**Parameters:**
- `m` (Model): Current application model state
- `msg` (tea.KeyMsg): Keyboard input message

**Returns:** Updated Model and optional command (tea.Cmd)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, cmd := stepAIBackend(currentModel, keyMsg)
```

**Edge Cases:**
- Nil model input
- Unhandled key messages
- State transition errors


</details>

<details>
<summary><b><code>stepAIModel(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Processes keyboard input for AI model interaction

**Parameters:**
- `m` (Model): Current model state
- `msg` (tea.KeyMsg): Keyboard event

**Returns:** Modified Model and possible command

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
newModel, cmd := stepAIModel(appModel, tea.KeyEnter)
```

**Edge Cases:**
- Invalid model state
- Unsupported key combinations
- Command generation failure


</details>

<details>
<summary><b><code>stepAIEndpoint(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Handles keyboard events for AI endpoint configuration

**Parameters:**
- `m` (Model): Application state container
- `msg` (tea.KeyMsg): Key press event

**Returns:** Updated state model and optional command

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, _ := stepAIEndpoint(configModel, tea.KeyTab)
```

**Edge Cases:**
- Missing required model fields
- Edge key cases (Esc, Ctrl+C)
- State validation errors


</details>

<details>
<summary><b><code>stepAPIKey(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Handles API key input updates in a TUI model

**Parameters:**
- `m` (Model): Current application model state
- `msg` (tea.KeyMsg): Key press event to process

**Returns:** Updated Model and optional tea.Cmd for side effects

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, cmd := stepAPIKey(currentModel, keyEvent)
```

**Edge Cases:**
- Empty API key input
- Special key combinations


</details>

<details>
<summary><b><code>stepOutputFolder(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Processes output folder selection key events

**Parameters:**
- `m` (Model): Current application model state
- `msg` (tea.KeyMsg): Key press event to process

**Returns:** Updated Model and optional tea.Cmd for side effects

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, cmd := stepOutputFolder(currentModel, keyEvent)
```

**Edge Cases:**
- Invalid folder paths
- Permission denied errors


</details>

<details>
<summary><b><code>viewProjectDir(m Model)</code></b></summary>

**Summary:** Renders project directory view in TUI

**Parameters:**
- `m` (Model): Current application model state

**Returns:** String representation of project directory view

**Complexity:**
- Time: O(n) where n is directory contents
- Space: O(n) for string buffer

**Example:**
```go
dirView := viewProjectDir(currentModel)
```

**Edge Cases:**
- Empty directory
- Unreadable files


</details>

<details>
<summary><b><code>viewAIBackend(m Model)</code></b></summary>

**Summary:** Generates a string representation of an AI backend model

**Parameters:**
- `m` (Model): The AI backend model to be represented

**Returns:** String representation of the AI backend model

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
repr := viewAIBackend(myModel) // returns a string describing the backend
```

**Edge Cases:**
- Nil or uninitialized Model input
- Model with missing or invalid fields


</details>

<details>
<summary><b><code>viewAIModel(m Model)</code></b></summary>

**Summary:** Generates a string representation of an AI model

**Parameters:**
- `m` (Model): The AI model to be represented

**Returns:** String representation of the AI model

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
repr := viewAIModel(myModel) // returns a string describing the model
```

**Edge Cases:**
- Nil or uninitialized Model input
- Model with missing or invalid fields


</details>

<details>
<summary><b><code>viewAIEndpoint(m Model)</code></b></summary>

**Summary:** Generates a string representation of an AI endpoint

**Parameters:**
- `m` (Model): The AI endpoint model to be represented

**Returns:** String representation of the AI endpoint

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
repr := viewAIEndpoint(myModel) // returns a string describing the endpoint
```

**Edge Cases:**
- Nil or uninitialized Model input
- Model with missing or invalid fields


</details>

<details>
<summary><b><code>viewAPIKey(m Model)</code></b></summary>

**Summary:** Returns API key from Model as string

**Parameters:**
- `m` (Model): Model containing API key

**Returns:** API key as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
key := viewAPIKey(model) // returns "sk_123abc"
```

**Edge Cases:**
- Empty API key in Model
- Non-string type in Model's API key field


</details>

<details>
<summary><b><code>viewOutputFolder(m Model)</code></b></summary>

**Summary:** Returns output folder path from Model

**Parameters:**
- `m` (Model): Model containing output folder path

**Returns:** Output folder path as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
path := viewOutputFolder(model) // returns "/tmp/output"
```

**Edge Cases:**
- Empty path in Model
- Invalid path characters


</details>


---

## 📄 File: `styles.go`

> 📍 `tui\styles.go`

## 📑 Contents



---

## 📄 File: `wizard.go`

> 📍 `tui\wizard.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>Init()</code></b></summary>

**Summary:** Initializes Model and returns tea.Cmd

**Parameters:**
- `receiver` (Model): Model instance to initialize

**Returns:** tea.Cmd for BubbleTea framework

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
cmd := model.Init() // returns initial command
```

**Edge Cases:**
- Uninitialized Model fields
- Nil receiver


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


---

## 📄 File: `model.go`

> 📍 `tui\model.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (1)](#-functions)

## 🧱 Structs

### `Model`

```go
type Model struct {
	step int 
	ProjectDir string 
	AiBackend string 
	AiModel string 
	AiEndpoint string 
	AiAPIKey string 
	OutputFolder string 
	verbose bool 
	inputValue string 
	choices []string 
	cursor int 
}
```

**Summary:** Processes batch AI requests using a client

**Parameters:**
- `requests` ([]AIRequest): Batch of requests to process
- `client` (*ai.Client): AI service client

**Returns:** None (handles requests asynchronously)

**Complexity:**
- Time: O(n) where n is number of requests
- Space: O(n) for request processing buffers

**Example:**
```go
processAIRequests(reqBatch, aiClient)
```

**Edge Cases:**
- Empty requests slice
- Nil client pointer
- Network timeouts


---

## 🔧 Functions

<details>
<summary><b><code>initialModel()</code></b></summary>

**Summary:** Initializes and returns a default Model instance.

**Returns:** Initialized Model instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model := initialModel()
```

**Edge Cases:**
- Model initialization failure
- Uninitialized fields in returned Model


</details>


---

## 📄 File: `steps.go`

> 📍 `tui\steps.go`

## 📑 Contents

- [🔧 Functions (13)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>handleTextInput(m Model, msg tea.KeyMsg)</code></b></summary>

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
<summary><b><code>stepProjectDir(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Updates Model state based on key message and returns updated Model with command.

**Parameters:**
- `m` (Model): Current model state
- `msg` (tea.KeyMsg): Key press message

**Returns:** Updated Model and optional tea.Cmd

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
newModel, cmd := stepProjectDir(currentModel, keyMsg)
```

**Edge Cases:**
- Nil or invalid key message
- Model state corruption during update


</details>

<details>
<summary><b><code>stepAIBackend(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Updates AI backend model state based on keyboard input

**Parameters:**
- `m` (Model): Current application model state
- `msg` (tea.KeyMsg): Keyboard input message

**Returns:** Updated Model and optional command (tea.Cmd)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, cmd := stepAIBackend(currentModel, keyMsg)
```

**Edge Cases:**
- Nil model input
- Unhandled key messages
- State transition errors


</details>

<details>
<summary><b><code>stepAIModel(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Processes keyboard input for AI model interaction

**Parameters:**
- `m` (Model): Current model state
- `msg` (tea.KeyMsg): Keyboard event

**Returns:** Modified Model and possible command

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
newModel, cmd := stepAIModel(appModel, tea.KeyEnter)
```

**Edge Cases:**
- Invalid model state
- Unsupported key combinations
- Command generation failure


</details>

<details>
<summary><b><code>stepAIEndpoint(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Handles keyboard events for AI endpoint configuration

**Parameters:**
- `m` (Model): Application state container
- `msg` (tea.KeyMsg): Key press event

**Returns:** Updated state model and optional command

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, _ := stepAIEndpoint(configModel, tea.KeyTab)
```

**Edge Cases:**
- Missing required model fields
- Edge key cases (Esc, Ctrl+C)
- State validation errors


</details>

<details>
<summary><b><code>stepAPIKey(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Handles API key input updates in a TUI model

**Parameters:**
- `m` (Model): Current application model state
- `msg` (tea.KeyMsg): Key press event to process

**Returns:** Updated Model and optional tea.Cmd for side effects

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, cmd := stepAPIKey(currentModel, keyEvent)
```

**Edge Cases:**
- Empty API key input
- Special key combinations


</details>

<details>
<summary><b><code>stepOutputFolder(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Processes output folder selection key events

**Parameters:**
- `m` (Model): Current application model state
- `msg` (tea.KeyMsg): Key press event to process

**Returns:** Updated Model and optional tea.Cmd for side effects

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, cmd := stepOutputFolder(currentModel, keyEvent)
```

**Edge Cases:**
- Invalid folder paths
- Permission denied errors


</details>

<details>
<summary><b><code>viewProjectDir(m Model)</code></b></summary>

**Summary:** Renders project directory view in TUI

**Parameters:**
- `m` (Model): Current application model state

**Returns:** String representation of project directory view

**Complexity:**
- Time: O(n) where n is directory contents
- Space: O(n) for string buffer

**Example:**
```go
dirView := viewProjectDir(currentModel)
```

**Edge Cases:**
- Empty directory
- Unreadable files


</details>

<details>
<summary><b><code>viewAIBackend(m Model)</code></b></summary>

**Summary:** Generates a string representation of an AI backend model

**Parameters:**
- `m` (Model): The AI backend model to be represented

**Returns:** String representation of the AI backend model

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
repr := viewAIBackend(myModel) // returns a string describing the backend
```

**Edge Cases:**
- Nil or uninitialized Model input
- Model with missing or invalid fields


</details>

<details>
<summary><b><code>viewAIModel(m Model)</code></b></summary>

**Summary:** Generates a string representation of an AI model

**Parameters:**
- `m` (Model): The AI model to be represented

**Returns:** String representation of the AI model

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
repr := viewAIModel(myModel) // returns a string describing the model
```

**Edge Cases:**
- Nil or uninitialized Model input
- Model with missing or invalid fields


</details>

<details>
<summary><b><code>viewAIEndpoint(m Model)</code></b></summary>

**Summary:** Generates a string representation of an AI endpoint

**Parameters:**
- `m` (Model): The AI endpoint model to be represented

**Returns:** String representation of the AI endpoint

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
repr := viewAIEndpoint(myModel) // returns a string describing the endpoint
```

**Edge Cases:**
- Nil or uninitialized Model input
- Model with missing or invalid fields


</details>

<details>
<summary><b><code>viewAPIKey(m Model)</code></b></summary>

**Summary:** Returns API key from Model as string

**Parameters:**
- `m` (Model): Model containing API key

**Returns:** API key as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
key := viewAPIKey(model) // returns "sk_123abc"
```

**Edge Cases:**
- Empty API key in Model
- Non-string type in Model's API key field


</details>

<details>
<summary><b><code>viewOutputFolder(m Model)</code></b></summary>

**Summary:** Returns output folder path from Model

**Parameters:**
- `m` (Model): Model containing output folder path

**Returns:** Output folder path as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
path := viewOutputFolder(model) // returns "/tmp/output"
```

**Edge Cases:**
- Empty path in Model
- Invalid path characters


</details>


---

## 📄 File: `styles.go`

> 📍 `tui\styles.go`

## 📑 Contents



---

## 📄 File: `wizard.go`

> 📍 `tui\wizard.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>Init()</code></b></summary>

**Summary:** Initializes Model and returns tea.Cmd

**Parameters:**
- `receiver` (Model): Model instance to initialize

**Returns:** tea.Cmd for BubbleTea framework

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
cmd := model.Init() // returns initial command
```

**Edge Cases:**
- Uninitialized Model fields
- Nil receiver


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


---

## 📄 File: `model.go`

> 📍 `tui\model.go`

## 📑 Contents

- [🧱 Structs (1)](#-structs)
- [🔧 Functions (1)](#-functions)

## 🧱 Structs

### `Model`

```go
type Model struct {
	step int 
	ProjectDir string 
	AiBackend string 
	AiModel string 
	AiEndpoint string 
	AiAPIKey string 
	OutputFolder string 
	verbose bool 
	inputValue string 
	choices []string 
	cursor int 
}
```

**Summary:** Processes batch AI requests using a client

**Parameters:**
- `requests` ([]AIRequest): Batch of requests to process
- `client` (*ai.Client): AI service client

**Returns:** None (handles requests asynchronously)

**Complexity:**
- Time: O(n) where n is number of requests
- Space: O(n) for request processing buffers

**Example:**
```go
processAIRequests(reqBatch, aiClient)
```

**Edge Cases:**
- Empty requests slice
- Nil client pointer
- Network timeouts


---

## 🔧 Functions

<details>
<summary><b><code>initialModel()</code></b></summary>

**Summary:** Initializes and returns a default Model instance.

**Returns:** Initialized Model instance

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model := initialModel()
```

**Edge Cases:**
- Model initialization failure
- Uninitialized fields in returned Model


</details>


---

## 📄 File: `steps.go`

> 📍 `tui\steps.go`

## 📑 Contents

- [🔧 Functions (13)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>handleTextInput(m Model, msg tea.KeyMsg)</code></b></summary>

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
<summary><b><code>stepProjectDir(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Updates Model state based on key message and returns updated Model with command.

**Parameters:**
- `m` (Model): Current model state
- `msg` (tea.KeyMsg): Key press message

**Returns:** Updated Model and optional tea.Cmd

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
newModel, cmd := stepProjectDir(currentModel, keyMsg)
```

**Edge Cases:**
- Nil or invalid key message
- Model state corruption during update


</details>

<details>
<summary><b><code>stepAIBackend(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Updates AI backend model state based on keyboard input

**Parameters:**
- `m` (Model): Current application model state
- `msg` (tea.KeyMsg): Keyboard input message

**Returns:** Updated Model and optional command (tea.Cmd)

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, cmd := stepAIBackend(currentModel, keyMsg)
```

**Edge Cases:**
- Nil model input
- Unhandled key messages
- State transition errors


</details>

<details>
<summary><b><code>stepAIModel(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Processes keyboard input for AI model interaction

**Parameters:**
- `m` (Model): Current model state
- `msg` (tea.KeyMsg): Keyboard event

**Returns:** Modified Model and possible command

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
newModel, cmd := stepAIModel(appModel, tea.KeyEnter)
```

**Edge Cases:**
- Invalid model state
- Unsupported key combinations
- Command generation failure


</details>

<details>
<summary><b><code>stepAIEndpoint(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Handles keyboard events for AI endpoint configuration

**Parameters:**
- `m` (Model): Application state container
- `msg` (tea.KeyMsg): Key press event

**Returns:** Updated state model and optional command

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, _ := stepAIEndpoint(configModel, tea.KeyTab)
```

**Edge Cases:**
- Missing required model fields
- Edge key cases (Esc, Ctrl+C)
- State validation errors


</details>

<details>
<summary><b><code>stepAPIKey(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Handles API key input updates in a TUI model

**Parameters:**
- `m` (Model): Current application model state
- `msg` (tea.KeyMsg): Key press event to process

**Returns:** Updated Model and optional tea.Cmd for side effects

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, cmd := stepAPIKey(currentModel, keyEvent)
```

**Edge Cases:**
- Empty API key input
- Special key combinations


</details>

<details>
<summary><b><code>stepOutputFolder(m Model, msg tea.KeyMsg)</code></b></summary>

**Summary:** Processes output folder selection key events

**Parameters:**
- `m` (Model): Current application model state
- `msg` (tea.KeyMsg): Key press event to process

**Returns:** Updated Model and optional tea.Cmd for side effects

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
model, cmd := stepOutputFolder(currentModel, keyEvent)
```

**Edge Cases:**
- Invalid folder paths
- Permission denied errors


</details>

<details>
<summary><b><code>viewProjectDir(m Model)</code></b></summary>

**Summary:** Renders project directory view in TUI

**Parameters:**
- `m` (Model): Current application model state

**Returns:** String representation of project directory view

**Complexity:**
- Time: O(n) where n is directory contents
- Space: O(n) for string buffer

**Example:**
```go
dirView := viewProjectDir(currentModel)
```

**Edge Cases:**
- Empty directory
- Unreadable files


</details>

<details>
<summary><b><code>viewAIBackend(m Model)</code></b></summary>

**Summary:** Generates a string representation of an AI backend model

**Parameters:**
- `m` (Model): The AI backend model to be represented

**Returns:** String representation of the AI backend model

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
repr := viewAIBackend(myModel) // returns a string describing the backend
```

**Edge Cases:**
- Nil or uninitialized Model input
- Model with missing or invalid fields


</details>

<details>
<summary><b><code>viewAIModel(m Model)</code></b></summary>

**Summary:** Generates a string representation of an AI model

**Parameters:**
- `m` (Model): The AI model to be represented

**Returns:** String representation of the AI model

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
repr := viewAIModel(myModel) // returns a string describing the model
```

**Edge Cases:**
- Nil or uninitialized Model input
- Model with missing or invalid fields


</details>

<details>
<summary><b><code>viewAIEndpoint(m Model)</code></b></summary>

**Summary:** Generates a string representation of an AI endpoint

**Parameters:**
- `m` (Model): The AI endpoint model to be represented

**Returns:** String representation of the AI endpoint

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
repr := viewAIEndpoint(myModel) // returns a string describing the endpoint
```

**Edge Cases:**
- Nil or uninitialized Model input
- Model with missing or invalid fields


</details>

<details>
<summary><b><code>viewAPIKey(m Model)</code></b></summary>

**Summary:** Returns API key from Model as string

**Parameters:**
- `m` (Model): Model containing API key

**Returns:** API key as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
key := viewAPIKey(model) // returns "sk_123abc"
```

**Edge Cases:**
- Empty API key in Model
- Non-string type in Model's API key field


</details>

<details>
<summary><b><code>viewOutputFolder(m Model)</code></b></summary>

**Summary:** Returns output folder path from Model

**Parameters:**
- `m` (Model): Model containing output folder path

**Returns:** Output folder path as string

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
path := viewOutputFolder(model) // returns "/tmp/output"
```

**Edge Cases:**
- Empty path in Model
- Invalid path characters


</details>


---

## 📄 File: `styles.go`

> 📍 `tui\styles.go`

## 📑 Contents



---

## 📄 File: `wizard.go`

> 📍 `tui\wizard.go`

## 📑 Contents

- [🔧 Functions (4)](#-functions)

## 🔧 Functions

<details>
<summary><b><code>Init()</code></b></summary>

**Summary:** Initializes Model and returns tea.Cmd

**Parameters:**
- `receiver` (Model): Model instance to initialize

**Returns:** tea.Cmd for BubbleTea framework

**Complexity:**
- Time: O(1)
- Space: O(1)

**Example:**
```go
cmd := model.Init() // returns initial command
```

**Edge Cases:**
- Uninitialized Model fields
- Nil receiver


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

