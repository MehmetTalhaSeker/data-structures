package tree

import (
	"data-structures/stack"
	"slices"
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

func (b *BinaryTree[T]) PostOrderTraversal() []T {
	if b.Root == nil {
		return nil
	}

	s := stack.NewStack(make([]Node[T], 0))
	s.Push(*b.Root)

	result := make([]T, 0)

	for {
		_, err := s.Peek()
		if err != nil {
			break
		}

		curr, err := s.Pop()
		if err == nil {
			result = append(result, curr.Value)
		}

		if curr.Left != nil {
			s.Push(*curr.Left)
		}

		if curr.Right != nil {
			s.Push(*curr.Right)
		}

	}

	slices.Reverse(result)

	return result
}

func (b *BinaryTree[T]) InOrderTraversal() []T {
	if b.Root == nil {
		return nil
	}

	s := stack.NewStack(make([]Node[T], 0))

	result := make([]T, 0)

	curr := b.Root

	for {
		_, err := s.Peek()
		if err != nil && curr == nil {
			break
		}

		for curr != nil {
			s.Push(*curr)
			curr = curr.Left
		}

		curr, err = s.Pop()
		if err != nil {
			break
		}

		result = append(result, curr.Value)

		curr = curr.Right
	}

	return result
}

func PreOrderTraversalRecursive[T any](r *Node[T]) []T {
	if r == nil {
		return []T{}
	}

	return append([]T{r.Value}, append(PreOrderTraversalRecursive(r.Left), PreOrderTraversalRecursive(r.Right)...)...)
}

func PostOrderTraversalRecursive[T any](r *Node[T]) []T {
	if r == nil {
		return []T{}
	}

	return append(PostOrderTraversalRecursive(r.Left), append(PostOrderTraversalRecursive(r.Right), r.Value)...)
}

func InOrderTraversalRecursive[T any](r *Node[T]) []T {
	if r == nil {
		return []T{}
	}

	return append(InOrderTraversalRecursive(r.Left), append([]T{r.Value}, InOrderTraversalRecursive(r.Right)...)...)
}
