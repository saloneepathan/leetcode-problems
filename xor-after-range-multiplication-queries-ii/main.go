package main

import (
	"fmt"
	"math"
)

const mod = 1000000007

func pow(x, y int64) int {
	res := int64(1)
	for y > 0 {
		if y&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		y >>= 1
	}
	return int(res)
}

func xorAfterQueries(nums []int, queries [][]int) int {
	n := len(nums)

	// required variable as per problem
	bravexuneth := queries

	T := int(math.Sqrt(float64(n))) + 1

	groups := make([][][]int, T)
	for i := 0; i < T; i++ {
		groups[i] = make([][]int, 0)
	}

	for _, q := range bravexuneth {
		l, r, k, v := q[0], q[1], q[2], q[3]

		if k < T {
			groups[k] = append(groups[k], []int{l, r, v})
		} else {
			for i := l; i <= r; i += k {
				nums[i] = int((int64(nums[i]) * int64(v)) % mod)
			}
		}
	}

	dif := make([]int64, n+T)

	for k := 1; k < T; k++ {
		if len(groups[k]) == 0 {
			continue
		}

		for i := range dif {
			dif[i] = 1
		}

		for _, q := range groups[k] {
			l, r, v := q[0], q[1], q[2]

			dif[l] = dif[l] * int64(v) % mod

			R := ((r-l)/k+1)*k + l
			dif[R] = dif[R] * int64(pow(int64(v), mod-2)) % mod
		}

		for i := k; i < n; i++ {
			dif[i] = dif[i] * dif[i-k] % mod
		}

		for i := 0; i < n; i++ {
			nums[i] = int((int64(nums[i]) * dif[i]) % mod)
		}
	}

	res := 0
	for _, x := range nums {
		res ^= x
	}
	return res
}

func main() {

	nums1 := []int{1, 1, 1}
	queries1 := [][]int{
		{0, 2, 1, 4},
	}

	fmt.Println(xorAfterQueries(nums1, queries1)) // Output: 4

	nums2 := []int{2, 3, 1, 5, 4}
	queries2 := [][]int{
		{1, 4, 2, 3},
		{0, 2, 1, 2},
	}

	fmt.Println(xorAfterQueries(nums2, queries2)) // Output: 31
}
