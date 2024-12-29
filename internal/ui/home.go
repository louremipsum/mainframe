package ui

import (
	"mainframe/pkg/styles"

	tea "github.com/charmbracelet/bubbletea"
)

type HomeModel struct {
	BaseModel
	choices []string
	cursor  int
	quit    bool
}

func NewHomeModel() *HomeModel {
	return &HomeModel{
		choices: []string{
			"Start Lesson",
			"Sandbox Mode",
			"Challenges",
			"Settings",
			"Exit",
		},
		cursor: 0,
		quit:   false,
	}
}

func (m *HomeModel) Init() tea.Cmd {
	return nil
}

func (m *HomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case WindowSizeMsg:
		m.UpdateSize(msg.Width, msg.Height)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			switch m.choices[m.cursor] {
			case "Exit":
				m.quit = true
				return m, tea.Quit
			case "Start Lesson":
				// TODO: Implement transition to lesson view
				return m, nil
			case "Sandbox Mode":
				// TODO: Implement transition to sandbox view
				return m, nil
			case "Challenges":
				// TODO: Implement transition to challenges view
				return m, nil
			case "Settings":
				return NewSettingsModel(), nil
			}
		}
	}
	return m, nil
}

func (m *HomeModel) View() string {
	var menuContent string
	for i, choice := range m.choices {
		cursor := "  "
		if m.cursor == i {
			cursor = "> "
		}

		if m.cursor == i {
			menuContent += styles.HighlightedOption.Render(cursor+choice) + "\n"
		} else {
			menuContent += styles.MenuOption.Render(cursor+choice) + "\n"
		}
	}

	content := styles.AppTitle.Render("Mainframe") + "\n" +
		styles.SubTitle.Render("An immersive terminal-based learning environment") + "\n\n" +
		styles.MenuBox.Render(menuContent) + "\n" +
		styles.Description.Render(
			"Welcome to Mainframe, your gateway to mastering terminal commands\n"+
				"and system administration through interactive learning.\n\n"+
				"• Gamified lessons with progressive difficulty\n"+
				"• Real-world scenarios in a safe environment\n"+
				"• AI-powered guidance and assistance\n",
		) + "\n" +
		styles.PageFooter.Render("↑/↓ to move • enter to select • q to quit")

	return m.CenterView(content)
}
