package main

import (
	"encoding/json"
	"log"
	"time"
)

//Connection presents an interface for finding, inserting, editing, and deleting Documents in a Collection.
type Connection struct {
	id       int64
	opened   int64
	lastUsed int64
	open     bool
	request  chan request
	response chan response
	buffer   [][]byte
}

//listen for responses, remove them from the channel and append them to a local buffer to be read when needed.
func (conn *Connection) listen() {
	go func() {
		for response := range conn.response {
			for _, resp := range response.json {
				bytes, err := json.Marshal(resp)
				if err != nil {
					log.Println(err)
				}
				conn.buffer = append(conn.buffer, bytes)
			}
		}
	}()
}

//Dispatch an action type and json to the connection Collection.
func (conn *Connection) Dispatch(action string, json ...map[string]interface{}) {
	r := request{
		action,
		time.Now().Unix(),
		json,
		conn.response,
	}
	conn.request <- r
}

//GetBytes returns the oldest unread response as an array of bytes.
func (conn *Connection) GetBytes() []byte {
	if len(conn.buffer) > 0 {
		bytes := make([]byte, len(conn.buffer[0]))
		copy(bytes, conn.buffer[0])
		conn.buffer = append(conn.buffer[:0], conn.buffer[0+1:]...)
		return bytes
	}
	return nil
}

type request struct {
	action  string
	time    int64
	json    []map[string]interface{}
	respond chan response
}

type response struct {
	action string
	time   int64
	json   []map[string]interface{}
	err    error
}
