package main

import "fmt"

func checkStrings(s1 string, s2 string) bool {
	evenFreq, oddFreq := [26]int{}, [26]int{}

	for i := range len(s1) {
		if i%2 == 0 {
			// process for even
			evenFreq[s1[i]-'a']++
			evenFreq[s2[i]-'a']--
		} else {
			oddFreq[s1[i]-'a']++
			oddFreq[s2[i]-'a']--
			// process for odd
		}
	}

	for i := range 26 {
		if evenFreq[i] != 0 || oddFreq[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkStrings("abcdba", "cabdab"))
}
