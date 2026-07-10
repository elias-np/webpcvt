package main

import (
	"fmt"
	"os"

	"webpcvt/internal/cli"
)

const version = "0.1.0"

func main() {
	if err := cli.Run(os.Args[1:], version); err != nil {
		if err == cli.ErrVersion {
			return
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
