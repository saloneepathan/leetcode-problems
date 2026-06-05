package main

import (
	"fmt"
)

type State struct {
	prev, curr, tight, lead int
	cnt, sum                int64
}

func solve(num int64) int64 {
	// Numbers with fewer than 3 digits have waviness = 0
	if num < 100 {
		return 0
	}

	s := fmt.Sprintf("%d", num)
	n := len(s)

	// prev=10 and curr=10 denote "invalid/not yet started"
	currStates := []State{
		{10, 10, 1, 1, 1, 0},
	}

	for pos := 0; pos < n; pos++ {
		limit := int(s[pos] - '0')

		var cnt [2][2][11][11]int64
		var sum [2][2][11][11]int64

		for _, st := range currStates {
			maxDigit := 9
			if st.tight == 1 {
				maxDigit = limit
			}

			for digit := 0; digit <= maxDigit; digit++ {
				newLead := 0
				if st.lead == 1 && digit == 0 {
					newLead = 1
				}

				newPrev := st.curr
				newCurr := digit

				if newLead == 1 {
					newCurr = 10
				}

				newTight := 0
				if st.tight == 1 && digit == limit {
					newTight = 1
				}

				add := int64(0)

				// Check if st.curr is a peak or valley
				if newLead == 0 && st.prev != 10 && st.curr != 10 {
					if (st.prev < st.curr && st.curr > digit) ||
						(st.prev > st.curr && st.curr < digit) {
						add = st.cnt
					}
				}

				cnt[newTight][newLead][newPrev][newCurr] += st.cnt
				sum[newTight][newLead][newPrev][newCurr] += st.sum + add
			}
		}

		nextStates := make([]State, 0)

		for tight := 0; tight < 2; tight++ {
			for lead := 0; lead < 2; lead++ {
				for prev := 0; prev <= 10; prev++ {
					for curr := 0; curr <= 10; curr++ {
						c := cnt[tight][lead][prev][curr]
						if c == 0 {
							continue
						}

						nextStates = append(nextStates, State{
							prev:  prev,
							curr:  curr,
							tight: tight,
							lead:  lead,
							cnt:   c,
							sum:   sum[tight][lead][prev][curr],
						})
					}
				}
			}
		}

		currStates = nextStates
	}

	var ans int64
	for _, st := range currStates {
		ans += st.sum
	}

	return ans
}

func totalWaviness(num1 int64, num2 int64) int64 {
	return solve(num2) - solve(num1-1)
}

func main() {
	var num1, num2 int64

	num1 = 100
	num2 = 999

	fmt.Println(totalWaviness(num1, num2))
}
