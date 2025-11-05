package main

import "fmt"

type OrderedStream struct {
	ptr    int
	stream []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		ptr:    1,
		stream: make([]string, n),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey-1] = value

	if idKey == this.ptr {
		cnt := 0
		start := this.ptr - 1
		i := this.ptr - 1
		for i < len(this.stream) && this.stream[i] != "" {
			i++
			cnt++
			this.ptr++
		}
		return this.stream[start : start+cnt]

	} else {
		return []string{}
	}
}

func main() {
	os := Constructor(5)
	ids := []int{3, 1, 2, 5, 4}
	vals := []string{"ccccc", "aaaaa", "bbbbb", "eeeee", "ddddd"}
	for i, id := range ids {
		fmt.Println(os.Insert(id, vals[i]))
	}
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
