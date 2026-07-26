package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	// Split the version strings into revisions
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	// Get max length of revisions
	n := len(v1)
	if len(v2) > n {
		n = len(v2)
	}

	// Compare revisions
	for i := 0; i < n; i++ {
		var num1, num2 int

		if i < len(v1) {
			num1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			num2, _ = strconv.Atoi(v2[i])
		}

		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}

	// If all revisions are equal
	return 0
}

func main() {
	fmt.Println(compareVersion("1.2", "1.10"))    // -1
	fmt.Println(compareVersion("1.01", "1.001"))  // 0
	fmt.Println(compareVersion("1.0", "1.0.0.0")) // 0
	fmt.Println(compareVersion("2.3.4", "2.3.3")) // 1
}
