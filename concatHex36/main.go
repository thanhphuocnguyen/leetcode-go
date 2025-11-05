package main

import (
	"fmt"
)

func concatHex36(n int) string {
	squared := n * n
	cubed := n * n * n
	bs1 := make([]byte, 0)
	for squared > 0 {
		remainder := squared % 16
		if remainder < 10 {
			bs1 = append(bs1, (byte(remainder + '0')))
		} else {
			bs1 = append(bs1, byte(remainder-10+'A'))
		}
		squared /= 16
	}
	for i, j := 0, len(bs1)-1; i < j; i, j = i+1, j-1 {
		bs1[i], bs1[j] = bs1[j], bs1[i]
	}

	bs2 := make([]byte, 0)
	for cubed > 0 {
		remainder := cubed % 36
		if remainder < 10 {
			bs2 = append(bs2, byte(remainder+'0'))
		} else {
			bs2 = append(bs2, byte(remainder-10+'A'))
		}
		cubed /= 36
	}
	for i, j := 0, len(bs2)-1; i < j; i, j = i+1, j-1 {
		bs2[i], bs2[j] = bs2[j], bs2[i]
	}
	return string(bs1) + string(bs2)
}

func main() {
	fmt.Println(concatHex36(36))
	fmt.Println(concatHex36(13))
}
