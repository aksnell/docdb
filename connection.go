package main

type connection struct {
	id       int64
	opened   int64
	lastUsed int64
	open     bool
	in       chan<- *event
	out      chan result
}

type event struct {
	goal       int
	time       int64
	json       []map[string]interface{}
	connection *connection
}

type result struct {
	time int64
	data []map[string]interface{}
	err  error
}
