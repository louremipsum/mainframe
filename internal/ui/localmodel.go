package ui

import (
	"mainframe/pkg/config"
	"mainframe/pkg/styles"

	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type LocalModelModel struct {
	BaseModel
	choices     []string
	cursor      int
	config      *config.Config
	quit        bool
	showHelp    bool
	pathInput   textinput.Model
	showInput   bool
	errorMsg    string
	currentStep int
}

func NewLocalModelModel(cfg *config.Config) *LocalModelModel {
	pathInput := textinput.New()
	pathInput.Placeholder = "Enter path to model weights"
	pathInput.Width = 50

	return &LocalModelModel{
		choices: []string{
			"1. Download Model",
			"2. Configure Model Path",
			"3. Test Model",
			"Back to Settings",
		},
		cursor:      0,
		config:      cfg,
		pathInput:   pathInput,
		currentStep: 1,
	}
}

func (m *LocalModelModel) Init() tea.Cmd {
	return nil
}

func (m *LocalModelModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case WindowSizeMsg:
		m.UpdateSize(msg.Width, msg.Height)
		return m, nil

	case tea.KeyMsg:
		if m.showInput {
			m.pathInput, cmd = m.pathInput.Update(msg)

			switch msg.String() {
			case "enter":
				// TODO: Validate model path
				m.showInput = false
				m.currentStep = 3
				return m, nil
			case "esc":
				m.showInput = false
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
			case 0: // Download Model
				m.currentStep = 2
			case 1: // Configure Model Path
				m.showInput = true
				m.pathInput.Focus()
				return m, textinput.Blink
			case 2: // Test Model
				if m.currentStep < 3 {
					m.errorMsg = "Please complete previous steps first"
				} else {
					// TODO: Implement model testing
					m.errorMsg = ""
				}
			case 3: // Back to Settings
				return NewSettingsModel(), nil
			}
		case "?":
			m.showHelp = !m.showHelp
		case "esc":
			return NewSettingsModel(), nil
		}
	}

	return m, nil
}

func (m *LocalModelModel) View() string {
	if m.showHelp {
		return m.CenterView(
			styles.DialogBox.Render(
				styles.AppTitle.Render("Local Model Setup Help") + "\n\n" +
					"Steps to setup your local model:\n\n" +
					"1. Download a compatible model (e.g., Llama2)\n" +
					"2. Configure the path to model weights\n" +
					"3. Test the model connection\n\n" +
					"Navigation:\n" +
					"• Up/Down or j/k: Move cursor\n" +
					"• Enter/Space: Select option\n" +
					"• ?: Toggle help\n" +
					"• Esc: Back to settings\n\n" +
					styles.PageFooter.Render("Press ? to close help"),
			),
		)
	}

	if m.showInput {
		return m.CenterView(
			styles.DialogBox.Render(
				styles.AppTitle.Render("Configure Model Path") + "\n\n" +
					styles.MenuOption.Render("Enter the path to your model weights:") + "\n" +
					styles.InputBox.Render(m.pathInput.View()) + "\n\n" +
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

		status := getStepStatus(m.currentStep, i+1)
		icon := getStepIcon(m.currentStep, i+1)

		option := icon + " " + choice + " " + status
		if m.cursor == i {
			menuContent += styles.HighlightedOption.Render(cursor+option) + "\n"
		} else {
			menuContent += styles.MenuOption.Render(cursor+option) + "\n"
		}
	}

	menuView := styles.MenuBox.Render(
		styles.SectionTitle.Render("Setup Steps") + "\n\n" +
			menuContent + "\n\n" +
			styles.PageFooter.Render("? for help • esc to go back"),
	)

	// Right panel - Detailed content
	var detailContent string
	switch m.cursor {
	case 0: // Download Model
		detailContent = styles.MainTitle.Render("Download Compatible Model") + "\n\n" +
			styles.Description.Render(
				"Choose and download one of these models:\n\n"+
					"🤖 Llama 2\n"+
					"  • Size: 7B/13B/70B parameters\n"+
					"  • License: Meta AI Research License\n"+
					"  • URL: huggingface.co/meta-llama\n\n"+
					"🤖 GPT-J\n"+
					"  • Size: 6B parameters\n"+
					"  • License: Apache 2.0\n"+
					"  • URL: huggingface.co/EleutherAI\n\n"+
					"🤖 BLOOM\n"+
					"  • Size: 7B parameters\n"+
					"  • License: OpenRAIL-M\n"+
					"  • URL: huggingface.co/bigscience\n",
			)

	case 1: // Configure Model Path
		detailContent = styles.MainTitle.Render("Model Path Configuration") + "\n\n" +
			styles.Description.Render(
				"Set up the path to your downloaded model:\n\n"+
					"1. Locate your downloaded model files\n"+
					"2. Copy the full path to the weights file\n"+
					"3. Press ENTER to open the path input\n"+
					"4. Paste or type the path\n\n"+
					"Example paths:\n"+
					"• ~/.cache/huggingface/llama2-7b\n"+
					"• ~/models/gpt-j-6B/weights\n"+
					"• /opt/models/bloom-7b1\n",
			)

	case 2: // Test Model
		detailContent = styles.MainTitle.Render("Model Testing") + "\n\n" +
			styles.Description.Render(
				"Verify your model configuration:\n\n"+
					"• Check model file accessibility\n"+
					"• Validate model format\n"+
					"• Test basic inference\n"+
					"• Measure performance\n\n"+
					"Status: "+getTestStatus(m.currentStep),
			)
	}

	if m.errorMsg != "" {
		detailContent += "\n\n" + styles.ErrorText.Render(m.errorMsg)
	}

	// Add progress bar
	progress := float64(m.currentStep-1) / 3.0
	detailContent += "\n\n" + styles.SectionTitle.Render("Setup Progress") + "\n" +
		renderProgressBar(progress, 40)

	detailView := styles.ContentBox.Render(detailContent)

	// Combine views
	return m.SplitView(menuView, detailView)
}

func getStepIcon(currentStep, step int) string {
	if step == 4 { // Back button
		return "←"
	}
	if currentStep > step {
		return "✓"
	} else if currentStep == step {
		return "►"
	}
	return "○"
}

func getTestStatus(currentStep int) string {
	if currentStep < 3 {
		return styles.WarningText.Render("NOT READY")
	}
	return styles.SuccessText.Render("READY TO TEST")
}

func getStepStatus(currentStep, step int) string {
	if currentStep > step {
		return styles.SuccessText.Render("COMPLETED")
	} else if currentStep == step {
		return styles.WarningText.Render("IN PROGRESS")
	}
	return "NOT STARTED"
}

func renderProgressBar(progress float64, width int) string {
	filled := int(progress * float64(width))
	empty := width - filled

	bar := styles.SuccessText.Render(strings.Repeat("█", filled))
	bar += styles.Description.Render(strings.Repeat("░", empty))
	percentage := int(progress * 100)

	return bar + styles.Description.Render(
		"  "+styles.SuccessText.Render(
			"["+styles.HighlightedOption.Render(
				""+string(rune('0'+percentage/10))+string(rune('0'+percentage%10))+"%",
			)+"]",
		),
	)
}
