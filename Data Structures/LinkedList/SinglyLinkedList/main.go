package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	size int
}

func (l *LinkedList) Prepend(value int) {
	newNode := &Node{Value: value}
	newNode.Next = l.Head
	l.Head = newNode
	l.size++
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	l.size++

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (l *LinkedList) Find(value int) bool {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Print() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList{}

	fmt.Print("Initial: ")
	list.Print()

	list.Append(10)
	list.Append(20)
	list.Append(30)
	fmt.Println("After append 10, 20, 30:")
	list.Print()

	list.Prepend(5)
	fmt.Println("After prepend 5:")
	list.Print()

	fmt.Println("Size:", list.Size())
	fmt.Println("Find 20:", list.Find(20))
	fmt.Println("Find 99:", list.Find(99))
}
