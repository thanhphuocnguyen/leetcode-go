package main

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	monoStack := make([][2]int, 0)
	result := make([]int, len(queries))
	newQueries := make([][][2]int, len(heights))
	for i := range heights {
		result[i] = -1
		newQueries = append(newQueries, [][2]int{})
	}

	for i, q := range queries {
		a, b := q[0], q[1]
	}
}
func binarySearch(arr [][]int, target int) int {
	left, right := 0, len(arr)-1
	ans := -1
	for left <= right {
		mid := (right - left) / 2
		if arr[mid][0] > target {
			ans = max(ans, mid)
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
