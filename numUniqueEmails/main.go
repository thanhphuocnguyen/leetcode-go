package main

import "fmt"

func numUniqueEmails(emails []string) int {
	mp := make(map[string]bool)
	ans := 0
	for _, email := range emails {
		seenAt := false
		skip := false
		bs := make([]rune, 0)
		for _, c := range email {
			if c == '@' {
				bs = append(bs, c)
				seenAt = true
				skip = false
			} else {
				if !seenAt {
					if c == '.' {
						continue
					} else if c == '+' {
						skip = true
					} else if !skip {
						bs = append(bs, c)
					}
				} else {
					bs = append(bs, c)
				}
			}
		}
		s := string(bs)
		if !mp[s] {
			mp[s] = true
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}
