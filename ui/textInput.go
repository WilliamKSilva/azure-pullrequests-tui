package ui

import (
	"fmt"

	"github.com/WilliamKSilva/azure-pullrequests-cli/cmd"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
  textInput textinput.Model
  list list.Model
  mode string
  err error
}

type (
  errMsg error
)


var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
  title, desc string
}

func (i item) Title() string { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func fetchProjects(token string) (string, error) {
  url := "https://dev.azure.com/{organization}/_apis/projects?api-version=7.0"  
  body, err := cmd.GetRequest(url, token) 

  if err != nil {
    return "", err
  }

  return body, nil
}

func InitialModel() model {
  ti := textinput.New()
  ti.Placeholder = "PAT (Personal Access Token)"
  ti.Focus()
  ti.CharLimit = 150
  ti.Width = 25

  items := []list.Item{
    item{title: "Test", desc: "Test"},
  }

  li := list.New(items, list.NewDefaultDelegate(), 0, 0)

  return model{
    textInput: ti,
    list: li,
    mode: "ti",
    err: nil,
  }
}

func (m model) Init() tea.Cmd {
  return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd
  var patToken string

  switch msg := msg.(type) {
    case tea.KeyMsg:
      switch msg.Type {
        case tea.KeyCtrlC, tea.KeyEsc:
          return m, tea.Quit
        case tea.KeyEnter:
          m.mode = "li"
          patToken = msg.String()
      }

    case errMsg:
      m.err = msg
      return m, nil
  }

  if (m.mode == "ti") {
    m.textInput, cmd = m.textInput.Update(msg)
  } 

  _, err := fetchProjects(patToken)

  if err != nil {
    return m, tea.Quit
  }

  return m, cmd
}

func (m model) View() string {
  if m.mode == "ti" {
    return fmt.Sprintf(
      "Enter your Personal Access Token from Azure DevOps\n\n%s\n\n%s",
      m.textInput.View(),
      "esc exits the terminal",
    ) + "\n"
  }

  return docStyle.Render(m.list.View())
}
