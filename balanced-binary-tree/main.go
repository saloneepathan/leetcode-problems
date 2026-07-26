package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := height(node.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := height(node.Right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	// Example 1: Balanced tree
	// [3,9,20,null,null,15,7]
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}

	fmt.Println("Example 1:", isBalanced(root1)) // true

	// Example 2: Unbalanced tree
	// [1,2,2,3,3,null,null,4,4]
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 2}
	root2.Left.Left = &TreeNode{Val: 3}
	root2.Left.Right = &TreeNode{Val: 3}
	root2.Left.Left.Left = &TreeNode{Val: 4}
	root2.Left.Left.Right = &TreeNode{Val: 4}

	fmt.Println("Example 2:", isBalanced(root2)) // false

	// Example 3: Empty tree
	fmt.Println("Example 3:", isBalanced(nil)) // true
}
