//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func findGoodIntegers(n int) []int {
	if n < 1729 {
		return []int{}
	}
	res := []int{}
	count := []int{}

	for i := 1; i*i*i < n; i++ {
		i3 := (i * i * i)
		for j := i + 1; ; j++ {
			sum := i3 + (j * j * j)
			if sum > n {
				break
			}
			count = append(count, sum)
		}
	}

	sort.Ints(count)

	for i := 1; i < len(count); i++ {
		if count[i] == count[i-1] {
			if len(res) == 0 || res[len(res)-1] != count[i] {
				res = append(res, count[i])
			}
		}
	}

	return res
}

func main() {
	fmt.Println(findGoodIntegers(4108))
}
