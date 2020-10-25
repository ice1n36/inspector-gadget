package main

import (
	"fmt"
	"github.com/ice1n36/inspector-gadget/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
