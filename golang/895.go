//go:build ignore

package main

type FreqStack struct {
	count    map[int]int
	maxCount int
	group    map[int][]int
}

func Constructor() FreqStack {
	return FreqStack{
		count:    map[int]int{},
		maxCount: 0,
		group:    map[int][]int{},
	}
}

func (this *FreqStack) Push(val int) {
	valCnt := 1 + this.count[val]
	this.count[val] = valCnt

	if valCnt > this.maxCount {
		this.maxCount = valCnt
		this.group[valCnt] = []int{}
	}
	this.group[valCnt] = append(this.group[valCnt], val)
}

func (this *FreqStack) Pop() int {
	res := this.group[this.maxCount][len(this.group[this.maxCount])-1]
	this.group[this.maxCount] = this.group[this.maxCount][:len(this.group[this.maxCount])-1]
	this.count[res] -= 1
	if len(this.group[this.maxCount]) == 0 {
		this.maxCount -= 1
	}

	return res
}

func main() {
}
