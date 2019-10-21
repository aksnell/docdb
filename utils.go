package main

func matchFilter(doc *Document, filter map[string]interface{}) bool {
	for k, v := range filter {
		if doc.get(k) != v {
			return false
		}
	}
	return true
}
