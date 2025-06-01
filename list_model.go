package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func newDelegate() list.DefaultDelegate {
	// keymap := newKeymap()
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		// switch msg := msg.(type) {
		// case tea.KeyMsg:
		// 	switch {
		// 	}
		// }
		return nil
	}

	return d
}

func newListModel(title string, items []list.Item, width, height int) list.Model {
	m := list.New(items, newDelegate(), width, height)
	m.Title = title
	m.SetShowHelp(false)
	return m
}
