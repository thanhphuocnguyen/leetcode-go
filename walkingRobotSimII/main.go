package main

import "fmt"

type Robot struct {
	width     int
	height    int
	perimeter int
	direction int // 0,1,2,3 as North, East, South, West
	x         int
	y         int
}

func Constructor(width int, height int) Robot {
	width--
	height--
	return Robot{width, height, 2 * (width + height), 1, 0, 0}
}

func (this *Robot) Step(num int) {
	if num >= this.perimeter {
		num = num % this.perimeter
		if num == 0 && this.x == 0 && this.y == 0 {
			this.direction = 2
		}
	}
	if this.direction%2 == 0 {
		ty := this.y
		ty += num * (1 - this.direction)
		if ty > this.height {
			this.direction = (this.direction + 3) % 4
			this.y = this.height
			this.Step(ty - this.height)
		} else if ty < 0 {
			this.direction = (this.direction + 3) % 4
			this.y = 0
			this.Step(-ty)
		} else {
			this.y = ty
		}
	} else {
		tx := this.x
		tx += num * (2 - this.direction)
		if tx > this.width {
			this.direction = (this.direction + 3) % 4
			this.x = this.width
			this.Step(tx - this.width)
		} else if tx < 0 {
			this.direction = (this.direction + 3) % 4
			this.x = 0
			this.Step(-tx)
		} else {
			this.x = tx
		}
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.x, this.y}
}

func (this *Robot) GetDir() string {
	switch this.direction {
	case 0:
		return "North"
	case 1:
		return "East"
	case 2:
		return "South"
	case 3:
		return "West"
	default:
		return ""
	}
}

func main() {
	var rb Robot
	steps := []string{"Robot", "step", "getPos", "getDir", "step", "getPos", "getDir", "step", "step", "getPos", "getDir", "step", "step", "step", "getPos", "getDir", "step", "step", "step", "getPos", "getDir", "step"}
	values := [][]int{{8, 2}, {17}, {}, {}, {21}, {}, {}, {22}, {34}, {}, {}, {1}, {46}, {35}, {}, {}, {44}, {14}, {31}, {}, {}, {50}}
	for i, step := range steps {
		switch step {
		case "Robot":
			rb = Constructor(values[i][0], values[i][1])
		case "step":
			rb.Step(values[i][0])
			fmt.Printf("%s: %v, ", rb.GetDir(), rb.GetPos())
		default:
			continue
		}
	}
}
