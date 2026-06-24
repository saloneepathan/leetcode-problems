package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int64 = 1_000_000_007

func zigZagArrays(n int, l int, r int) int {
	m := r - l + 1
	size := 2 * m

	// State layout:
	// [0 ... m-1]       => last move was UP, ending at value index i
	// [m ... 2m-1]      => last move was DOWN, ending at value index i
	//
	// Value index i represents actual value (l + i).

	// For arrays of length 2:
	// up[i]   = number of pairs (a, b) where a < b and b = i => i choices
	// down[i] = number of pairs (a, b) where a > b and b = i => m-1-i choices
	initial := make([]int64, size)

	for i := 0; i < m; i++ {
		initial[i] = int64(i)           // UP ending at i
		initial[m+i] = int64(m - 1 - i) // DOWN ending at i
	}

	// Transition matrix:
	//
	// Previous UP ending at prev:
	// next must be smaller -> becomes DOWN ending at next
	//
	// Previous DOWN ending at prev:
	// next must be larger -> becomes UP ending at next
	transition := make([][]int64, size)
	for i := 0; i < size; i++ {
		transition[i] = make([]int64, size)
	}

	for prev := 0; prev < m; prev++ {
		// UP(prev) -> DOWN(next), next < prev
		for next := 0; next < prev; next++ {
			transition[m+next][prev] = 1
		}

		// DOWN(prev) -> UP(next), next > prev
		for next := prev + 1; next < m; next++ {
			transition[next][m+prev] = 1
		}
	}

	// Initial state represents arrays of length 2.
	// To reach length n, apply transition n-2 times.
	poweredTransition := matrixPower(transition, int64(n-2))
	finalState := multiplyMatrixVector(poweredTransition, initial)

	var answer int64
	for _, ways := range finalState {
		answer = (answer + ways) % MOD
	}

	return int(answer)
}

func matrixPower(base [][]int64, exponent int64) [][]int64 {
	size := len(base)

	result := make([][]int64, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int64, size)
		result[i][i] = 1
	}

	for exponent > 0 {
		if exponent&1 == 1 {
			result = multiplyMatrices(result, base)
		}

		base = multiplyMatrices(base, base)
		exponent >>= 1
	}

	return result
}

func multiplyMatrices(a, b [][]int64) [][]int64 {
	size := len(a)

	result := make([][]int64, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int64, size)
	}

	for i := 0; i < size; i++ {
		for k := 0; k < size; k++ {
			if a[i][k] == 0 {
				continue
			}

			for j := 0; j < size; j++ {
				if b[k][j] == 0 {
					continue
				}

				result[i][j] = (result[i][j] + a[i][k]*b[k][j]) % MOD
			}
		}
	}

	return result
}

func multiplyMatrixVector(matrix [][]int64, vector []int64) []int64 {
	size := len(matrix)
	result := make([]int64, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[i][j] == 0 || vector[j] == 0 {
				continue
			}

			result[i] = (result[i] + matrix[i][j]*vector[j]) % MOD
		}
	}

	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l, r int
	fmt.Fscan(in, &n, &l, &r)

	fmt.Println(zigZagArrays(n, l, r))
}
