package main

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"os"
)

var ROWS, COLS int
var offsetX, offsetY int

var textBuffer = [][]rune{
	{'h', 'e', 'l', 'l', 'o'},
	{'w', 'o', 'r', 'l', 'd'},
}

func main() {
	runEditor()
}

func runEditor() {
	if err := termbox.Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		COLS, ROWS = termbox.Size()
		ROWS--

		if COLS < 80 {
			COLS = 80
		}

		err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		if err != nil {
			return
		}

		displayTextBuffer()

		err = termbox.Flush()
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

func displayTextBuffer() {
	var row, col int

	for row = 0; row < ROWS; row++ {
		textBufferRow := row + offsetY
		for col = 0; col < COLS; col++ {
			textBufferCol := col + offsetX
			if textBufferRow >= 0 && textBufferRow < len(textBuffer) && textBufferCol < len(textBuffer[textBufferRow]) {
				if textBuffer[textBufferRow][textBufferCol] != '\t' {
					termbox.SetChar(col, row, textBuffer[textBufferRow][textBufferCol])
				} else {
					termbox.SetCell(col, row, rune(' '), termbox.ColorDefault, termbox.ColorGreen)
				}
			} else if row+offsetY > len(textBuffer)-1 {
				termbox.SetCell(0, row, rune('*'), termbox.ColorBlue, termbox.ColorDefault)
			}
		}

		termbox.SetChar(col, row, rune('\n'))
	}
}
