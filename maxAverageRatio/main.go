package main

import "container/heap"

type Item struct {
	passes   int
	students int
	gain     float64
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].gain > pq[j].gain
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, class := range classes {
		passes, students := class[0], class[1]
		heap.Push(&pq, &Item{passes, students, calculateGain(passes, students)})
	}

	for extraStudents > 0 {
		extraStudents--
		potential := heap.Pop(&pq).(*Item)
		newPasses, newStudents := potential.passes+1, potential.students+1
		newGain := calculateGain(newPasses, newStudents)
		heap.Push(&pq, &Item{newPasses, newStudents, newGain})
	}
	var ans float64
	for pq.Len() > 0 {

	}
	return ans / float64(len(classes))

}

func calculateGain(passes, students int) float64 {
	return float64(passes+1)/float64(students+1) - float64(passes)/float64(students)
}

func main() {
	classes := [][]int{{1, 2}, {3, 5}, {2, 2}}
	extraStudents := 2
	println(maxAverageRatio(classes, extraStudents))
}
