package main

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	ans := make([]string, 0)
	mp := make(map[string]int)
	for _, str := range cpdomains {
		pair := strings.Split(str, " ")
		numStr := pair[0]
		num := 0
		for _, r := range numStr {
			num = num*10 + int(r-'0')
		}
		arr := strings.Split(pair[1], ".")
		combine := ""
		for i := len(arr) - 1; i >= 0; i-- {
			if len(combine) > 0 {
				combine = arr[i] + "." + combine
			} else {
				combine = arr[i]
			}
			mp[combine] += num
		}
	}
	for k, v := range mp {
		ans = append(ans, strconv.Itoa(v)+" "+k)
	}
	return ans
}

func main() {
	subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"})
}
