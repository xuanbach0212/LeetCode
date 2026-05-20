//go:build ignore

package main

import "fmt"

func validDigit(n int, x int) bool {
	isValid := false
	for n > 0 {
		if n == x {
			return false
		}
		remain := n % 10
		if remain == x {
			isValid = true
		}
		n = n / 10
	}

	return isValid
}

func main() {
	fmt.Println(validDigit(101, 0))
	fmt.Println(validDigit(232, 2))
	fmt.Println(validDigit(5, 1))
}
