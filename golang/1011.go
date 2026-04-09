//go:build ignore

package main

import "fmt"

func shipWithinDays(weights []int, days int) int {
	l, r := 0, 0
	for _, w := range weights {
		r += w
		if w > l {
			l = w
		}
	}

	res := r

	for l <= r {
		mid := l + (r-l)/2

		// calculate the days if valid
		canShip := true
		ships, currCap := 1, mid
		for _, w := range weights {
			if currCap-w < 0 {
				ships++
				currCap = mid
				if ships > days {
					canShip = false
					break
				}
			}
			currCap -= w
		}

		if !canShip {
			l = mid + 1
		} else {
			res = min(res, mid)
			r = mid - 1
		}
	}

	return res
}

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
