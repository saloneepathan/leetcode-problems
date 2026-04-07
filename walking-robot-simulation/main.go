package main

import (
	"fmt"
)

func robotSim(commands []int, obstacles [][]int) int {

	// directions: North, East, South, West
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	// store obstacles in set
	obstacleSet := make(map[string]bool)

	for _, obs := range obstacles {
		key := fmt.Sprintf("%d,%d", obs[0], obs[1])
		obstacleSet[key] = true
	}

	x, y := 0, 0
	dir := 0
	maxDist := 0

	for _, cmd := range commands {

		if cmd == -1 {
			// turn right
			dir = (dir + 1) % 4

		} else if cmd == -2 {
			// turn left
			dir = (dir + 3) % 4

		} else {

			for step := 0; step < cmd; step++ {

				nx := x + dx[dir]
				ny := y + dy[dir]

				key := fmt.Sprintf("%d,%d", nx, ny)

				if obstacleSet[key] {
					break
				}

				x = nx
				y = ny

				dist := x*x + y*y
				if dist > maxDist {
					maxDist = dist
				}
			}
		}
	}

	return maxDist
}

func main() {

	commands := []int{4, -1, 4, -2, 4}
	obstacles := [][]int{{2, 4}}

	fmt.Println(robotSim(commands, obstacles)) // Output: 65
}
