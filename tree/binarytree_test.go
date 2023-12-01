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
		rn *tree.Node[int]
	)

	Describe("Non recursive", func() {
		BeforeEach(func() {
			root := &tree.Node[int]{Value: 1}
			root.Left = &tree.Node[int]{Value: 2}
			root.Right = &tree.Node[int]{Value: 3}

			root.Left.Left = &tree.Node[int]{Value: 4}
			root.Left.Right = &tree.Node[int]{Value: 5}

			root.Right.Left = &tree.Node[int]{Value: 6}
			root.Right.Right = &tree.Node[int]{Value: 7}

			bt = *tree.NewBinaryTree[int](root)
		})

		Describe("PreOrderTraversal", func() {
			Context("when the tree is not empty", func() {
				It("should return elements in pre-order traversal order", func() {
					result := bt.PreOrderTraversal()
					Expect(result).To(Equal([]int{1, 2, 4, 5, 3, 6, 7}))
				})
				It("should return elements in pre-order traversal order", func() {
					bt.Root.Left.Left.Left = &tree.Node[int]{Value: 555}

					result := bt.PreOrderTraversal()
					Expect(result).To(Equal([]int{1, 2, 4, 555, 5, 3, 6, 7}))
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
		Describe("PostOrderTraversal", func() {
			Context("when the tree is not empty", func() {
				It("should return elements in post-order traversal order", func() {
					result := bt.PostOrderTraversal()
					Expect(result).To(Equal([]int{4, 5, 2, 6, 7, 3, 1}))
				})
				It("should return elements in post-order traversal order", func() {
					bt.Root.Left.Left.Left = &tree.Node[int]{Value: 555}

					result := bt.PostOrderTraversal()
					Expect(result).To(Equal([]int{555, 4, 5, 2, 6, 7, 3, 1}))
				})
			})

			Context("when the tree is empty", func() {
				BeforeEach(func() {
					bt = *tree.NewBinaryTree(nn)
				})

				It("should return nil", func() {
					result := bt.PostOrderTraversal()
					Expect(result).To(BeNil())
				})
			})
		})
		Describe("InOrderTraversal", func() {
			Context("when the tree is not empty", func() {
				It("should return elements in in-order traversal order", func() {
					result := bt.InOrderTraversal()
					Expect(result).To(Equal([]int{4, 2, 5, 1, 6, 3, 7}))
				})
				It("should return elements in in-order traversal order", func() {
					bt.Root.Left.Left.Left = &tree.Node[int]{Value: 555}

					result := bt.InOrderTraversal()
					Expect(result).To(Equal([]int{555, 4, 2, 5, 1, 6, 3, 7}))
				})
			})

			Context("when the tree is empty", func() {
				BeforeEach(func() {
					bt = *tree.NewBinaryTree(nn)
				})

				It("should return nil", func() {
					result := bt.InOrderTraversal()
					Expect(result).To(BeNil())
				})
			})
		})
	})

	Describe("recursive", func() {
		BeforeEach(func() {
			rn = &tree.Node[int]{Value: 1}
			rn.Left = &tree.Node[int]{Value: 2}
			rn.Right = &tree.Node[int]{Value: 3}

			rn.Left.Left = &tree.Node[int]{Value: 4}
			rn.Left.Right = &tree.Node[int]{Value: 5}

			rn.Right.Left = &tree.Node[int]{Value: 6}
			rn.Right.Right = &tree.Node[int]{Value: 7}
		})

		Describe("PreOrderTraversal Recursive", func() {
			Context("when the tree is not empty", func() {
				It("should return elements in pre-order traversal order", func() {
					result := tree.PreOrderTraversalRecursive(rn)
					Expect(result).To(Equal([]int{1, 2, 4, 5, 3, 6, 7}))
				})
				It("should return elements in pre-order traversal order", func() {
					rn.Left.Left.Left = &tree.Node[int]{Value: 555}

					result := tree.PreOrderTraversalRecursive(rn)
					Expect(result).To(Equal([]int{1, 2, 4, 555, 5, 3, 6, 7}))
				})
			})
		})
		Describe("PostOrderTraversal Recursive", func() {
			Context("when the tree is not empty", func() {
				It("should return elements in post-order traversal order", func() {
					result := tree.PostOrderTraversalRecursive(rn)
					Expect(result).To(Equal([]int{4, 5, 2, 6, 7, 3, 1}))
				})
				It("should return elements in post-order traversal order", func() {
					rn.Left.Left.Left = &tree.Node[int]{Value: 555}

					result := tree.PostOrderTraversalRecursive(rn)
					Expect(result).To(Equal([]int{555, 4, 5, 2, 6, 7, 3, 1}))
				})
			})
		})
		Describe("InOrderTraversal Recursive", func() {
			Context("when the tree is not empty", func() {
				It("should return elements in in-order traversal order", func() {
					result := tree.InOrderTraversalRecursive(rn)
					Expect(result).To(Equal([]int{4, 2, 5, 1, 6, 3, 7}))
				})
				It("should return elements in in-order traversal order", func() {
					rn.Left.Left.Left = &tree.Node[int]{Value: 555}

					result := tree.InOrderTraversalRecursive(rn)
					Expect(result).To(Equal([]int{555, 4, 2, 5, 1, 6, 3, 7}))
				})
			})
		})
	})

})

func TestBinaryTree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BinaryTree Suite")
}
