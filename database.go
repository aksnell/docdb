package main

import (
	"math/rand"
	"time"
)

//Database defines the interface for creating, deleting and handeling connections to Collections.
type Database struct {
	collections map[string]Collection
	connections []*connection
}

//Connect a channel to the database able to interface to a specific collection.
func (d *Database) Connect(collection string, out chan result) {
	time := time.Now().Unix()
	c := &connection{
		rand.Int63(),
		time,
		time,
		true,
		d.collections[collection].queue,
		out,
	}
	d.connections = append(d.connections, c)
}
