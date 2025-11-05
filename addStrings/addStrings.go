package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	if len(num2) > len(num1) {
		temp := num1
		num1 = num2
		num2 = temp
	}
	var carry rune = 0
	nums := make([]rune, len(num1))
	j := len(num2) - 1
	for i := len(num1) - 1; i >= 0; i-- {
		var r2 rune = 0
		if j >= 0 {
			r2 = rune(num2[j] - '0')
			j--
		}
		sum := rune(num1[i]-'0') + r2 + carry
		if sum >= 10 {
			carry = sum / 10
		} else {
			carry = 0
		}
		nums[i] = sum%10 + '0'
	}

	if carry > 0 {
		return string('0'+carry) + string(nums)
	}
	return string(nums)
}

func main() {
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("3876620623801494171", "6529364523802684779"))
}
