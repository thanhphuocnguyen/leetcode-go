package main

import "fmt"

func convertTime(current string, correct string) int {
	currMin, correctMin := 0, 0

	for i := 3; i < 5; i++ {
		currMin = currMin*10 + int(current[i]-'0')
		correctMin = correctMin*10 + int(correct[i]-'0')
	}
	currHour, correctHour := 0, 0
	for i := 0; i < 2; i++ {
		currHour = currHour*10 + int(current[i]-'0')
		correctHour = correctHour*10 + int(correct[i]-'0')
	}
	remain := (correctHour*60 + correctMin) - (currHour*60 + currMin)
	ans := 0
	for _, min := range []int{60, 15, 5, 1} {
		if remain >= min {
			ans += remain / min
			remain %= min
		}
		if remain == 0 {
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(convertTime("02:30", "04:35"))
}
