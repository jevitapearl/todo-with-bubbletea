package app

import (
	"encoding/json"
	"os"

	tea "charm.land/bubbletea/v2"
)

func executeCommand(m *Model, cmd Command) {

	cmd.Execute(m)

	m.UndoStack = append(m.UndoStack, cmd)

	m.RedoStack = nil
}

func undo(m *Model) {

	if len(m.UndoStack) == 0 {
		return
	}

	cmd := m.UndoStack[len(m.UndoStack)-1]

	m.UndoStack = m.UndoStack[:len(m.UndoStack)-1]

	cmd.Undo(m)

	m.RedoStack = append(m.RedoStack, cmd)
}

func redo(m *Model) {

	if len(m.RedoStack) == 0 {
		return
	}

	cmd := m.RedoStack[len(m.RedoStack)-1]

	m.RedoStack = m.RedoStack[:len(m.RedoStack)-1]

	cmd.Execute(m)

	m.UndoStack = append(m.UndoStack, cmd)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.Editing {

		var cmd tea.Cmd

		m.Editor, cmd = m.Editor.Update(msg)

		switch msg := msg.(type) {

		case tea.KeyPressMsg:

			switch msg.String() {

			case "enter":

				executeCommand(
					&m,
					NewEditTaskCommand(
						m.Cursor,
						m.Editor.Value(),
					),
				)

				m.Editing = false

			case "esc":
				m.Editing = false
			}
		}

		return m, cmd
	}

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "q", "ctrl+c":
			m.Quitting = true
			jsonData, err := json.MarshalIndent(m.Tasks, "", "	")
			if err != nil {
				panic(err)
			}

			err = os.WriteFile("Tasks.json", jsonData, 0644)
			if err != nil {
				panic(err)
			}


			return m, tea.Quit

		case "up", "k":

			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down", "j":

			if m.Cursor < len(m.Tasks)-1 {
				m.Cursor++
			}

		case "ctrl+n":

			executeCommand(
				&m,
				NewAddTaskCommand(),
			)

			m.Editor.SetValue("")
			m.Editor.Focus()

			m.Editing = true

		case "ctrl+x":

			executeCommand(
				&m,
				NewDeleteTaskCommand(),
			)

		case "space", "enter":

			if len(m.Tasks) > 0 {

				executeCommand(
					&m,
					NewToggleTaskCommand(m.Cursor),
				)
			}

		case "ctrl+z":
			undo(&m)

		case "ctrl+y":
			redo(&m)

		case "e":

			if len(m.Tasks) > 0 {

				m.Editing = true

				m.Editor.SetValue(
					m.Tasks[m.Cursor].Title,
				)

				m.Editor.Focus()
			}
		}
	}

	return m, nil
}
