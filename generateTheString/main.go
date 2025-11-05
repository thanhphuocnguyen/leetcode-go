package main

import "math/rand"

func generateTheString(n int) string {
	bs := make([]byte, n)
	randomNumber := rand.Intn(25)
	if n%2 == 0 {
		for i := 0; i < n-1; i++ {
			bs[i] = byte('a' + randomNumber)
		}
		randomNumber = rand.Intn(25)
		bs[n-1] = byte('a' + randomNumber)
	} else {
		for i := 0; i < n; i++ {
			bs[i] = byte('a' + randomNumber)
		}
	}
	return string(bs)
}

func main() {
	generateTheString(4)
}
