package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := make(map[int]*TreeNode)
	hasParent := make(map[int]bool)

	for _, d := range descriptions {
		parentVal := d[0]
		childVal := d[1]
		isLeft := d[2]

		if _, exists := nodes[parentVal]; !exists {
			nodes[parentVal] = &TreeNode{Val: parentVal}
		}

		if _, exists := nodes[childVal]; !exists {
			nodes[childVal] = &TreeNode{Val: childVal}
		}

		parent := nodes[parentVal]
		child := nodes[childVal]

		if isLeft == 1 {
			parent.Left = child
		} else {
			parent.Right = child
		}

		hasParent[childVal] = true
	}

	for val, node := range nodes {
		if !hasParent[val] {
			return node
		}
	}

	return nil
}

// Helper: Level-order traversal for testing
func levelOrder(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}

	var result []interface{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
			continue
		}

		result = append(result, node.Val)
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	// Remove trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

func main() {
	descriptions := [][]int{
		{20, 15, 1},
		{20, 17, 0},
		{50, 20, 1},
		{50, 80, 0},
		{80, 19, 1},
	}

	root := createBinaryTree(descriptions)

	fmt.Println(levelOrder(root))
	// Output: [50 20 80 15 17 19]
}
