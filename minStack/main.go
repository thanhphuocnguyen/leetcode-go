package main

import "fmt"

type MinStack struct {
	st     []int
	minArr []int
}

func Constructor() MinStack {
	return MinStack{
		st:     []int{},
		minArr: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.minArr) == 0 || val <= this.minArr[len(this.minArr)-1] {
		this.minArr = append(this.minArr, val)
	}
	this.st = append(this.st, val)
}

func (this *MinStack) Pop() {
	popped := this.st[len(this.st)-1]
	if popped == this.minArr[len(this.minArr)-1] {
		this.minArr = this.minArr[:len(this.minArr)-1]
	}
	this.st = this.st[:len(this.st)-1]
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	return this.minArr[len(this.minArr)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	st := Constructor()
	st.Push(-2)
	st.Push(0)
	st.Push(-3)
	fmt.Println(st.GetMin())
	st.Pop()
	fmt.Println(st.Top())
	fmt.Println(st.GetMin())
}
