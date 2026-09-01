package main

import "fmt"

func minMoves(classroom []string, energy int) int {
	m, n := len(classroom), len(classroom[0])

	// Assign an ID to every litter cell.
	litterID := make([][]int, m)
	for r := 0; r < m; r++ {
		litterID[r] = make([]int, n)
		for c := 0; c < n; c++ {
			litterID[r][c] = -1
		}
	}

	startR, startC := 0, 0
	litterCount := 0

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			switch classroom[r][c] {
			case 'S':
				startR, startC = r, c
			case 'L':
				litterID[r][c] = litterCount
				litterCount++
			}
		}
	}

	// No litter to collect.
	if litterCount == 0 {
		return 0
	}

	allMask := (1 << litterCount) - 1

	type State struct {
		r, c   int
		mask   int
		energy int
		dist   int
	}

	/*
		visited[r][c][mask][energy]

		Whether we've already reached this exact state.
	*/
	visited := make([][][][]bool, m)

	for r := 0; r < m; r++ {
		visited[r] = make([][][]bool, n)

		for c := 0; c < n; c++ {
			visited[r][c] = make([][]bool, 1<<litterCount)

			for mask := 0; mask <= allMask; mask++ {
				visited[r][c][mask] = make([]bool, energy+1)
			}
		}
	}

	// BFS queue.
	queue := make([]State, 0)

	start := State{
		r:      startR,
		c:      startC,
		mask:   0,
		energy: energy,
		dist:   0,
	}

	queue = append(queue, start)
	visited[startR][startC][0][energy] = true

	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}

	head := 0

	for head < len(queue) {
		cur := queue[head]
		head++

		for d := 0; d < 4; d++ {
			nr := cur.r + dr[d]
			nc := cur.c + dc[d]

			// Outside the classroom.
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			// Cannot walk through obstacles.
			if classroom[nr][nc] == 'X' {
				continue
			}

			// Moving costs one energy.
			if cur.energy == 0 {
				continue
			}

			nextEnergy := cur.energy - 1
			nextMask := cur.mask

			// Collect litter.
			if id := litterID[nr][nc]; id != -1 {
				nextMask |= 1 << id
			}

			// Reset energy after entering an R cell.
			if classroom[nr][nc] == 'R' {
				nextEnergy = energy
			}

			nextDist := cur.dist + 1

			// All litter collected.
			if nextMask == allMask {
				return nextDist
			}

			// Avoid revisiting the same state.
			if !visited[nr][nc][nextMask][nextEnergy] {
				visited[nr][nc][nextMask][nextEnergy] = true

				queue = append(queue, State{
					r:      nr,
					c:      nc,
					mask:   nextMask,
					energy: nextEnergy,
					dist:   nextDist,
				})
			}
		}
	}

	return -1
}

func main() {
	// Example 1
	classroom1 := []string{
		"S.",
		"XL",
	}

	fmt.Println(minMoves(classroom1, 2))
	// Output: 2

	// Example 2
	classroom2 := []string{
		"LS",
		"RL",
	}

	fmt.Println(minMoves(classroom2, 4))
	// Output: 3

	// Example 3
	classroom3 := []string{
		"L.S",
		"RXL",
	}

	fmt.Println(minMoves(classroom3, 3))
	// Output: -1
}