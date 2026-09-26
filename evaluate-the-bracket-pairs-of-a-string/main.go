package main

import (
	"fmt"
	"strings"
)

func evaluate(s string, knowledge [][]string) string {
	mp := make(map[string]string, len(knowledge))

	for _, pair := range knowledge {
		mp[pair[0]] = pair[1]
	}

	var res strings.Builder

	for i := 0; i < len(s); {
		if s[i] != '(' {
			res.WriteByte(s[i])
			i++
			continue
		}

		// Find closing bracket.
		j := i + 1
		for s[j] != ')' {
			j++
		}

		key := s[i+1 : j]

		if value, ok := mp[key]; ok {
			res.WriteString(value)
		} else {
			res.WriteByte('?')
		}

		i = j + 1
	}

	return res.String()
}

func main() {
	s := "(name)is(age)yearsold"
	knowledge := [][]string{
		{"name", "bob"},
		{"age", "two"},
	}

	fmt.Println(evaluate(s, knowledge))
}
