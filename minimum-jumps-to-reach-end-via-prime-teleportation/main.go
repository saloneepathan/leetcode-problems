package main

import (
	"fmt"
)

func minJumps(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	// Smallest Prime Factor sieve
	spf := make([]int, maxVal+1)
	for i := 2; i <= maxVal; i++ {
		if spf[i] == 0 {
			for j := i; j <= maxVal; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}

	// prime -> indices having nums[idx] divisible by prime
	divisible := make(map[int][]int)

	// Build mapping using unique prime factors
	for i, v := range nums {
		x := v
		used := map[int]bool{}

		for x > 1 {
			p := spf[x]
			if !used[p] {
				divisible[p] = append(divisible[p], i)
				used[p] = true
			}
			for x%p == 0 {
				x /= p
			}
		}
	}

	visited := make([]bool, n)
	visited[0] = true

	// To avoid processing same prime teleport multiple times
	usedPrime := make(map[int]bool)

	queue := []int{0}
	steps := 0

	for len(queue) > 0 {
		size := len(queue)

		for s := 0; s < size; s++ {
			i := queue[0]
			queue = queue[1:]

			if i == n-1 {
				return steps
			}

			// Adjacent moves
			if i-1 >= 0 && !visited[i-1] {
				visited[i-1] = true
				queue = append(queue, i-1)
			}

			if i+1 < n && !visited[i+1] {
				visited[i+1] = true
				queue = append(queue, i+1)
			}

			// Prime teleportation
			val := nums[i]

			// teleport allowed only if nums[i] itself is prime
			if val >= 2 && spf[val] == val && !usedPrime[val] {
				usedPrime[val] = true

				for _, nxt := range divisible[val] {
					if !visited[nxt] {
						visited[nxt] = true
						queue = append(queue, nxt)
					}
				}
			}
		}

		steps++
	}

	return -1
}

func main() {
	fmt.Println(minJumps([]int{1, 2, 4, 6}))    // 2
	fmt.Println(minJumps([]int{2, 3, 4, 7, 9})) // 2
	fmt.Println(minJumps([]int{4, 6, 5, 8}))    // 3
}
