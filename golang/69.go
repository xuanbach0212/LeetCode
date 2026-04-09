//go:build ignore

package main

func mySqrt(x int) int {
	l, r := 0, x
	res := 0

	for l <= r {
		mid := l + (r-l)/2

		if mid*mid < x {
			l = mid + 1
			res = mid
		} else if mid*mid > x {
			r = mid - 1
		} else {
			return mid
		}
	}

	return res
}

func main() {
}
