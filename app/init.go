package app

import (
	"encoding/json"
	"errors"
	"os"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"

	"todo/styles"
	"todo/types"
)

func InitModel() Model {

	ti := textinput.New()

	ti.Prompt = ""
	ti.Placeholder = "Edit task"
	ti.SetWidth(30)

	s := ti.Styles()

	s.Cursor = textinput.CursorStyle{
		Color: styles.SecondaryColor,
		Blink: true,
	}

	ti.SetStyles(s)

	var tasks []types.Task

	_, err := os.Stat("Tasks.json")

	if errors.Is(err, os.ErrNotExist) {
		tasks = []types.Task{}

	} else {
		data, err := os.ReadFile("Tasks.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &tasks)
	}

	return Model{
		Tasks:    tasks,
		Editor:   ti,
		Quitting: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
