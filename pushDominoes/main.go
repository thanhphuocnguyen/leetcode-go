package main

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	ans := []byte(dominoes)
	rightForce := make([]int, n)
	force := 0
	for l := 0; l < n; l++ {
		if ans[l] == 'L' {
			force = 0
		} else if ans[l] == 'R' {
			force = n
		} else {
			force = max(force-1, 0)
		}
		rightForce[l] = force
	}

	force = 0
	for r := n - 1; r >= 0; r-- {
		if ans[r] == 'L' {
			force = n
		} else if ans[r] == 'R' {
			force = 0
		} else {
			force = max(force-1, 0)
		}

		if rightForce[r] > force {
			ans[r] = 'R'
		} else if rightForce[r] < force {
			ans[r] = 'L'
		} else {
			ans[r] = '.'
		}
	}
	return string(ans)
}

func main() {
	// Example usage
	dominoes := "RR.L"
	result := pushDominoes(dominoes)
	println(result) // Output: "RR.L"
	dominoes = ".L.R...LR..L.."
	result = pushDominoes(dominoes)
	println(result) // Output: "RR.L"
}
