package main

import "fmt"

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	personNeedToLearn := make(map[int]bool)
	for _, fs := range friendships {
		u, v := fs[0]-1, fs[1]-1
		need := true
		langMp := make(map[int]bool)
		for _, langU := range languages[u] {
			langMp[langU] = true
		}
		for _, langV := range languages[v] {
			if langMp[langV] {
				need = false
				break
			}
		}
		if need {
			personNeedToLearn[u] = true
			personNeedToLearn[v] = true
		}
	}
	ans := len(languages) + 1
	for i := 1; i <= n; i++ {
		currStudent := 0
		for person := range personNeedToLearn {
			foundLang := false
			for _, lang := range languages[person] {
				if i == lang {
					foundLang = true
					break
				}
			}
			if foundLang {
				currStudent++
			}
		}
		ans = min(ans, currStudent)
	}
	return ans
}

func main() {
	fmt.Println(minimumTeachings(2, [][]int{{1}, {2}, {1, 2}}, [][]int{{1, 2}, {1, 3}, {2, 3}}))
}
