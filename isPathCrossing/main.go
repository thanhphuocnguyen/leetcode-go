package main

import "fmt"

func isPathCrossing(path string) bool {
	mp := make(map[string]bool)
	mp["0-0"] = true
	currX, currY := 0, 0
	for _, c := range path {
		switch c {
		case 'N':
			currY += 1
		case 'E':
			currX += 1
		case 'S':
			currY -= 1
		case 'W':
			currX -= 1
		}
		coord := fmt.Sprintf("%d-%d", currX, currY)

		if mp[coord] {
			return true
		}
		mp[coord] = true
	}

	return false
}

func main() {
	fmt.Println(isPathCrossing("ENNNNNNNNNNNEEEEEEEEEESSSSSSSSSS"))
	fmt.Println(isPathCrossing("NES"))
	fmt.Println(isPathCrossing("NESWW"))
}
