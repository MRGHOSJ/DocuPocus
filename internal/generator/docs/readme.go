package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MRGHOSJ/docupocus/internal/analyzer"
	cfg "github.com/MRGHOSJ/docupocus/internal/generator/types"
	generator "github.com/MRGHOSJ/docupocus/internal/generator/utils"
)

func GenerateProjectReadme(result *analyzer.AnalyzerResult, cfg cfg.GeneratorConfig) error {
	var b strings.Builder
	readmePath := filepath.Join(cfg.OutputDir, "README.md")

	// Header
	b.WriteString(fmt.Sprintf("# %s\n\n", cfg.Project.Name))
	b.WriteString(fmt.Sprintf("> %s\n\n", cfg.Project.Description))

	if cfg.Project.RepoURL != "" {
		b.WriteString(fmt.Sprintf(
			"[![Go](https://img.shields.io/badge/Go-%E2%9D%A4%EF%B8%8F-blue)](%s) "+
				"[![GitHub](https://img.shields.io/badge/GitHub-Repository-lightgrey)](%s)\n\n",
			cfg.Project.RepoURL, cfg.Project.RepoURL,
		))
	}

	// Overview
	b.WriteString("## 🧭 Overview\n\n")
	b.WriteString("| Feature | Description |\n|---------|-------------|\n")
	for _, f := range cfg.Project.Features {
		b.WriteString(fmt.Sprintf("| %s | %s |\n", f.Title, f.Description))
	}
	b.WriteString("\n")

	if len(cfg.Project.TechStack) > 0 {
		b.WriteString("**🛠 Tech Stack:** ")
		for i, tech := range cfg.Project.TechStack {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(fmt.Sprintf("`%s`", tech))
		}
		b.WriteString("\n\n")
	}

	// Packages
	b.WriteString("## 📦 Packages\n\n")
	b.WriteString("> Explore each documented package below:\n\n")
	b.WriteString("<table>\n<tr>\n")
	count := 0
	for _, file := range result.Files {
		for _, pkg := range file.Packages {
			docPath := filepath.Join(pkg.Name, "README.md")
			b.WriteString("<td valign=\"top\" width=\"33%\">\n\n")
			b.WriteString(fmt.Sprintf("### [%s](%s)\n", pkg.Name, docPath))
			b.WriteString(fmt.Sprintf("<small>`%s`</small><br/>\n", file.Path))
			b.WriteString(fmt.Sprintf("📘 %d structs<br/>\n", len(pkg.Structs)))
			b.WriteString(fmt.Sprintf("🛠 %d functions<br/>\n", len(pkg.Funcs)))
			b.WriteString(fmt.Sprintf("📊 %d%% documented\n", generator.CalculateDocCompletion(pkg)))
			b.WriteString("</td>\n")

			count++
			if count%3 == 0 {
				b.WriteString("</tr>\n<tr>\n")
			}
		}
	}
	b.WriteString("</tr>\n</table>\n\n")

	// Quick Start (optional, dynamic from cfg.Project.QuickStart)
	if len(cfg.Project.QuickStart) > 0 {
		b.WriteString("## 🚀 Quick Start\n\n")
		for _, qs := range cfg.Project.QuickStart {
			b.WriteString(fmt.Sprintf("<details>\n<summary><b>%s</b></summary>\n\n", qs.Title))
			b.WriteString(fmt.Sprintf("```%s\n%s\n```\n", qs.Shell, qs.Command))
			b.WriteString("</details>\n\n")
		}
	}

	// Best Practices
	if len(cfg.Project.BestPractices.Do) > 0 || len(cfg.Project.BestPractices.Dont) > 0 {
		b.WriteString("## ✅ Best Practices\n\n")
		b.WriteString("| 👍 Do | 👎 Don’t |\n|-------|-----------|\n")
		max := len(cfg.Project.BestPractices.Do)
		if len(cfg.Project.BestPractices.Dont) > max {
			max = len(cfg.Project.BestPractices.Dont)
		}
		for i := 0; i < max; i++ {
			do := ""
			dont := ""
			if i < len(cfg.Project.BestPractices.Do) {
				do = cfg.Project.BestPractices.Do[i]
			}
			if i < len(cfg.Project.BestPractices.Dont) {
				dont = cfg.Project.BestPractices.Dont[i]
			}
			b.WriteString(fmt.Sprintf("| %s | %s |\n", do, dont))
		}
		b.WriteString("\n")
	}

	return os.WriteFile(readmePath, []byte(b.String()), 0644)
}
