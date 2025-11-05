package main

func decodeMessage(key string, message string) string {
	var dict [26]byte
	var currByte byte = 'a'
	for _, k := range key {
		if k != ' ' && dict[k-'a'] == 0 {
			dict[k-'a'] = currByte
			currByte++
		}
	}
	ans := make([]byte, len(message))
	for i, c := range message {
		if c == ' ' {
			ans[i] = ' '
		} else {
			ans[i] = dict[c-'a']
		}
	}
	return string(ans)
}

func main() {
	println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
	println(decodeMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"))
}
