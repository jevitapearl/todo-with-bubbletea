package app

type ToggleTaskCommand struct {
	Index int
}

func NewToggleTaskCommand(index int) *ToggleTaskCommand {
	return &ToggleTaskCommand{
		Index: index,
	}
}

func (c *ToggleTaskCommand) Execute(m *Model) {

	if len(m.Tasks) == 0 {
		return
	}

	m.Tasks[c.Index].Status =
		!m.Tasks[c.Index].Status
}

func (c *ToggleTaskCommand) Undo(m *Model) {

	if len(m.Tasks) == 0 {
		return
	}

	m.Tasks[c.Index].Status =
		!m.Tasks[c.Index].Status
}