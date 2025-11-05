package main

import "fmt"

func replaceNonCoprimes(nums []int) []int {
	ans := make([]int, 0)
	for i := range len(nums) {
		num := nums[i]
		for len(ans) > 0 {
			top := ans[len(ans)-1]
			d := gcd(num, top)
			if d == 1 {
				break
			}
			ans = ans[:len(ans)-1]
			num = lcm(num, top, d)
		}
		ans = append(ans, num)
	}
	return ans
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func lcm(a, b, gcd int) int {
	return (a / gcd) * b
}

func main() {
	fmt.Println(replaceNonCoprimes([]int{6, 4, 3, 2, 7, 6, 2}))
}
