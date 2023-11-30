package linkedlist

import (
	"errors"
	"reflect"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

var (
	ErrEmptyLinkedList = errors.New("linked list is empty")
)

type LinkedList[T any] interface {
	Add(value T)
	Remove(value T) error
	Size() int
}

type linkedList[T any] struct {
	head *Node[T]
	// we can add tail if we want to.
}

func NewLinkedList[T any](head *Node[T], more ...*Node[T]) LinkedList[T] {
	ll := &linkedList[T]{
		head: head,
	}

	n := head
	for i := 0; i < len(more); i++ {
		n.Next = more[i]
		n = n.Next
	}

	return ll
}

func (l *linkedList[T]) Add(value T) {
	if l.head == nil {
		l.head = &Node[T]{Value: value}
		return
	}

	n := l.head

	for n.Next != nil {
		n = n.Next
	}

	n.Next = &Node[T]{Value: value}
}

func (l *linkedList[T]) Remove(value T) error {
	if l.head == nil {
		return ErrEmptyLinkedList
	}

	if reflect.DeepEqual(l.head.Value, value) {
		l.head = l.head.Next
		return nil
	}

	n := l.head
	for n.Next != nil {
		if reflect.DeepEqual(n.Next.Value, value) {
			n.Next = n.Next.Next
			return nil
		}
		n = n.Next
	}

	return nil
}

func (l *linkedList[T]) Size() int {
	if l.head == nil {
		return 0
	}

	i := 1

	n := l.head
	for n.Next != nil {
		i++
		n = n.Next
	}

	return i
}
