package stack

import "errors"

type StackIntf interface {
	Push(int)
	Pop() (int, error)
	Size() int
}

type StackImpl struct {
	dataStore []int
}

func New() StackIntf {
	return &StackImpl{make([]int, 0)}
}

func (s *StackImpl) Push(a int) {
	s.dataStore = append(s.dataStore, a)
}

func (s *StackImpl) Pop() (int, error) {
	size := len(s.dataStore)
	if size == 0 {
		return -1, errors.New("Stack Empty")
	}
	last := s.dataStore[size-1]
	s.dataStore = s.dataStore[:size-1]
	return last, nil

}

func (s *StackImpl) Size() int {
	return len(s.dataStore)
}
