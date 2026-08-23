package main

import "fmt"

func sumGame(num string) bool {
	n := len(num)

	get := func(s string) (int, int) {
		sum, q := 0, 0

		for _, ch := range s {
			if ch == '?' {
				q++
			} else {
				sum += int(ch - '0')
			}
		}

		return sum, q
	}

	n0, q0 := get(num[:n/2])
	n1, q1 := get(num[n/2:])

	return ((q0+q1)%2 == 1) || (n0-n1 != (q1-q0)*9/2)
}

func main() {
	tests := []struct {
		num  string
		want bool
	}{
		{"5023", false},
		{"25??", true},
		{"?3295???", false},
	}

	for _, test := range tests {
		got := sumGame(test.num)
		fmt.Printf("num = %q, Alice wins = %v, expected = %v\n",
			test.num, got, test.want)
	}
}