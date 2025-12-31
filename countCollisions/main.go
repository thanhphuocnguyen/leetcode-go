package main

import "fmt"

func countCollisions(directions string) int {
	st := make([]rune, 0)
	n := len(directions)
	i, j := 0, n-1
	for i < n && (directions[i] == 'L') {
		i++
	}
	for j >= 0 && directions[j] == 'R' {
		j--
	}
	ans := 0
	for i <= j {
		if len(st) > 0 {
			switch directions[i] {
			case 'R':
				for len(st) > 0 && st[len(st)-1] != 'R' {
					st = st[:len(st)-1]
				}
				st = append(st, 'R')
			case 'L':
				cnt := 0
				for len(st) > 0 {
					if st[len(st)-1] == 'R' {
						cnt++
					} else {
						break
					}
					st = st[:len(st)-1]
				}
				ans += cnt + 1

			case 'S':
				for len(st) > 0 && st[len(st)-1] == 'R' {
					ans++
					st = st[:len(st)-1]
				}
			}

		} else {
			st = append(st, rune(directions[i]))
		}
		i++
	}

	return ans
}

func main() {
	fmt.Println(countCollisions("SSRSSRLLRSLLRSRSSRLRRRRLLRRLSSRR"))
	fmt.Println(countCollisions("RLRSLL"))
}
