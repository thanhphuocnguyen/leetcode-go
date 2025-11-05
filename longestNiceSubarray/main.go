package main

func longestNiceSubarray(nums []int) int {
	currentBitSet := nums[0]
	left, right := 0, 0
	ans := 1
	for right < len(nums) {
		if (nums[right] & currentBitSet) == 0 {
			currentBitSet ^= nums[right]
			ans = max(ans, 1+right-left)
		} else {
			left = right
			currentBitSet = nums[left]
		}
		right++
	}
	return ans
}

func main() {
	println((longestNiceSubarray([]int{84139415, 693324769, 614626365, 497710833, 615598711, 264, 65552, 50331652, 1, 1048576, 16384, 544, 270532608, 151813349, 221976871, 678178917, 845710321, 751376227, 331656525, 739558112, 267703680})))
	println((longestNiceSubarray([]int{3, 1, 5, 11, 13})))
}
