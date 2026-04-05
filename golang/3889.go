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
	freqc := make(map[byte]int, 36)

	for i := 0; i < len(s); i++ {
		c := s[i]
		freqc[c]++
	}

	for c := range freqc {
		if c >= 'a' && c <= 'z' {
			sum += abs(freqc[c] - freqc['z'-(c-'a')])
			delete(freqc, 'z'-(c-'a'))
		} else {
			sum += abs(freqc[c] - freqc['9'-(c-'0')])
			delete(freqc, '9'-(c-'0'))
		}
		delete(freqc, c)
	}

	return sum
}

func main() {
	fmt.Println(mirrorFrequency("ab1z9"))
}
