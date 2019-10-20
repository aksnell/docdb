package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Run err: %v", err)
	}
}

func run() error {
	return nil
}
