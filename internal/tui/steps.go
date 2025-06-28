package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Helper to update Model based on key input for text input steps
func handleTextInput(m Model, msg tea.KeyMsg) Model {
	switch msg.String() {
	case "backspace":
		if len(m.inputValue) > 0 {
			m.inputValue = m.inputValue[:len(m.inputValue)-1]
		}
	default:
		if len(msg.String()) == 1 {
			m.inputValue += msg.String()
		}
	}
	return m
}

// Handle step 0: projectDir input
func stepProjectDir(m Model, msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		m.ProjectDir = strings.TrimSpace(m.inputValue)
		if m.ProjectDir == "" {
			m.ProjectDir = "."
		}
		m.inputValue = ""
		m.choices = []string{"openrouter", "ollama"}
		m.cursor = 0
		m.step = 1
	case "ctrl+c", "esc":
		return m, tea.Quit
	default:
		m = handleTextInput(m, msg)
	}
	return m, nil
}

// Handle step 1: aiBackend selection
func stepAIBackend(m Model, msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down":
		if m.cursor < len(m.choices)-1 {
			m.cursor++
		}
	case "enter":
		m.AiBackend = m.choices[m.cursor]
		m.inputValue = ""
		m.cursor = 0
		m.step = 2
	case "ctrl+c", "esc":
		return m, tea.Quit
	}
	return m, nil
}

// Handle step 2: aiModel input
func stepAIModel(m Model, msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		m.AiModel = strings.TrimSpace(m.inputValue)
		m.inputValue = ""
		if m.AiBackend == "ollama" {
			m.step = 3
		} else {
			m.step = 4
		}
	case "ctrl+c", "esc":
		return m, tea.Quit
	default:
		m = handleTextInput(m, msg)
	}
	return m, nil
}

// Handle step 3: aiEndpoint input (ollama)
func stepAIEndpoint(m Model, msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		m.AiEndpoint = strings.TrimSpace(m.inputValue)
		m.inputValue = ""
		m.step = 5
	case "ctrl+c", "esc":
		return m, tea.Quit
	default:
		m = handleTextInput(m, msg)
	}
	return m, nil
}

// Handle step 4: aiAPIKey input (openrouter)
func stepAPIKey(m Model, msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		m.AiAPIKey = strings.TrimSpace(m.inputValue)
		m.inputValue = ""
		m.step = 5
	case "ctrl+c", "esc":
		return m, tea.Quit
	default:
		m = handleTextInput(m, msg)
	}
	return m, nil
}

// Handle step 5: outputFolder input
func stepOutputFolder(m Model, msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		m.OutputFolder = strings.TrimSpace(m.inputValue)
		if m.OutputFolder == "" {
			m.OutputFolder = "docs"
		}
		return m, tea.Quit
	case "ctrl+c", "esc":
		return m, tea.Quit
	default:
		m = handleTextInput(m, msg)
	}
	return m, nil
}

// Step views (return string to render)

func viewProjectDir(m Model) string {
	return fmt.Sprintf(
		"%s\n%s|\n\n%s",
		QuestionStyle.Render("1. Project Directory (default '.')"),
		InputStyle.Render(m.inputValue),
		"  Type path and press Enter • Ctrl+C to Quit",
	)
}

func viewAIBackend(m Model) string {
	var b strings.Builder
	b.WriteString(QuestionStyle.Render("2. Select AI Backend") + "\n\n")
	for i, choice := range m.choices {
		cursor := " "
		style := InputStyle
		if m.cursor == i {
			cursor = ">"
			style = CursorStyle
		}
		b.WriteString(fmt.Sprintf("%s %s\n", cursor, style.Render(choice)))
	}
	b.WriteString("\n  Use ↑/↓ to navigate, Enter to select, Ctrl+C to Quit")
	return b.String()
}

func viewAIModel(m Model) string {
	return fmt.Sprintf(
		"%s\n%s|\n\n%s\n%s",
		QuestionStyle.Render("3. Enter AI Model"),
		InputStyle.Render(m.inputValue),
		"  Example: deepseek/deepseek-chat-v3-0324:free",
		"  Type your Model and press Enter • Ctrl+C to Quit",
	)
}

func viewAIEndpoint(m Model) string {
	return fmt.Sprintf(
		"%s\n%s|\n\n%s",
		QuestionStyle.Render("4. Enter AI Endpoint (Ollama backend)"),
		InputStyle.Render(m.inputValue),
		"  Type custom AI endpoint URL and press Enter • Ctrl+C to Quit",
	)
}

func viewAPIKey(m Model) string {
	return fmt.Sprintf(
		"%s\n%s|\n\n%s",
		QuestionStyle.Render("4. Enter AI API Key (OpenRouter backend)"),
		InputStyle.Render(m.inputValue),
		"  Type your API key and press Enter • Ctrl+C to Quit",
	)
}

func viewOutputFolder(m Model) string {
	return fmt.Sprintf(
		"%s\n%s|\n\n%s",
		QuestionStyle.Render("5. Output Folder (default 'docs')"),
		InputStyle.Render(m.inputValue),
		"  Type output folder and press Enter to finish • Ctrl+C to Quit",
	)
}
