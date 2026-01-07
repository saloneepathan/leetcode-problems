package main

import "fmt"

const MOD int64 = 1e9 + 7

// TreeNode definition for local environment
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	totalSum := treeSum(root)
	var maxProd int64 = 0

	var dfs func(*TreeNode) int64
	dfs = func(node *TreeNode) int64 {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		currSum := int64(node.Val) + left + right

		product := currSum * (totalSum - currSum)
		if product > maxProd {
			maxProd = product
		}

		return currSum
	}

	dfs(root)
	return int(maxProd % MOD)
}

func treeSum(node *TreeNode) int64 {
	if node == nil {
		return 0
	}
	return int64(node.Val) + treeSum(node.Left) + treeSum(node.Right)
}

func main() {
	/*
	   Example 1:
	   Tree:
	           1
	          / \
	         2   3
	        / \   \
	       4   5   6
	*/

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	result := maxProduct(root)
	fmt.Println("Maximum Product:", result) // Expected: 110
}
