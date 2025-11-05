package main

import "fmt"

func totalFruit(fruits []int) int {
	if len(fruits) <= 2 {
		return len(fruits)
	}
	mp := make(map[int]int)
	l, r := 0, 0
	ans := 0
	for r < len(fruits) {
		mp[fruits[r]]++
		if len(mp) > 2 {
			ans = max(ans, r-l)
			// remove til current type equal 0
			for {
				mp[fruits[l]]--
				if mp[fruits[l]] == 0 {
					delete(mp, fruits[l])
					break
				}
				l++
			}
			l++
		}
		r++
	}
	ans = max(ans, r-l)

	return ans
}

func main() {
	fmt.Println(totalFruit([]int{1, 0, 1, 4, 1, 4, 1, 2, 3}))
	fmt.Println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))
	fmt.Println(totalFruit([]int{1, 2, 1}))
}
