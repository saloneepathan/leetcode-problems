package main

import "fmt"

func maxDotProduct(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)

	// Very small number to represent -infinity
	const NEG_INF = -1 << 60

	// dp[i][j] = max dot product using non-empty subsequences
	// from nums1[0..i-1] and nums2[0..j-1]
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = NEG_INF
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			product := nums1[i-1] * nums2[j-1]

			dp[i][j] = max(
				product,              // start new subsequence
				dp[i-1][j-1]+product, // extend subsequence
				dp[i-1][j],           // skip nums1[i-1]
				dp[i][j-1],           // skip nums2[j-1]
			)
		}
	}

	return dp[n][m]
}

func max(vals ...int) int {
	res := vals[0]
	for _, v := range vals[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {
	fmt.Println(maxDotProduct([]int{2, 1, -2, 5}, []int{3, 0, -6})) // 18
	fmt.Println(maxDotProduct([]int{3, -2}, []int{2, -6, 7}))       // 21
	fmt.Println(maxDotProduct([]int{-1, -1}, []int{1, 1}))          // -1
}
