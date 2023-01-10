package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (i item) FilterValue() string { return i.title }
func (i item) Title() string { return i.title }
func (i item) Description() string { return i.desc}

type model struct {
    listProjects listModel
    listPullRequests listModel 
    inputPatToken inputModel
    inputOrganization inputModel 
    mode Mode
    err error
}

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
        listProjects: listModel{
            list: list.New(nil, list.NewDefaultDelegate(), 0, 0),
            data: "",
        },
        listPullRequests: listModel{
            list: list.New(nil, list.NewDefaultDelegate(), 0, 0),
            data: "",
        },
        mode: inputOrganization,
    }

    m.listProjects.list.Title = "Select an Azure DevOps project"
    m.listPullRequests.list.Title = "Azure DevOps project Pull Requests"

    return m
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
                projects, err := getProjects(m.inputPatToken.data, m.inputOrganization.data)

                if err != nil {
                    return m, tea.Quit
                }

                var projectsItems []list.Item

                for _, s := range projects.Value {
                    newProject := item{title: s.Name, desc: s.Description }
                    projectsItems = append(projectsItems, newProject)
                }

                m.listProjects.list.SetItems(projectsItems)


                m.mode = listProjects 
            case listProjects:
                selectedItem := m.listProjects.list.SelectedItem()
                m.listProjects.data = selectedItem.FilterValue() 

                pullRequests, err := getPullRequests(m.inputPatToken.data, m.inputOrganization.data, m.listProjects.data)

                if err != nil {
                    return m, tea.Quit
                }

                var pullRequestsItems []list.Item

                for _, s := range pullRequests.Value {
                    desc := fmt.Sprintf("Repo: %s/Status: %s", s.Repository.Name, s.Status)
                    newPullRequest := item{title: s.Title, desc: desc }
                    pullRequestsItems = append(pullRequestsItems, newPullRequest)
                }

                m.listPullRequests.list.SetItems(pullRequestsItems)

                m.mode = listPullRequests
            }
        case "r":
            switch m.mode {
            case listPullRequests:
                pullRequests, err := getPullRequests(m.inputPatToken.data, m.inputOrganization.data, m.listProjects.data)

                if err != nil {
                    return m, tea.Quit
                }

                var pullRequestsItems []list.Item

                for _, s := range pullRequests.Value {
                    desc := fmt.Sprintf("Repo: %s/Status: %s", s.Repository.Name, s.Status)
                    newPullRequest := item{title: s.Title, desc: desc }
                    pullRequestsItems = append(pullRequestsItems, newPullRequest)
                }

                m.listPullRequests.list.SetItems(pullRequestsItems) 
                cmd := m.listPullRequests.list.NewStatusMessage(statusMessageStyle("Refresh"))
                return m, cmd
            }
        }
    case tea.WindowSizeMsg:
        h, v := docStyle.GetFrameSize()
        m.listProjects.list.SetSize(msg.Width-h, msg.Height-v)
        m.listPullRequests.list.SetSize(msg.Width-h, msg.Height-v)
    }

    switch m.mode {
    case inputOrganization:
        m.inputOrganization.input, cmd = m.inputOrganization.input.Update(msg)
    case inputPatToken:
        m.inputPatToken.input, cmd = m.inputPatToken.input.Update(msg)
    case listProjects:
        m.listProjects.list, cmd = m.listProjects.list.Update(msg)
    case listPullRequests:
        m.listPullRequests.list, cmd = m.listPullRequests.list.Update(msg)
    }

    return m, cmd
}

func (m model) View() string {
    switch m.mode {
    case inputOrganization:
        return fmt.Sprintf(
            "\n\nEnter your Azure Devops organization name\n\n%s\n\n%s",
            m.inputOrganization.input.View(),
            "esc exits the terminal",
        ) + "\n"
    case inputPatToken:
        return fmt.Sprintf(
            "\n\nEnter your Personal Access Token from Azure DevOps\n\n%s\n\n%s",
            m.inputPatToken.input.View(),
            "esc exits the terminal",
        ) + "\n"
    case listProjects:
        return docStyle.Render(m.listProjects.list.View())
    case listPullRequests:
        return docStyle.Render(m.listPullRequests.list.View())
    }

    return ""
}

