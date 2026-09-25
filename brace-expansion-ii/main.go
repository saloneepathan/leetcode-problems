package main

import (
	"fmt"
	"sort"
)

func braceExpansionII(expression string) []string {
	idx := 0
	expr := expression

	isLetter := func(c byte) bool {
		return c >= 'a' && c <= 'z'
	}

	// item -> letter | { expr }
	var item func() map[string]bool
	var term func() map[string]bool
	var exprFunc func() map[string]bool

	item = func() map[string]bool {
		ret := make(map[string]bool)

		if expr[idx] == '{' {
			idx++
			ret = exprFunc()
		} else {
			ret[string(expr[idx])] = true
		}

		idx++
		return ret
	}

	// term -> item | item term
	term = func() map[string]bool {
		// Start with an empty string.
		ret := map[string]bool{"": true}

		// Continue while the next character can start an item.
		for idx < len(expr) && (expr[idx] == '{' || isLetter(expr[idx])) {
			sub := item()
			tmp := make(map[string]bool)

			// Cartesian product / concatenation.
			for left := range ret {
				for right := range sub {
					tmp[left+right] = true
				}
			}

			ret = tmp
		}

		return ret
	}

	// expr -> term | term, expr
	exprFunc = func() map[string]bool {
		ret := make(map[string]bool)

		for {
			// Union with the result of term().
			for k := range term() {
				ret[k] = true
			}

			// If there's a comma, parse the next term.
			if idx < len(expr) && expr[idx] == ',' {
				idx++
				continue
			}

			break
		}

		return ret
	}

	retMap := exprFunc()

	result := make([]string, 0, len(retMap))
	for k := range retMap {
		result = append(result, k)
	}

	sort.Strings(result)
	return result
}

func main() {
	tests := []string{
		"{a,b}{c,{d,e}}",
		"{{a,z},a{b,c},{ab,z}}",
		"{a,b,c}",
		"abc",
		"{a,b}c{d,e}f",
	}

	for _, expression := range tests {
		fmt.Printf("Input:  %s\n", expression)
		fmt.Printf("Output: %v\n\n", braceExpansionII(expression))
	}
}
