//go:build ignore

package main

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := [][2]int{} // index, height
	res := 0
	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			lastIdx, lastHeight := stack[len(stack)-1][0], stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
			res = max(res, lastHeight*(i-lastIdx))
			start = lastIdx
		}
		stack = append(stack, [2]int{start, h})
	}
	return res
}

func main() {
}
