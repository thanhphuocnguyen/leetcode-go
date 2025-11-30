package main

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	num, ans := 0, 0
	for {
		num = (num*10 + 1) % k
		ans++
		if num == 0 {
			return ans
		}
	}
}

func main() {
}
