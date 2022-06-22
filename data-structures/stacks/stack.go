package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top      int
	capacity uint
	array    []interface{}
}

// Returns an initialized stack
func NewStack(capacity uint) *Stack {
	return &Stack{
		top:      -1,
		capacity: capacity,
		array:    make([]interface{}, capacity),
	}
}

// Stack is full when top is equal the last index
func (s *Stack) IsFull() bool {
	return s.top == int(s.capacity)-1
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Size() uint {
	return uint(s.top + 1)
}

func (s *Stack) Push(data interface{}) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}
	s.top++
	s.array[s.top] = data
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	temp := s.array[s.top]
	s.top--
	return temp, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	temp := s.array[s.top]
	return temp, nil
}

// Remove all elements in the stack
func (s *Stack) Drain() {
	s.array = nil
	s.top = -1
}

func main() {
	stack := NewStack(10)
	stack.Push(5)
	stack.Push(10)
	stack.Push(15)
	stack.Push(20)
	fmt.Println("Stack size: ", stack.Size())

	data, _ := stack.Peek()
	fmt.Println("Top element: ", data)
	data, _ = stack.Pop()
	fmt.Println("Poped data from stack: ", data)
	fmt.Println("Stack size: ", stack.Size())

	stack.Drain()
	fmt.Println("Stack size: ", stack.Size()) // prints 0 as stack size
}
