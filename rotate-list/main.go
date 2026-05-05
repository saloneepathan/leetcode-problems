package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Step 1: Find length and tail
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	// Step 2: Optimize k
	k = k % length
	if k == 0 {
		return head
	}

	// Step 3: Make circular
	tail.Next = head

	// Step 4: Find new tail
	stepsToNewTail := length - k
	newTail := head
	for i := 1; i < stepsToNewTail; i++ {
		newTail = newTail.Next
	}

	// Step 5: Break the circle
	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}

// Helper: Create linked list from slice
func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, v := range nums[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

// Helper: Print linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example 1
	head1 := buildList([]int{1, 2, 3, 4, 5})
	k1 := 2
	result1 := rotateRight(head1, k1)
	printList(result1) // Expected: 4 -> 5 -> 1 -> 2 -> 3

	// Example 2
	head2 := buildList([]int{0, 1, 2})
	k2 := 4
	result2 := rotateRight(head2, k2)
	printList(result2) // Expected: 2 -> 0 -> 1
}
