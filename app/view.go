package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Width(70)

	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true).
			MarginBottom(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginTop(1)

	selectedTaskStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("229")).
				Background(lipgloss.Color("")).
				Bold(true).
				PaddingLeft(1)

	completedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Strikethrough(true)

	pendingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))
)

func (m Model) View() tea.View {

	var b strings.Builder

	header := headerStyle.Render("TODO APP")

	b.WriteString(header)
	b.WriteString("\n")

	if m.Quitting {
		b.WriteString("\nSaving tasks and quitting...\n\n\n")
	}

	if len(m.Tasks) == 0 {

		empty := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Render("No tasks yet. Press Ctrl+N to create one.")

		b.WriteString(empty)
		b.WriteString("\n")

	} else {

		for i, task := range m.Tasks {

			cursor := " "

			if i == m.Cursor {
				cursor = "▸"
			}

			check := "○"

			title := pendingStyle.Render(task.Title)

			if task.Status {
				check = "✔"
				title = completedStyle.Render(task.Title)
			}

			line := fmt.Sprintf(
				"%s %s %s",
				cursor,
				check,
				title,
			)

			if i == m.Cursor {
				line = selectedTaskStyle.Render(line)
			}

			if m.Editing && i == m.Cursor {

				line = selectedTaskStyle.Render(
					fmt.Sprintf(
						"%s %s %s",
						cursor,
						check,
						m.Editor.View(),
					),
				)
			}

			b.WriteString(line)
			b.WriteString("\n")
		}
	}

	help := helpStyle.Render(
		"• Ctrl+N Add • Ctrl+X Delete • e Edit \n• Ctrl+Z Undo • Ctrl+Y Redo • q Quit",
	)

	b.WriteString(help)

	content := appStyle.Render(b.String())

	v := tea.NewView(content)
	v.WindowTitle = "TODO App"
	v.AltScreen = true

	return v
}
