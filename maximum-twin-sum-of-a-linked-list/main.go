package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// Find the middle of the linked list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half
	var prev *ListNode
	curr := slow

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// Compute maximum twin sum
	maxSum := 0
	p1, p2 := head, prev

	for p2 != nil {
		sum := p1.Val + p2.Val
		if sum > maxSum {
			maxSum = sum
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return maxSum
}

// Helper function to create a linked list from a slice
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}

	return dummy.Next
}

func main() {
	head1 := createList([]int{5, 4, 2, 1})
	fmt.Println(pairSum(head1)) // 6

	head2 := createList([]int{4, 2, 2, 3})
	fmt.Println(pairSum(head2)) // 7

	head3 := createList([]int{1, 100000})
	fmt.Println(pairSum(head3)) // 100001
}
