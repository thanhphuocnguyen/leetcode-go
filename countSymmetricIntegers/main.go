package main

func countSymmetricIntegers(low int, high int) int {
	ans := 0
	for num := max(10, low); num <= high; num++ {
		if num > 10 && num < 100 {
			if num%11 == 0 {
				ans++
				num += 9
			}
		} else if num > 1000 {
			leftHalf := num/1000 + (num%1000)/100
			rightHalf := (num%100)/10 + (num%100)%10
			if leftHalf == rightHalf {
				ans++
			}
		}
	}
	return ans
}

func main() {
	println(countSymmetricIntegers(1, 1000))    // Example usage
	println(countSymmetricIntegers(1200, 1230)) // Example usage
}
