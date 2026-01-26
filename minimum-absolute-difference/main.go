import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	minDiff := int(1e9)
	n := len(arr)

	// Step 1: Find minimum difference
	for i := 1; i < n; i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	// Step 2: Collect pairs with minDiff
	result := [][]int{}
	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] == minDiff {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}

	return result
}
