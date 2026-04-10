package main

import (
	"fmt"
)

type criteria struct {
	lcase bool
	ucase bool
	digit bool
	spec  bool
}

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	crit := criteria{}
	for i, r := range password {
		if i > 0 {
			if r == rune(password[i-1]) {
				return false
			}
		}
		if r >= 'a' && r <= 'z' {
			crit.lcase = true
		} else if r >= 'A' && r <= 'Z' {
			crit.ucase = true
		} else if r >= '0' && r <= '9' {
			crit.digit = true
		} else if !crit.spec {
			for _, sr := range "!@#$%^&*()-+" {
				if sr == r {
					crit.spec = true
					break
				}
			}
		}
	}

	return crit.digit && crit.lcase && crit.ucase && crit.spec
}

func main() {
	fmt.Println(strongPasswordCheckerII("IloveLe3tcode!"))
}
