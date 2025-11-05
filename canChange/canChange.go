package main

func canChange(start string, target string) bool {
	n := len(start)
	start += "+"
	target += "+"
	for i, j := 0, 0; i < n || j < n; i, j = i+1, j+1 {
		for i < n && start[i] == '_' {
			i++
		}

		for j < n && target[j] == '_' {
			j++
		}
		c := start[i]
		if c != target[j] {
			return false
		}

		if start[i] == 'L' && i < j {
			return false
		}

		if start[i] == 'R' && i > j {
			return false
		}
	}
	return true
}

func main() {
	start := "_L__R__R_"
	target := "L______RR"
	println(canChange(start, target))
}
