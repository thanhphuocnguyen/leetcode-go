package main

import "fmt"

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	// monotonic stack
	mnStack := make([]int, 0, n)
	for i := range n {
		for len(mnStack) > 0 && heights[mnStack[len(mnStack)-1]] >= heights[i] {
			// if peek of stack larger than current keep pop til it is not anymore
			// it means we can extend the width of current height to the left
			mnStack = mnStack[0 : len(mnStack)-1]
		}
		// set left as current peek or -1 as
		if len(mnStack) > 0 {
			left[i] = mnStack[len(mnStack)-1]
		} else {

			left[i] = -1
		}
		mnStack = append(mnStack, i)
	}
	mnStack = mnStack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(mnStack) > 0 && heights[mnStack[len(mnStack)-1]] >= heights[i] {
			mnStack = mnStack[0 : len(mnStack)-1]
		}
		if len(mnStack) > 0 {
			right[i] = mnStack[len(mnStack)-1]
		} else {
			right[i] = n
		}
		mnStack = append(mnStack, i)
	}
	ans := 0
	for i := range n {
		width := right[i] - left[i] - 1
		ans = max(ans, width*heights[i])
	}
	return ans
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 4}))
}
