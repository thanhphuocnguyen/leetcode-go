package main

func primeSubOperation(nums []int) bool {
	maxNum := -1
	for _, num := range nums {
		maxNum = max(num, maxNum)
	}
	sieve := make([]bool, maxNum+1)
	for i := 0; i <= maxNum; i++ {
		sieve[i] = true
	}
	sieve[1] = false
	for p := 2; p*p <= maxNum; p++ {
		if sieve[p] {
			for i := p * p; i <= maxNum; i += p {
				sieve[i] = false
			}
		}
	}

	curVal := 1
	i := 0
	for i < len(nums) {
		diff := nums[i] - curVal
		if diff < 0 {
			return false
		}

		if sieve[diff] || diff == 0 {
			i++
		}
		curVal++
	}
	return true
}

func main() {
	nums := []int{998, 2}
	primeSubOperation(nums)
}
