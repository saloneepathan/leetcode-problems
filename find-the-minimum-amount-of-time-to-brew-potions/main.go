package main

import "fmt"

func minTime(skill []int, mana []int) int64 {
	n := len(skill)
	m := len(mana)
	if n == 0 || m == 0 {
		return 0
	}

	// prefPrev[i] = prefix sum up to i (inclusive) of processing times for the previous job
	prefPrev := make([]int64, n)
	var acc int64 = 0
	for i := 0; i < n; i++ {
		acc += int64(skill[i]) * int64(mana[0])
		prefPrev[i] = acc
	}

	var Sprev int64 = 0

	for j := 1; j < m; j++ {
		// Compute S_j = max_i ( Sprev + prefPrev[i] - prefBefore[i] )
		// Here prefBefore[i] is the prefix up to i-1 for current job j,
		// we compute it on the fly with cumulativeBefore.
		Sj := int64(-1 << 63) // MinInt64 typed as int64
		var cumulativeBefore int64 = 0
		for i := 0; i < n; i++ {
			cand := Sprev + prefPrev[i] - cumulativeBefore
			if cand > Sj {
				Sj = cand
			}
			cumulativeBefore += int64(skill[i]) * int64(mana[j])
		}

		// Update prefPrev to be the inclusive prefix sums for job j
		var acc2 int64 = 0
		for i := 0; i < n; i++ {
			acc2 += int64(skill[i]) * int64(mana[j])
			prefPrev[i] = acc2
		}

		Sprev = Sj
	}

	// makespan = start time of last job + its total processing time
	return Sprev + prefPrev[n-1]
}

func main() {
	fmt.Println(minTime([]int{1, 5, 2, 4}, []int{5, 1, 4, 2})) // 110
	fmt.Println(minTime([]int{1, 1, 1}, []int{1, 1, 1}))       // 5
	fmt.Println(minTime([]int{1, 2, 3, 4}, []int{1, 2}))       // 21
}
