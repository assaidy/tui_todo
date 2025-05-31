package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	ColorNormal = "#4E4E4E"
	ColorActive = "#FFAF00"
)

var (
	LipglossColorNormal = lipgloss.Color(ColorNormal)
	LipglossColorActive = lipgloss.Color(ColorActive)
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

type Pane int

const (
	Pane_Todo Pane = iota
	Pane_InProgress
	Pane_Done
)

type Model struct {
	width, height         int
	paneWidth, paneHeight int
	todos                 map[TodoStatus][]Todo
	focusedPane           Pane
}

func newModel() Model {
	return Model{
		todos:       make(map[TodoStatus][]Todo),
		focusedPane: Pane_Todo,
	}
}

func (me Model) Init() tea.Cmd { return nil }

func (me Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		me.updateSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return me, tea.Quit
		case "h":
			me.focusLeftPane()
		case "l":
			me.focusRightPane()
		}
	}
	return me, nil
}

func (me *Model) updateSize(newW, newH int) {
	me.width, me.height = newW, newH
	me.paneWidth = me.width/3 - 2
	me.paneHeight = me.height - 2
}

func (me *Model) focusLeftPane() {
	if me.focusedPane > 0 {
		me.focusedPane -= 1
	}

}
func (me *Model) focusRightPane() {
	if me.focusedPane < 2 {
		me.focusedPane += 1
	}
}

func (me Model) View() string {
	todoPane := me.renderTodoPane()
	inProgresPane := me.renderInProgresPane()
	donePane := me.renderDonePane()
	return lipgloss.JoinHorizontal(lipgloss.Center, todoPane, inProgresPane, donePane)
}

func (me *Model) renderTodoPane() string {
	style := lipgloss.NewStyle().
		Width(me.paneWidth).
		Height(me.paneHeight).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(LipglossColorNormal)
	titleStyle := lipgloss.NewStyle().
		Width(me.paneWidth).
		Bold(true).
		Foreground(LipglossColorNormal).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(LipglossColorNormal)
	if me.focusedPane == Pane_Todo {
		style = style.Foreground(LipglossColorActive).BorderForeground(LipglossColorActive)
		titleStyle = titleStyle.Foreground(LipglossColorActive).BorderForeground(LipglossColorActive)
	}
	return style.Render(titleStyle.Render("TODO"))
}

func (me *Model) renderInProgresPane() string {
	paneStyle := lipgloss.NewStyle().
		Width(me.paneWidth).
		Height(me.paneHeight).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(LipglossColorNormal)
	titleStyle := lipgloss.NewStyle().
		Width(me.paneWidth).
		Bold(true).
		Foreground(LipglossColorNormal).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(LipglossColorNormal)
	if me.focusedPane == Pane_InProgress {
		paneStyle = paneStyle.Foreground(LipglossColorActive).BorderForeground(LipglossColorActive)
		titleStyle = titleStyle.Foreground(LipglossColorActive).BorderForeground(LipglossColorActive)
	}
	return paneStyle.Render(titleStyle.Render("IN PROGRESS"))
}

func (me *Model) renderDonePane() string {
	style := lipgloss.NewStyle().
		Width(me.width - me.paneWidth*2 - 6). // make last pane fill the remaining width
		Height(me.paneHeight).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(LipglossColorNormal)
	titleStyle := lipgloss.NewStyle().
		Width(me.width-me.paneWidth*2-6). // make last pane fill the remaining width
		Bold(true).
		Foreground(LipglossColorNormal).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(LipglossColorNormal)
	if me.focusedPane == Pane_Done {
		style = style.Foreground(LipglossColorActive).BorderForeground(LipglossColorActive)
		titleStyle = titleStyle.Foreground(LipglossColorActive).BorderForeground(LipglossColorActive)
	}
	return style.Render(titleStyle.Render("DONE"))
}

func main() {
	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running tea program: %+v", err)
		os.Exit(1)
	}
}
