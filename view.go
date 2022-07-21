package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var (
	style  = lipgloss.NewStyle().PaddingLeft(2)
	red    = style.Copy().Foreground(lipgloss.Color("9"))
	green  = style.Copy().Foreground(lipgloss.Color("10"))
	yellow = style.Copy().Foreground(lipgloss.Color("11"))
)

func generateStyleForeground(color string) lipgloss.Style {
	return style.Copy().Foreground(lipgloss.Color(color))
}

func generateStyleBackground(color string) lipgloss.Style {
	return style.Copy().Background(lipgloss.Color(color))
}

func printStatus(c *config) {
	selectedColor := strconv.Itoa(25 + c.Selected*20)
	selectedString := generateStyleForeground(selectedColor).Render(fmt.Sprintf("Selected: %v ", c.Selected))

	colorBar := generateStyleBackground(selectedColor).Render(strings.Repeat(" ", 10))
	fmt.Println(selectedString + colorBar + "\n")

	var colorIndex = 25
	for i, v := range c.Hosts {
		hostString := generateStyleForeground(strconv.Itoa(colorIndex)).Render(fmt.Sprintf("Index:%v, Created:%v, Server:%v", i, v.Created.Format(time.RFC1123), v.Server))
		fmt.Println(hostString)
		colorIndex += 20
	}
}

func printNewHost(c *config, newHost host) {
	selectedColor := lipgloss.Color(strconv.Itoa(25 + c.Selected*20))
	fmt.Println(style.PaddingLeft(2).Foreground(selectedColor).Render(fmt.Sprintf("AddHost: %v\n", newHost.Server)))
}

func printUpdateHost(updateHost host) {
	fmt.Println(green.Copy().PaddingLeft(2).Render(fmt.Sprintf("updateHost: %v\n", updateHost.Server)))
}

func printServerExists(server string) {
	serverExistsString := yellow.Copy().PaddingLeft(2).Render(fmt.Sprintf("serverExists: %v", server))
	fmt.Println(serverExistsString)
}

func printPuneHost(pruneHost host) {
	pruneHostString := red.Copy().PaddingLeft(2).Render(fmt.Sprintf("pruneHost: %v\n", pruneHost.Server))
	fmt.Println(pruneHostString)
}
