package main

import (
	"fmt"
	"os"
)

//TODO:: Database will be hosted on a server, contains Database, Collections
//TODO:: Model be local interface to interact with Database.

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Run err: %v", err)
	}
}

func run() error {
	return nil
}
