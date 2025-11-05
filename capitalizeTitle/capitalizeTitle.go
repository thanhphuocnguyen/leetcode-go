package main

import "fmt"

func capitalizeTitle(title string) string {
	bs := make([]byte, len(title))
	start, cnt := 0, 0
	for i := range len(title) {
		cnt++
		b := title[i]
		if b == ' ' {
			cnt = 0
			start = i + 1
			bs[i] = b
			continue
		}

		// lower case all character is upper case
		if b >= 'A' && b <= 'Z' {
			b += 'a' - 'A'
		}
		bs[i] = b
		if cnt > 2 && bs[start] >= 'a' && bs[start] <= 'z' {
			// upper case
			bs[start] -= 'a' - 'A'
		}
	}

	return string(bs)
}

func main() {
	fmt.Println(capitalizeTitle("First leTTeR of EACH Word"))
}
