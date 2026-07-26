package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next

	return head
}

// Helper: create linked list from slice
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	curr := head

	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}

	return head
}

// Helper: print linked list
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
	head := createList([]int{1, 3, 4, 7, 1, 2, 6})

	fmt.Println("Original List:")
	printList(head)

	head = deleteMiddle(head)

	fmt.Println("After Deleting Middle:")
	printList(head)

	// Example 2
	head2 := createList([]int{1, 2, 3, 4})
	fmt.Println("\nOriginal List:")
	printList(head2)

	head2 = deleteMiddle(head2)

	fmt.Println("After Deleting Middle:")
	printList(head2)

	// Example 3
	head3 := createList([]int{2, 1})
	fmt.Println("\nOriginal List:")
	printList(head3)

	head3 = deleteMiddle(head3)

	fmt.Println("After Deleting Middle:")
	printList(head3)
}
