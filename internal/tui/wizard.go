package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)
	questionStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#04B575"))
	inputStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4"))
)

type model struct {
	step       int
	projectDir string
	template   string
	aiEnabled  bool
	inputs     []string
	cursor     int
}

func initialModel() model {
	return model{
		inputs: make([]string, 3),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.step < 2 {
				m.step++
			} else {
				// Start documentation generation
				return m, tea.Quit
			}
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < 2 {
				m.cursor++
			}
		default:
			// Handle text input
			if m.step == 0 {
				m.projectDir += msg.String()
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("📄 DocuGen - Interactive Documentation Generator") + "\n\n"

	switch m.step {
	case 0:
		s += questionStyle.Render("1. Project Path :") + "\n"
		s += inputStyle.Render(m.projectDir) + "|" + "\n\n"
		s += "  ↑/↓: Navigate • Enter: Confirm • Ctrl+C: Quit"

	case 1:
		s += questionStyle.Render("2. Select Template:") + "\n"
		options := []string{"Default", "Modern", "Minimal"}
		for i, option := range options {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, option)
		}
		s += "\n  ↑/↓: Navigate • Enter: Select • Ctrl+C: Quit"

	case 2:
		s += questionStyle.Render("3. AI Enhancement:") + "\n"
		options := []string{"Enable AI (Recommended)", "Disable AI"}
		for i, option := range options {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, option)
		}
		s += "\n\n" + "  ↑/↓: Navigate • Enter: Generate • Ctrl+C: Quit"
	}

	return s
}

func StartWizard() error {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
