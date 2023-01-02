package ui

import (
    "fmt"

    "github.com/WilliamKSilva/azure-pullrequests-cli/cmd"
    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

func (i item) Title() string { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func fetchProjects(token string, organization string) (string, error) {
    url := "https://dev.azure.com/{organization}/_apis/projects?api-version=7.0"  
    body, err := cmd.GetRequest(url, token) 

    if err != nil {
        return "", err
    }

    return body, nil
}

func InitialModel() model {
    items := []list.Item{
        item{title: "Test", desc: "Test"},
    }

    li := list.New(items, list.NewDefaultDelegate(), 0, 0)

    m := model{
        inputPatToken: inputModel{
            newInput("Enter your PAT (Personal Access Token)"),
            "",
        }, 
        inputOrganization: inputModel{
            newInput("Enter your organization name"),
            "",
        },
        list: li,
        err: nil,
        mode: inputOrganization,
    }

    return m
}

func newInput (placeholder string) (textinput.Model) {
    var t textinput.Model

    t = textinput.New()
    t.CharLimit = 32

    t.Placeholder = placeholder
    t.Focus()

    t.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

    return t
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
                m.inputOrganization.data = msg.String()
                m.mode = inputPatToken
            case inputPatToken:
               m.inputPatToken.data  = msg.String()
            }
        }
    }

    switch m.mode {
    case inputOrganization:
        m.inputOrganization.input, cmd = m.inputOrganization.input.Update(msg)
    case inputPatToken:
        m.inputOrganization.input, cmd = m.inputOrganization.input.Update(msg)
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
    }

    return ""
}
