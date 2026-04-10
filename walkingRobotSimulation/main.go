package main

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {
	obstMp := make(map[[2]int]bool)
	for _, o := range obstacles {
		obstMp[[2]int{o[0], o[1]}] = true
	}
	// init is at 0,0 and facing north
	x, y := 0, 0
	// n -> e -> s -> w
	// 0,1,2,3
	direction := 0
	ans := 0
	for _, cmd := range commands {
		if cmd == -1 {
			// go right
			direction += 1
			direction %= 4
		} else if cmd == -2 {
			// go left
			direction += 3
			direction %= 4
		} else {
			for i := 0; i < cmd; i++ {
				tx, ty := x, y
				tx += (direction % 2) * (2 - direction)
				ty += ((direction + 1) % 2) * (1 - direction)
				if obstMp[[2]int{tx, ty}] {
					break
				}
				x, y = tx, ty
				ans = max(ans, (x*x)+(y*y))
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
	fmt.Println(robotSim([]int{4, -1, 3}, [][]int{}))
}
