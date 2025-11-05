package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}

func largestNumber(nums []int) string {
	sortedNums := mergeSort(nums, 0, len(nums)-1)
	var buffer bytes.Buffer
	for _, num := range sortedNums {
		buffer.WriteString(strconv.Itoa(num))
	}
	ans := buffer.String()
	if ans[0] == '0' {
		return "0"
	}
	return ans
}

func mergeSort(nums []int, left int, right int) []int {
	if left >= right {
		return []int{nums[left]}
	}
	var mid int = left + (right-left)/2
	n1 := mergeSort(nums, left, mid)
	n2 := mergeSort(nums, mid+1, right)
	return merge(n1, n2)
}

func merge(n1 []int, n2 []int) []int {
	i, j := 0, 0
	merged := []int{}
	for i < len(n1) && j < len(n2) {
		if compare(n1[i], n2[j]) {
			merged = append(merged, n1[i])
			i++
		} else {
			merged = append(merged, n2[j])
			j++
		}
	}

	for i < len(n1) {
		merged = append(merged, n1[i])
		i++
	}

	for j < len(n2) {
		merged = append(merged, n2[j])
		j++
	}

	return merged
}

func compare(firstNum int, secondNum int) bool {
	a, b := strconv.Itoa(firstNum), strconv.Itoa(secondNum)
	return strings.Compare(a+b, b+a) > 0
}
