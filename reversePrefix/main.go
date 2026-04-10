package main

import "fmt"

func reversePrefix(s string, k int) string {
	bs := []byte(s)
	for i := range k / 2 {
		j := k - i - 1
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}

func main() {
	fmt.Println(reversePrefix("abcd", 2))
	fmt.Println(reversePrefix("xyz", 3))
}
