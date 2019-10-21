package main

//Document presents an interface for a Collection to manage a JSON document.
type Document struct {
	id           int64
	created      int64
	lastModified int64
	data         map[string]interface{}
}

func (d *Document) get(key string) interface{} {
	if value, exists := d.data[key]; exists {
		return value
	}
	return nil
}

func (d *Document) update(values map[string]interface{}) {

}
