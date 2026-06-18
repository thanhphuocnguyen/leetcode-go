package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	digits := make([]int, len(num1)+len(num2))

	for i := len(num2) - 1; i >= 0; i-- {
		d1 := int(num2[i] - '0')
		for j := len(num1) - 1; j >= 0; j-- {
			d2 := int(num1[j] - '0')
			mul := (d1 * d2) + digits[i+j+1]
			digits[i+j] += mul / 10
			digits[i+j+1] = mul % 10

		}
	}
	bs := make([]byte, 0, len(num1)+len(num2))
	i := 0
	for digits[i] == 0 {
		i++
	}
	for ; i < len(digits); i++ {
		bs = append(bs, byte(digits[i]+'0'))
	}
	return string(bs)
}
func main() {
	fmt.Println(multiply("123", "456"))
}
