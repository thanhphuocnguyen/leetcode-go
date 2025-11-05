package main

func stringSequence(target string) []string {
	ans := make([]string, 0)
	str := ""
	for _, r := range target {
		c := 'a'
		for r != c {
			ans = append(ans, str+string(c))
			c++
		}
		ans = append(ans, str+string(c))
		str = ans[len(ans)-1]
	}

	return ans
}

func main() {

}
