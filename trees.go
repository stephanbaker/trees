package trees

import (
	"fmt"
	"strings"
)

type BinarySearchTree struct {
	Left  *BinarySearchTree
	Value string
	Right *BinarySearchTree
}

func (t *BinarySearchTree) Walk(ch chan string) {
	defer close(ch)

	if t == nil || (t.Left == nil && t.Right == nil) {
		return
	}

	var walk func(tree *BinarySearchTree)
	walk = func(tree *BinarySearchTree) {
		if tree.Left != nil {
			walk(tree.Left)
		}

		ch <- tree.Value

		if tree.Right != nil {
			walk(tree.Right)
		}
	}
	walk(t)
}

func (t *BinarySearchTree) PrintTree() {
	if t == nil {
		fmt.Println("Tree is empty")
		return
	}

	var printTree func(tree *BinarySearchTree, level int) string
	printTree = func(tree *BinarySearchTree, level int) string {
		if tree == nil {
			return ""
		}

		s := fmt.Sprintf("\"%v\"", tree.Value)
		if tree.Left == nil && tree.Right == nil {
			return s
		}

		var spacer string
		if level > 0 {
			spacer = strings.Repeat("│  ", level)
		}
		for _, child := range []*BinarySearchTree{tree.Left, tree.Right} {
			s += fmt.Sprintf("\n%s├──%s", spacer, printTree(child, level+1))
		}
		return s
	}

	fmt.Println(printTree(t, 0))
}

func (t *BinarySearchTree) Size() int {
	if t == nil {
		return 0
	}

	children := 1
	if t.Left != nil {
		children += t.Left.Size()
	}
	if t.Right != nil {
		children += t.Right.Size()
	}
	return children
}

func (t *BinarySearchTree) Height() int {
	if t == nil || (t.Left == nil && t.Right == nil) {
		return 0
	}

	var leftHeight, rightHeight int
	if t.Left != nil {
		leftHeight = t.Left.Height()
	}
	if t.Right != nil {
		rightHeight = t.Right.Height()
	}

	if leftHeight > rightHeight {
		return 1 + leftHeight
	} else {
		return 1 + rightHeight
	}
}

func (t *BinarySearchTree) Insert(input string) {
	var insert func(tree *BinarySearchTree, val string)
	insert = func(tree *BinarySearchTree, val string) {
		if input < tree.Value {
			if tree.Left == nil {
				tree.Left = &BinarySearchTree{Value: val}
			} else {
				tree.Left.Insert(val)
			}
		} else {
			if tree.Right == nil {
				tree.Right = &BinarySearchTree{Value: val}
			} else {
				tree.Right.Insert(val)
			}
		}
	}
	insert(t, input)
	t.Balance()
}

func (t *BinarySearchTree) RotateLeft() {
	if t == nil || t.Right == nil {
		return
	}

	original := *t
	newRoot := *t.Right
	if newRoot.Left == nil {
		original.Right = nil
	} else {
		original.Right = newRoot.Left
	}
	newRoot.Left = &original
	*t = newRoot
}

func (t *BinarySearchTree) RotateRight() {
	if t == nil || t.Left == nil {
		return
	}

	original := *t
	newRoot := *t.Left
	if newRoot.Right == nil {
		original.Left = nil
	} else {
		original.Left = newRoot.Right
	}
	newRoot.Right = &original
	*t = newRoot
}

func (t *BinarySearchTree) IsBalanced() bool {
	bf := t.BalanceFactor()
	return bf >= -1 && bf <= 1
}

func (t *BinarySearchTree) BalanceFactor() int {
	if t == nil || (t.Left == nil && t.Right == nil) {
		return 0
	}

	return t.Left.Height() - t.Right.Height()
}

func (t *BinarySearchTree) Balance() {
	for !t.IsBalanced() {
		if t.BalanceFactor() < 0 {
			t.RotateLeft()
		} else {
			t.RotateRight()
		}
	}
}
