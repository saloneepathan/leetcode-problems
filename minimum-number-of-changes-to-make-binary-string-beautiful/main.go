package main

import "fmt"

// minChanges counts the pairs (s[i], s[i+1]) at even i whose characters differ.
// Every even-length block can be split into pairs, so a string is beautiful
// exactly when every such pair is equal; each unequal pair costs one change.
func minChanges(s string) int {
	changes := 0
	for i := 0; i+1 < len(s); i += 2 {
		if s[i] != s[i+1] {
			changes++
		}
	}
	return changes
}

func main() {
	tests := []struct {
		s        string
		expected int
	}{
		{"1001", 2},
		{"10", 1},
		{"0000", 0},
		{"01", 1},   // smallest length, mismatched
		{"11", 0},   // smallest length, already fine
		{"0110", 2}, // pairs 01 and 10 both differ
		{"110011", 0},
		{"101010", 3},
	}
	allPass := true
	for _, t := range tests {
		got := minChanges(t.s)
		ok := got == t.expected
		allPass = allPass && ok
		fmt.Printf("input=%q expected=%d actual=%d pass=%v\n", t.s, t.expected, got, ok)
	}
	fmt.Println("ALL PASS:", allPass)
}
