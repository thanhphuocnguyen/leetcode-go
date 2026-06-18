package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	prev := 0
	kth := 0

	for _, num := range arr {
		gap := num - prev
		if gap > 1 {
			for i := 1; i < gap; i++ {
				kth++
				if kth == k {
					return prev + i
				}
			}
		}
		prev = num
	}
	return arr[len(arr)-1] + (k - kth)
}

func main() {
	fmt.Println(findKthPositive([]int{5, 6, 7, 8, 9}, 9))
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
}
