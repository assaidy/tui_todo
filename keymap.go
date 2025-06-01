package main

import "github.com/charmbracelet/bubbles/key"

type Keymap struct {
	nextPane key.Binding
	prevPane key.Binding
	next     key.Binding
	prev     key.Binding
	quit     key.Binding
	delete   key.Binding
	filter   key.Binding
}

func newKeymap() *Keymap {
	return &Keymap{
		quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		nextPane: key.NewBinding(
			key.WithKeys("l"),
			key.WithHelp("l", "navigate next pane"),
		),
		prevPane: key.NewBinding(
			key.WithKeys("h"),
			key.WithHelp("h", "navigate previous pane"),
		),
		next: key.NewBinding(
			key.WithKeys("j"),
			key.WithHelp("j", "next"),
		),
		prev: key.NewBinding(
			key.WithKeys("k"),
			key.WithHelp("k", "previous"),
		),
		delete: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "delete"),
		),
		filter: key.NewBinding(
			key.WithKeys("f"),
			key.WithHelp("f", "filter"),
		),
	}
}
