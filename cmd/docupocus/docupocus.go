package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/ai"
	aibackend "github.com/MRGHOSJ/docupocus/internal/ai/backend"
	"github.com/MRGHOSJ/docupocus/internal/analyzer"
	"github.com/MRGHOSJ/docupocus/internal/generator"
	docTypes "github.com/MRGHOSJ/docupocus/internal/generator/types"
	"github.com/MRGHOSJ/docupocus/internal/tui"
	"github.com/MRGHOSJ/docupocus/internal/utils"
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
	projectDirFlag := flag.String("project-dir", ".", "Project directory to analyze")
	aiBackendFlag := flag.String("ai-backend", "openrouter", "AI backend (ollama or openrouter)")
	aiModelFlag := flag.String("ai-model", "deepseek/deepseek-chat-v3-0324:free", "AI Model to use")
	aiEndpointFlag := flag.String("ai-endpoint", "", "Custom AI endpoint URL")
	aiAPIKeyFlag := flag.String("ai-api-key", "", "API key for OpenRouter")
	outputFolderFlag := flag.String("output", "docs", "Output file path")
	generateSummaryFlag := flag.Bool("summary", false, "Generate a PR change summary")
	baseBranchFlag := flag.String("base-branch", "main", "Base branch to compare against")
	verboseFlag := flag.Bool("verbose", true, "Enable verbose logging")

	flag.Parse()

	projectDir := *projectDirFlag
	aiBackend := *aiBackendFlag
	aiModel := *aiModelFlag
	aiEndpoint := *aiEndpointFlag
	aiAPIKey := *aiAPIKeyFlag
	outputFolder := *outputFolderFlag
	verbose := *verboseFlag
	generateSummary := *generateSummaryFlag
	baseBranch := *baseBranchFlag

	// If interactive, run wizard to get values instead of flags
	if !*nonInteractive {
		if verbose {
			fmt.Println("✨ Starting interactive documentation wizard...")
		}

		m, err := tui.StartWizard()
		if err != nil {
			return fmt.Errorf("failed to run wizard: %w", err)
		}

		// Overwrite with wizard results
		if m.ProjectDir != "" {
			projectDir = m.ProjectDir
		}
		if m.AiBackend != "" {
			aiBackend = m.AiBackend
		}
		if m.AiModel != "" {
			aiModel = m.AiModel
		}
		if aiBackend == "ollama" && m.AiEndpoint != "" {
			aiEndpoint = m.AiEndpoint
			// Clear API key if switching backend
			aiAPIKey = ""
		}
		if aiBackend == "openrouter" && m.AiAPIKey != "" {
			aiAPIKey = m.AiAPIKey
			// Clear endpoint if switching backend
			aiEndpoint = ""
		}
		if m.OutputFolder != "" {
			outputFolder = m.OutputFolder
		}
	}

	// Validate project directory
	absProjectDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("invalid project directory: %w", err)
	}

	// Setup AI client
	aiClient, err := setupAIClient(aiBackend, aiModel, aiEndpoint, aiAPIKey, verbose)
	if err != nil {
		return fmt.Errorf("AI setup failed: %w", err)
	}

	if generateSummary {
		if verbose {
			fmt.Println("🧠 Generating pull request summary...")
		}
		if err := generatePRSummary(absProjectDir, baseBranch, aiClient); err != nil {
			return fmt.Errorf("failed to generate PR summary: %w", err)
		}
		return nil
	}

	// Generate docs
	if verbose {
		fmt.Println("🚀 Starting documentation generation...")
	}
	return generateDocs(absProjectDir, outputFolder, aiClient, verbose)
}

func setupAIClient(backend, Model, endpoint, apiKey string, verbose bool) (*ai.Client, error) {
	// Create backend configuration
	cfg := aibackend.BackendConfig{
		Model:    Model,
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
			fmt.Println("✅ Using Ollama backend with Model:", Model)
		}
	case "openrouter":
		backendImpl = aibackend.NewOpenRouterBackend(cfg)
		if verbose {
			fmt.Println("✅ Using OpenRouter backend with Model:", Model)
			if !strings.Contains(Model, ":free") {
				fmt.Println("⚠️  Warning: Using non-free OpenRouter Model may incur costs")
			}
		}
	default:
		return nil, fmt.Errorf("unsupported backend: %s", backend)
	}

	// Create the AI client
	client := ai.NewClient(backendImpl, cfg)
	client.ApplyDefaults()

	return client, nil

}

func generateDocs(projectDir, outputFolder string, aiClient *ai.Client, verbose bool) error {
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

	langs := utils.DetectLanguages(projectDir)
	if verbose {
		fmt.Printf("🧩 Detected languages: %v\n", langs)
	}

	projectName := ""
	projectDescription := ""

	// Try Go mod for name
	projectName = utils.ParseGoMod(projectDir)

	// Try package.json if no Go module name
	if projectName == "" {
		n, _ := utils.ParsePackageJSON(projectDir)
		projectName = n
	}

	// Fallback to folder name
	if projectName == "" {
		projectName = filepath.Base(projectDir)
	}

	projectDescription = utils.ExtractDescriptionFromReadme(projectDir)

	repoURL := utils.GetGitRemoteURL(projectDir)

	techStack := utils.DetectTechStack(projectDir, langs)

	features, quickstarts := utils.DetectFeaturesAndQuickstart(projectDir, langs)

	cfg := docTypes.GeneratorConfig{
		AIClient:  aiClient,
		OutputDir: outputFolder,
		Project: docTypes.ProjectMeta{
			Name:        projectName,
			Description: projectDescription,
			RepoURL:     repoURL,
			Features:    features,
			TechStack:   techStack,
			QuickStart:  quickstarts,
			BestPractices: docTypes.BestPractices{
				Do: []string{
					"Keep functions small and focused",
					"Write clear comments and documentation",
					"Validate user input",
				},
				Dont: []string{
					"Use global state unnecessarily",
					"Ignore error handling",
					"Hardcode configuration values",
				},
			},
		},
	}

	if err = generator.GeneratePackageDocs(result, cfg); err != nil {
		return fmt.Errorf("document generation failed: %w", err)
	}

	if verbose {
		fmt.Printf("✅ Documentation generated successfully\n")
	}

	return nil
}

func getAllDiff(projectDir, baseBranch string) (string, error) {
	// Fetch base branch first
	cmdFetch := exec.Command("git", "fetch", "origin", baseBranch)
	cmdFetch.Dir = projectDir
	if err := cmdFetch.Run(); err != nil {
		return "", fmt.Errorf("failed to fetch base branch: %w", err)
	}

	// Diff committed changes against base branch
	cmdDiffCommitted := exec.Command("git", "diff", fmt.Sprintf("origin/%s...HEAD", baseBranch))
	cmdDiffCommitted.Dir = projectDir
	committedDiff, err := cmdDiffCommitted.Output()
	if err != nil {
		return "", fmt.Errorf("git diff committed failed: %w", err)
	}

	// Diff staged but uncommitted changes
	cmdDiffStaged := exec.Command("git", "diff", "--cached")
	cmdDiffStaged.Dir = projectDir
	stagedDiff, err := cmdDiffStaged.Output()
	if err != nil {
		return "", fmt.Errorf("git diff staged failed: %w", err)
	}

	// Diff unstaged changes
	cmdDiffUnstaged := exec.Command("git", "diff")
	cmdDiffUnstaged.Dir = projectDir
	unstagedDiff, err := cmdDiffUnstaged.Output()
	if err != nil {
		return "", fmt.Errorf("git diff unstaged failed: %w", err)
	}

	// Combine all diffs
	fullDiff := string(committedDiff) + "\n" + string(stagedDiff) + "\n" + string(unstagedDiff)
	return fullDiff, nil
}

func generatePRSummary(projectDir, baseBranch string, aiClient *ai.Client) error {
	// Get the full diff string using your helper
	diff, err := getAllDiff(projectDir, baseBranch)
	if err != nil {
		return fmt.Errorf("failed to get diff: %w", err)
	}

	if diff == "" {
		fmt.Println("✅ No changes detected.")
		return nil
	}

	fmt.Println("🔍 Files and changes in this PR:\n", diff)

	// Call AI summary API with diff content
	ctx := context.Background()
	summary, err := aiClient.CallSummaryAPI(ctx, diff)
	if err != nil {
		return fmt.Errorf("failed to generate AI summary: %w", err)
	}

	fmt.Println("\n📋 **PR Summary Preview:**")
	fmt.Println(summary)

	return nil
}
