package main

import "fmt"

type MyQueue struct {
	tmpSt []int
	st    []int
}

func Constructor() MyQueue {
	return MyQueue{tmpSt: []int{}, st: []int{}}
}

func (this *MyQueue) Push(x int) {
	for len(this.st) > 0 {
		this.tmpSt = append(this.tmpSt, this.st[len(this.st)-1])
		this.st = this.st[:len(this.st)-1]
	}
	this.tmpSt = append(this.tmpSt, x)
	for len(this.tmpSt) > 0 {
		this.st = append(this.st, this.tmpSt[len(this.tmpSt)-1])
		this.tmpSt = this.tmpSt[:len(this.tmpSt)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.st) > 0 {
		temp := this.st[len(this.st)-1]
		this.st = this.st[:len(this.st)-1]
		return temp
	}
	return -1
}

func (this *MyQueue) Peek() int {
	if len(this.st) > 0 {
		return this.st[len(this.st)-1]
	}
	return -1
}

func (this *MyQueue) Empty() bool {
	return len(this.st) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	mq := Constructor()
	mq.Push(1)
	mq.Push(2)
	fmt.Println(mq.Peek())
	fmt.Println(mq.Pop())
	fmt.Println(mq.Empty())
}
