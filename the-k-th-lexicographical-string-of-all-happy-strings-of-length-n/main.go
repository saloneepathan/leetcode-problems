package main

import (
	"fmt"
)

func getHappyString(n int, k int) string {
	count := 0
	result := ""

	var dfs func(path []byte)
	dfs = func(path []byte) {
		if len(result) > 0 {
			return
		}

		if len(path) == n {
			count++
			if count == k {
				result = string(path)
			}
			return
		}

		for _, c := range []byte{'a', 'b', 'c'} {
			if len(path) > 0 && path[len(path)-1] == c {
				continue
			}
			dfs(append(path, c))
		}
	}

	dfs([]byte{})
	return result
}

func main() {
	fmt.Println(getHappyString(1, 3)) // c
	fmt.Println(getHappyString(1, 4)) // ""
	fmt.Println(getHappyString(3, 9)) // cab
}
