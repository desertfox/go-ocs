package ocs

import "github.com/charmbracelet/lipgloss"

var (
	style  = lipgloss.NewStyle().PaddingLeft(2)
	red    = style.Foreground(lipgloss.Color("9"))
	green  = style.Foreground(lipgloss.Color("10"))
	yellow = style.Foreground(lipgloss.Color("11"))
)

func generateStyleForeground(color string) lipgloss.Style {
	return style.Foreground(lipgloss.Color(color))
}

func generateStyleBackground(color string) lipgloss.Style {
	return style.Background(lipgloss.Color(color))
}
