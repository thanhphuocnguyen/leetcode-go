package main

import "fmt"

type MyLinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyLinkedList {
	fmt.Print("null, ")
	return MyLinkedList{
		Len: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Len {
		fmt.Printf("-1, ")
		return -1
	}

	temp := this.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	fmt.Print(temp.Val, ", ")
	return temp.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newHead := &Node{val, nil}
	if this.Len == 0 {
		this.Head = newHead
		this.Tail = this.Head
	} else {
		newHead.Next = this.Head
		this.Head = newHead
	}
	this.Len++
	fmt.Print("null, ")
}

func (this *MyLinkedList) AddAtTail(val int) {
	newTail := &Node{val, nil}
	if this.Len == 0 {
		this.Head = newTail
		this.Tail = this.Head
	} else {
		this.Tail.Next = newTail
		this.Tail = newTail
	}
	this.Len++
	fmt.Print("null, ")
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Len {
		return
	} else if index == this.Len {
		this.AddAtTail(val)
	} else if index == 0 {
		this.AddAtHead(val)
	} else {
		tmp := this.Head
		newNode := &Node{val, nil}
		for i := 0; i < index-1; i++ {
			tmp = tmp.Next
		}
		newNode.Next = tmp.Next
		tmp.Next = newNode
		this.Len++
	}
	fmt.Print("null, ")
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Len || this.Len == 0 {
		return
	}
	if index == 0 {
		this.Head = this.Head.Next
	} else {
		tmp := this.Head
		for i := 0; i < index-1; i++ {
			tmp = tmp.Next
		}

		tmp.Next = tmp.Next.Next
		if index == this.Len-1 {
			this.Tail = tmp
		}
	}
	this.Len--
	fmt.Print("null, ")
}

func main() {
	// Your code here!
	obj := Constructor()
	obj.AddAtHead(0)
	obj.AddAtIndex(1, 4)
	obj.AddAtTail(8)
	obj.AddAtHead(5)
	obj.AddAtIndex(4, 3)
	obj.AddAtTail(0)
	obj.AddAtTail(5)
	obj.AddAtIndex(6, 3)
	obj.DeleteAtIndex(7)
	obj.DeleteAtIndex(5)
	obj.AddAtTail(4)
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

// Test case
/*
["MyLinkedList","addAtHead","addAtIndex","addAtTail","addAtHead","addAtIndex","addAtTail","addAtTail","addAtIndex","deleteAtIndex","deleteAtIndex","addAtTail"]

[[],[0],[1,4],[8],[5],[4,3],[0],[5],[6,3],[7],[5],[4]]
*/
