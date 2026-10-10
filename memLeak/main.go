package main

import "fmt"

func memLeak(memory1 int, memory2 int) []int {
	time := 1
	for {
		if memory1 >= memory2 {
			if memory1 >= time {
				memory1 -= time
			} else {
				break
			}
		} else {
			if memory2 >= time {
				memory2 -= time
			} else {
				break
			}
		}
		time++
	}
	return []int{time, memory1, memory2}
}

func main() {
	fmt.Println(memLeak(2, 2))
	fmt.Println(memLeak(8, 11))
}
