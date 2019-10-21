package main

import (
	"math/rand"
	"time"
)

//Collection manages the creating, updating, and removing Documents via connections managed by its Database.
type Collection struct {
	name      string
	documents []Document
	docMap    map[int64]int
	queue     chan request
}

func (c *Collection) listen() {
	go func() {
		for e := range c.queue {
			time := time.Now().Unix()
			switch e.action {
			case "INSERT":
				c.Insert(e.json...)
				e.respond <- response{
					time: time,
					err:  nil,
				}
			case "FIND":
				e.respond <- response{
					time: time,
					json: c.Find(e.json...),
				}
			}
		}
	}()
}

func (c *Collection) registerID() int64 {
	id := rand.Int63()
	for _, exists := c.docMap[id]; exists; id = rand.Int63() {
	}
	c.docMap[id] = len(c.documents)
	return id
}

func (c *Collection) createResponse(action string, conn *Connection, json ...map[string]interface{}) response {
	return response{
		action,
		time.Now().Unix(),
		json,
		nil,
	}
}

//Insert a series of JSON objects into the collection.
func (c *Collection) Insert(json ...map[string]interface{}) error {
	time := time.Now().Unix()
	for _, data := range json {
		doc := Document{
			c.registerID(),
			time,
			time,
			data,
		}
		c.documents = append(c.documents, doc)
	}
	return nil
}

//Find a series of JSON objects based on the provided filters.
//Key/Values within Filters are logically matched as "AND" and seperate filters are logically matched as "OR",
func (c *Collection) Find(filters ...map[string]interface{}) []map[string]interface{} {
	results := make([]map[string]interface{}, 0)
	for _, doc := range c.documents {
		for _, filter := range filters {
			if matchFilter(&doc, filter) {
				results = append(results, doc.data)
			}
		}
	}
	return results
}

//Update a series of Documents matching provided filters.
func (c *Collection) Update(update map[string]interface{}, filters ...map[string]interface{}) {

}

//Delete a series of Documents matching provided filters.
func (c *Collection) Delete(filters ...map[string]interface{}) {

}
