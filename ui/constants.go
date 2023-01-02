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
    listProjects Mode = "list"
)

type inputData struct {
    inputPatToken string
    inputOrganization string
}

type inputModel struct {
    input textinput.Model
    data string
}


type model struct {
    list list.Model
    inputPatToken inputModel
    inputOrganization inputModel 
    mode Mode
    err error
}

type (
    errMsg error
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
    title, desc string
}
