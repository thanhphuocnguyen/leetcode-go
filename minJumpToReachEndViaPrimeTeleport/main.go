package main

import "fmt"

func minJumps(nums []int) int {
	n := len(nums)
	maxNum := 0
	for _, num := range nums {
		maxNum = max(num, maxNum)
	}
	bucket := make(map[int][]int)
	visited := make([]bool, n)
	prime := sieve(maxNum + 1)
	for i, num := range nums {
		for _, p := range prime[num] {
			if bucket[p] == nil {
				bucket[p] = []int{}
			}
			bucket[p] = append(bucket[p], i)
		}
	}
	// bfs
	q := [][2]int{{0, 0}}
	visited[0] = true
	for len(q) > 0 {
		curr := q[0]
		idx, dist := curr[0], curr[1]
		q = q[1:]
		if idx == n-1 {
			return dist
		}
		// back
		if idx > 0 && !visited[idx-1] {
			q = append(q, [2]int{idx - 1, dist + 1})
			visited[idx-1] = true
		}
		num := nums[idx]
		// forward
		if !visited[idx+1] {
			q = append(q, [2]int{idx + 1, dist + 1})
			visited[idx+1] = true
		}
		if len(bucket[num]) > 0 {
			for _, idx := range bucket[num] {
				if !visited[idx] {
					q = append(q, [2]int{idx, dist + 1})
					visited[idx] = true
				}
			}
			bucket[num] = []int{}
		}
	}

	return n - 1
}

func sieve(n int) [][]int {
	prime := make([][]int, n+1)
	for i := range n {
		prime[i] = []int{}
	}

	for p := 2; p <= n; p++ {
		if len(prime[p]) == 0 {
			for j := p; j <= n; j += p {
				prime[j] = append(prime[j], p)
			}
		}
	}

	return prime
}

func main() {
	fmt.Println(minJumps([]int{893, 786, 607, 137, 69, 381, 790, 233, 15, 42, 7, 764, 890, 269, 84, 262, 870, 514, 514, 650, 269, 485, 760, 181, 489, 107, 585, 428, 862, 563}))
	fmt.Println(minJumps([]int{5, 5, 3}))
	fmt.Println(minJumps([]int{4, 6, 5, 8}))
	fmt.Println(minJumps([]int{1, 2, 4, 6}))
	fmt.Println(minJumps([]int{2, 3, 4, 7, 9}))
}
