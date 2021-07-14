package queue

import "errors"

type QueueIntf interface {
	Push(int)
	Pop() (int, error)
	Size() int
}

type QueueImpl struct {
	dataStore []int
}

func (q *QueueImpl) Push(a int) {
	q.dataStore = append(q.dataStore, a)
}

func (q *QueueImpl) Pop() (int, error) {
	if len(q.dataStore) == 0 {
		return -1, errors.New("Queue Empty")
	}

	first := q.dataStore[0]
	q.dataStore = q.dataStore[1:]
	return first, nil
}

func (q *QueueImpl) Size() int {
	return len(q.dataStore)
}

func New() *QueueImpl {
	return &QueueImpl{make([]int, 0, 10)}
}
