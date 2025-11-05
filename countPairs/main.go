package main

func countPairs(nums []int, k int) int {
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}
	return -1
}
func main() {
	countPairs([]int{3, 1, 2, 2, 2, 1, 3}, 2)
}
