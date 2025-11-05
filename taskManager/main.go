package main

import (
	"container/heap"
	"fmt"
)

type Task struct {
	UserId   int
	TaskId   int
	priority int
	Removed  bool
	index    int
}
type PriorityQueue []*Task

// Len implements heap.Interface.
func (p PriorityQueue) Len() int {
	return len(p)
}

// Less implements heap.Interface.
func (p PriorityQueue) Less(i int, j int) bool {
	if p[i].priority == p[j].priority {
		return p[i].TaskId > p[j].TaskId
	}
	return p[i].priority > p[j].priority
}

// Pop implements heap.Interface.
func (p *PriorityQueue) Pop() any {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*p = old[:n-1]
	return item
}

func (p PriorityQueue) Peek() *Task {
	if p.Len() == 0 {
		return nil // Or handle an empty queue error
	}
	return p[0]
}

// Push implements heap.Interface.
func (p *PriorityQueue) Push(x any) {
	n := len(*p)
	item := x.(*Task)
	item.index = n
	*p = append(*p, item)
}

// Swap implements heap.Interface.
func (p PriorityQueue) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Task, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

type TaskManager struct {
	taskIdMp  map[int]*Task
	taskQueue PriorityQueue
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		taskIdMp: map[int]*Task{},
	}
	tm.taskQueue = make(PriorityQueue, len(tasks))
	for i, task := range tasks {
		userId, taskId, priority := task[0], task[1], task[2]
		task := &Task{
			priority: priority,
			UserId:   userId,
			TaskId:   taskId,
			index:    i,
		}
		tm.taskIdMp[taskId] = task

		tm.taskQueue[i] = task
	}

	heap.Init(&tm.taskQueue)
	return tm
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	if _, ok := this.taskIdMp[taskId]; !ok {
		n := this.taskQueue.Len()
		task := &Task{
			priority: priority,
			index:    n,
			UserId:   userId,
			TaskId:   taskId,
		}
		this.taskIdMp[taskId] = task
		heap.Push(&this.taskQueue, task)
	}
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	if _, ok := this.taskIdMp[taskId]; ok {
		item := this.taskIdMp[taskId]
		this.taskQueue.update(item, newPriority)
	}
}

func (this *TaskManager) Rmv(taskId int) {
	if _, ok := this.taskIdMp[taskId]; ok {
		item := this.taskIdMp[taskId]
		item.Removed = true
		delete(this.taskIdMp, taskId)
	}
}

func (this *TaskManager) ExecTop() int {
	if this.taskQueue.Len() == 0 {
		return -1
	}

	peek := this.taskQueue.Peek()
	for peek != nil && peek.Removed {
		heap.Pop(&this.taskQueue)
		peek = this.taskQueue.Peek()
	}

	if peek == nil {
		return -1
	}

	heap.Pop(&this.taskQueue)
	delete(this.taskIdMp, peek.TaskId)
	return peek.UserId
}

func main() {
	tm := Constructor([][]int{{1, 101, 8}, {2, 102, 20}, {3, 103, 5}})
	tm.Add(4, 104, 5)
	tm.Edit(102, 9)
	fmt.Println(tm.ExecTop())
	tm.Rmv(101)
	tm.Add(50, 101, 8)
	fmt.Println(tm.ExecTop())
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
