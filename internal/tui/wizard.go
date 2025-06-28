package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.step {
		case 0:
			return stepProjectDir(m, msg)
		case 1:
			return stepAIBackend(m, msg)
		case 2:
			return stepAIModel(m, msg)
		case 3:
			return stepAIEndpoint(m, msg)
		case 4:
			return stepAPIKey(m, msg)
		case 5:
			return stepOutputFolder(m, msg)
		}
	}
	return m, nil
}

func (m Model) View() string {
	header := TitleStyle.Render("📄 DocuGen - Interactive Documentation Generator") + "\n\n"
	switch m.step {
	case 0:
		return header + viewProjectDir(m)
	case 1:
		return header + viewAIBackend(m)
	case 2:
		return header + viewAIModel(m)
	case 3:
		return header + viewAIEndpoint(m)
	case 4:
		return header + viewAPIKey(m)
	case 5:
		return header + viewOutputFolder(m)
	default:
		return header + "Unknown step"
	}
}

func StartWizard() (Model, error) {
	p := tea.NewProgram(initialModel())
	finalModel, err := p.Run()
	if err != nil {
		return Model{}, err
	}

	m, ok := finalModel.(Model)
	if !ok {
		return Model{}, fmt.Errorf("failed to cast final Model")
	}
	return m, nil
}
