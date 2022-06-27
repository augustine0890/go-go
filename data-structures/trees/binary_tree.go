package main

import (
	"fmt"
	"math"
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

/**
Inorder Traversal
- Traverse the left subtree in Inorder
- Visit the root
- Traverse the right subtree in Inorder
**/
func InOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	temp := root
	stack := NewStack(1)
	for stack.Size() > 0 || temp != nil {
		if temp != nil {
			stack.Push(temp)
			temp = temp.left
		} else {
			obj, _ := stack.Pop()
			temp = obj.(*BinaryTreeNode)
			fmt.Printf("%d", temp.data)
			temp = temp.right
		}
	}
}

// InOrderWalk traverses
func InOrderWalk(root *BinaryTreeNode, ch chan int) {
	if root == nil {
		return
	}

	InOrderWalk(root.left, ch)
	ch <- root.data
	InOrderWalk(root.right, ch)
}

// InOrderWalker
func InOrderWalker(root *BinaryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		InOrderWalk(root, ch)
		close(ch)
	}()
	return ch
}

/**
PostOrder Traversal
- Traverse the left subtree in Inorder
- Traverse the right subtree in Inorder
- Visit the root
**/
func PostOrder(root *BinaryTreeNode) []int {
	var result []int
	stack := NewStack(1)
	stack.Push(root)

	for !stack.IsEmpty() {
		node, _ := stack.Pop()
		temp := node.(*BinaryTreeNode)
		result = append(result, temp.data)
		if temp.left != nil {
			stack.Push(temp.left)
		}
		if temp.right != nil {
			stack.Push(temp.right)
		}
	}
	n := len(result)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		result[i], result[j] = result[j], result[i]
	}
	fmt.Println(result)
	return result
}

/**
Level Order Traversal
- Visit the root
- While traversing level l, keep all the elements at level l + 1 in queue
- Go to the next level and visit all the nodes at that level
- Repeat this until all levels are completed
**/
func LevelOrder(root *BinaryTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		var level []int
		for i := 0; i < qlen; i++ {
			node := queue[0]
			level = append(level, node.data)
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		result = append(result, level)
	}
	return result
}

// Find the maximum element in a binary tree
func findMax(root *BinaryTreeNode) int {
	max := math.MinInt32
	if root == nil {
		return max
	}
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			if node.data > max {
				max = node.data
			}
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return max
}

// Find the minimum element in binary tree
func findMin(root *BinaryTreeNode) int {
	min := math.MaxInt32
	if root == nil {
		return min
	}

	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			if node.data < min {
				min = node.data
			}
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return min
}

// Search the element in binary tree
func find(root *BinaryTreeNode, data int) *BinaryTreeNode {
	if root == nil {
		return root
	}
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			if node.data == data {
				return node
			}
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return nil
}

// getHeight return the height of a binary tree as an integer
func getHeight(root *BinaryTreeNode) int {
	height := -1
	if root == nil {
		return -1
	}
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		height++
	}
	return height
}

func main() {
	/*
						10
				   /  \
				 20	   30
				/ \      \
			40  50     60
			/
		70
	*/
	root := &BinaryTreeNode{nil, 10, nil}
	root.left = &BinaryTreeNode{nil, 20, nil}
	root.right = &BinaryTreeNode{nil, 30, nil}
	root.left.left = &BinaryTreeNode{nil, 40, nil}
	root.left.right = &BinaryTreeNode{nil, 50, nil}
	root.right.right = &BinaryTreeNode{nil, 60, nil}
	root.left.left.left = &BinaryTreeNode{nil, 70, nil}

	fmt.Println("Preorder Traversal - Iterative Solution: ")
	PreOrder(root)
	fmt.Println()

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

	fmt.Println("Inorder Traversal - Iterative Solution: ")
	InOrder(root)
	fmt.Println()

	fmt.Println("Inorder Traversal - Iterative Solution: ")
	PostOrder(root)

	levelOrder := LevelOrder(root)
	fmt.Println("Level Order Traversal:", levelOrder)

	max := findMax(root)
	fmt.Println("Maximum value:", max)
	fmt.Println("Minimum value:", findMin(root))
	fmt.Println("Find the value (50):", find(root, 50))
	fmt.Println("Height of the tree:", getHeight(root))
}
