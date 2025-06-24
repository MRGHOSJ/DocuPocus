# DocuPocus

Generates Markdown documentation for Go, Python, JavaScript, and YAML projects using AI enhancements.

## Features
- Analyzes code and configurations (Go, Python, JS, YAML).
- Generates detailed documentation with AI-powered explanations.
- Supports two AI backends:
  - **Ollama** (local, private, default): Runs Gemma 2B in Docker.
  - **OpenRouter** (cloud, free tier): Accesses xAI’s Grok 3 Mini via free credits.

## Requirements
- **Go**: 1.21+
- **Docker**: For Ollama (default backend).
- **RAM**: 4GB (Ollama requires ~1.5-2GB).
- **Internet**: For OpenRouter API or initial Ollama setup.

## Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/MRGHOSJ/docupocus
   cd docupocus
   ```
2. Install Go dependencies:
   ```bash
   go mod tidy
   ```
3. Make `run.sh` executable:
   ```bash
   chmod +x run.sh
   ```

## Usage
Run the script to analyze your project and generate `DOCUMENTATION.md`.

### Ollama (Local, Default)
```bash
bash run.sh --project-dir .
```
- Automatically starts Docker, pulls `gemma:2b`, and generates documentation.
- **Privacy**: Fully local, no data sent to external servers.

### OpenRouter (Free Cloud API)
1. Sign up at `https://openrouter.ai` and get a free API key (~$1-$5 credits/month).
2. Run:
   ```bash
   bash run.sh --project-dir . --ai-backend=openrouter --ai-model=xai/grok-3-mini --ai-api-key=<your-openrouter-api-key>
   ```
- **Privacy Warning**: OpenRouter sends code to cloud servers. Use Ollama for sensitive projects.

### Flags
- `--project-dir`: Project directory (default: `.`).
- `--ai-enabled`: Enable AI enhancements (default: `true`).
- `--ai-backend`: `ollama` or `openrouter` (default: `ollama`).
- `--ai-model`: Model name (e.g., `gemma:2b` for Ollama, `xai/grok-3-mini` for OpenRouter).
- `--ai-api-key`: API key for OpenRouter.

## Example
Create test files:
```go
// test.go
package main

type Person struct {
    Name string
    Age  int
}

func Greet(name string) string {
    return "Hello, " + name + "!"
}
```
```bash
bash run.sh --project-dir .
```

Output (`DOCUMENTATION.md`):
```markdown
# Project Documentation

## Package: main

### Structs:
- **Person**: Represents a user with a name and age. (AI: This struct defines a Person with a string field `Name` and an integer field `Age`, used to store basic user information.)

### Functions:
- **Greet**(name string): Returns a greeting message. (AI: The `Greet` function takes a `name` parameter and returns a string concatenating "Hello, " with the name and an exclamation mark, e.g., "Hello, Alice!".)
```

## Privacy
- **Ollama**: Local processing, no data leaves your machine.
- **OpenRouter**: Cloud-based, code is sent to OpenRouter servers. Use Ollama for sensitive projects.

## License
MIT

## Contributing
Issues and PRs welcome at `https://github.com/MRGHOSJ/docupocus`.