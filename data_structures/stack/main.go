package main

import (
	"fmt"
)

//Structure for element
type element struct {
	name string
}

//Stack is a slice of elements
type Stack []element

//Check if the Stack is empty
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

//Push using Go's Builtin function for Slice we can append elements to Stack
func (s *Stack) push(elem element) {
	*s = append(*s, elem)
}

//Pop using Go's builtin feature of slicing [start(default=0):end]
func (s *Stack) pop() element {
	if s.isEmpty() {
		return element{""}
	} else {
		stackPtr := len(*s) - 1
		elem := (*s)[stackPtr]
		*s = (*s)[:stackPtr]
		return elem
	}
}

//Return true or false based upon the first occurence of the element
func (s *Stack) searchElement(elem element) bool {
	stackPtr := len(*s) - 1
	for stackPtr >= 0 {
		if (*s)[stackPtr] == elem {
			return true
		}
		stackPtr--
	}
	return false
}

//Count the total elements in Stack
func (s *Stack) countElements() int {
	return len(*s)
}

//Print all the elements in Stack
func (s *Stack) displayElements() {
	stackPtr := len(*s) - 1
	for stackPtr >= 0 {
		fmt.Printf("%v ", (*s)[stackPtr])
		stackPtr--
	}
	fmt.Printf("\n")
}

//Entry Point
func main() {
	//Create a Stack Data type
	var stack Stack

	//Pushing elements to Stack
	stack.push(element{"Homer"})
	stack.push(element{"Marge"})
	stack.push(element{"Lisa"})
	stack.push(element{"Bart"})
	stack.push(element{"Maggie"})

	//Printing total elements in the Stack
	fmt.Println(stack.countElements())

	//Displaying the contents of the Stack
	stack.displayElements()

	//Checking if a element is present
	fmt.Println(stack.searchElement(element{"Simpsons"}))

	//Poping elements from the Stack
	for i := stack.countElements(); i > 0; i-- {
		fmt.Println(stack.pop())
	}
}
