package main

import "fmt"

func pancakeSort(arr []int) []int {
	n := len(arr)
	ans := make([]int, 0, 10)
	end := n - 1
	for i := 0; i < 10; i++ {
		k := findMax(arr, end)
		if k == end {
			continue
		}
		reverse(arr, k)
		reverse(arr, end)
		ans = append(ans, end+1)
		end--
		if end == 0 {
			break
		}
	}

	return ans
}

func findMax(arr []int, end int) int {
	currMax := -1
	idx := -1
	for i := 0; i <= end; i++ {
		if arr[i] > currMax {
			idx = i
			currMax = arr[i]
		}
	}

	return idx
}

func reverse(arr []int, end int) {
	for i := 0; i <= end/2; i++ {
		arr[i], arr[end-i] = arr[end-i], arr[i]
	}
}

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
}
