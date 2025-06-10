package stack_queue

import (
	"container/list"
)

type Queue struct {
	l *list.List
}

func NewQueue() *Queue {
	return &Queue{l: list.New()}
}

func (queue *Queue) Init() {
	queue.l = list.New()
}

func (queue *Queue) Push(v interface{}) {
	queue.l.PushFront(v)
}

func (queue *Queue) Pop() interface{} {
	if queue.l.Len() == 0 {
		return nil
	}
	v := queue.l.Back()
	return queue.l.Remove(v)
}

func (queue *Queue) IsEmpty() bool {
	return queue.l.Len() == 0
}

func (queue *Queue) Len() int {
	return queue.l.Len()
}
