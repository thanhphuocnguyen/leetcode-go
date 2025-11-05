package main

func maxChunksToSorted(arr []int) int {
	chunk, maxEl := 0, 0
	for i, num := range arr {
		maxEl = max(maxEl, num)
		if i == maxEl {
			chunk++
		}
	}
	return chunk
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}
	maxChunksToSorted(arr)
}
