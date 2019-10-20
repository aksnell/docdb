package main

//Database
type Database struct {
	collections map[string]Collection
}

//Collection wraps a slice of Documents of a particular model.
type Collection struct {
	Type string
}

//Model wraps a document Schema
type Model struct {
}

//Document wraps a key/value store based on a particular model.
type Document struct {
	Type string
}
