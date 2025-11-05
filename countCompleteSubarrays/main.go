package main

func countCompleteSubarrays(nums []int) int {
	distinctCnt := 0
	numMp := make(map[int]int)
	for _, num := range nums {
		if _, ok := numMp[num]; !ok {
			distinctCnt++
		}
		numMp[num]++
	}

	for k := range numMp {
		numMp[k] = 0
	}
	cnt := 0
	l, r := 0, 0
	ans := 0
	for r < len(nums) {
		if numMp[nums[r]] == 0 {
			cnt++
		}
		numMp[nums[r]]++
		for cnt == distinctCnt {
			ans += len(nums) - r
			numMp[nums[l]]--
			if numMp[nums[l]] == 0 {
				cnt--
			}
			l++
		}

		r++
	}
	return ans
}

func main() {
	println(countCompleteSubarrays([]int{1, 3, 1, 2, 2}))
}
