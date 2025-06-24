package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/ai"
	aibackend "github.com/MRGHOSJ/docupocus/internal/ai/backend"
	"github.com/MRGHOSJ/docupocus/internal/analyzer"
	"github.com/MRGHOSJ/docupocus/internal/generator"
	"github.com/MRGHOSJ/docupocus/internal/tui"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Parse command line flags
	nonInteractive := flag.Bool("non-interactive", false, "Run in CI mode")
	projectDir := flag.String("project-dir", ".", "Project directory to analyze")
	template := flag.String("template", "default", "Template name for documentation")
	aiEnabled := flag.Bool("ai", true, "Enable AI enhancements")
	aiBackend := flag.String("ai-backend", "openrouter", "AI backend (ollama or openrouter)")
	aiModel := flag.String("ai-model", "deepseek/deepseek-chat-v3-0324:free", "AI model to use")
	aiEndpoint := flag.String("ai-endpoint", "", "Custom AI endpoint URL")
	aiAPIKey := flag.String("ai-api-key", "", "API key for OpenRouter")
	outputFile := flag.String("output", "DOCUMENTATION.md", "Output file path")
	verbose := flag.Bool("verbose", true, "Enable verbose logging")

	flag.Parse()

	// Validate project directory
	absProjectDir, err := filepath.Abs(*projectDir)
	if err != nil {
		return fmt.Errorf("invalid project directory: %w", err)
	}

	// Initialize AI client if enabled
	var aiClient *ai.Client
	if *aiEnabled {
		client, err := setupAIClient(*aiBackend, *aiModel, *aiEndpoint, *aiAPIKey, *verbose)
		if err != nil {
			return fmt.Errorf("AI setup failed: %w", err)
		}
		aiClient = client
	}

	// Run in appropriate mode
	if *nonInteractive {
		if *verbose {
			fmt.Println("🚀 Starting non-interactive documentation generation...")
		}
		return generateDocs(absProjectDir, *template, *outputFile, aiClient, *verbose)
	}

	if *verbose {
		fmt.Println("✨ Starting interactive documentation wizard...")
	}
	return tui.StartWizard()
}

func setupAIClient(backend, model, endpoint, apiKey string, verbose bool) (*ai.Client, error) {
	// Create backend configuration
	cfg := aibackend.BackendConfig{
		Model:    model,
		Endpoint: endpoint,
		APIKey:   apiKey,
	}

	// Create the appropriate backend
	var backendImpl aibackend.Backend
	var err error

	switch strings.ToLower(backend) {
	case "ollama":
		backendImpl, err = aibackend.NewOllamaBackend(cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create Ollama backend: %w", err)
		}
		if verbose {
			fmt.Println("✅ Using Ollama backend with model:", model)
		}
	case "openrouter":
		backendImpl = aibackend.NewOpenRouterBackend(cfg)
		if verbose {
			fmt.Println("✅ Using OpenRouter backend with model:", model)
			if !strings.Contains(model, ":free") {
				fmt.Println("⚠️  Warning: Using non-free OpenRouter model may incur costs")
			}
		}
	default:
		return nil, fmt.Errorf("unsupported backend: %s", backend)
	}

	// Create the AI client
	return ai.NewClient(backendImpl, cfg), nil
}

func generateDocs(projectDir, template, outputFile string, aiClient *ai.Client, verbose bool) error {
	if verbose {
		fmt.Printf("🔍 Analyzing project at: %s\n", projectDir)
	}

	// Analyze project
	result, err := analyzer.AnalyzeProject(projectDir)
	if err != nil {
		return fmt.Errorf("project analysis failed: %w", err)
	}

	if verbose {
		fmt.Printf("✅ Found %d files with documentation\n", len(result.Files))
	}

	// Generate documentation
	cfg := generator.GeneratorConfig{
		AIClient: aiClient,
	}

	doc, err := generator.GenerateMarkdown(result, template, cfg)
	if err != nil {
		return fmt.Errorf("document generation failed: %w", err)
	}

	// Write output file
	outputPath := filepath.Join(projectDir, outputFile)
	if err := os.WriteFile(outputPath, []byte(doc), 0644); err != nil {
		return fmt.Errorf("failed to write documentation file: %w", err)
	}

	if verbose {
		fmt.Printf("📄 Documentation generated at: %s\n", outputPath)
	}
	return nil
}
