package main

const MOD = 1_000_000_007

func countTrapezoids(points [][]int) int {
	// Group x-values by y-value
	yGroups := make(map[int][]int)
	for _, p := range points {
		x, y := p[0], p[1]
		yGroups[y] = append(yGroups[y], x)
	}

	// For each y, compute number of horizontal pairs
	segCounts := []int64{}
	for _, xs := range yGroups {
		k := int64(len(xs))
		if k >= 2 {
			seg := (k * (k - 1) / 2) % MOD
			segCounts = append(segCounts, seg)
		}
	}

	if len(segCounts) < 2 {
		return 0
	}

	// Sum of all segCounts
	total := int64(0)
	for _, v := range segCounts {
		total = (total + v) % MOD
	}

	// Compute sum of seg[i] * (total - seg[i])
	ans := int64(0)
	for _, v := range segCounts {
		ans = (ans + v*(total-v)%MOD) % MOD
	}

	// Each pair counted twice, divide by 2
	ans = ans * ((MOD + 1) / 2) % MOD

	return int(ans)
}
