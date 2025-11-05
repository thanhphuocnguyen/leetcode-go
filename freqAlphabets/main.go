package main

import (
	"fmt"
	"strings"
)

func freqAlphabets(s string) string {
	bs := make([]byte, 0)
	meetSharp := false
	i := len(s) - 1
	for i >= 0 {
		if s[i] == '#' {
			meetSharp = true
			i--
			continue
		}

		if meetSharp {
			num := int(s[i-1]-'0')*10 + int(s[i]-'0')
			bs = append(bs, byte('a'+num-1))
			meetSharp = false
			i--
		} else {
			bs = append(bs, byte('a'+s[i]-'1'))
		}
		i--
	}
	var sb strings.Builder
	for j := len(bs) - 1; j >= 0; j-- {
		sb.WriteByte(bs[j])
	}
	return sb.String()
}

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}
