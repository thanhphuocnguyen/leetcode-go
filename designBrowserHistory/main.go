package main

import "fmt"

type BrowserHistory struct {
	currPage string
	backward []string
	forward  []string
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		backward: []string{},
		currPage: homepage,
		forward:  []string{},
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.backward = append(this.backward, this.currPage)
	this.forward = this.forward[:0]
	this.currPage = url
}

func (this *BrowserHistory) Back(steps int) string {
	ans := this.currPage
	n := len(this.backward)
	for i := 0; i < min(steps, n); i++ {
		this.forward = append(this.forward, this.currPage)
		this.currPage = this.backward[len(this.backward)-1]
		ans = this.currPage
		this.backward = this.backward[:len(this.backward)-1]
	}
	return ans
}

func (this *BrowserHistory) Forward(steps int) string {
	ans := this.currPage
	n := len(this.forward)
	for i := 0; i < min(steps, n); i++ {
		this.backward = append(this.backward, this.currPage)
		this.currPage = this.forward[len(this.forward)-1]
		ans = this.currPage
		this.forward = this.forward[:len(this.forward)-1]
	}
	return ans
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */

func main() {
	/*
		obj := Constructor("leetcode.com")
		obj.Visit("google.com")
		obj.Visit("facebook.com")
		obj.Visit("youtube.com")
		fmt.Println(obj.Back(1))
		fmt.Println(obj.Back(1))
		fmt.Println(obj.Forward(1))
		obj.Visit("linkedin.com")
		fmt.Println(obj.Forward(2))
		fmt.Println(obj.Back(2))
		fmt.Println(obj.Back(7))
	*/
	obj := Constructor("zav.com")
	obj.Visit("kni.com")
	fmt.Println(obj.Back(7))
	fmt.Println(obj.Back(7))
	fmt.Println(obj.Forward(5))
	fmt.Println(obj.Forward(1))
	obj.Visit("pwrrbnw.com")
	obj.Visit("mosohif.com")
	fmt.Println(obj.Back(9))
}
