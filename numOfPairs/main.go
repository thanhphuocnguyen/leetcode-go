package main

import (
	"fmt"
	"strings"
)

func numOfPairs(nums []string, target string) int {
	mp := make(map[string]int)

	ans := 0
	for _, num := range nums {
		if strings.HasPrefix(target, num) {
			// get suffix
			suffix := target[len(num):]
			if _, ok := mp[suffix]; ok {
				ans += mp[suffix]
			}
		}
		if strings.HasSuffix(target, num) {
			// get prefix
			prefix := target[:len(target)-len(num)]
			if _, ok := mp[prefix]; ok {
				ans += mp[prefix]
			}
		}
		mp[num]++
	}
	return ans
}

func main() {
	fmt.Println(numOfPairs([]string{"123", "4", "12", "34"}, "1234"))
	fmt.Println(numOfPairs([]string{"74", "1", "67", "1", "74"}, "174"))
	fmt.Println(numOfPairs([]string{"1", "1", "1"}, "11"))
	fmt.Println(numOfPairs([]string{"777", "7", "77", "77"}, "7777"))
	fmt.Println(numOfPairs([]string{"1", "111"}, "11"))
}
