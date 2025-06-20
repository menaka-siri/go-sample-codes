package functions

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	Next *List[T]
	Val  T
}

func (list *List[T]) Display() {
	for p := list; p != nil; p = p.Next {
		fmt.Printf("%v ", p.Val)
	}
}