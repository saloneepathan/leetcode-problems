package main

import (
	"fmt"
)

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to remove nodes whose values exist in nums
func modifiedList(nums []int, head *ListNode) *ListNode {
	// Step 1: Create a set for quick lookup
	removeSet := make(map[int]bool)
	for _, num := range nums {
		removeSet[num] = true
	}

	// Step 2: Use a dummy node to simplify edge cases
	dummy := &ListNode{Next: head}
	curr := dummy

	// Step 3: Iterate and remove nodes
	for curr.Next != nil {
		if removeSet[curr.Next.Val] {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return dummy.Next
}

// Helper function: Build linked list from slice
func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, val := range nums[1:] {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return head
}

// Helper function: Print linked list
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	// Example 1
	nums1 := []int{1, 2, 3}
	head1 := buildList([]int{1, 2, 3, 4, 5})
	fmt.Print("Input:  [1,2,3,4,5], nums=[1,2,3] => Output: ")
	printList(modifiedList(nums1, head1))

	// Example 2
	nums2 := []int{1}
	head2 := buildList([]int{1, 2, 1, 2, 1, 2})
	fmt.Print("Input:  [1,2,1,2,1,2], nums=[1] => Output: ")
	printList(modifiedList(nums2, head2))

	// Example 3
	nums3 := []int{5}
	head3 := buildList([]int{1, 2, 3, 4})
	fmt.Print("Input:  [1,2,3,4], nums=[5] => Output: ")
	printList(modifiedList(nums3, head3))
}
