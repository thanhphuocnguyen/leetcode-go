package main

import "fmt"

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		pings: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	lb := lowerBound(this.pings, t-3000)
	return len(this.pings) - lb
}

func lowerBound(arr []int, target int) int {
	lo, hi := 0, len(arr)
	ans := len(arr)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if arr[mid] >= target {
			ans = mid
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))
	fmt.Println(obj.Ping(100))
	fmt.Println(obj.Ping(3001))
	fmt.Println(obj.Ping(3002))
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
