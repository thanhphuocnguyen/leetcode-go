package main

import "fmt"

func processStr(s string) string {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '*':
			if len(bs) > 0 {
				bs = bs[:len(bs)-1]
			}
		case '#':
			n := len(bs)
			for j := 0; j < n; j++ {
				bs = append(bs, bs[j])
			}
		case '%':
			rev(bs)
		default:
			bs = append(bs, s[i])
		}
	}
	return string(bs)
}

func rev(bs []byte) {
	n := len(bs)
	for i := 0; i < n/2; i++ {
		bs[i], bs[n-i-1] = bs[n-i-1], bs[i]
	}
}

func main() {
	fmt.Println(processStr("***"))
}
