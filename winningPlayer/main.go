package main

import "fmt"

func winningPlayer(x int, y int) string {
	turns := min(x, y/4)
	if turns%2 == 0 {
		return "Bob"
	}
	return "Alice"
}

func main() {
	fmt.Println(winningPlayer(2, 7))
	fmt.Println(winningPlayer(4, 11))
}
