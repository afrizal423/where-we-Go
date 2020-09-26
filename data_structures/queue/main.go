package main

import "fmt"

type element struct {
	npm      string
	name     string
	semester int
}

// Node represent a unit of list
type Node struct {
	value element
	next  *Node
}

type Queue struct {
	node   *Node
	first  *Node
	last   *Node
	length int
}

func (q Queue) display() {
	cursor := q.node
	for cursor != nil {
		fmt.Println(cursor.value)
		cursor = cursor.next
	}
}

func (q Queue) firstElement() element {
	return q.first.value
}

func (q Queue) lastElement() element {
	return q.last.value
}

func (q Queue) count() int {
	return q.length
}

func (q *Queue) insert(newElement element) {
	newNode := &Node{
		value: newElement,
		next:  nil,
	}
	if q.node == nil {
		q.node = newNode
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	q.length++
}

func main() {
	q := Queue{}

	e1 := element{"NPM#1", "whoami", 3}
	e2 := element{"NPM#2", "root", 7}
	e3 := element{"NPM#3", "anonymous", 9}

	q.insert(e1)
	q.insert(e2)
	q.insert(e3)

	q.display()

	fmt.Println(q.count())
	fmt.Println(q.firstElement())
	fmt.Println(q.lastElement())
}
