package main

import "fmt"

func maxPartitionsAfterOperations(s string, k int) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if k >= 26 {
		return 1
	}

	// --- Optimized case: k == 1 ---
	if k == 1 {
		// count partitions (number of runs)
		partitions := 1
		for i := 1; i < n; i++ {
			if s[i] != s[i-1] {
				partitions++
			}
		}
		best := partitions

		// try changing one character
		for i := 0; i < n; i++ {
			loss := 0
			gain := 0

			// before change: boundaries possibly lost
			if i > 0 && s[i] != s[i-1] {
				loss++
			}
			if i < n-1 && s[i] != s[i+1] {
				loss++
			}

			// after change: we can pick a char different from both neighbors
			if i > 0 {
				gain++
			}
			if i < n-1 {
				gain++
			}

			newPartitions := partitions - loss + gain
			if newPartitions > best {
				best = newPartitions
			}
		}

		if best > n {
			best = n
		}
		return best
	}

	// --- General case: k > 1 ---
	best := getPartitionsFast(s, k)
	n = len(s)
	for i := 0; i < n; i++ {
		orig := s[i]
		for c := byte('a'); c <= 'z'; c++ {
			if c == orig {
				continue
			}
			modified := make([]byte, n)
			copy(modified, s)
			modified[i] = c
			p := getPartitionsFast(string(modified), k)
			if p > best {
				best = p
			}
		}
	}
	return best
}

// Efficient single-pass partition counter
func getPartitionsFast(s string, k int) int {
	freq := [26]int{}
	distinct := 0
	partitions := 1
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if freq[idx] == 0 {
			if distinct == k {
				partitions++
				for j := range freq {
					freq[j] = 0
				}
				distinct = 0
			}
			distinct++
		}
		freq[idx]++
	}
	return partitions
}

func main() {
	fmt.Println(maxPartitionsAfterOperations("accca", 2))  // 3
	fmt.Println(maxPartitionsAfterOperations("aabaab", 3)) // 1
	fmt.Println(maxPartitionsAfterOperations("xxyz", 1))   // 4
	fmt.Println(maxPartitionsAfterOperations("aaabc", 1))  // ✅ 5
}
