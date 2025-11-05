package main

func robotWithString(s string) string {
	freq := make([]int, 26)
	for _, r := range s {
		freq[int(r-'a')]++
	}
	ans := make([]byte, 0)
	st := make([]byte, 0)
	for i := range s {
		st = append(st, s[i])
		freq[int(s[i]-'a')]--
		startIdx := findStartIdx(freq)
		for len(st) > 0 && int(st[len(st)-1]-'a') <= startIdx {
			ans = append(ans, st[len(st)-1])
			st = st[:len(st)-1]
		}
	}
	for len(st) > 0 {
		ans = append(ans, st[len(st)-1])
		st = st[:len(st)-1]
	}
	return string(ans)
}

func findStartIdx(freq []int) int {
	for i, f := range freq {
		if f > 0 {
			return i
		}
	}
	return 0
}

func main() {
	println(robotWithString("zza"))
	println(robotWithString("bac"))
	println(robotWithString("bdda"))
	println(robotWithString("bydizfve"))
}
