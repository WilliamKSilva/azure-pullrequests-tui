package ui

import (
    "fmt"
    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

func (i item) Title() string { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
    list list.Model
    inputPatToken inputModel
    inputOrganization inputModel 
    mode Mode
    err error
}

var projects *Projects

func InitialModel() model {
    m := model{
        inputPatToken: inputModel{
            newInput("Enter your PAT (Personal Access Token)"),
            "",
        }, 
        inputOrganization: inputModel{
            newInput("Enter your organization name"),
            "",
        },
        list: list.New(nil, list.NewDefaultDelegate(), 0, 0),
        err: nil,
        mode: inputOrganization,
    }

    m.list.Title = "Azure DevOps Projects"

    return m
}


func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "esc":
            return m, tea.Quit
        case "enter":
            switch m.mode {
            case inputOrganization:
                m.inputOrganization.data = string(Mode(m.inputOrganization.input.Value())) 
                m.mode = inputPatToken
            case inputPatToken:
                m.inputPatToken.data = string(Mode(m.inputPatToken.input.Value()))


                err := projects.getProjects(m.inputPatToken.data, m.inputOrganization.data)

                if err != nil {
                    return m, tea.Quit
                }

                var projectsItems []list.Item

                for _, s := range projects.Value {
                    newProject := item{title: s.Name, desc: s.Description}
                    projectsItems = append(projectsItems, newProject)
                }

                m.list = list.New(projectsItems, list.NewDefaultDelegate(), 0, 0)
                m.list.Title = "Select a Azure DevOps project to track"
                m.list.SetSize(100, 50)

                m.mode = listProjects
            }
        }
    }

    switch m.mode {
    case inputOrganization:
        m.inputOrganization.input, cmd = m.inputOrganization.input.Update(msg)
    case inputPatToken:
        m.inputPatToken.input, cmd = m.inputPatToken.input.Update(msg)
    case listProjects:
        m.list, cmd = m.list.Update(msg)
    }
    return m, cmd
}

func (m model) View() string {
    switch m.mode {
    case inputOrganization:
        return fmt.Sprintf(
            "Enter your Azure Devops organization name\n\n%s\n\n%s",
            m.inputOrganization.input.View(),
            "esc exits the terminal",
        ) + "\n"
    case inputPatToken:
        return fmt.Sprintf(
            "Enter your Personal Access Token from Azure DevOps\n\n%s\n\n%s",
            m.inputPatToken.input.View(),
            "esc exits the terminal",
        ) + "\n"
    case listProjects:
        return docStyle.Render(m.list.View())
    }

    return ""
}

func newInput (placeholder string) (textinput.Model) {
    var t textinput.Model

    t = textinput.New()
    t.CharLimit = 70

    t.Placeholder = placeholder
    t.Focus()

    t.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

    return t
}
