//go:build ignore

package main

import "fmt"

func xorAfterQueries(nums []int, queries [][]int) int {
	MOD := 1000000007

	for _, q := range queries {
		l, r, k, v := q[0], q[1], q[2], q[3]
		for i := l; i <= r; i += k {
			nums[i] = int((int64(nums[i]) * int64(v)) % int64(MOD))
		}
	}

	res := 0
	for _, x := range nums {
		res ^= x
	}

	return res
}

func main() {
	fmt.Println(xorAfterQueries([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, [][]int{}))
}
