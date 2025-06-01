package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Pane int

const (
	Pane_Todo Pane = iota
	Pane_InProgress
	Pane_Done
)

type Model struct {
	width, height         int
	paneWidth, paneHeight int
	todoList              list.Model
	inProgressList        list.Model
	doneList              list.Model
	focusedPane           Pane
}

func newModel() Model {
	m := Model{focusedPane: Pane_Todo}
	m.todoList = newListModel("TODO", []list.Item{
		NewTodo("todo 1"),
		NewTodo("todo 2"),
		NewTodo("todo 3"),
		NewTodo("todo 4"),
		NewTodo("todo 5"),
		NewTodo("todo 6"),
		NewTodo("todo 7"),
		NewTodo("todo 8"),
		NewTodo("todo 9"),
		NewTodo("todo 10"),
	}, m.width, m.height)
	m.inProgressList = newListModel("IN PROGRESS", []list.Item{}, m.width, m.height)
	m.doneList = newListModel("DONE", []list.Item{}, m.width, m.height)
	return m
}

func (me Model) Init() tea.Cmd { return nil }

func (me Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keymap := newKeymap()

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		me.resize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.quit):
			return me, tea.Quit
		case key.Matches(msg, keymap.nextPane):
			me.focusNextPane()
		case key.Matches(msg, keymap.prevPane):
			me.focusPrevPane()
		case key.Matches(msg, keymap.next):
			me.nextItem()
		case key.Matches(msg, keymap.prev):
			me.prevItem()
		case key.Matches(msg, keymap.delete):
			return me, me.delete(msg)
		}
	}

	return me, nil
}

func (me *Model) delete(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	switch me.focusedPane {
	case Pane_Todo:
		me.todoList.RemoveItem(me.todoList.Index())
		me.todoList, cmd = me.todoList.Update(msg)
	case Pane_InProgress:
		me.inProgressList.RemoveItem(me.inProgressList.Index())
		me.inProgressList, cmd = me.inProgressList.Update(msg)
	case Pane_Done:
		me.doneList.RemoveItem(me.doneList.Index())
		me.doneList, cmd = me.doneList.Update(msg)
	}
	return cmd
}

func (me *Model) nextItem() {
	switch me.focusedPane {
	case Pane_Todo:
		me.todoList.CursorDown()
	case Pane_InProgress:
		me.inProgressList.CursorDown()
	case Pane_Done:
		me.doneList.CursorDown()
	}
}

func (me *Model) prevItem() {
	switch me.focusedPane {
	case Pane_Todo:
		me.todoList.CursorUp()
	case Pane_InProgress:
		me.inProgressList.CursorUp()
	case Pane_Done:
		me.doneList.CursorUp()
	}
}

func (me *Model) resize(newW, newH int) {
	me.width, me.height = newW, newH
	me.paneWidth = me.width/3 - 2
	me.paneHeight = me.height - 2
	me.todoList.SetSize(me.paneWidth, me.paneHeight)
	me.inProgressList.SetSize(me.paneWidth, me.paneHeight)
	me.doneList.SetSize(me.paneWidth, me.paneHeight)
}

func (me *Model) focusPrevPane() {
	if me.focusedPane > 0 {
		me.focusedPane -= 1
	}

}
func (me *Model) focusNextPane() {
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
		BorderForeground(ColorInactiveBorderFg)
	if me.focusedPane == Pane_Todo {
		style = style.BorderForeground(ColorActiveBorderFg)
	}
	return style.Render(me.todoList.View())
}

func (me *Model) renderInProgresPane() string {
	paneStyle := lipgloss.NewStyle().
		Width(me.paneWidth).
		Height(me.paneHeight).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorInactiveBorderFg)
	if me.focusedPane == Pane_InProgress {
		paneStyle = paneStyle.BorderForeground(ColorActiveBorderFg)
	}
	return paneStyle.Render(me.inProgressList.View())
}

func (me *Model) renderDonePane() string {
	style := lipgloss.NewStyle().
		Width(me.paneWidth).
		Height(me.paneHeight).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorInactiveBorderFg)
	if me.focusedPane == Pane_Done {
		style = style.BorderForeground(ColorActiveBorderFg)
	}
	return style.Render(me.doneList.View())
}
