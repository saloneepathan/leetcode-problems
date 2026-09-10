package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	ans := 0

	var dfs func(*TreeNode) (int, int)

	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		// Get sum and count from left subtree
		leftSum, leftCount := dfs(node.Left)

		// Get sum and count from right subtree
		rightSum, rightCount := dfs(node.Right)

		// Include current node
		sum := leftSum + rightSum + node.Val
		count := leftCount + rightCount + 1

		// Check if current node equals subtree average
		if node.Val == sum/count {
			ans++
		}

		return sum, count
	}

	dfs(root)

	return ans
}

func main() {
	/*
	        4
	       / \
	      8   5
	     / \   \
	    0   1   6
	*/

	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	result := averageOfSubtree(root)

	fmt.Println("Number of nodes equal to their subtree average:", result)
}