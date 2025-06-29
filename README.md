# 📚 DocuPocus

Generate intelligent Markdown documentation and pull request summaries for Go, Python, JavaScript, and YAML projects — all powered by AI.

---

## ✨ Features

- ✅ **AI-enhanced documentation** for Go, YAML, Python, and JavaScript code  
- 🤖 **Automated PR summaries** that describe what changed and why  
- 🔄 **GitHub Actions integration** for CI-based doc generation and PR commenting  
- 🧠 Support for both **Ollama** (local AI backend) and **OpenRouter** (cloud-based API)  
- ⚙️ YAML structure breakdown with field types, best practices, usage, and defaults  

---

## 🚀 Quick Start

### 1. Clone and build the CLI

```bash
git clone https://github.com/MRGHOSJ/DocuPocus
cd DocuPocus
go build -o docupocus ./cmd/docupocus
```

### 2. Run the interactive wizard

```bash
./docupocus
```

Or run non-interactively:

```bash
./docupocus --non-interactive --ai-backend ollama --output docs
```

---

## 🧠 AI Backends

| Backend        | Type   | Description                                  |
|----------------|--------|----------------------------------------------|
| **Ollama**     | Local  | Fast & private. Runs `gemma:2b` locally.     |
| **OpenRouter** | Cloud  | Access powerful models like `grok-3-mini`.   |

> 💡 Use Ollama for full privacy. OpenRouter sends code to cloud servers.

---

## 🛠️ Flags

| Flag              | Description                                               |
|-------------------|-----------------------------------------------------------|
| `--project-dir`   | Project directory to analyze (default: `.`)              |
| `--output`        | Output folder for docs (default: `docs`)                 |
| `--ai-backend`    | `ollama` or `openrouter`                                  |
| `--ai-model`      | e.g., `gemma:2b` or `deepseek/deepseek-chat-v3-0324:free`|
| `--ai-api-key`    | API key for OpenRouter                                    |
| `--ai-endpoint`   | Custom endpoint for Ollama (default: `http://localhost`) |
| `--summary`       | Generate summary of pull request changes                 |
| `--base-branch`   | Base branch to compare PR diffs against (`main`, etc.)   |

---

## 🧪 Example Output

### Code Example

```go
type Person struct {
    Name string
    Age  int
}

func Greet(name string) string {
    return "Hello, " + name
}
```

### Generated Documentation

```md
## Struct: Person
- Represents a user with name and age.
- Fields:
  - `Name`: string — person's name
  - `Age`: int — person's age

## Function: Greet
- Returns a greeting for the given name.
- Time complexity: O(1)
```

---

## 💬 Pull Request Summary (via GitHub Actions)

DocuPocus can post a comment like this on every pull request:

```md
### 🤖 Automated PR Summary

**Changed Files:**
- `docupocus.go`: Added `--summary` flag
- `prompt.go`: New templates for PR reasoning

**Key Improvements:**
- New PR summary generation feature
- Supports AI diff analysis and comment automation
```

See `.github/workflows/pr-summary.yaml` for automation setup.

---

## 🔄 GitHub Actions Usage

Supports:

- ✅ **Documentation preview** per PR  
- ✅ **AI-generated summary** comments on PRs

### Basic Setup

```yaml
uses: ./.github/actions/docupocus
with:
  ai-backend: openrouter
  ai-api-key: ${{ secrets.OPENROUTER_API_KEY }}
```

To post automated PR summaries, create a separate workflow that runs:

```bash
./docupocus --non-interactive --summary --base-branch main
```

Then pipe that into a PR comment using `actions/github-script`.

---

## 🔒 Privacy Notes

- **Ollama**: Fully local and private — ideal for sensitive code.  
- **OpenRouter**: Cloud-based — your code is sent to third-party APIs.

---

## 📄 License

[MIT License](./LICENSE)

---

## 🤝 Contributing

Contributions are welcome!  
Open an issue or submit a pull request at:  
👉 [https://github.com/MRGHOSJ/DocuPocus](https://github.com/MRGHOSJ/DocuPocus)
