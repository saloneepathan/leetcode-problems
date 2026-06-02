package main

import "fmt"

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	ans := int(1e9)

	for i := 0; i < len(landStartTime); i++ {
		landFinish := landStartTime[i] + landDuration[i]

		for j := 0; j < len(waterStartTime); j++ {
			// Land -> Water
			finish1 := max(landFinish, waterStartTime[j]) + waterDuration[j]
			if finish1 < ans {
				ans = finish1
			}

			// Water -> Land
			waterFinish := waterStartTime[j] + waterDuration[j]
			finish2 := max(waterFinish, landStartTime[i]) + landDuration[i]
			if finish2 < ans {
				ans = finish2
			}
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
	landStartTime1 := []int{2, 8}
	landDuration1 := []int{4, 1}
	waterStartTime1 := []int{6}
	waterDuration1 := []int{3}

	fmt.Println(earliestFinishTime(
		landStartTime1,
		landDuration1,
		waterStartTime1,
		waterDuration1,
	)) // 9

	landStartTime2 := []int{5}
	landDuration2 := []int{3}
	waterStartTime2 := []int{1}
	waterDuration2 := []int{10}

	fmt.Println(earliestFinishTime(
		landStartTime2,
		landDuration2,
		waterStartTime2,
		waterDuration2,
	)) // 14
}
