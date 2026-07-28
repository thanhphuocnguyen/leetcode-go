package main

func maximumLength(nums []int) int {
	ans := 0
	oneCnt := 0
	freq := make(map[int]int)
	for _, num := range nums {
		if num != 1 {
			freq[num]++
		} else {
			oneCnt++
		}
	}
	if oneCnt%2 == 0 {
		oneCnt--
	}

	for k := range freq {
		curr := k
		length := 1
		for freq[curr] >= 2 && freq[curr*curr] > 0 {
			length += 2
			curr = curr * curr
		}
		ans = max(ans, length)
	}
	return max(ans, oneCnt)
}
