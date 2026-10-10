package main

import "fmt"

func remainingMethods(n int, k int, invocations [][]int) []int {
	G := make([][]int, n)
	for i := range n {
		G[i] = make([]int, 0)
	}
	for _, ivc := range invocations {
		u, v := ivc[0], ivc[1]
		G[u] = append(G[u], v)
	}
	suspicious := make([]bool, n)
	var markSuspicious func(u int)
	markSuspicious = func(u int) {
		suspicious[u] = true
		for _, v := range G[u] {
			if !suspicious[v] {
				markSuspicious(v)
			}
		}
	}
	markSuspicious(k)
	// visited := make([]bool, n)
	// var checkInvoked func(u int) bool
	// checkInvoked = func(u int) bool {
	// 	if suspicious[u] {
	// 		return true
	// 	}
	// 	visited[u] = true
	// 	for _, v := range G[u] {
	// 		if !visited[v] && checkInvoked(v) {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }

	ans := make([]int, 0)
	// for i, v := range suspicious {
	// 	if !v {
	// 		ans = append(ans, i)
	// 		check := checkInvoked(i)
	// 		if check {
	// 			all := make([]int, n)
	// 			for i := range n {
	// 				all[i] = i
	// 			}
	// 			return all
	// 		}
	// 	}
	// }

	for _, ivc := range invocations {
		u, v := ivc[0], ivc[1]
		if !suspicious[u] && suspicious[v] {
			all := make([]int, n)
			for i := range n {
				all[i] = i
			}
			return all
		}
	}
	for i, sus := range suspicious {
		if !sus {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Println(remainingMethods(3, 2, [][]int{{1, 0}, {2, 0}}))
	fmt.Println(remainingMethods(4, 1, [][]int{{1, 2}, {0, 1}, {3, 2}}))
}
