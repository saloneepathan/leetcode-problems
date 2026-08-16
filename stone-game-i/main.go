package main

import "fmt"

func stoneGameIX(stones []int) bool {
	cnt := [3]int{}

	for _, stone := range stones {
		cnt[stone%3]++
	}

	// If the number of remainder-0 stones is even,
	// Alice wins iff both remainder-1 and remainder-2 stones exist.
	if cnt[0]%2 == 0 {
		return cnt[1] > 0 && cnt[2] > 0
	}

	// If the number of remainder-0 stones is odd,
	// Alice wins iff the counts of remainder-1 and remainder-2
	// differ by more than 2.
	diff := cnt[1] - cnt[2]
	if diff < 0 {
		diff = -diff
	}

	return diff > 2
}

func main() {
	tests := [][]int{
		{2, 1},
		{2},
		{5, 1, 2, 4, 3},
	}

	for _, stones := range tests {
		fmt.Println(stoneGameIX(stones))
	}
}
