package main

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {

	result := []string{}

	// Loop through all possible hours and minutes
	for hour := 0; hour < 12; hour++ {

		for minute := 0; minute < 60; minute++ {

			// Count total LEDs ON
			if bits.OnesCount(uint(hour))+bits.OnesCount(uint(minute)) == turnedOn {

				// Format time properly
				time := fmt.Sprintf("%d:%02d", hour, minute)

				result = append(result, time)
			}
		}
	}

	return result
}

func main() {

	fmt.Println("turnedOn = 1")
	fmt.Println(readBinaryWatch(1))

	fmt.Println("\nturnedOn = 9")
	fmt.Println(readBinaryWatch(9))

	fmt.Println("\nturnedOn = 2")
	fmt.Println(readBinaryWatch(2))

}
