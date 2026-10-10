package main

import "fmt"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := make([][]int, 0)
	m, n := len(firstList), len(secondList)
	if m == 0 || n == 0 {
		return ans
	}
	// for _, fist := range firstList {
	// 	s1, e1 := fist[0], fist[1]
	// 	for _, second := range secondList {
	// 		s2, e2 := second[0], second[1]
	// 		if s2 > e1 {
	// 			break
	// 		}
	// 		if e1 <= e2 && e1 >= s2 || e2 <= e1 && e2 >= s1 {
	// 			ans = append(ans, []int{max(s1, s2), min(e1, e2)})
	// 		}
	// 	}
	// }

	i, j := 0, 0
	for i < m && j < n {
		first, second := firstList[i], secondList[j]
		maxS, minE := max(first[0], second[0]), min(first[1], second[1])
		if maxS <= minE {
			ans = append(ans, []int{maxS, minE})
		}

		if second[1] > minE {
			i++
		} else if first[1] > minE {
			j++
		} else {
			i++
			j++
		}
	}

	return ans
}

func main() {
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
}
