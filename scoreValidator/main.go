package main

import "fmt"

func scoreValidator(events []string) []int {
	ans := []int{0, 0}
	for _, evt := range events {
		switch evt {
		case "W":
			ans[1]++
		case "WD":
			fallthrough
		case "NB":
			ans[0]++
		default:
			ans[0] += int(evt[0] - '0')
		}
		if ans[1] == 10 {
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(scoreValidator([]string{"1", "4", "W", "6", "WD"}))
}
