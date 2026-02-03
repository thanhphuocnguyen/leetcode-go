package main

import "fmt"

func maxIceCream(costs []int, coins int) int {
	ans := 0
	sorted := countingSort(costs)
	for _, c := range sorted {
		if coins < c {
			break
		}
		ans++
		coins -= c
	}
	return ans
}

func countingSort(arr []int) []int {
	// mini := arr[0]
	maxi := arr[0]
	for _, num := range arr {
		maxi = max(maxi, num)
	}
	freq := make([]int, maxi+1)
	for _, num := range arr {
		freq[num]++
	}
	ans := make([]int, len(arr))
	i := 0
	for num := range freq {
		f := freq[num]
		for f > 0 {
			ans[i] = num
			i++
			f--
		}
	}
	return ans
}

func main() {
	fmt.Println(maxIceCream([]int{1, 3, 2, 4, 1}, 7))
	fmt.Println(maxIceCream([]int{1, 6, 3, 1, 2, 5}, 20))
}
