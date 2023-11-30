package tree

import (
	"data-structures/stack"
)

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type BinaryTree[T any] struct {
	Root *Node[T]
}

func NewBinaryTree[T any](root *Node[T]) *BinaryTree[T] {
	return &BinaryTree[T]{
		Root: root,
	}
}

func (b *BinaryTree[T]) PreOrderTraversal() []T {
	if b.Root == nil {
		return nil
	}

	s := stack.NewStack(make([]Node[T], 0))
	s.Push(*b.Root)

	a := make([]T, 0)

	for {
		_, err := s.Peek()
		if err != nil {
			break
		}

		curr, err := s.Pop()
		if err == nil {
			a = append(a, curr.Value)
		}

		if curr.Right != nil {
			s.Push(*curr.Right)
		}

		if curr.Left != nil {
			s.Push(*curr.Left)
		}
	}

	return a
}
