package main

import "fmt"

type ListNode struct {
	data interface{}
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) Display() error {
	if ll.head == nil {
		return fmt.Errorf("Display: List is empty")
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v ->", current.data)
		current = current.next
	}
	fmt.Println()
	return nil
}

func (ll *LinkedList) Len() int {
	return ll.size
}

// 0(1) time, O(1) space
func (ll *LinkedList) InsertBeginning(data interface{}) {
	node := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = node
	} else {
		node.next = ll.head
		ll.head = node
	}
	ll.size++
	return
}

// O(n) time for scanning the list of size n, O(1) space.
func (ll *LinkedList) InsertEnd(data interface{}) {
	node := &ListNode{
		data: data,
	}
	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current != nil {
			current = current.next
		}
		current.next = node
	}
	ll.size++
	return
}

// O(n) time for worst case, O(1) space.
func (ll *LinkedList) Insert(position int, data interface{}) error {
	if position < 1 || position > ll.size+1 {
		return fmt.Errorf("Insert: Index out of bounds")
	}
	newNode := &ListNode{data, nil}
	var prev, current *ListNode
	prev = nil
	current = ll.head
	for position > 1 {
		prev = current
		current = current.next
		position = position - 1
	}
	if prev != nil {
		prev.next = newNode
		newNode.next = current
	} else {
		newNode.next = current
		ll.head = newNode
	}

	ll.size++
	return nil
}

// O(1) time for worst case, O(1) space.
func (ll *LinkedList) DeleteFirst() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("Delete Front: List is empty")
	}
	data := ll.head.data
	ll.head = ll.head.next
	ll.size--
	return data, nil
}

// O(n) time for worst case, O(1) space.
func (ll *LinkedList) DeleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("Delete Last: List is empty")
	}
	var prev *ListNode
	current := ll.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		ll.head = nil
	}

	ll.size--
	return current.data, nil
}

// O(n) time for worst case, O(1) space.
func (ll *LinkedList) Delete(position int) (interface{}, error) {
	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("Delete: Index out of bounds")
	}
	var prev, current *ListNode
	prev = nil
	current = ll.head
	pos := 0
	if position == 1 {
		ll.head = ll.head.next
	} else {
		for pos != position-1 {
			pos = pos + 1
			prev = current
			current = current.next
		}
		if current != nil {
			prev.next = current.next
		}
	}

	ll.size--
	return current.data, nil
}
func main() {}
