package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TodoStatus int

const (
	TodoStatus_Todo TodoStatus = iota
	TodoStatus_InProgress
	TodoStatus_Done
)

type Todo struct {
	title     string
	createdAt time.Time
	status    TodoStatus
}

func NewTodo(title string) Todo {
	return Todo{
		title:     title,
		createdAt: time.Now(),
		status:    TodoStatus_Todo,
	}
}

func main() {
	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running tea program: %+v", err)
		os.Exit(1)
	}
}
