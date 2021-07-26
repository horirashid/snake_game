package main

type Node struct {
	prev  *Node
	next  []*Node
	name  string
	id    string
	value string
}

func NewNode(prev *Node, next []*Node, name string, id string, value string) *Node {
	node := &Node{
		prev:  prev,
		next:  next,
		name:  name,
		id:    id,
		value: value,
	}
	return node
}
