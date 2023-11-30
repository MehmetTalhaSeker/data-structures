package stack

import "errors"

type Stack[T any] interface {
	Push(T)
	Pop() (*T, error)
	Peek() (*T, error)
	// get stack elements for testing.
	getter() []T
}

type stack[T any] struct {
	els []T
}

var (
	ErrEmptyStack = errors.New("stack is empty")
)

func NewStack[T any](els []T) Stack[T] {
	return &stack[T]{
		els: els,
	}
}

func (s *stack[T]) Push(el T) {
	s.els = append(s.els, el)
}

func (s *stack[T]) Pop() (*T, error) {
	l := len(s.els)
	if l < 1 {
		return nil, ErrEmptyStack
	}

	le := s.els[l-1]
	s.els = s.els[:l-1]

	return &le, nil
}

func (s *stack[T]) Peek() (*T, error) {
	l := len(s.els)
	if l < 1 {
		return nil, ErrEmptyStack
	}

	return &s.els[len(s.els)-1], nil
}

func (s *stack[T]) getter() []T {
	return s.els
}
