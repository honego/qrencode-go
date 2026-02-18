package main

import (
	"os"
	"strings"

	"github.com/honeok/qrencode-go/qrencode"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	text := strings.Join(os.Args[1:], " ")

	grid, err := qrencode.Encode(text, qrencode.ECLevelQ)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	grid.TerminalOutput(os.Stdout)
}
