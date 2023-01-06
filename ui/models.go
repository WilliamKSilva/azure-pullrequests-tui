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

type item struct {
    title, desc string
}
