package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 1 {
		return bits[0] == 0
	}
	if bits[len(bits)-1] == 1 {
		return false
	}
	i, j := 0, 1
	for j < len(bits) {
		if bits[i] == 1 {
			i += 2
			j += 2
		} else {
			i++
			j++
		}
	}
	if i >= len(bits) {
		return false
	}
	return bits[i] == 0
}

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))
}
