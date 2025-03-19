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

	for {
		printMessage(25, 11, termbox.ColorDefault, termbox.ColorDefault, "EGO Text Editor")

		err := termbox.Flush()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
			termbox.Close()
			break
		}
	}

}

func printMessage(col, row int, fg, bg termbox.Attribute, message string) {
	for _, ch := range message {
		termbox.SetCell(col, row, ch, fg, bg)
		col += runewidth.RuneWidth(ch)
	}
}
