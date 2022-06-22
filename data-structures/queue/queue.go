package main

import (
	"bytes"
	"errors"
	"fmt"
)

type ListNode struct {
	data interface{}
	next *ListNode
}

// Queue implements a queue by a linked list
type Queue struct {
	front *ListNode
	rear  *ListNode
	size  int
}

// enQueue adds an element to the end of the queue
func (q *Queue) enQueue(data interface{}) {
	rear := new(ListNode)
	rear.data = data
	if q.isEmpty() {
		q.front = rear
	} else {
		oldLast := q.rear
		oldLast.next = rear
	}
	q.rear = rear
	q.size++
}

func (q *Queue) deQueue() (interface{}, error) {
	if q.isEmpty() {
		q.rear = nil
		return nil, errors.New("unable to dequeue element, queue is empty")
	}
	data := q.front.data
	q.front = q.front.next
	q.size--
	return data, nil
}

// Return from front element without removing
func (q *Queue) fromElement() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("unable to peek element, queue is empty")
	}
	data := q.front.data
	return data, nil
}

func (q *Queue) isEmpty() bool {
	return q.front == nil
}

func (q *Queue) length() int {
	return q.size
}

// Representation of queue q formatted
func (q *Queue) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := q.front
	for i := 0; i < q.size; i++ {
		result.WriteString(fmt.Sprintf("%v", j.data))
		if i < q.size-1 {
			result.WriteByte(' ')
		}
		j = j.next
	}
	result.WriteByte(']')
	return result.String()
}

func main() {
	queue := new(Queue)
	queue.enQueue(1)
	queue.enQueue(2)
	queue.enQueue(3)
	queue.enQueue(4)
	queue.enQueue(5)
	fmt.Println(queue.String())

	fmt.Println(queue.deQueue())
	fmt.Println(queue.deQueue())
	fmt.Println(queue.deQueue())
	fmt.Println(queue.deQueue())
	fmt.Println(queue.deQueue())
	fmt.Println(queue.String())

}
