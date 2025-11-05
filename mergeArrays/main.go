package main

import "fmt"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	ans := make([][]int, 0)
	m, n := len(nums1), len(nums2)
	i, j := 0, 0
	for i < m && j < n {
		id1, id2 := nums1[i][0], nums2[j][0]
		val1, val2 := nums1[i][1], nums2[j][1]
		if id1 == id2 {
			ans = append(ans, []int{id1, val1 + val2})
			i++
			j++
		} else if id1 < id2 {
			ans = append(ans, []int{id1, val1})
			i++
		} else {
			ans = append(ans, []int{id2, val2})
			j++
		}
	}
	for i < m {
		id1, val1 := nums1[i][0], nums1[i][1]
		ans = append(ans, []int{id1, val1})
		i++
	}
	for j < n {
		id2, val2 := nums2[j][0], nums2[j][1]
		ans = append(ans, []int{id2, val2})
		j++
	}
	return ans
}

func main() {
	nums1 := [][]int{{1, 2}, {2, 3}, {4, 5}}
	nums2 := [][]int{{1, 4}, {3, 2}, {4, 1}}
	fmt.Println(mergeArrays(nums1, nums2))
}
