package main

import "sort"

func maxDistance(side int, points [][]int, k int) int {
	arr := make([]int64, 0)
	sz := int64(side)
	for _, p := range points {
		x, y := int64(p[0]), int64(p[1])
		if x == 0 {
			arr = append(arr, y)
		} else if y == sz {
			arr = append(arr, sz+x)
		} else if x == sz {
			arr = append(arr, 3*sz-y)
		} else {
			arr = append(arr, 4*sz-x)
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var lo, hi int64 = 1, sz

	ans := 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if check(arr, int64(sz), k, mid) {
			lo = mid + 1
			ans = int(mid)
		} else {
			hi = mid - 1
		}
	}

	return ans
}

func check(arr []int64, size int64, k int, limit int64) bool {
	perimeter := size * 4
	for _, start := range arr {
		end := start + perimeter - limit
		curr := start
		for i := 0; i < k-1; i++ {
			idx := lowerBound(arr, curr+limit)
			if idx == len(arr) || arr[idx] > end {
				curr = -1
				break
			}
			curr = arr[idx]
		}
		if curr >= 0 {
			return true
		}
	}
	return false
}

func lowerBound(arr []int64, target int64) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
