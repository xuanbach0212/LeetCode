//go:build ignore

package main

import "fmt"

func averageWaitingTime(customers [][]int) float64 {
	waitingTime := 0
	currentTime := customers[0][0]
	for _, cus := range customers {
		currentTime = max(currentTime, cus[0]) + cus[1]
		waitingTime += (currentTime - cus[0])
	}
	return float64(waitingTime) / float64(len(customers))
}

func main() {
	fmt.Println(averageWaitingTime([][]int{{1, 2}, {2, 5}, {4, 3}}))
	fmt.Println(averageWaitingTime([][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}))
}
