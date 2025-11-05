package main

func findLengthOfShortestSubarray(arr []int) int {
	left, right := 0, len(arr)-1
	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}
	ans := right
	for left < right && (left == 0 || arr[left-1] <= arr[left]) {
		for right < len(arr) && arr[left] > arr[right] {
			right++
		}
		ans = min(ans, right-left-1)
		left++
	}
	return ans
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	findLengthOfShortestSubarray(arr)
}
