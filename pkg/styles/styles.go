package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	primaryColor   = lipgloss.Color("#7D56F4")
	secondaryColor = lipgloss.Color("#5B5B5B")
	bgColor        = lipgloss.Color("#1A1B26")
	textColor      = lipgloss.Color("#FFFFFF")
	errorColor     = lipgloss.Color("#FF0000")
	successColor   = lipgloss.Color("#00FF00")
	warningColor   = lipgloss.Color("#FFA500")
	accentColor    = lipgloss.Color("#FF79C6")

	// Layout styles
	DocStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Background(bgColor)

	SplitLeft = lipgloss.NewStyle().
			Width(30).
			Height(30).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1).
			MarginRight(2)

	SplitRight = lipgloss.NewStyle().
			Width(90).
			Height(30).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1)

	// Title styles
	AppTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			Background(bgColor).
			Padding(1, 4).
			MarginBottom(2).
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(primaryColor).
			Align(lipgloss.Center)

	MainTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			Background(bgColor).
			Padding(1, 4).
			MarginTop(1).
			MarginBottom(2).
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(primaryColor).
			Align(lipgloss.Center)

	SubTitle = lipgloss.NewStyle().
			Foreground(textColor).
			Background(bgColor).
			Padding(0, 2).
			MarginTop(1).
			MarginBottom(2).
			Align(lipgloss.Center)

	// Menu styles
	MenuBox = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(primaryColor).
		Padding(1).
		MarginRight(2)

	MenuOption = lipgloss.NewStyle().
			Foreground(textColor).
			Padding(0, 2).
			MarginTop(0).
			MarginBottom(0)

	HighlightedOption = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true).
				Padding(0, 2).
				MarginTop(0).
				MarginBottom(0)

	// Content styles
	ContentBox = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1).
			MarginLeft(2)

	// Footer styles
	PageFooter = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#999999")).
			Padding(1, 2).
			MarginTop(1).
			Align(lipgloss.Center)

	// Text styles
	ErrorText = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)

	SuccessText = lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true)

	WarningText = lipgloss.NewStyle().
			Foreground(warningColor).
			Bold(true)

	// Input styles
	InputBox = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1, 2)

	// Dialog styles
	DialogBox = lipgloss.NewStyle().
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(primaryColor).
			Padding(1, 2).
			MarginTop(1).
			MarginBottom(1)

	// Section styles
	SectionTitle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			MarginBottom(1).
			MarginTop(1)

	// Description styles
	Description = lipgloss.NewStyle().
			Foreground(textColor).
			MarginTop(1).
			MarginBottom(1)

	// Status styles
	StatusIndicator = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true).
			Padding(0, 1)
)

// UpdateSplitSizes updates the split view sizes based on terminal dimensions
func UpdateSplitSizes(width, height int) {
	leftWidth := width / 4
	rightWidth := width - leftWidth - 4 // Account for margins and borders
	viewHeight := height - 4            // Account for margins and borders

	SplitLeft = SplitLeft.Width(leftWidth).Height(viewHeight)
	SplitRight = SplitRight.Width(rightWidth).Height(viewHeight)

	// Update dependent styles
	MenuBox = MenuBox.Width(leftWidth - 2)        // Account for padding
	ContentBox = ContentBox.Width(rightWidth - 2) // Account for padding
	AppTitle = AppTitle.Width(width - 4)          // Account for margins
	MainTitle = MainTitle.Width(rightWidth - 4)   // Account for margins
	SubTitle = SubTitle.Width(rightWidth - 4)     // Account for margins
	PageFooter = PageFooter.Width(width - 4)      // Account for margins
}
