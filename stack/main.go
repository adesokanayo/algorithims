package main

import "fmt"

func main() {

	var myStack Stack

	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}

type Stack struct {
	items []int
}

//Push will add an item to the stack
func (s *Stack) Push(k int) {
	s.items = append(s.items, k)
}

func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:len(s.items)-1]
	return toRemove

}
