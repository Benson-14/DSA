package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	lastIdx := len(s.items) - 1
	item := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return item, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}

	lastItem := s.items[len(s.items)-1]
	return lastItem, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println("Pushed: 10, 20, 30")

	fmt.Println("Size:", stack.Size())

	if val, ok := stack.Peek(); ok {
		fmt.Println("Peek:", val)
	}

	for !stack.IsEmpty() {
		if val, ok := stack.Pop(); ok {
			fmt.Println("Pop:", val)
		}
	}

	fmt.Println("IsEmpty:", stack.IsEmpty())
}
