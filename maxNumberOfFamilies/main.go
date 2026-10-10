package main

import "fmt"

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	mp := make(map[int][10]bool, len(reservedSeats))
	for _, rs := range reservedSeats {
		row, seat := rs[0]-1, rs[1]-1
		if _, ok := mp[row]; !ok {
			mp[row] = [10]bool{}
		}
		arr := mp[row]
		arr[seat] = true
		mp[row] = arr
	}

	ans := 2 * (n - len(mp))
	for _, v := range mp {
		cnt := 0

		// case 2
		valid := true
		for i := 1; i < 4; i++ {
			if v[i] {
				valid = false
				break
			}
		}
		if valid {
			cnt++
		}
		valid = true
		for i := 5; i < 9; i++ {
			if v[i] {
				valid = false
				break
			}
		}
		if valid {
			cnt++
		}

		// case 1
		if cnt == 0 {
			valid = true
			for i := 3; i < 7; i++ {
				if v[i] {
					valid = false
					break
				}
			}
			if valid {
				cnt++
			}
		}
		ans += cnt
	}
	return ans
}

func main() {
	fmt.Println(maxNumberOfFamilies(3, [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}))
	fmt.Println(maxNumberOfFamilies(4, [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}}))
	fmt.Println(maxNumberOfFamilies(2, [][]int{{2, 1}, {1, 8}, {2, 6}}))
}
