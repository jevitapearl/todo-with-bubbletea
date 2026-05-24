package app

import "todo/types"

type DeleteTaskCommand struct {
	Task  types.Task
	Index int
}

func NewDeleteTaskCommand() *DeleteTaskCommand {
	return &DeleteTaskCommand{}
}

func (c *DeleteTaskCommand) Execute(m *Model) {

	if len(m.Tasks) == 0 {
		return
	}

	c.Task = m.Tasks[m.Cursor]
	c.Index = m.Cursor

	m.Tasks = append(
		m.Tasks[:m.Cursor],
		m.Tasks[m.Cursor+1:]...,
	)

	if m.Cursor > len(m.Tasks)-1 {
		m.Cursor = len(m.Tasks) - 1
	}
}

func (c *DeleteTaskCommand) Undo(m *Model) {

	m.Tasks = append(
		m.Tasks[:c.Index],
		append([]types.Task{c.Task}, m.Tasks[c.Index:]...)...,
	)

	m.Cursor = c.Index
}
