package main

import (
	"fmt"
	"os"

	"github.com/danieljmanningdev/djm-cli/cmd"
)

func main() {
	if err := cmd.Command(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "djm:", err)
		os.Exit(1)
	}
}
