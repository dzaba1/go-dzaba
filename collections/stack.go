package collections

import (
	"dzaba/go-dzaba/utils"
	"errors"
	"sync"
)

type Stack[T any] struct {
	lock sync.Mutex
	list []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		lock: sync.Mutex{},
		list: []T{},
	}
}

func (s *Stack[T]) Push(v T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.list = append(s.list, v)
}

func (s *Stack[T]) Pop() (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	length := len(s.list)
	if length == 0 {
		return utils.DefaultGeneric[T](), errors.New("Empty Stack.")
	}

	res := s.list[length-1]
	s.list = s.list[:length-1]
	return res, nil
}

func (s *Stack[T]) Peek() (T, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	length := len(s.list)
	if length == 0 {
		return utils.DefaultGeneric[T](), errors.New("Empty Stack.")
	}

	res := s.list[length-1]
	return res, nil
}

func (s *Stack[T]) Count() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.list)
}

func (s *Stack[T]) GetList() []T {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.list
}
