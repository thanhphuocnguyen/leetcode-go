package main

import (
	"fmt"
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	ans := 0
	for _, num := range arr1 {
		l, r := 0, len(arr2)-1
		from, to := num-d, num+d
		found := false
		for l <= r {
			mid := (l + r) / 2
			if arr2[mid] >= from && arr2[mid] <= to {
				found = true
				break
			} else if arr2[mid] < from {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if !found {
			ans++
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(findTheDistanceValue(
		[]int{-803, 715, -224, 909, 121, -296, 872, 807, 715, 407, 94, -8, 572, 90, -520, -867, 485, -918, -827, -728, -653, -659, 865, 102, -564, -452, 554, -320, 229, 36, 722, -478, -247, -307, -304, -767, -404, -519, 776, 933, 236, 596, 954, 464},
		[]int{817, 1, -723, 187, 128, 577, -787, -344, -920, -168, -851, -222, 773, 614, -699, 696, -744, -302, -766, 259, 203, 601, 896, -226, -844, 168, 126, -542, 159, -833, 950, -454, -253, 824, -395, 155, 94, 894, -766, -63, 836, -433, -780, 611, -907, 695, -395, -975, 256, 373, -971, -813, -154, -765, 691, 812, 617, -919, -616, -510, 608, 201, -138, -669, -764, -77, -658, 394, -506, -675, 523, 730, -790, -109, 865, 975, -226, 651, 987, 111, 862, 675, -398, 126, -482, 457, -24, -356, -795, -575, 335, -350, -919, -945, -979, 611},
		37))
	fmt.Println(findTheDistanceValue([]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3))
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
}
