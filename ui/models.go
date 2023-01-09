package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type Mode string

const (
    inputPatToken Mode = "inputPatToken"
    inputOrganization Mode = "inputOrganization"
    listProjects Mode = "listProjects"
    listPullRequests Mode = "listPullRequests"
    loading Mode = "loading"
)

type inputModel struct {
    input textinput.Model
    data string
}

type listModel struct {
    list list.Model
    data string
}

type (
    errMsg error
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var statusMessageStyle = lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("205")).Background(lipgloss.NoColor{}).Render

type item struct {
    title, desc string
}
