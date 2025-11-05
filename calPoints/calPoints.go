package main

import "fmt"

func calPoints(operations []string) int {
	records := make([]int, 0)
	for _, op := range operations {
		switch op {
		case "C":
			records = records[:len(records)-1]
		case "D":
			double := records[len(records)-1] * 2
			records = append(records, double)
		case "+":
			prev1, prev2 := records[len(records)-1], records[len(records)-2]
			records = append(records, prev1+prev2)
		default:
			num := 0
			sign := 1
			for _, r := range op {
				if r == '-' {
					sign = -1
					continue
				}
				num = num*10 + int(r-'0')
			}
			records = append(records, sign*num)
		}
	}
	ans := 0
	for _, rc := range records {
		ans += rc
	}
	return ans
}

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}
