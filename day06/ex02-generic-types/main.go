package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Append(v T) *List[T] {
	if l == nil {
		return &List[T]{val: v}
	}

	head := l
	for head.next != nil {
		head = head.next
	}
	head.next = &List[T]{val: v}
	return l
}

func (l *List[T]) Length() int {
	count := 0
	for curr := l; curr != nil; curr = curr.next {
		count++
	}
	return count
}

func main() {
	var list *List[int]

	list = list.Append(1)
	list = list.Append(2)
	list = list.Append(3)

	fmt.Println(list)
	fmt.Println("Length:", list.Length())
}
