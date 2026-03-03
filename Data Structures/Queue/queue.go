package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}

	firstItem := q.items[0]
	q.items = q.items[1:]
	return firstItem, true
}

func (q *Queue) Front() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[0], true
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func main() {
	queue := &Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	fmt.Println("Enqueued: 10, 20, 30")

	fmt.Println("Size:", queue.Size())

	if val, ok := queue.Front(); ok {
		fmt.Println("Front:", val)
	}

	for !queue.IsEmpty() {
		if val, ok := queue.Dequeue(); ok {
			fmt.Println("Dequeue:", val)
		}
	}

	fmt.Println("IsEmpty:", queue.IsEmpty())
}
