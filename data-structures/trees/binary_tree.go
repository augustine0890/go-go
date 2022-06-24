package main

import (
	"fmt"
	"math/rand"
)

// Binary Tree Node is a binary tree with integer values
type BinaryTreeNode struct {
	left  *BinaryTreeNode
	data  int
	right *BinaryTreeNode
}

// PreOrder traverses a tree in pre-order
// func PreOrder(root *BinaryTreeNode) {
// if root == nil {
// return
// }
// fmt.Printf("%d", root.data)
// PreOrder(root.left)
// PreOrder(root.right)
// }

/**
Preorder traversal
- Visit the root
- Traverse the left subtree in Preorder
- Traverse the right subtree in Preorder
**/
// O(n) time, O(n) space
func PreOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	current := root
	stack := NewStack(1)
	stack.Push(current)

	for stack.Size() > 0 {
		temp, _ := stack.Pop()
		current = temp.(*BinaryTreeNode)
		fmt.Printf("%d", current.data)
		if current.right != nil {
			stack.Push(current.right)
		}
		if current.left != nil {
			stack.Push(current.left)
		}
	}
}

// PreOrderWalk traverses
func PreOrderWalk(root *BinaryTreeNode, ch chan int) {
	if root == nil {
		return
	}
	ch <- root.data
	PreOrderWalk(root.left, ch)
	PreOrderWalk(root.right, ch)
}

// PreOrderWalker launches Walk in a new goroutine, and return a read-only channel of value
func PreOrderWalker(root *BinaryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		PreOrderWalk(root, ch)
		close(ch)
	}()
	return ch
}

// NewBinaryTree returns a new, randam binary tree holding the values 1k, 2k, ..., nk
func NewBinaryTree(n, k int) *BinaryTreeNode {
	var root *BinaryTreeNode
	for _, v := range rand.Perm(n) {
		root = insert(root, (1+v)*k)
	}
	return root
}

func insert(root *BinaryTreeNode, v int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{nil, v, nil}
	}
	if v < root.data {
		root.left = insert(root.left, v)
		return root
	}
	root.right = insert(root.right, v)
	return root
}

func main() {
	t1 := NewBinaryTree(10, 1)
	// Pre-Order walk with print statements
	PreOrder(t1)
	fmt.Println()

	// Pre-Order walk with channel
	c := PreOrderWalker(t1)
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("%d", v)
	}
	fmt.Println()
}
