package main

import "fmt"

func checkStrings(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	even1 := make([]int, 26)
	odd1 := make([]int, 26)

	even2 := make([]int, 26)
	odd2 := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		if i%2 == 0 {
			even1[s1[i]-'a']++
			even2[s2[i]-'a']++
		} else {
			odd1[s1[i]-'a']++
			odd2[s2[i]-'a']++
		}
	}

	for i := 0; i < 26; i++ {
		if even1[i] != even2[i] || odd1[i] != odd2[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkStrings("abcdba", "cabdab")) // true
	fmt.Println(checkStrings("abe", "bea"))       // false
}
