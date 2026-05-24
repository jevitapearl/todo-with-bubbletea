package app

import (
	"time"

	"github.com/google/uuid"

	"todo/types"
)

type AddTaskCommand struct {
	Task  types.Task
	Index int
}

func NewAddTaskCommand() *AddTaskCommand {
	return &AddTaskCommand{
		Task: types.Task{
			ID:        uuid.NewString(),
			CreatedAt: time.Now(),
		},
	}
}

func (c *AddTaskCommand) Execute(m *Model) {

	m.Tasks = append(m.Tasks, c.Task)

	c.Index = len(m.Tasks) - 1

	m.Cursor = c.Index
}

func (c *AddTaskCommand) Undo(m *Model) {

	if len(m.Tasks) == 0 {
		return
	}

	m.Tasks = append(
		m.Tasks[:c.Index],
		m.Tasks[c.Index+1:]...,
	)

	if m.Cursor > len(m.Tasks)-1 {
		m.Cursor = len(m.Tasks) - 1
	}
}
