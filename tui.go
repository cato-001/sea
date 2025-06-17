package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	DuckDuckGoModel struct {
		Links []Weblink
		List list.Model
	}
)

func NewDuckDuckGoModel(links []Weblink) *DuckDuckGoModel {
	items := make([]list.Item, len(links))
	for index, link := range links {
		items[index] = link
	}
	return &DuckDuckGoModel {
		Links: links,
		List: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
}

func (m *DuckDuckGoModel) Init() tea.Cmd {
	return nil
}

func (m *DuckDuckGoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			cmd = tea.Quit
		case "j":
			m.List.CursorDown()
		case "k":
			m.List.CursorUp()
		case "enter":
			index := m.List.Cursor()
			SelectedLink = m.Links[index].Link()
			cmd = tea.Batch(tea.ClearScreen, tea.Quit)
		}
	case tea.WindowSizeMsg:
		m.List.SetSize(msg.Width, msg.Height / 2)
	}

	return m, cmd
}

func (m *DuckDuckGoModel) View() string {
	return m.List.View()
}
