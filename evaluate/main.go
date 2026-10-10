package main

import "fmt"

func evaluate(s string, knowledge [][]string) string {
	dict := make(map[string]string)
	for _, k := range knowledge {
		dict[k[0]] = k[1]
	}

	bs := make([]rune, 0, len(s))
	opened := false
	temp := make([]rune, 0)
	for _, r := range s {
		if r == '(' {
			opened = true
			temp = make([]rune, 0)
		} else if r == ')' {
			opened = false
			if val, ok := dict[string(temp)]; ok {
				bs = append(bs, []rune(val)...)
			} else {
				bs = append(bs, '?')
			}
		} else if opened {
			temp = append(temp, r)
		} else {
			bs = append(bs, r)
		}
	}
	return string(bs)
}

func main() {
	fmt.Println(evaluate("(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}}))
}
