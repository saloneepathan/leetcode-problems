package main

import (
	"container/heap"
	"fmt"
)

// Task represents a single task with userId, taskId, and priority.
type Task struct {
	userId   int
	taskId   int
	priority int
	valid    bool // for lazy deletion
}

// MaxHeap implements a max-heap of tasks.
type MaxHeap []*Task

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].priority == h[j].priority {
		return h[i].taskId > h[j].taskId // tie-breaker: higher taskId first
	}
	return h[i].priority > h[j].priority
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*Task))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	task := old[n-1]
	*h = old[:n-1]
	return task
}

// TaskManager manages tasks across users.
type TaskManager struct {
	heap    MaxHeap
	taskMap map[int]*Task // taskId -> *Task
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		heap:    MaxHeap{},
		taskMap: make(map[int]*Task),
	}
	for _, t := range tasks {
		tm.add(t[0], t[1], t[2])
	}
	return tm
}

func (tm *TaskManager) add(userId, taskId, priority int) {
	task := &Task{userId: userId, taskId: taskId, priority: priority, valid: true}
	heap.Push(&tm.heap, task)
	tm.taskMap[taskId] = task
}

func (tm *TaskManager) edit(taskId, newPriority int) {
	if old, ok := tm.taskMap[taskId]; ok && old.valid {
		old.valid = false // invalidate old task
		newTask := &Task{userId: old.userId, taskId: taskId, priority: newPriority, valid: true}
		heap.Push(&tm.heap, newTask)
		tm.taskMap[taskId] = newTask
	}
}

func (tm *TaskManager) rmv(taskId int) {
	if task, ok := tm.taskMap[taskId]; ok && task.valid {
		task.valid = false
		delete(tm.taskMap, taskId)
	}
}

func (tm *TaskManager) execTop() int {
	for tm.heap.Len() > 0 {
		top := heap.Pop(&tm.heap).(*Task)
		if top.valid {
			top.valid = false
			delete(tm.taskMap, top.taskId)
			return top.userId
		}
	}
	return -1
}

// ------------------ Example Test ------------------
func main() {
	taskManager := Constructor([][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}})
	taskManager.add(4, 104, 5)
	taskManager.edit(102, 8)
	fmt.Println(taskManager.execTop()) // 3
	taskManager.rmv(101)
	taskManager.add(5, 105, 15)
	fmt.Println(taskManager.execTop()) // 5
}
