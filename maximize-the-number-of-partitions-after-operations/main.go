package main

import (
	"fmt"
	"math/bits"
)

type Data struct {
	mask     uint32
	segs     uint16
	fullpref bool
	fullsuff bool
}

var DP [10000]Data

func max(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}

func maxPartitionsAfterOperations(s string, k int) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	// Edge cases
	if n == 1 || k == 26 {
		return 1
	}

	// check if k > number of distinct letters of the string
	var uniqueletters uint32
	for i := 0; i < n; i++ {
		uniqueletters |= 1 << (s[i] - 'a')
	}
	if bits.OnesCount32(uniqueletters) < k {
		return 1
	}

	// prefix processing
	for i := range DP {
		DP[i] = Data{}
	}

	var mask uint32
	var segments uint16
	var count int

	for i := 0; i < n-1; i++ {
		c := s[i] - 'a'
		newbit := 1 &^ int(mask>>c)
		count += newbit
		bit := uint32(1 << c)
		mask |= bit
		if count > k {
			segments++
			count = 1
			mask = bit
		}
		DP[i+1] = Data{mask, segments, count == k, false}
	}

	// suffix processing
	mask, segments, count = 0, 0, 0
	for i := n - 1; i > 0; i-- {
		c := s[i] - 'a'
		newbit := 1 &^ int(mask>>c)
		count += newbit
		bit := uint32(1 << c)
		mask |= bit
		if count > k {
			segments++
			count = 1
			mask = bit
		}
		DP[i-1].mask |= mask
		DP[i-1].segs += segments
		DP[i-1].fullsuff = count == k
	}

	// calculate result
	var result uint16
	const LETTER_MASK = (1 << 26) - 1
	for i := 0; i < n; i++ {
		full := DP[i].fullpref && DP[i].fullsuff
		mask, seg := DP[i].mask, DP[i].segs
		if full && mask != LETTER_MASK {
			seg += 2
		} else if bits.OnesCount32(mask) >= k {
			seg += 1
		}
		result = max(result, seg+1)
	}

	return int(result)
}

func main() {
	fmt.Println(maxPartitionsAfterOperations("accca", 2))  // ✅ expected 3
	fmt.Println(maxPartitionsAfterOperations("aabaab", 3)) // ✅ expected 1
	fmt.Println(maxPartitionsAfterOperations("xxyz", 1))   // ✅ expected 4
	fmt.Println(maxPartitionsAfterOperations("aaabc", 1))  // ✅ expected 4
}
