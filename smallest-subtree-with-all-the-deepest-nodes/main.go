package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Main solution
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	node, _ := dfs(root)
	return node
}

// DFS returns (subtreeRoot, depth)
func dfs(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	leftNode, leftDepth := dfs(root.Left)
	rightNode, rightDepth := dfs(root.Right)

	if leftDepth > rightDepth {
		return leftNode, leftDepth + 1
	}
	if rightDepth > leftDepth {
		return rightNode, rightDepth + 1
	}

	// Equal depth → current node is LCA
	return root, leftDepth + 1
}

// ---------- Helpers for Testing ----------

// Build tree from level-order array (-1 represents null)
func buildTree(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		if arr[i] != -1 {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(arr) && arr[i] != -1 {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// Print subtree (preorder)
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// Example 1
	root1 := buildTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})
	ans1 := subtreeWithAllDeepest(root1)
	fmt.Print("Output 1: ")
	printTree(ans1)
	fmt.Println()

	// Example 2
	root2 := buildTree([]int{1})
	ans2 := subtreeWithAllDeepest(root2)
	fmt.Print("Output 2: ")
	printTree(ans2)
	fmt.Println()

	// Example 3
	root3 := buildTree([]int{0, 1, 3, -1, 2})
	ans3 := subtreeWithAllDeepest(root3)
	fmt.Print("Output 3: ")
	printTree(ans3)
	fmt.Println()
}
