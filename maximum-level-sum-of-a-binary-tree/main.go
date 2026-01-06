package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	level := 1
	maxSum := root.Val
	answerLevel := 1

	for len(queue) > 0 {
		size := len(queue)
		levelSum := 0

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if levelSum > maxSum {
			maxSum = levelSum
			answerLevel = level
		}

		level++
	}

	return answerLevel
}

func main() {
	// Example 1
	// Tree: [1,7,0,7,-8,null,null]
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 7}
	root1.Right = &TreeNode{Val: 0}
	root1.Left.Left = &TreeNode{Val: 7}
	root1.Left.Right = &TreeNode{Val: -8}

	fmt.Println(maxLevelSum(root1)) // Output: 2

	// Example 2
	// Tree: [989,null,10250,98693,-89388,null,null,null,-32127]
	root2 := &TreeNode{Val: 989}
	root2.Right = &TreeNode{Val: 10250}
	root2.Right.Left = &TreeNode{Val: 98693}
	root2.Right.Right = &TreeNode{Val: -89388}
	root2.Right.Left.Right = &TreeNode{Val: -32127}

	fmt.Println(maxLevelSum(root2)) // Output: 2
}
