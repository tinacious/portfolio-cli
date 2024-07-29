package lib

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// TinaciousDesignTheme custom huh Theme for the form
func TinaciousDesignTheme() *huh.Theme {
	t := huh.ThemeBase()

	var (
		background = lipgloss.Color("#1d1d26")
		foreground = lipgloss.Color("#ffffff")
		grey       = lipgloss.Color("#b3b3d4")
		green      = lipgloss.Color("#00d364")
		blue       = lipgloss.Color("#00bfff")
		purple     = lipgloss.Color("#cc66ff")
		pink       = lipgloss.Color("#ff3399")
		turquoise  = lipgloss.Color("#00ced1")
	)

	t.Focused.Base = t.Focused.Base.BorderForeground(blue)
	t.Focused.Title = t.Focused.Title.Foreground(pink)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(pink)
	t.Focused.Description = t.Focused.Description.Foreground(grey)
	t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(pink)
	t.Focused.Directory = t.Focused.Directory.Foreground(pink)
	t.Focused.File = t.Focused.File.Foreground(foreground)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(pink)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(purple)
	t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(purple)
	t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(purple)
	t.Focused.Option = t.Focused.Option.Foreground(foreground)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(purple)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(green)
	t.Focused.SelectedPrefix = t.Focused.SelectedPrefix.Foreground(green)
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(foreground)
	t.Focused.UnselectedPrefix = t.Focused.UnselectedPrefix.Foreground(grey)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(purple).Background(pink).Bold(true)
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(foreground).Background(background)

	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(turquoise)
	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(grey)
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(purple)

	t.Blurred = t.Focused
	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}
