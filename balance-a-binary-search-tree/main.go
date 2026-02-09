package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	vals := []int{}
	inorder(root, &vals)
	return buildBalancedBST(vals, 0, len(vals)-1)
}

// Inorder traversal to collect sorted values
func inorder(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, vals)
	*vals = append(*vals, root.Val)
	inorder(root.Right, vals)
}

// Build balanced BST from sorted array
func buildBalancedBST(vals []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	node := &TreeNode{Val: vals[mid]}

	node.Left = buildBalancedBST(vals, left, mid-1)
	node.Right = buildBalancedBST(vals, mid+1, right)

	return node
}

// Helper function to print inorder traversal
func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Print(root.Val, " ")
	printInorder(root.Right)
}

func main() {
	// Example: Unbalanced BST
	/*
	   1
	    \
	     2
	      \
	       3
	        \
	         4
	*/

	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 3}
	root.Right.Right.Right = &TreeNode{Val: 4}

	fmt.Println("Inorder of original BST:")
	printInorder(root)
	fmt.Println()

	balancedRoot := balanceBST(root)

	fmt.Println("Inorder of balanced BST:")
	printInorder(balancedRoot)
	fmt.Println()
}
