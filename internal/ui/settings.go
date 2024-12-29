package ui

import (
	"mainframe/pkg/config"
	"mainframe/pkg/styles"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SettingsModel struct {
	BaseModel
	choices      []string
	cursor       int
	apiKeyInput  textinput.Model
	showAPIInput bool
	modelChoice  string
	config       *config.Config
	quit         bool
	errorMsg     string
	showHelp     bool
}

func NewSettingsModel() *SettingsModel {
	apiKey := textinput.New()
	apiKey.Placeholder = "Enter your API key"
	apiKey.Width = 50
	apiKey.EchoMode = textinput.EchoPassword
	apiKey.CharLimit = 100

	cfg, err := config.Load()
	if err != nil {
		cfg = &config.Config{
			AIModel: "local",
		}
	}

	return &SettingsModel{
		choices: []string{
			"AI Model",
			"Model Configuration",
			"Developer Options",
			"Back to Main Menu",
		},
		cursor:       0,
		apiKeyInput:  apiKey,
		showAPIInput: false,
		modelChoice:  cfg.AIModel,
		config:       cfg,
		quit:         false,
	}
}

func (m *SettingsModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *SettingsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case WindowSizeMsg:
		m.UpdateSize(msg.Width, msg.Height)
		return m, nil

	case tea.KeyMsg:
		if m.showAPIInput {
			m.apiKeyInput, cmd = m.apiKeyInput.Update(msg)

			switch msg.String() {
			case "enter":
				if m.validateAPIKey() {
					m.config.APIKey = m.apiKeyInput.Value()
					m.showAPIInput = false
					m.errorMsg = ""
					config.Save(m.config)
				}
				return m, nil
			case "esc":
				m.showAPIInput = false
				m.errorMsg = ""
				return m, nil
			}
			return m, cmd
		}

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
			switch m.cursor {
			case 0: // AI Model
				if m.modelChoice == "local" {
					m.modelChoice = "gpt"
				} else {
					m.modelChoice = "local"
				}
				m.config.AIModel = m.modelChoice
				config.Save(m.config)
			case 1: // Model Configuration
				if m.modelChoice == "local" {
					return NewLocalModelModel(m.config), nil
				} else {
					m.showAPIInput = true
					m.apiKeyInput.Focus()
					return m, textinput.Blink
				}
			case 2: // Developer Options
				return NewDeveloperModel(m.config), nil
			case 3: // Back to Main Menu
				return NewHomeModel(), nil
			}
		case "?":
			m.showHelp = !m.showHelp
		case "esc":
			return NewHomeModel(), nil
		}
	}

	return m, nil
}

func (m *SettingsModel) validateAPIKey() bool {
	key := m.apiKeyInput.Value()
	if len(key) < 32 {
		m.errorMsg = "API key must be at least 32 characters"
		return false
	}
	if !strings.HasPrefix(key, "sk-") {
		m.errorMsg = "Invalid API key format"
		return false
	}
	return true
}

func (m *SettingsModel) View() string {
	if m.showHelp {
		return m.CenterView(
			styles.DialogBox.Render(
				styles.AppTitle.Render("Settings Help") + "\n\n" +
					"Navigation:\n" +
					"• Up/Down or j/k: Move cursor\n" +
					"• Enter/Space: Select option\n" +
					"• ?: Toggle help\n" +
					"• Esc: Back to main menu\n\n" +
					styles.PageFooter.Render("Press ? to close help"),
			),
		)
	}

	if m.showAPIInput {
		return m.CenterView(
			styles.DialogBox.Render(
				styles.AppTitle.Render("API Key Configuration") + "\n\n" +
					styles.MenuOption.Render("Enter your OpenAI API key:") + "\n" +
					styles.InputBox.Render(m.apiKeyInput.View()) + "\n" +
					(func() string {
						if m.errorMsg != "" {
							return "\n" + styles.ErrorText.Render(m.errorMsg)
						}
						return ""
					})() + "\n\n" +
					styles.PageFooter.Render("enter to save • esc to cancel"),
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

		if m.cursor == i {
			menuContent += styles.HighlightedOption.Render(cursor+choice) + "\n"
		} else {
			menuContent += styles.MenuOption.Render(cursor+choice) + "\n"
		}
	}

	menuView := styles.MenuBox.Render(
		styles.SectionTitle.Render("Settings") + "\n\n" +
			menuContent + "\n\n" +
			styles.PageFooter.Render("? for help • esc to go back"),
	)

	// Right panel - Detailed content
	var detailContent string
	switch m.cursor {
	case 0: // AI Model
		detailContent = styles.MainTitle.Render("AI Model Selection") + "\n\n" +
			styles.Description.Render(
				"Choose the AI model that powers your learning experience:\n\n"+
					"• Local Model\n"+
					"  Run models directly on your machine\n"+
					"  Supports Llama 2, GPT-J, and BLOOM\n"+
					"  Complete privacy and offline usage\n\n"+
					"• GPT Model\n"+
					"  Use OpenAI's powerful GPT models\n"+
					"  Requires internet and API key\n"+
					"  State-of-the-art performance\n\n",
			) +
			styles.StatusIndicator.Render("Current Model: "+strings.ToUpper(m.modelChoice))

	case 1: // Model Configuration
		if m.modelChoice == "local" {
			detailContent = styles.MainTitle.Render("Local Model Setup") + "\n\n" +
				styles.Description.Render(
					"Configure your local model installation:\n\n"+
						"1. Download a compatible model\n"+
						"2. Set up the model path\n"+
						"3. Test the connection\n\n"+
						"Press ENTER to start the setup process",
				)
		} else {
			detailContent = styles.MainTitle.Render("OpenAI Configuration") + "\n\n" +
				styles.Description.Render(
					"Configure your OpenAI API access:\n\n"+
						"• Set up your API key\n"+
						"• Manage model preferences\n"+
						"• Test API connectivity\n\n"+
						"Press ENTER to configure your API key",
				)
		}

	case 2: // Developer Options
		detailContent = styles.MainTitle.Render("Developer Options") + "\n\n" +
			styles.Description.Render(
				"Advanced settings for development and debugging:\n\n"+
					"• Debug logging\n"+
					"• Performance monitoring\n"+
					"• Experimental features\n"+
					"• Network diagnostics\n\n"+
					"Press ENTER to access developer settings",
			)
	}

	detailView := styles.ContentBox.Render(detailContent)

	// Combine views
	return m.SplitView(menuView, detailView)
}
