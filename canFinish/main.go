package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	for i := range numCourses {
		g[i] = make([]int, 0)
	}
	for _, p := range prerequisites {
		a, b := p[0], p[1]
		g[a] = append(g[a], b)
	}
	st := make([]bool, numCourses)
	visited := make([]bool, numCourses)

	var isCyclic func(node int) bool
	isCyclic = func(node int) bool {
		if st[node] {
			return true
		}
		if visited[node] {
			return false
		}

		st[node] = true
		visited[node] = true
		for _, nei := range g[node] {
			if isCyclic(nei) {
				return true
			}
		}

		st[node] = false
		return false
	}

	for i := range numCourses {
		if isCyclic(i) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}
