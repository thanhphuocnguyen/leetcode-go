package main

import "fmt"

func minimumAverage(nums []int) float64 {
	quickSort(nums, 0, len(nums)-1)
	ans := 50.0
	for i, j := 0, len(nums)-1; i < len(nums)/2; i, j = i+1, j-1 {
		ans = min(ans, float64(nums[i]+nums[j])/2.0)
	}
	return ans

}

func quickSort(nums []int, left, right int) {
	pivot := nums[(left+right)/2]
	i, j := left, right
	for i <= j {
		for nums[i] < pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	if i < right {
		quickSort(nums, i, right)
	}
	if j > left {
		quickSort(nums, left, j)
	}
}

func main() {
	fmt.Println(minimumAverage([]int{4, 7, 5, 5}))
}
