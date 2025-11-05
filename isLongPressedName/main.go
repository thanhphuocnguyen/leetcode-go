package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for i < len(name) || j < len(typed) {
		if i < len(name) && j < len(typed) && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && j < len(typed) && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isLongPressedName("alexd", "ale"))
	fmt.Println(isLongPressedName("laiden", "laiden"))
	fmt.Println(isLongPressedName("alex", "aaleexa"))
	fmt.Println(isLongPressedName("saeed", "ssaaedd"))
	fmt.Println(isLongPressedName("alex", "aaleex"))
}
