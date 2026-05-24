package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"todo/app"
)

func main() {

	p := tea.NewProgram(app.InitModel())

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
