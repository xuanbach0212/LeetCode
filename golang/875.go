//go:build ignore

package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for _, p := range piles {
		if p > r {
			r = p
		}
	}

	res := r

	for l <= r {
		k := l + (r-l)/2

		totalTime := 0
		for _, p := range piles {
			totalTime += int(math.Ceil(float64(p) / float64(k)))
		}

		if totalTime <= h {
			res = k
			r = k - 1
		} else {
			l = k + 1
		}
	}

	return res
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}
