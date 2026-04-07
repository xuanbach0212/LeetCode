//go:build ignore

package main

import "strings"

func decodeString(s string) string {
	countStack := []int{}
	stringStack := []string{}
	currStr := ""
	k := 0

	for _, char := range s {
		if char >= '0' && char <= '9' {
			k = k*10 + int(char-'0')
		} else if char == ']' {
			repeatTimes := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			prevString := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]

			currStr = prevString + strings.Repeat(currStr, repeatTimes)

		} else if char == '[' {
			countStack = append(countStack, k)
			stringStack = append(stringStack, currStr)
			k = 0
			currStr = ""

		} else {
			currStr += string(char)
		}
	}

	return currStr
}

func main() {
}
