package main

import "fmt"

const MODULO int = 1000000007

func productQueries(n int, queries [][]int) []int {
	prefixProd := []int{1}
	currProd := 1
	for n > 0 {
		if n&1 != 0 {
			prev := prefixProd[len(prefixProd)-1]
			prefixProd = append(prefixProd, (prev*currProd)%MODULO)
		}
		n >>= 1
		currProd <<= 1
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		ans[i] = (prefixProd[r+1] / prefixProd[l])
	}

	return ans
}

func main() {
	fmt.Println(productQueries(919, [][]int{{5, 5}, {4, 4}, {0, 1}, {1, 5}, {4, 6}, {6, 6}, {5, 6}, {0, 3}, {5, 5}, {5, 6}, {1, 2}, {3, 5}, {3, 6}, {5, 5}, {4, 4}, {1, 1}, {2, 4}, {4, 5}, {4, 4}, {5, 6}, {0, 4}, {3, 3}, {0, 4}, {0, 5}, {4, 4}, {5, 5}, {4, 6}, {4, 5}, {0, 4}, {6, 6}, {6, 6}, {6, 6}, {2, 2}, {0, 5}, {1, 4}, {0, 3}, {2, 4}, {5, 5}, {6, 6}, {2, 2}, {2, 3}, {5, 5}, {0, 6}, {3, 3}, {6, 6}, {4, 4}, {0, 0}, {0, 2}, {6, 6}, {6, 6}, {3, 6}, {0, 4}, {6, 6}, {2, 2}, {4, 6}}))
	fmt.Println(productQueries(2, [][]int{{0, 0}}))
	fmt.Println(productQueries(15, [][]int{{0, 1}, {2, 2}, {0, 3}}))
}
