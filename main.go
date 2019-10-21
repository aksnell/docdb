package main

//TODO:: Handle database close, connection monitoring and closing.

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Run err: %v", err)
	}
}

func run() error {
	//Create new database
	db := Database{
		make(map[string]*Collection),
		make([]*Connection, 0),
	}
	//Some nonsense documents.
	Alex := map[string]interface{}{
		"Name": "Alex",
		"Age":  99,
	}
	Bob := map[string]interface{}{
		"Name": "Bob",
		"Age":  27,
	}
	Steven := map[string]interface{}{
		"Name": "Steven",
		"Age":  9,
	}
	Chris := map[string]interface{}{
		"Name": "Chris",
		"Age":  14,
	}
	//Create a new collection named "TEST".
	db.Create("TEST")

	//Create a connection to the "TEST" collection.
	conn := db.Connect("TEST")

	//Insert these documents a whole bunch.
	conn.Dispatch("INSERT", Alex, Bob, Steven, Chris)
	conn.Dispatch("INSERT", Alex, Bob, Steven, Chris)
	conn.Dispatch("INSERT", Alex, Bob, Steven, Chris)
	conn.Dispatch("INSERT", Alex, Bob, Steven, Chris)
	conn.Dispatch("INSERT", Alex, Bob, Steven, Chris)
	conn.Dispatch("INSERT", Alex, Bob, Steven, Chris)
	conn.Dispatch("INSERT", Alex, Bob, Steven, Chris)
	conn.Dispatch("INSERT", Alex, Bob, Steven, Chris)
	conn.Dispatch("INSERT", Alex, Bob, Steven, Chris)
	conn.Dispatch("INSERT", Alex, Bob, Steven, Chris)
	conn.Dispatch("INSERT", Alex, Bob, Steven, Chris)

	//Find all documents with a "Name" key matching "Alex" OR "Age" key matching "27"
	conn.Dispatch("FIND", map[string]interface{}{"Name": "Alex"}, map[string]interface{}{"Age": 27})

	//Just printing some stuff out.
	for {
		if bytes := conn.GetBytes(); bytes != nil {
			resp := make(map[string]interface{})
			err := json.Unmarshal(bytes, &resp)
			if err != nil {
				log.Printf("%v\n%v\n", err, bytes)
			} else {
				for k, v := range resp {
					log.Println(k, v)
				}
			}
		}
	}
	return nil
}
