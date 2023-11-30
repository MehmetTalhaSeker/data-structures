package queue

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Queue", func() {
	var (
		q Queue[int]
	)

	Describe("Enqueue", func() {
		BeforeEach(func() {
			q = NewQueue([]int{1, 2, 3})
		})
		It("should enqueue an element into the queue", func() {
			q.Enqueue(4)

			l := len(q.getter())
			els := q.getter()

			Expect(els[l-1]).To(Equal(4))
		})
		It("should enqueue an element into the queue", func() {
			q.Enqueue(21)

			l := len(q.getter())
			els := q.getter()

			Expect(els[l-1]).To(Equal(21))
		})
	})

	Describe("Dequeue", func() {
		Context("when the queue is not empty", Ordered, func() {
			BeforeAll(func() {
				q = NewQueue([]int{1, 2, 3})
			})
			It("should dequeue the front element from the queue", func() {
				value, err := q.Dequeue()
				Expect(err).ToNot(HaveOccurred())
				Expect(*value).To(Equal(1))
			})
			It("should dequeue the front element from the queue", func() {
				value, err := q.Dequeue()
				Expect(err).ToNot(HaveOccurred())
				Expect(*value).To(Equal(2))
			})
		})

		Context("when the queue is empty", func() {
			BeforeEach(func() {
				q = NewQueue([]int{})
			})
			It("should return an error", func() {
				value, err := q.Dequeue()
				Expect(err).To(MatchError(ErrEmptyQueue))
				Expect(value).To(BeNil())
			})
		})
	})

	Describe("Peek", func() {
		Context("when the queue is not empty", Ordered, func() {
			BeforeAll(func() {
				q = NewQueue([]int{1, 2, 3})
			})
			It("should return the front element from the queue", func() {
				value, err := q.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(*value).To(Equal(1))
			})
			It("should return the front element from the queue", func() {
				_, _ = q.Dequeue()

				value, err := q.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(*value).To(Equal(2))
			})
		})

		Context("when the queue is empty", func() {
			BeforeEach(func() {
				q = NewQueue([]int{})
			})
			It("should return an error", func() {
				value, err := q.Peek()
				Expect(err).To(MatchError(ErrEmptyQueue))
				Expect(value).To(BeNil())
			})
		})
	})
})

func TestQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queue Suite")
}
