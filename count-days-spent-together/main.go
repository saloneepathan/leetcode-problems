package main

import "fmt"

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Convert MM-DD into the day number of the year.
	toDay := func(date string) int {
		month := int(date[0]-'0')*10 + int(date[1]-'0')
		day := int(date[3]-'0')*10 + int(date[4]-'0')

		for i := 0; i < month-1; i++ {
			day += daysInMonth[i]
		}
		return day
	}

	startAlice := toDay(arriveAlice)
	endAlice := toDay(leaveAlice)
	startBob := toDay(arriveBob)
	endBob := toDay(leaveBob)

	// Find the overlapping interval.
	start := max(startAlice, startBob)
	end := min(endAlice, endBob)

	if start > end {
		return 0
	}

	return end - start + 1
}

func main() {
	fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19")) // 3
	fmt.Println(countDaysTogether("10-01", "10-31", "11-01", "12-31")) // 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
