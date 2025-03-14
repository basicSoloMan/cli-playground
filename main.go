package main

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"os"
)

func main() {
	runEditor()
}

func runEditor() {
	if err := termbox.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer termbox.Close()

	printMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "EGO Text Editor")

	err := termbox.Flush()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	termbox.PollEvent()
}

func printMessage(col, row int, fg, bg termbox.Attribute, message string) {
	for _, ch := range message {
		termbox.SetCell(col, row, ch, fg, bg)
		col += runewidth.RuneWidth(ch)
	}
}
