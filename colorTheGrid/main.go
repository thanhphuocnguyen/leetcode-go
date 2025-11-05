package main

import "math"

const MOD = int(10e9) + 7

func colorTheGrid(m int, n int) int {
	valid := make(map[int][]int)
	maskEnd := int(math.Pow(3, float64(m)))
	for mask := range maskEnd {
		color := make([]int, m)
		mm := mask
		for i := range m {
			color[i] = mm % 3
			mm /= 3
		}
		check := true
		for i := range m - 1 {
			if color[i] == color[i+1] {
				check = false
				break
			}
		}
		if check {
			valid[mask] = color
		}
	}
	adjacent := make(map[int][]int)
	for mask1, color1 := range valid {
		for mask2, color2 := range valid {
			check := true
			for i := range m {
				if color1[i] == color2[i] {
					check = false
					break
				}
			}
			if check {
				if _, ok := adjacent[mask1]; !ok {
					adjacent[mask1] = make([]int, 0)
				}
				adjacent[mask1] = append(adjacent[mask1], mask2)
			}
		}
	}

	f := make(map[int]int)
	for mask := range valid {
		f[mask] = 1
	}

	for i := 1; i < n; i++ {
		g := make(map[int]int)
		for mask2 := range valid {
			if _, ok := adjacent[mask2]; ok {
				for _, mask1 := range adjacent[mask2] {
					if _, ok := g[mask2]; !ok {
						g[mask2] = 0
					}
					g[mask2] = (g[mask2] + f[mask1]) % MOD
				}
			}
		}
		f = g
	}

	ans := 0
	for _, val := range f {
		ans = (ans + val) % MOD
	}
	return ans
}
