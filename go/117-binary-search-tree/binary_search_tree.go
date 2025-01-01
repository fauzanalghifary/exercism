package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{
		left:  nil,
		data:  i,
		right: nil,
	}
}

func (bst *BinarySearchTree) Insert(i int) {
	if i <= bst.data {
		if bst.left == nil {
			bst.left = NewBst(i)
		} else {
			bst.left.Insert(i)
		}
	} else {
		if bst.right == nil {
			bst.right = NewBst(i)
		} else {
			bst.right.Insert(i)
		}
	}
}

func (bst *BinarySearchTree) SortedData() []int {
	var sorted []int
	Traverse(bst, &sorted)
	return sorted
}

func Traverse(bst *BinarySearchTree, sorted *[]int) {
	if bst.left != nil {
		Traverse(bst.left, sorted)
	}

	*sorted = append(*sorted, bst.data)

	if bst.right != nil {
		Traverse(bst.right, sorted)
	}
}

//func (bst *BinarySearchTree) SortedData() []int {
//	var sorted []int
//
//	if bst.left != nil {
//		sorted = append(sorted, bst.left.SortedData()...)
//	}
//	sorted = append(sorted, bst.data)
//	if bst.right != nil {
//		sorted = append(sorted, bst.right.SortedData()...)
//	}
//
//	return sorted
//}
