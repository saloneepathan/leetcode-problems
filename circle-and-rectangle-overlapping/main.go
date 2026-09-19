package main

import "fmt"

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	// Find the closest x-coordinate in the rectangle to the circle center.
	closestX := xCenter
	if closestX < x1 {
		closestX = x1
	} else if closestX > x2 {
		closestX = x2
	}

	// Find the closest y-coordinate in the rectangle to the circle center.
	closestY := yCenter
	if closestY < y1 {
		closestY = y1
	} else if closestY > y2 {
		closestY = y2
	}

	// Check if the closest point is inside the circle.
	dx := closestX - xCenter
	dy := closestY - yCenter

	return dx*dx+dy*dy <= radius*radius
}

func main() {
	fmt.Println(checkOverlap(1, 0, 0, 1, -1, 3, 1)) // true
	fmt.Println(checkOverlap(1, 1, 1, 1, -3, 2, -1)) // false
	fmt.Println(checkOverlap(1, 0, 0, -1, 0, 0, 1))  // true
}