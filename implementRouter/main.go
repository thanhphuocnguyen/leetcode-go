package main

import "fmt"

type Packet struct {
	Src       int
	Dest      int
	TimeStamp int
	Next      *Packet
}

type Destination struct {
	start      int
	packetSet  map[[2]int]bool
	timestamps []int
}

type Queue struct {
	len  int
	head *Packet
	tail *Packet
}

func (q *Queue) Enqueue(source int, destination int, timestamp int) {
	node := &Packet{
		Src:       source,
		Dest:      destination,
		TimeStamp: timestamp,
	}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.Next = node
		q.tail = q.tail.Next
	}
	q.len++
}

func (q *Queue) Deque() (source int, destination int, timestamp int) {
	head := q.head
	q.head = q.head.Next
	q.len--
	return head.Src, head.Dest, head.TimeStamp
}

func (q *Queue) Len() int {
	return q.len
}

type Router struct {
	limit       int
	packetQueue *Queue
	destMap     map[int]*Destination
}

func Constructor(memoryLimit int) Router {
	return Router{
		limit: memoryLimit,
		packetQueue: &Queue{
			len: 0,
		},
		destMap: make(map[int]*Destination),
	}
}

func (r *Router) AddPacket(source int, destination int, timestamp int) bool {
	if _, ok := r.destMap[destination]; !ok {
		r.destMap[destination] = &Destination{
			start:      0,
			timestamps: make([]int, 0),
			packetSet:  make(map[[2]int]bool),
		}
	}

	key := [2]int{source, timestamp}

	if r.destMap[destination].packetSet[key] {
		return false
	}

	if r.packetQueue.Len() == r.limit {
		// remove first!
		outdatedSrc, outdatedDst, outdatedTime := r.packetQueue.Deque()
		key := [2]int{outdatedSrc, outdatedTime}
		r.destMap[outdatedDst].packetSet[key] = false
		r.destMap[outdatedDst].start++

	}

	r.packetQueue.Enqueue(source, destination, timestamp)

	r.destMap[destination].timestamps = append(r.destMap[destination].timestamps, timestamp)
	r.destMap[destination].packetSet[key] = true
	return true
}

func (r *Router) ForwardPacket() []int {
	if r.packetQueue.Len() == 0 {
		return []int{}
	}

	src, dest, timestamp := r.packetQueue.Deque()
	res := []int{src, dest, timestamp}
	key := [2]int{src, timestamp}
	r.destMap[dest].packetSet[key] = false
	r.destMap[dest].start++
	if r.destMap[dest].start == len(r.destMap[dest].timestamps) {
		r.destMap[dest].cleanup()
	}

	return res
}

func (r *Router) GetCount(destination int, startTime int, endTime int) int {
	arr := r.destMap[destination].timestamps
	start := r.destMap[destination].start

	left := lowerBound(arr, start, startTime)
	right := upperBound(arr, start, endTime)
	return right - left
}
func (r *Destination) cleanup() {
	r.timestamps = []int{}
}

func lowerBound(arr []int, start, target int) int {
	left, right := start, len(arr)
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

func upperBound(arr []int, start, target int) int {
	left, right := start, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */

func main() {
	// router := Constructor(4)
	// fmt.Println(router.AddPacket(4, 5, 1)) // true
	// fmt.Println(router.GetCount(5, 1, 1))  // true

	// router := Constructor(3)
	// fmt.Println(router.AddPacket(1, 4, 90))   // true
	// fmt.Println(router.AddPacket(2, 5, 90))   // true
	// fmt.Println(router.AddPacket(1, 4, 90))   // false
	// fmt.Println(router.AddPacket(3, 5, 95))   // true
	// fmt.Println(router.AddPacket(4, 5, 105))  // true
	// fmt.Println(router.ForwardPacket())       // true
	// fmt.Println(router.AddPacket(5, 2, 110))  // true
	// fmt.Println(router.GetCount(5, 100, 110)) // true

	router := Constructor(2)
	fmt.Println(router.AddPacket(3, 2, 1)) // true
	fmt.Println(router.AddPacket(4, 1, 1)) // true
	fmt.Println(router.AddPacket(4, 2, 1)) // false
	fmt.Println(router.GetCount(2, 1, 1))  // false
}
