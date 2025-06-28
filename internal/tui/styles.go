package tui

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#7D56F4")).Padding(0, 1)
	QuestionStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#04B575"))
	InputStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4"))
	CursorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4")).Bold(true)
)
