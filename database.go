package main

//Database defines the interface for creating, updating, and editing Models and Collections.
type Database struct {
	collections map[string]Collection
}

//Model returns a Model struct which presents an interface to a Database collection.
func (d *Database) Model(key string) *Model {
	return nil
}

//Collection wraps a slice of Documents of a particular model.
type Collection struct {
	Type      string
	Documents []Document
}

func (c *Collection) insert() {
}

//Document wraps a key/value store.
type Document map[string]interface{}

//Filter is a type that wraps maps used to query collections.
type Filter map[string]interface{}

//Model presents an interface to a collection in the database.
type Model struct {
}

//Insert documents into the model's store to be inserted into the database later.
//If there is any error, none of the documents will be inserted.
func (m *Model) Insert() {
}

//Save all documents stored in the model to the Database.
func (m *Model) Save() {
}

//Find documents in the Model's associated collection matching a set of Filters.
func (m *Model) Find() {
}
