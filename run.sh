#!/bin/bash
# .\main.exe --project-dir .  --ai-backend=openrouter --ai-model=google/gemma-7b-it:free  --ai-api-key=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
# Parse flags
PROJECT_DIR="."
AI_BACKEND="ollama"
AI_MODEL="gemma:2b"
AI_API_KEY="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"
AI_ENABLED="true"

while [[ $# -gt 0 ]]; do
    case "$1" in
        --project-dir=*)
            PROJECT_DIR="${1#*=}"
            ;;
        --ai-backend=*)
            AI_BACKEND="${1#*=}"
            ;;
        --ai-model=*)
            AI_MODEL="${1#*=}"
            ;;
        --ai-api-key=*)
            AI_API_KEY="${1#*=}"
            ;;
        --ai-enabled=*)
            AI_ENABLED="${1#*=}"
            ;;
        *)
            echo "Unknown flag: $1"
            exit 1
            ;;
    esac
    shift
done

# Build flags for the Go binary
FLAGS="--non-interactive --project-dir=\"$PROJECT_DIR\""
FLAGS+=" --ai-enabled=$AI_ENABLED"
FLAGS+=" --ai-backend=$AI_BACKEND"
FLAGS+=" --ai-model=\"$AI_MODEL\""

# Add API key only if using OpenRouter
if [ "$AI_BACKEND" = "openrouter" ] && [ -n "$AI_API_KEY" ]; then
    FLAGS+=" --ai-api-key=\"$AI_API_KEY\""
fi

# Start Docker for Ollama if needed
if [ "$AI_BACKEND" = "ollama" ] && [ "$AI_ENABLED" = "true" ]; then
    if ! docker info &> /dev/null; then
        echo "🚀 Starting Docker..."
        open --background -a Docker
        while ! docker info &> /dev/null; do sleep 1; done
    fi

    if ! docker ps | grep -q "ollama"; then
        echo "🚀 Starting Ollama container..."
        docker run -d --name ollama -p 11434:11434 ollama/ollama
    fi

    echo "🔍 Checking for model: $AI_MODEL"
    if ! docker exec ollama ollama list | grep -q "$AI_MODEL"; then
        echo "📥 Downloading model: $AI_MODEL"
        docker exec ollama ollama pull "$AI_MODEL"
    fi
fi

# Run the Go binary
echo "🚀 Generating documentation with DocuPocus..."
eval "go run cmd/docupocus/main.go $FLAGS"

echo "✅ Documentation generated at $PROJECT_DIR/DOCUMENTATION.md"