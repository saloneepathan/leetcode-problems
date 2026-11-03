
package main

import "fmt"

func minMoves(nums []int) int {
    maxNum := nums[0]

    // Find maximum element
    for _, num := range nums {
        if num > maxNum {
            maxNum = num
        }
    }

    // Calculate minimum moves
    moves := 0

    for _, num := range nums {
        moves += maxNum - num
    }

    return moves
}

func main() {
    nums := []int{2, 1, 3}

    result := minMoves(nums)

    fmt.Println(result) // Output: 3
}
