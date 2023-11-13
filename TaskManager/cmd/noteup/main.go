package main

import (
	"flag"
	"github.com/pranav77/noteup"
	"fmt"
	"os"
	"github.com/charmbracelet/lipgloss"
)
const (
	noteUp = ".noteup.json"
)

func main() {
	add := flag.Bool("add", false, "add a new item")
	update := flag.Int("update", 0, "update an existing item")
	del := flag.Int("del", 0, "delete an existing item")
	list := flag.Bool("list", false, "list all items")
	flag.Parse()

	noteUps := &noteup.Passd{}
	if err := noteUps.Load(noteUp); err != nil {
		panic(err)
	}

	switch {
	case *add:
		var account, password string
		account = styledInput("Enter account: ")
		password = styledInput("Enter password: ")

		noteUps.Add(account, password)
		err := noteUps.Store(noteUp)
		if err != nil {
			panic(err)
		}
	case *update>0:
		var password string
		password = styledInput("Enter password: ")
		err:=noteUps.Update(*update, password)
		if err != nil {
			panic(err)
		}
		err = noteUps.Store(noteUp)
		if err != nil {
			panic(err)
		}
	case *del>0:
		err := noteUps.Delete(*del)
		if err != nil {
			panic(err)
		}
		err = noteUps.Store(noteUp)
		if err != nil {
			panic(err)
		}
	case *list:	
		tab:=noteUps.List()
		fmt.Println(tab)

	default:
		fmt.Fprintf(os.Stderr, "invalid")
		os.Exit(0)

}

}

func styledInput(prompt string) string {
	styledPrompt := lipgloss.NewStyle().
		SetString(prompt). // Set the text to the value of the variable 'prompt'.
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("63")). // Set text color to gold
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).

		Bold(true)

	fmt.Print(styledPrompt)
	var input string
	fmt.Scanln(&input)
	return input
}






