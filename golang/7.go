//go:build ignore

package main

import "math"

// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
//
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
//
//
// Example 1:
//
// Input: x = 123
// Output: 321
// Example 2:
//
// Input: x = -123
// Output: -321
// Example 3:
//
// Input: x = 120
// Output: 21
//
//
// Constraints:
//
// -231 <= x <= 231 - 1

func reverse(x int) int {
	res := 0
	remain := x
	for remain != 0 {
		digit := remain % 10
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}
		res = res * 10
		res = res + digit
		remain /= 10
	}
	return res
}

func main() {
}
