package main

import (
	"fmt"
)

type Pair struct {
	A, B, C int // line form Ax + By + C = 0 (normalized)
}

type Slope struct {
	dx, dy int // reduced slope dy/dx
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcd3(a, b, c int) int {
	return gcd(gcd(a, b), c)
}

func countTrapezoids(points [][]int) int {
	n := len(points)

	// slope → map intercept → count of segments
	slopeToIntercept := make(map[Slope]map[Pair]int)

	// midpoint → map slope → count of segments
	midToSlope := make(map[[2]int]map[Slope]int)

	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dx := x2 - x1
			dy := y2 - y1

			// slope dy/dx reduced
			if dx == 0 {
				dy = 1 // vertical slope = (1,0) special handling
			} else if dy == 0 {
				dx = 1 // horizontal slope = (1,0)
			} else {
				g := gcd(dx, dy)
				dx /= g
				dy /= g
				// normalize dx > 0
				if dx < 0 {
					dx = -dx
					dy = -dy
				}
			}
			s := Slope{dx, dy}

			// Intercept (use normalized line equation)
			A := y2 - y1
			B := x1 - x2
			C := x2*y1 - x1*y2

			g := gcd3(A, B, C)
			if g != 0 {
				A /= g
				B /= g
				C /= g
			}

			// normalize sign
			if A < 0 || (A == 0 && B < 0) {
				A = -A
				B = -B
				C = -C
			}

			line := Pair{A, B, C}

			if slopeToIntercept[s] == nil {
				slopeToIntercept[s] = make(map[Pair]int)
			}
			slopeToIntercept[s][line]++

			// midpoint reduced (exact integers)
			mx := x1 + x2
			my := y1 + y2
			mid := [2]int{mx, my}

			if midToSlope[mid] == nil {
				midToSlope[mid] = make(map[Slope]int)
			}
			midToSlope[mid][s]++
		}
	}

	ans := 0

	// Count parallel but non-collinear segment pairs
	for _, interceptMap := range slopeToIntercept {
		total := 0
		for _, cnt := range interceptMap {
			ans += total * cnt
			total += cnt
		}
	}

	// Subtract parallelograms (each counted twice)
	for _, slopeMap := range midToSlope {
		total := 0
		for _, cnt := range slopeMap {
			ans -= total * cnt
			total += cnt
		}
	}

	return ans
}

func main() {
	fmt.Println(countTrapezoids([][]int{
		{82, 7}, {82, -9}, {82, -52}, {82, 78},
	})) // expected 0

	fmt.Println(countTrapezoids([][]int{
		{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3},
	})) // expected 2

	fmt.Println(countTrapezoids([][]int{
		{0, 0}, {1, 0}, {0, 1}, {2, 1},
	})) // expected 1
}
