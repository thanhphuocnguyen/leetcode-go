package main

func smallestNumber(n int, t int) int {
	for i := 0; i < 10; i++ {
		x := i + n
		prod := 1
		for x > 0 {
			prod *= x % 10
			x /= 10
		}
		if prod%t == 0 {
			return i + n
		}
	}
	return -1
}
