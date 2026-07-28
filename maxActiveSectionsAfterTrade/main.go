package main

import "fmt"

func maxActiveSectionsAfterTrade(s string) int {
	firstZeros := 0
	oneCnt := 0
	secondZeros := 0
	totalActive := 0
	maxWindow := 0
	for _, r := range s {
		if r == '0' {
			// enter the window that have patter '0..1..'
			if oneCnt > 0 {
				// now it will be '..0..1..0..'
				secondZeros++
			} else {
				// still '00...'
				firstZeros++
			}
		} else {
			totalActive++
			// flip when the window is invalid and reset window
			if secondZeros > 0 {
				if firstZeros+secondZeros > maxWindow {
					maxWindow = firstZeros + secondZeros
				}
				firstZeros = secondZeros
				secondZeros = 0
				// start new window
				oneCnt = 1
			} else if firstZeros > 0 {
				oneCnt++
			}
		}
	}
	if secondZeros > 0 && firstZeros+secondZeros > maxWindow {
		maxWindow = firstZeros + secondZeros
	}
	if maxWindow > 0 {
		return totalActive + maxWindow
	}
	return totalActive
}

func main() {
	fmt.Println(maxActiveSectionsAfterTrade("01"))
	fmt.Println(maxActiveSectionsAfterTrade("0100"))
	fmt.Println(maxActiveSectionsAfterTrade("1000100"))
	fmt.Println(maxActiveSectionsAfterTrade("01010"))
	fmt.Println(maxActiveSectionsAfterTrade("01101"))
	fmt.Println(maxActiveSectionsAfterTrade("01100101"))
	fmt.Println(maxActiveSectionsAfterTrade("01101001"))
}
