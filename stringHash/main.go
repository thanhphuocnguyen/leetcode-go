package main

func stringHash(s string, k int) string {
	bs := make([]byte, len(s)/2)
	cnt, sum, j := 0, 0, 0
	for i := range len(s) {
		sum += int(s[i] - 'a')
		cnt++
		if k == cnt {
			bs[j] = byte((sum % 26) + 'a')
			j++
			cnt, sum = 0, 0
		}
	}
	return string(bs)
}

func main()
