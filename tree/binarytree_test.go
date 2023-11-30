package tree_test

import (
	"data-structures/tree"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("BinaryTree", func() {
	var (
		bt tree.BinaryTree[int]
		nn *tree.Node[int]
	)

	BeforeEach(func() {
		root := &tree.Node[int]{Value: 1}
		root.Left = &tree.Node[int]{Value: 2}
		root.Right = &tree.Node[int]{Value: 3}
		root.Left.Left = &tree.Node[int]{Value: 4}
		root.Left.Right = &tree.Node[int]{Value: 5}

		bt = *tree.NewBinaryTree[int](root)
	})

	Describe("PreOrderTraversal", func() {
		Context("when the tree is not empty", func() {
			It("should return elements in pre-order traversal order", func() {
				result := bt.PreOrderTraversal()
				Expect(result).To(Equal([]int{1, 2, 4, 5, 3}))
			})
			It("should return elements in pre-order traversal order", func() {
				bt.Root.Left.Left.Left = &tree.Node[int]{Value: 555}

				result := bt.PreOrderTraversal()
				Expect(result).To(Equal([]int{1, 2, 4, 555, 5, 3}))
			})
		})

		Context("when the tree is empty", func() {
			BeforeEach(func() {
				bt = *tree.NewBinaryTree(nn)
			})

			It("should return nil", func() {
				result := bt.PreOrderTraversal()
				Expect(result).To(BeNil())
			})
		})
	})
})

func TestBinaryTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BinaryTree Suite")
}
