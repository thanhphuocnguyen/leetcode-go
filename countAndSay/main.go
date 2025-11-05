package main

import (
	"strconv"
)

func countAndSay(n int) string {
	ans := "1"
	for i := 1; i < n; i++ {
		ans = pToStr(count(ans))
	}
	return ans
}

func count(nums string) [][2]int {
	ans := make([][2]int, 0)
	prev := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			cnt++
		} else {
			num := int(prev - '0')
			ans = append(ans, [2]int{cnt, num})
			prev = nums[i]
			cnt = 1
		}
	}
	ans = append(ans, [2]int{cnt, int(nums[len(nums)-1] - '0')})
	return ans
}

func pToStr(pairs [][2]int) string {
	var ans string = ""
	for _, p := range pairs {
		cnt, num := p[0], p[1]
		ans += strconv.Itoa(cnt)
		ans += strconv.Itoa(num)
	}
	return ans
}

func main() {
	println(countAndSay(6))
}
