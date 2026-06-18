package main

func minMoves(nums []int, limit int) int {
	diff := make(map[int]int)
	for i := 0; i < len(nums)/2; i++ {
		diff[nums[i]+nums[len(nums)-1-i]]++
	}

	res := len(nums)
	for k, v := range diff {
		if v+min(k-1, limit-k) < res {
			res = v + min(k-1, limit-k)
		}
	}

	return res
}

func main() {

}
