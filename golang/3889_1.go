//go:build ignore

package main

import (
	"fmt"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func mirrorFrequency(s string) int {
	sum := 0
	freqc := make([]int, 36)

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			freqc[c-'a']++
		} else {
			freqc[c-'0'+26]++
		}
	}

	for c := 0; c < len(freqc); c++ {
		if c < 26 {
			char := c + 'a'
			mirror := 'z' - (char - 'a')
			sum += abs(freqc[char-'a'] - freqc[mirror-'a'])

			freqc[char-'a'] = 0
			freqc[mirror-'a'] = 0
		} else {
			char := c + '0' - 26
			mirror := '9' - (char - '0')
			sum += abs(freqc[char-'0'+26] - freqc[mirror-'0'+26])

			freqc[char-'0'+26] = 0
			freqc[mirror-'0'+26] = 0
		}
	}

	return sum
}

func main() {
	fmt.Println(mirrorFrequency("ab1z9"))
}
