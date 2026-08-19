package main

import "fmt"

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	// Store reserved seats for each row as a bitmask.
	// Only seats 2 through 9 matter.
	rows := make(map[int]int)

	for _, seat := range reservedSeats {
		row := seat[0]
		col := seat[1]

		if col >= 2 && col <= 9 {
			rows[row] |= 1 << col
		}
	}

	// Every completely empty row can accommodate 2 families.
	ans := 2 * n

	for _, mask := range rows {
		// Seats:
		// Left   = 2,3,4,5
		// Middle = 4,5,6,7
		// Right  = 6,7,8,9

		leftFree := (mask & ((1 << 2) | (1 << 3) | (1 << 4) | (1 << 5))) == 0

		middleFree := (mask & ((1 << 4) | (1 << 5) | (1 << 6) | (1 << 7))) == 0

		rightFree := (mask & ((1 << 6) | (1 << 7) | (1 << 8) | (1 << 9))) == 0

		groups := 0

		// Left and right blocks don't overlap,
		// so they can both be used.
		if leftFree {
			groups++
		}

		if rightFree {
			groups++
		}

		// If neither left nor right works,
		// check whether the middle block works.
		if groups == 0 && middleFree {
			groups = 1
		}

		// We initially assumed this row could fit 2 groups.
		// Replace that assumption with the actual number.
		ans -= 2 - groups
	}

	return ans
}

func main() {
	// Example 1
	n1 := 3
	reserved1 := [][]int{
		{1, 2},
		{1, 3},
		{1, 8},
		{2, 6},
		{3, 1},
		{3, 10},
	}

	fmt.Println(maxNumberOfFamilies(n1, reserved1))
	// Output: 4

	// Example 2
	n2 := 2
	reserved2 := [][]int{
		{2, 1},
		{1, 8},
		{2, 6},
	}

	fmt.Println(maxNumberOfFamilies(n2, reserved2))
	// Output: 2

	// Example 3
	n3 := 4
	reserved3 := [][]int{
		{4, 3},
		{1, 4},
		{4, 6},
		{1, 7},
	}

	fmt.Println(maxNumberOfFamilies(n3, reserved3))
	// Output: 4
}
