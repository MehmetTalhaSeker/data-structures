package stack

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Stack", func() {
	var (
		s Stack[int]
	)

	Describe("Push", func() {
		BeforeEach(func() {
			s = NewStack([]int{1, 2, 3})
		})
		It("should push an element onto the stack", func() {
			s.Push(4)

			l := len(s.getter())
			els := s.getter()

			Expect(els[l-1]).To(Equal(4))
		})
		It("should push an element onto the stack", func() {
			s.Push(21)

			l := len(s.getter())
			els := s.getter()

			Expect(els[l-1]).To(Equal(21))
		})
	})

	Describe("Pop", func() {
		Context("when the stack is not empty", Ordered, func() {
			BeforeAll(func() {
				s = NewStack([]int{1, 2, 3})
			})
			It("should pop the top element from the stack", func() {
				value, err := s.Pop()
				Expect(err).ToNot(HaveOccurred())
				Expect(*value).To(Equal(3))
			})
			It("should pop the top element from the stack", func() {
				value, err := s.Pop()
				Expect(err).ToNot(HaveOccurred())
				Expect(*value).To(Equal(2))
			})
			It("should pop the top element from the stack", func() {
				value, err := s.Pop()
				Expect(err).ToNot(HaveOccurred())
				Expect(*value).To(Equal(1))
			})
			It("should return an error", func() {
				value, err := s.Pop()
				Expect(err).To(MatchError(ErrEmptyStack))
				Expect(value).To(BeNil())
			})
		})

		Context("when the stack is empty", func() {
			BeforeEach(func() {
				s = NewStack([]int{})
			})
			It("should return an error", func() {
				value, err := s.Pop()
				Expect(err).To(MatchError(ErrEmptyStack))
				Expect(value).To(BeNil())
			})
		})
	})

	Describe("Peek", func() {
		Context("when the stack is not empty", Ordered, func() {
			BeforeAll(func() {
				s = NewStack([]int{1, 2, 3})
			})
			It("should return the top element from the stack", func() {
				value, err := s.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(*value).To(Equal(3))
			})
			It("should return the top element from the stack", func() {
				_, _ = s.Pop()

				value, err := s.Peek()
				Expect(err).ToNot(HaveOccurred())
				Expect(*value).To(Equal(2))
			})
		})

		Context("when the stack is empty", func() {
			BeforeEach(func() {
				s = NewStack([]int{})
			})
			It("should return an error", func() {
				value, err := s.Peek()
				Expect(err).To(MatchError(ErrEmptyStack))
				Expect(value).To(BeNil())
			})
		})
	})
})

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stack Suite")
}
