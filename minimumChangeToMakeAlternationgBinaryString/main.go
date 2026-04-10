package main

import "fmt"

func minOperations(s string) int {
	cnt := 0
	for i, r := range s {
		cnt += (int(r) ^ i) & 1
	}
	return min(cnt, len(s)-cnt)
}

func main() {
	fmt.Println(minOperations("1111"))
}
