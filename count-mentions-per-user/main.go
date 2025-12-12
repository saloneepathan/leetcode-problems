package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	mentions := make([]int, numberOfUsers)

	// offlineUntil[i] > currentTime  => user i is offline
	// offlineUntil[i] = 0 means online
	offlineUntil := make([]int, numberOfUsers)

	// Sort events by timestamp
	idx := make([]int, len(events))
	for i := range idx {
		idx[i] = i
	}

	sort.Slice(idx, func(a, b int) bool {
		ta, _ := strconv.Atoi(events[idx[a]][1])
		tb, _ := strconv.Atoi(events[idx[b]][1])
		if ta != tb {
			return ta < tb
		}
		// Tie-breaker: OFFLINE before MESSAGE in same timestamp grouping
		return events[idx[a]][0] < events[idx[b]][0]
	})

	i := 0
	for i < len(idx) {

		// Current timestamp
		tStr := events[idx[i]][1]
		t, _ := strconv.Atoi(tStr)

		// Find all events with same timestamp
		j := i + 1
		for j < len(idx) {
			tj, _ := strconv.Atoi(events[idx[j]][1])
			if tj != t {
				break
			}
			j++
		}

		// Before handling any events at timestamp t,
		// bring online users whose offline period expired.
		for u := 0; u < numberOfUsers; u++ {
			if offlineUntil[u] <= t {
				offlineUntil[u] = 0
			}
		}

		// 1. Apply all OFFLINE events at timestamp t
		for k := i; k < j; k++ {
			ev := events[idx[k]]
			if ev[0] == "OFFLINE" {
				id, _ := strconv.Atoi(ev[2])
				offlineUntil[id] = t + 60
			}
		}

		// 2. Apply all MESSAGE events at timestamp t
		for k := i; k < j; k++ {
			ev := events[idx[k]]
			if ev[0] != "MESSAGE" {
				continue
			}

			tokens := strings.Split(ev[2], " ")
			for _, tok := range tokens {
				if tok == "ALL" {
					// Mention all users (offline or online)
					for u := 0; u < numberOfUsers; u++ {
						mentions[u]++
					}
				} else if tok == "HERE" {
					// Mention only online users
					for u := 0; u < numberOfUsers; u++ {
						if offlineUntil[u] == 0 {
							mentions[u]++
						}
					}
				} else if strings.HasPrefix(tok, "id") {
					// Mention specific user
					idStr := tok[2:]
					id, _ := strconv.Atoi(idStr)
					mentions[id]++
				}
			}
		}

		i = j
	}

	return mentions
}

func main() {
	numberOfUsers := 3
	events := [][]string{
		{"MESSAGE", "2", "HERE"},
		{"OFFLINE", "2", "1"},
		{"OFFLINE", "1", "0"},
		{"MESSAGE", "61", "HERE"},
	}

	fmt.Println(countMentions(numberOfUsers, events)) // Expected: [1,0,2]
}
