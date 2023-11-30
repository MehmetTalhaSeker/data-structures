package queue

import "errors"

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() (*T, error)
	Peek() (*T, error)
	// get queue elements for testing.
	getter() []T
}

type queue[T any] struct {
	els []T
}

var (
	ErrEmptyQueue = errors.New("queue is empty")
)

func NewQueue[T any](els []T) Queue[T] {
	return &queue[T]{
		els: els,
	}
}

func (q *queue[T]) Enqueue(el T) {
	q.els = append(q.els, el)
}

func (q *queue[T]) Dequeue() (*T, error) {
	l := len(q.els)
	if l < 1 {
		return nil, ErrEmptyQueue
	}

	fe := q.els[0]
	q.els = q.els[1:]

	return &fe, nil
}

func (q *queue[T]) Peek() (*T, error) {
	l := len(q.els)
	if l < 1 {
		return nil, ErrEmptyQueue
	}

	return &q.els[0], nil
}

func (q *queue[T]) getter() []T {
	return q.els
}
