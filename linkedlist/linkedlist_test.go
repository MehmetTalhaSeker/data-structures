package linkedlist_test

import (
	"data-structures/linkedlist"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("LinkedList", func() {
	var (
		ll  linkedlist.LinkedList[int]
		err error
		nn  *linkedlist.Node[int]
	)

	BeforeEach(func() {
		node1 := &linkedlist.Node[int]{Value: 1}
		node2 := &linkedlist.Node[int]{Value: 2}
		node3 := &linkedlist.Node[int]{Value: 3}

		ll = linkedlist.NewLinkedList(node1, node2, node3)
	})

	Describe("Add", func() {
		It("should add an element to the linked list", func() {
			ll.Add(4)
			Expect(ll.Size()).To(Equal(4))
		})
	})

	Describe("Remove", func() {
		Context("when the linked list is not empty", func() {
			It("should remove an element from the linked list", func() {
				err = ll.Remove(2)
				Expect(err).ToNot(HaveOccurred())
				Expect(ll.Size()).To(Equal(2))
			})

			It("should not remove an element from the linked list", func() {
				err = ll.Remove(4)
				Expect(err).ToNot(HaveOccurred())
				Expect(ll.Size()).To(Equal(3))
			})
		})

		Context("when the linked list is empty", func() {
			BeforeEach(func() {
				ll = linkedlist.NewLinkedList(nn)
			})

			It("should return an error", func() {
				err = ll.Remove(2)
				Expect(err).To(MatchError(linkedlist.ErrEmptyLinkedList))
			})
		})
	})

	Describe("Size", func() {
		Context("when the linked list is not empty", func() {
			It("should return the size of the linked list", func() {
				size := ll.Size()
				Expect(size).To(Equal(3))
			})
		})

		Context("when the linked list is empty", func() {
			BeforeEach(func() {
				ll = linkedlist.NewLinkedList(nn)
			})

			It("should return 0", func() {
				size := ll.Size()
				Expect(size).To(Equal(0))
			})
		})
	})
})

func TestLinkedList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LinkedList Suite")
}
