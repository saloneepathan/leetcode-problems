package main

import "fmt"

func removeInvalidParentheses(s string) []string {
	// Find minimum number of '(' and ')' to remove.
	leftRemove, rightRemove := 0, 0

	for _, ch := range s {
		if ch == '(' {
			leftRemove++
		} else if ch == ')' {
			if leftRemove > 0 {
				leftRemove--
			} else {
				rightRemove++
			}
		}
	}

	result := make(map[string]bool)

	var dfs func(
		index int,
		leftCount int,
		rightCount int,
		leftRemove int,
		rightRemove int,
		path []byte,
	)

	dfs = func(
		index int,
		leftCount int,
		rightCount int,
		leftRemove int,
		rightRemove int,
		path []byte,
	) {
		// Reached the end.
		if index == len(s) {
			if leftRemove == 0 &&
				rightRemove == 0 &&
				leftCount == rightCount {

				result[string(path)] = true
			}
			return
		}

		ch := s[index]

		// Case 1: '('
		if ch == '(' {
			// Option A: Remove this '('.
			if leftRemove > 0 {
				dfs(
					index+1,
					leftCount,
					rightCount,
					leftRemove-1,
					rightRemove,
					path,
				)
			}

			// Option B: Keep this '('.
			path = append(path, '(')

			dfs(
				index+1,
				leftCount+1,
				rightCount,
				leftRemove,
				rightRemove,
				path,
			)

			path = path[:len(path)-1]

			return
		}

		// Case 2: ')'
		if ch == ')' {
			// Option A: Remove this ')'.
			if rightRemove > 0 {
				dfs(
					index+1,
					leftCount,
					rightCount,
					leftRemove,
					rightRemove-1,
					path,
				)
			}

			// Option B: Keep this ')' only if it
			// doesn't create an invalid prefix.
			if rightCount < leftCount {
				path = append(path, ')')

				dfs(
					index+1,
					leftCount,
					rightCount+1,
					leftRemove,
					rightRemove,
					path,
				)

				path = path[:len(path)-1]
			}

			return
		}

		// Case 3: letter
		path = append(path, ch)

		dfs(
			index+1,
			leftCount,
			rightCount,
			leftRemove,
			rightRemove,
			path,
		)

		path = path[:len(path)-1]
	}

	dfs(0, 0, 0, leftRemove, rightRemove, []byte{})

	ans := make([]string, 0, len(result))

	for str := range result {
		ans = append(ans, str)
	}

	return ans
}

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("(a)())()"))
	fmt.Println(removeInvalidParentheses(")("))
}
