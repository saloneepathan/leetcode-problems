package main

import "fmt"

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// nodesBetweenCriticalPoints returns:
// [minimum distance between any two critical points,
//
//	maximum distance between the first and last critical points]
func nodesBetweenCriticalPoints(head *ListNode) []int {
	const INF = int(1e9)

	minDist := INF
	maxDist := -1

	first := -1
	prev := -1

	index := 1

	// We need previous, current, and next nodes.
	prevNode := head
	curr := head.Next

	for curr != nil && curr.Next != nil {
		next := curr.Next

		// Check if curr is a local maximum or minimum.
		isCritical := (curr.Val > prevNode.Val && curr.Val > next.Val) ||
			(curr.Val < prevNode.Val && curr.Val < next.Val)

		if isCritical {
			if first == -1 {
				// First critical point.
				first = index
			} else {
				// Distance from previous critical point.
				minDist = min(minDist, index-prev)

				// Distance from first critical point.
				maxDist = index - first
			}

			// Update previous critical point.
			prev = index
		}

		prevNode = curr
		curr = next
		index++
	}

	// Fewer than two critical points.
	if first == -1 || maxDist == -1 {
		return []int{-1, -1}
	}

	return []int{minDist, maxDist}
}

// min returns the smaller of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// createList creates a linked list from a slice of integers.
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	curr := head

	for i := 1; i < len(values); i++ {
		curr.Next = &ListNode{Val: values[i]}
		curr = curr.Next
	}

	return head
}

func main() {
	// Example:
	// [5, 3, 1, 2, 5, 1, 2]
	//
	// Critical points:
	// 1 -> local minimum
	// 2 -> local maximum
	// 1 -> local minimum
	//
	// Distances: 1 and 2
	// Answer: [1, 4]
	values := []int{5, 3, 1, 2, 5, 1, 2}

	head := createList(values)

	result := nodesBetweenCriticalPoints(head)

	fmt.Println(result)
}
