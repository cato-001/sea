package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var SelectedLink = ""

func main() {
	links, err := Query(strings.Join(os.Args[1:], " "))
	if err != nil {
		fmt.Println("error while sending query:\n", err)
		return
	}

	if len(links) == 0 {
		fmt.Println("no results found")
		return
	}

	model := NewDuckDuckGoModel(links)
	program := tea.NewProgram(model, tea.WithAltScreen(), tea.WithOutput(os.Stderr))
	if _, err := program.Run(); err != nil {
		fmt.Println("error while showing tui:\n", err)
		return
	}

	fmt.Println(SelectedLink)
}

