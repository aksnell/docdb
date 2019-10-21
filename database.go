package main

import (
	"log"
	"math/rand"
	"time"
)

//Database defines the interface for creating, deleting and handeling connections to Collection(s).
type Database struct {
	collections map[string]*Collection
	connections []*Connection
}

//Create establishes a new collection if the supplied key does not already exist.
func (db *Database) Create(key string) {
	if _, exists := db.collections[key]; !exists {
		col := &Collection{
			key,
			make([]Document, 0),
			make(map[int64]int),
			make(chan request, 1),
		}
		db.collections[key] = col
		log.Printf("New collection: %s created!\n", key)
		col.listen()
	}
}

//Connect a channel to the database able to communicate with a specific Collection.
func (db *Database) Connect(collection string) *Connection {
	time := time.Now().Unix()
	conn := &Connection{
		rand.Int63(),
		time,
		time,
		true,
		db.collections[collection].queue,
		make(chan response, 1),
		make([][]byte, 0),
	}
	db.connections = append(db.connections, conn)
	conn.listen()
	return conn
}
