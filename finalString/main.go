package main

func finalString(s string) string {
	bs := make([]byte, 0)
	for i := range len(s) {
		if s[i] == 'i' {
			for j := 0; j < len(bs)/2; j++ {
				bs[j], bs[len(bs)-j-1] = bs[len(bs)-j-1], bs[j]
			}
		} else {
			bs = append(bs, s[i])
		}
	}
	return string(bs)
}

func main() {
	finalString("string")
}
