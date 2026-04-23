//go:build ignore

package main

import (
	"fmt"
)

//Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle, and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".
//
//One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.
//
//Implement the MyCircularQueue class:
//
//MyCircularQueue(k) Initializes the object with the size of the queue to be k.
//int Front() Gets the front item from the queue. If the queue is empty, return -1.
//int Rear() Gets the last item from the queue. If the queue is empty, return -1.
//boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
//boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
//boolean isEmpty() Checks whether the circular queue is empty or not.
//boolean isFull() Checks whether the circular queue is full or not.
//You must solve the problem without using the built-in queue data structure in your programming language.
//
//
//
//Example 1:
//
//Input
//["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
//[[3], [1], [2], [3], [4], [], [], [], [4], []]
//Output
//[null, true, true, true, false, 3, true, true, true, 4]
//
//Explanation
//MyCircularQueue myCircularQueue = new MyCircularQueue(3);
//myCircularQueue.enQueue(1); // return True
//myCircularQueue.enQueue(2); // return True
//myCircularQueue.enQueue(3); // return True
//myCircularQueue.enQueue(4); // return False
//myCircularQueue.Rear();     // return 3
//myCircularQueue.isFull();   // return True
//myCircularQueue.deQueue();  // return True
//myCircularQueue.enQueue(4); // return True
//myCircularQueue.Rear();     // return 4
//
//
//Constraints:
//
//1 <= k <= 1000
//0 <= value <= 1000
//At most 3000 calls will be made to enQueue, deQueue, Front, Rear, isEmpty, and isFull.

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}
type MyCircularQueue struct {
	space int
	left  *Node
	right *Node
}

func Constructor(k int) MyCircularQueue {
	left := &Node{Val: 0}
	right := &Node{Val: 0, Prev: left}
	left.Next = right
	return MyCircularQueue{
		space: k,
		left:  left,
		right: right,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	curr := &Node{Val: value, Next: this.right, Prev: this.right.Prev}
	this.right.Prev.Next = curr
	this.right.Prev = curr
	this.space--
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.left.Next = this.left.Next.Next
	this.left.Next.Prev = this.left
	this.space++
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.left.Next.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.right.Prev.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.left.Next == this.right
}

func (this *MyCircularQueue) IsFull() bool {
	return this.space == 0
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

func main() {
	fmt.Println("removeNthFromEnd")
}
