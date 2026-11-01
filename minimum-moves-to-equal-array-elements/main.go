package main

import "fmt"

func minMoves(nums []int) int {
    minVal := nums[0]

    // Find the minimum element
    for _, num := range nums {
        if num < minVal {
            minVal = num
        }
    }

    // Calculate total moves
    moves := 0

    for _, num := range nums {
        moves += num - minVal
    }

    return moves
}

func main() {
    // Test Case 1
    nums1 := []int{1, 2, 3}
    fmt.Println("Test Case 1:", minMoves(nums1)) // Output: 3

    // Test Case 2
    nums2 := []int{1, 1, 1}
    fmt.Println("Test Case 2:", minMoves(nums2)) // Output: 0

    // Test Case 3
    nums3 := []int{5, 6, 8, 8}
    fmt.Println("Test Case 3:", minMoves(nums3)) // Output: 7
}