package main

import (
	"fmt"
)

type Robot struct {
	w, h int
	x, y int
	dir  int
	per  int
}

var dirs = []string{"East", "North", "West", "South"}

func Constructor(width int, height int) Robot {
	return Robot{
		w:   width,
		h:   height,
		x:   0,
		y:   0,
		dir: 0,
		per: 2*(width+height) - 4,
	}
}

func (this *Robot) Step(num int) {

	num %= this.per

	// Special edge case
	if num == 0 && this.x == 0 && this.y == 0 {
		this.dir = 3 // South
		return
	}

	for num > 0 {

		if this.dir == 0 { // East
			move := min(num, this.w-1-this.x)
			this.x += move
			num -= move
			if move == 0 {
				this.dir = (this.dir + 1) % 4
			}

		} else if this.dir == 1 { // North
			move := min(num, this.h-1-this.y)
			this.y += move
			num -= move
			if move == 0 {
				this.dir = (this.dir + 1) % 4
			}

		} else if this.dir == 2 { // West
			move := min(num, this.x)
			this.x -= move
			num -= move
			if move == 0 {
				this.dir = (this.dir + 1) % 4
			}

		} else { // South
			move := min(num, this.y)
			this.y -= move
			num -= move
			if move == 0 {
				this.dir = (this.dir + 1) % 4
			}
		}
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.x, this.y}
}

func (this *Robot) GetDir() string {
	return dirs[this.dir]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	robot := Constructor(6, 3)

	robot.Step(2)
	robot.Step(2)

	fmt.Println(robot.GetPos()) // [4 0]
	fmt.Println(robot.GetDir()) // East

	robot.Step(2)
	robot.Step(1)
	robot.Step(4)

	fmt.Println(robot.GetPos()) // [1 2]
	fmt.Println(robot.GetDir()) // West
}
