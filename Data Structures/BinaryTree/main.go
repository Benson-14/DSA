package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (t *BST) Insert(value int) {
	t.Root = insertNode(t.Root, value)
}

func insertNode(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}
	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else {
		node.Right = insertNode(node.Right, value)
	}
	return node
}

func (t *BST) Search(value int) bool {
	return searchNode(t.Root, value)
}

func searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return searchNode(node.Left, value)
	}
	return searchNode(node.Right, value)
}

func (t *BST) InOrder() {
	inOrder(t.Root)
	fmt.Println()
}

func inOrder(node *Node) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	fmt.Println(node.Value, " ")
	inOrder(node.Right)
}

func main() {
	tree := &BST{}

	values := []int{5, 3, 7, 1, 4, 6, 8}
	fmt.Println("Inserting:", values)
	for _, v := range values {
		tree.Insert(v)
	}

	fmt.Print("In-order: ")
	tree.InOrder()

	fmt.Println("Search 4:", tree.Search(4))
	fmt.Println("Search 9:", tree.Search(9))
}
