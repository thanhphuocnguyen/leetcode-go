package main

import "fmt"

func duplicateZeros(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			for j := n - 2; j > i; j-- {
				arr[j+1] = arr[j]
			}
			arr[i+1] = 0
			i++
		}
	}
}
func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
}
