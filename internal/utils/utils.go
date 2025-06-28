package utils

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	docTypes "github.com/MRGHOSJ/docupocus/internal/generator/types"
)

func GetGitRemoteURL(projectDir string) string {
	cmd := exec.Command("git", "-C", projectDir, "config", "--get", "remote.origin.url")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// --- detect languages based on file extensions ---
func DetectLanguages(projectDir string) []string {
	langs := make(map[string]bool)
	filepath.WalkDir(projectDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		switch {
		case strings.HasSuffix(path, ".go"):
			langs["Go"] = true
		case strings.HasSuffix(path, ".py"):
			langs["Python"] = true
		case strings.HasSuffix(path, ".js"):
			langs["JavaScript"] = true
		case strings.HasSuffix(path, ".yaml"), strings.HasSuffix(path, ".yml"):
			langs["YAML"] = true
		}
		return nil
	})
	var result []string
	for l := range langs {
		result = append(result, l)
	}
	return result
}

// --- parse go.mod for module name ---
func ParseGoMod(projectDir string) string {
	goModPath := filepath.Join(projectDir, "go.mod")
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}
	return ""
}

// --- parse package.json for JS project name & description ---
func ParsePackageJSON(projectDir string) (string, string) {
	packagePath := filepath.Join(projectDir, "package.json")
	data, err := os.ReadFile(packagePath)
	if err != nil {
		return "", ""
	}
	var pkg struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return "", ""
	}
	return pkg.Name, pkg.Description
}

// --- extract first non-empty line from README.md as description ---
func ExtractDescriptionFromReadme(projectDir string) string {
	readmePath := filepath.Join(projectDir, "README.md")
	data, err := os.ReadFile(readmePath)
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			return line
		}
	}
	return ""
}

// --- detect tech stack based on languages and files ---
func DetectTechStack(projectDir string, langs []string) []string {
	tech := make([]string, 0, len(langs))
	tech = append(tech, langs...)

	if _, err := os.Stat(filepath.Join(projectDir, "Dockerfile")); err == nil {
		tech = append(tech, "Docker")
	}
	if _, err := os.Stat(filepath.Join(projectDir, ".github", "workflows")); err == nil {
		tech = append(tech, "GitHub Actions")
	}
	// Add more detections as needed (Kubernetes manifests, Helm charts, etc)
	return tech
}

// --- detect features and quickstart commands ---
func DetectFeaturesAndQuickstart(projectDir string, langs []string) ([]docTypes.Feature, []docTypes.QuickStartBlock) {
	var features []docTypes.Feature
	var quickstarts []docTypes.QuickStartBlock

	for _, lang := range langs {
		switch lang {
		case "Go":
			// Docker support
			if _, err := os.Stat(filepath.Join(projectDir, "Dockerfile")); err == nil {
				features = append(features, docTypes.Feature{
					Title:       "🐳 Docker Support",
					Description: "Containerized deployment using Docker",
				})
				quickstarts = append(quickstarts, docTypes.QuickStartBlock{
					Title:   "🐳 Run with Docker",
					Shell:   "bash",
					Command: "docker build -t myapp .\ndocker run -p 8080:8080 myapp",
				})
			}
			// CLI app
			if _, err := os.Stat(filepath.Join(projectDir, "main.go")); err == nil {
				features = append(features, docTypes.Feature{
					Title:       "🚀 CLI Application",
					Description: "Command line interface application entry point",
				})
				quickstarts = append(quickstarts, docTypes.QuickStartBlock{
					Title:   "▶️ Run Locally",
					Shell:   "bash",
					Command: "go run main.go",
				})
			}

		case "JavaScript":
			pkgName, _ := ParsePackageJSON(projectDir)
			if pkgName != "" {
				features = append(features, docTypes.Feature{
					Title:       "📦 NPM Package",
					Description: "JavaScript project managed with npm",
				})
				quickstarts = append(quickstarts, docTypes.QuickStartBlock{
					Title:   "▶️ Run with npm",
					Shell:   "bash",
					Command: "npm install && npm start",
				})
			}

		case "Python":
			reqPath := filepath.Join(projectDir, "requirements.txt")
			if _, err := os.Stat(reqPath); err == nil {
				features = append(features, docTypes.Feature{
					Title:       "🐍 Python Dependencies",
					Description: "Dependencies managed via requirements.txt",
				})
				quickstarts = append(quickstarts, docTypes.QuickStartBlock{
					Title:   "▶️ Run Python App",
					Shell:   "bash",
					Command: "pip install -r requirements.txt\npython app.py",
				})
			}

		case "YAML":
			// Assume Kubernetes manifests or CI config
			features = append(features, docTypes.Feature{
				Title:       "☸️ Kubernetes Configs",
				Description: "Kubernetes manifests included",
			})
		}
	}

	return features, quickstarts
}
