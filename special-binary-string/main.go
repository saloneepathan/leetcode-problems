package main

import (
	"fmt"
	"slices"
)

var subtrees [15]uint64
var indices [15]uint8

func makeLargestSpecial(s string) string {

	var subsize, idxsize uint8

	for _, c := range s {

		if c == '1' {

			indices[idxsize] = subsize
			idxsize++

		} else {

			idxsize--
			childidx := indices[idxsize]

			subs := subtrees[childidx:subsize]

			if subsize-childidx > 1 {

				slices.SortFunc(subs, func(a, b uint64) int {

					if b > a {
						return 1
					} else if b < a {
						return -1
					}
					return 0
				})
			}

			mask := uint64(1 << 63)
			size := uint64(1)

			for _, sub := range subs {

				mask |= (sub >> size)
				size += sub & 0x3F
			}

			packed := (mask &^ 0x3F) | (size + 1)

			subtrees[childidx] = packed
			subsize = childidx + 1
		}
	}

	slices.SortFunc(subtrees[:subsize], func(a, b uint64) int {

		if b > a {
			return 1
		} else if b < a {
			return -1
		}
		return 0
	})

	// SAFE decoding
	buffer := make([]byte, len(s))

	b := 0

	for _, sub := range subtrees[:subsize] {

		length := sub & 0x3F

		for i := uint64(0); i < length; i++ {

			buffer[b] = byte((sub>>(63-i))&1) + '0'
			b++
		}
	}

	return string(buffer)
}

func main() {

	input := "11011000"

	result := makeLargestSpecial(input)

	fmt.Println("Input :", input)
	fmt.Println("Output:", result)
}
