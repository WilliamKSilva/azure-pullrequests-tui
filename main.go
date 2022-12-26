package main

import (
	"log"

	cmd "github.com/WilliamKSilva/azure-pullrequests-cli/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
    p := tea.NewProgram(cmd.InitialModel())
    if _, err := p.Run(); err != nil {
      log.Fatal(err)
    }
}
