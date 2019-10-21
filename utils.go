package main

//Returns true if all key/value pairs in a filter match for the supplied document.
func matchFilter(doc *Document, filter map[string]interface{}) bool {
	for k, v := range filter {
		if doc.get(k) != v {
			return false
		}
	}
	return true
}
