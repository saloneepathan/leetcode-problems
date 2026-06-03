package main

import "fmt"

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	const INF = int(1 << 60)

	minLandFinish := INF
	for i := range landStartTime {
		finish := landStartTime[i] + landDuration[i]
		if finish < minLandFinish {
			minLandFinish = finish
		}
	}

	minWaterFinish := INF
	for i := range waterStartTime {
		finish := waterStartTime[i] + waterDuration[i]
		if finish < minWaterFinish {
			minWaterFinish = finish
		}
	}

	ans := INF

	// Land -> Water
	for i := range waterStartTime {
		finish := max(minLandFinish, waterStartTime[i]) + waterDuration[i]
		if finish < ans {
			ans = finish
		}
	}

	// Water -> Land
	for i := range landStartTime {
		finish := max(minWaterFinish, landStartTime[i]) + landDuration[i]
		if finish < ans {
			ans = finish
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	landStartTime := []int{2, 8}
	landDuration := []int{4, 1}
	waterStartTime := []int{6}
	waterDuration := []int{3}

	fmt.Println(earliestFinishTime(
		landStartTime,
		landDuration,
		waterStartTime,
		waterDuration,
	)) // 9

	landStartTime = []int{5}
	landDuration = []int{3}
	waterStartTime = []int{1}
	waterDuration = []int{10}

	fmt.Println(earliestFinishTime(
		landStartTime,
		landDuration,
		waterStartTime,
		waterDuration,
	)) // 14
}
