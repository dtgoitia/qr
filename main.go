package main

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "must pass exactly one argument")
		os.Exit(-1)
	}

	content := os.Args[1]

	fmt.Fprintln(os.Stderr, "generating QR code, content:", content)

	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		WhiteChar: qrterminal.WHITE,
		BlackChar: qrterminal.BLACK,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig(content, config)
}
