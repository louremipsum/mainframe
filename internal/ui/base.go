package ui

import (
	"mainframe/pkg/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// WindowSizeMsg is sent when the terminal window is resized
type WindowSizeMsg struct {
	Width  int
	Height int
}

// BaseModel provides common functionality for all models
type BaseModel struct {
	width  int
	height int
}

func (m *BaseModel) Init() tea.Cmd {
	return nil
}

func (m *BaseModel) UpdateSize(width, height int) {
	m.width = width
	m.height = height
	styles.UpdateSplitSizes(width, height)
}

// SplitView renders content in a split view layout
func (m *BaseModel) SplitView(left, right string) string {
	return styles.DocStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.SplitLeft.Render(left),
			styles.SplitRight.Render(right),
		),
	)
}

// CenterView renders content in a centered layout
func (m *BaseModel) CenterView(content string) string {
	return styles.DocStyle.Render(
		lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			content,
		),
	)
}
