package app

type Command interface {
	Execute(m *Model)
	Undo(m *Model)
}