package main

import (
	"errors"
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

func (s *Stack) Resize() {
	if s.IsFull() {
		s.capacity *= 2
	} else {
		s.capacity /= 2
	}
	target := make([]interface{}, s.capacity)
	copy(target, s.array[:s.top+1])
	s.array = target
}

func (s *Stack) Push(data interface{}) error {
	if s.IsFull() {
		s.Resize()
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
	if s.Size() < s.capacity/2 {
		s.Resize()
	}
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
