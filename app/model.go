package app

import (
	"charm.land/bubbles/v2/textinput"

	"todo/types"
)

type Model struct {
	Cursor int
	Tasks  []types.Task

	UndoStack []Command
	RedoStack []Command

	Editing bool
	Editor  textinput.Model

	Quitting bool
}