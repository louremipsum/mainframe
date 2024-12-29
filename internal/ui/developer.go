package ui

import (
	"mainframe/pkg/config"
	"mainframe/pkg/styles"

	tea "github.com/charmbracelet/bubbletea"
)

type DeveloperModel struct {
	BaseModel
	choices     []string
	cursor      int
	config      *config.Config
	quit        bool
	showHelp    bool
	description string
}

func NewDeveloperModel(cfg *config.Config) *DeveloperModel {
	return &DeveloperModel{
		choices: []string{
			"Debug Mode",
			"Log Output",
			"Experimental Features",
			"Performance Metrics",
			"Network Diagnostics",
			"Back to Settings",
		},
		cursor:      0,
		config:      cfg,
		description: "Toggle development features and debugging tools",
	}
}

func (m *DeveloperModel) Init() tea.Cmd {
	return nil
}

func (m *DeveloperModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.updateDescription()
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
			m.updateDescription()
		case "enter", " ":
			switch m.cursor {
			case 0: // Debug Mode
				m.config.Debug = !m.config.Debug
			case 1: // Log Output
				m.config.Logs = !m.config.Logs
			case 2: // Experimental Features
				m.config.Experimental = !m.config.Experimental
			case 5: // Back to Settings
				return NewSettingsModel(), nil
			}
			config.Save(m.config)
		case "?":
			m.showHelp = !m.showHelp
		case "esc":
			return NewSettingsModel(), nil
		}
	}
	return m, nil
}

func (m *DeveloperModel) updateDescription() {
	switch m.cursor {
	case 0:
		m.description = "Enable detailed debug information and error reporting for troubleshooting"
	case 1:
		m.description = "Save detailed logs to ~/.mainframe/logs for system analysis"
	case 2:
		m.description = "Enable experimental features and updates (may be unstable)"
	case 3:
		m.description = "Monitor system performance, memory usage, and resource allocation"
	case 4:
		m.description = "Test network connectivity and API endpoint responsiveness"
	case 5:
		m.description = "Return to the settings menu"
	}
}

func getStatusIcon(enabled bool) string {
	if enabled {
		return "◉" // Filled circle for enabled
	}
	return "○" // Empty circle for disabled
}

func (m *DeveloperModel) View() string {
	if m.showHelp {
		return m.CenterView(
			styles.DialogBox.Render(
				styles.AppTitle.Render("Developer Options Help") + "\n\n" +
					"Navigation:\n" +
					"• Up/Down or j/k: Move cursor\n" +
					"• Enter/Space: Toggle option\n" +
					"• ?: Toggle help\n" +
					"• Esc: Back to settings\n" +
					"• Ctrl+c/q: Quit\n\n" +
					styles.PageFooter.Render("Press ? to close help"),
			),
		)
	}

	// Left panel - Menu options
	var menuContent string
	for i, choice := range m.choices {
		cursor := "  "
		if m.cursor == i {
			cursor = "> "
		}

		status := ""
		icon := ""
		switch i {
		case 0:
			icon = getStatusIcon(m.config.Debug)
			if m.config.Debug {
				status = " " + styles.SuccessText.Render("ON")
			} else {
				status = " OFF"
			}
		case 1:
			icon = getStatusIcon(m.config.Logs)
			if m.config.Logs {
				status = " " + styles.SuccessText.Render("ON")
			} else {
				status = " OFF"
			}
		case 2:
			icon = getStatusIcon(m.config.Experimental)
			if m.config.Experimental {
				status = " " + styles.WarningText.Render("ON")
			} else {
				status = " OFF"
			}
		case 3, 4:
			icon = "⊘" // Disabled icon
			status = " " + styles.ErrorText.Render("SOON")
		}

		option := icon + " " + choice + status
		if m.cursor == i {
			menuContent += styles.HighlightedOption.Render(cursor+option) + "\n"
		} else {
			menuContent += styles.MenuOption.Render(cursor+option) + "\n"
		}
	}

	menuView := styles.MenuBox.Render(
		styles.SectionTitle.Render("Developer Options") + "\n\n" +
			menuContent + "\n\n" +
			styles.PageFooter.Render("? for help • esc to go back"),
	)

	// Right panel - Detailed content
	detailContent := styles.MainTitle.Render("Developer Tools") + "\n\n"

	// Add feature description
	detailContent += styles.Description.Render(m.description) + "\n\n"

	// Add system information
	detailContent += styles.SectionTitle.Render("System Information") + "\n" +
		styles.Description.Render(
			"• Config Path: ~/.mainframe/config.json\n"+
				"• Log Path: ~/.mainframe/logs\n"+
				"• Debug Level: "+(func() string {
				if m.config.Debug {
					return styles.SuccessText.Render("VERBOSE")
				}
				return "NORMAL"
			})(),
		) + "\n\n"

	// Add status dashboard
	detailContent += styles.SectionTitle.Render("Status Dashboard") + "\n" +
		styles.Description.Render(
			"Debug Mode:          "+getStatusIndicator(m.config.Debug)+"\n"+
				"Log Output:          "+getStatusIndicator(m.config.Logs)+"\n"+
				"Experimental Mode:   "+getStatusIndicator(m.config.Experimental)+"\n"+
				"Performance Monitor: "+styles.ErrorText.Render("NOT AVAILABLE")+"\n"+
				"Network Diagnostics: "+styles.ErrorText.Render("NOT AVAILABLE"),
		)

	detailView := styles.ContentBox.Render(detailContent)

	// Combine views
	return m.SplitView(menuView, detailView)
}

func getStatusIndicator(enabled bool) string {
	if enabled {
		return styles.SuccessText.Render("[ACTIVE]")
	}
	return styles.WarningText.Render("[INACTIVE]")
}
