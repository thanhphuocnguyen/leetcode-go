package main

import "fmt"

func addBinary(a string, b string) string {
	var carry byte = '0'
	lenA, lenB := len(a), len(b)
	if lenB > lenA {
		lenA, lenB = lenB, lenA
		a, b = b, a
	}
	n := max(lenA, lenB)
	i, j := lenA-1, lenB-1
	bs := make([]byte, n)
	for i >= 0 && j >= 0 {
		if a[i] == '0' && b[j] == '0' {
			bs[i] = carry
			carry = '0'
		} else if (a[i] == '1' && b[j] == '0') || (a[i] == '0' && b[j] == '1') {
			if carry == '1' {
				bs[i] = '0'
			} else {
				bs[i] = '1'
			}
		} else {
			// 1 + 1 = 0 with carry = 1 => 1
			//
			if carry == '1' {
				bs[i] = '1'
			} else {
				bs[i] = '0'
			}
			carry = '1'
		}
		i--
		j--
	}

	for i >= 0 {
		if carry == '1' {
			if a[i] == '1' {
				bs[i] = '0'
			} else {
				bs[i] = '1'
				carry = 0
			}
		} else {
			bs[i] = a[i]
		}
		i--
	}
	for j >= 0 {
		if carry == '1' {
			if b[j] == '1' {
				bs[j] = '0'
			} else {
				bs[j] = '1'
				carry = 0
			}
		} else {
			bs[j] = b[j]
		}
		j--
	}
	if carry == '1' {
		return "1" + string(bs)
	}

	return string(bs)
}

func main() {
	fmt.Println(addBinary("101111", "10"))
	fmt.Println(addBinary("100", "110010"))
	fmt.Println(addBinary("1", "111"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("11", "1"))
}
