package ui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
    p := tea.NewProgram(InitialModel())
    if _, err := p.Run(); err != nil {
      log.Fatal(err)
    }
}
