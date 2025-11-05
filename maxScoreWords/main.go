package main

func maxScoreWords(words []string, letters []byte, score []int) int {
	freq := make([]int, 26)
	for _, r := range letters {
		freq[r-'a']++
	}

	var backtrack func(idx int) int
	backtrack = func(idx int) int {
		if idx == len(words) {
			return 0
		}
		skip := backtrack(idx + 1)
		w := words[idx]
		sc := 0
		ok := true
		for _, r := range w {
			cIdx := r - 'a'
			freq[cIdx]--
			sc += score[r-'a']
			if freq[cIdx] < 0 {
				ok = false
			}
		}
		take := 0
		if ok {
			take = sc + backtrack(idx+1)
		}
		for _, r := range w {
			cIdx := r - 'a'
			freq[cIdx]++
		}
		return max(skip, take)
	}

	return backtrack(1)
}
