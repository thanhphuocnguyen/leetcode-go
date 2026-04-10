package main

import "fmt"

func resultArray(nums []int) []int {
	arr1, arr2 := []int{nums[0]}, []int{nums[1]}
	for i := 2; i < len(nums); i++ {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}
	arr1 = append(arr1, arr2...)
	return arr1
}
func main() {
	fmt.Println(resultArray([]int{1, 10, 2, 14}))
	fmt.Println(resultArray([]int{5, 4, 3, 8}))
	fmt.Println(resultArray([]int{2, 1, 3}))
}
