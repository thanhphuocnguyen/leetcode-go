package main

import (
	"fmt"
	"sort"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	blOrder := map[string]int{"electronics": 1, "grocery": 2, "pharmacy": 3, "restaurant": 4}
	validCoupons := make([][2]string, 0, len(code))
	for i, cd := range code {
		bl := businessLine[i]
		active := isActive[i]
		if _, ok := blOrder[bl]; !ok || !active || !checkValidCode(cd) {
			continue
		}
		validCoupons = append(validCoupons, [2]string{cd, bl})
	}
	sort.Slice(validCoupons, func(i, j int) bool {
		cpi, cpj := validCoupons[i], validCoupons[j]
		if blOrder[cpi[1]] == blOrder[cpj[1]] {
			return cpi[0] < cpj[0]
		} else {
			return blOrder[cpi[1]] < blOrder[cpj[1]]
		}
	})
	ans := make([]string, len(validCoupons))
	for i, cp := range validCoupons {
		ans[i] = cp[0]
	}
	return ans
}

func checkValidCode(code string) bool {
	if len(code) == 0 {
		return false
	}
	for _, r := range code {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(validateCoupons([]string{"m", "A", "B", "P", "J", "P", "u", "W", "4", "J", "C", "9"}, []string{"electronics", "invalid", "invalid", "pharmacy", "invalid", "electronics", "restaurant", "grocery", "restaurant", "pharmacy", "pharmacy", "pharmacy"}, []bool{true, true, false, true, false, true, true, false, false, false, true, false}))
	fmt.Println(validateCoupons([]string{"1OFw", "0MvB"}, []string{"electronics", "pharmacy"}, []bool{true, true}))
	fmt.Println(validateCoupons([]string{"SAVE20", "", "PHARMA5", "SAVE@20"}, []string{"restaurant", "grocery", "pharmacy", "restaurant"}, []bool{true, true, true, true}))
}
