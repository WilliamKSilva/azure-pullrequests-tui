package cmd

import (
	"fmt"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
  textInput textinput.Model
  err error
}

type (
  errMsg error
)

type Value struct {
  Url string `json:"url"`
  Status string `json:"status"`
  Title string `json:"title"`
}

type ResponseBody struct {
  Count int `json:"count"`
  Value []Value `json:"value"`
}

func InitialModel() model {
  ti := textinput.New()
  ti.Placeholder = "PAT (Personal Access Token)"
  ti.Focus()
  ti.CharLimit = 150
  ti.Width = 25

  return model{
    textInput: ti,
    err: nil,
  }
}

func (m model) Init() tea.Cmd {
  return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

  switch msg := msg.(type) {
    case tea.KeyMsg:
      switch msg.Type {
        case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
          return m, tea.Quit
      }

    case errMsg:
      m.err = msg
      return m, nil
  }

  m.textInput, cmd = m.textInput.Update(msg)
  return m, cmd
}

func (m model) View() string {
  return fmt.Sprintf(
    "Enter your Personal Access Token from Azure DevOps\n\n%s\n\n%s",
    m.textInput.View(),
    "esc exits the terminal",
  ) + "\n"
}
