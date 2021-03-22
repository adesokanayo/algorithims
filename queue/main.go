package main

import "fmt"

func main() {

	var q Queue

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Enqueue(5)
	q.Enqueue(6)
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)

}

//Queue holds the data to be enqueued and dequeued
type Queue struct {
	items []int
}

// Enqueue add an item to the end of the queue
func (q *Queue) Enqueue(k int) {

	q.items = append(q.items, k)
}

// Dequeue removes first item on the queue
func (q *Queue) Dequeue() {

	item := &q.items[0]
	if item == nil {
		q.items = make([]int, 0)
	}
	q.items = q.items[1:]
}
