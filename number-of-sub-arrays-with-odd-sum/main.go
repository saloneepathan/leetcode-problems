package main

import (
	"fmt"
	"math/rand"
)

const mod = 1_000_000_007

// numOfSubarrays counts subarrays with an odd sum.
// A subarray (i, j] has an odd sum iff prefix parities at i and j differ.
// So for each new prefix, add the number of earlier prefixes with the opposite parity.
func numOfSubarrays(arr []int) int {
	even, odd := 1, 0 // empty prefix (sum 0) is even
	parity := 0
	ans := 0
	for _, x := range arr {
		parity = (parity + x) & 1
		if parity == 1 {
			ans += even
			odd++
		} else {
			ans += odd
			even++
		}
		ans %= mod
	}
	return ans
}

func brute(arr []int) int {
	cnt := 0
	for i := range arr {
		s := 0
		for j := i; j < len(arr); j++ {
			s += arr[j]
			if s%2 == 1 {
				cnt++
			}
		}
	}
	return cnt % mod
}

func main() {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 3, 5}, 4},
		{[]int{2, 4, 6}, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 16},
		{[]int{1}, 1},
		{[]int{2}, 0},
		{[]int{1, 1}, 2},
	}
	ok := true
	for _, t := range tests {
		got := numOfSubarrays(t.arr)
		fmt.Printf("input=%v expected=%d actual=%d\n", t.arr, t.want, got)
		if got != t.want {
			ok = false
		}
	}
	// large input: all ones, n=1e5 -> compare against closed form-free brute is too slow; check modulo path runs
	big := make([]int, 100000)
	for i := range big {
		big[i] = 1
	}
	fmt.Println("big all-ones result:", numOfSubarrays(big))
	// random cross-check
	r := rand.New(rand.NewSource(1))
	for k := 0; k < 2000; k++ {
		n := 1 + r.Intn(12)
		a := make([]int, n)
		for i := range a {
			a[i] = 1 + r.Intn(100)
		}
		if numOfSubarrays(a) != brute(a) {
			fmt.Println("MISMATCH", a)
			ok = false
		}
	}
	fmt.Println("all passed:", ok)
}
