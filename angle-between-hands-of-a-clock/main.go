package main

import (
	"fmt"
	"math"
)

func angleClock(hour int, minutes int) float64 {
	// Position of hour hand
	hourAngle := float64(hour%12)*30.0 + float64(minutes)*0.5

	// Position of minute hand
	minuteAngle := float64(minutes) * 6.0

	// Absolute difference
	diff := math.Abs(hourAngle - minuteAngle)

	// Return the smaller angle
	return math.Min(diff, 360.0-diff)
}

func main() {
	fmt.Println(angleClock(12, 30)) // 165
	fmt.Println(angleClock(3, 30))  // 75
	fmt.Println(angleClock(3, 15))  // 7.5
}
