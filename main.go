package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("------------")

	lineNumber := 1

	for {
		fmt.Printf("%d ", lineNumber)
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\r\n", "", -1)
		lineNumber++
	}
}
