//go:build ignore

package main

func guessNumber(n int) int {
	l, r := 0, n-1

	for l <= r {
		mid := l + (r-l)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == -1 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func guess(num int) int {
	return 0
}

func main() {
}
