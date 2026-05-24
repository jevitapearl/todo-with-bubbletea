package app

type EditTaskCommand struct {
	Index int

	OldTitle string
	NewTitle string
}

func NewEditTaskCommand(
	index int,
	newTitle string,
) *EditTaskCommand {

	return &EditTaskCommand{
		Index:    index,
		NewTitle: newTitle,
	}
}

func (c *EditTaskCommand) Execute(m *Model) {

	if len(m.Tasks) == 0 {
		return
	}

	c.OldTitle = m.Tasks[c.Index].Title

	m.Tasks[c.Index].Title = c.NewTitle
}

func (c *EditTaskCommand) Undo(m *Model) {

	if len(m.Tasks) == 0 {
		return
	}

	m.Tasks[c.Index].Title = c.OldTitle
}