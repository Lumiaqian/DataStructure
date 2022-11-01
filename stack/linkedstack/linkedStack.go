package linkedstack

import "datastructure/list"

type LinkedStack[T any] struct {
	top  *list.Node[T] // 使用头插法，栈顶为头节点
	size int
}

// 初始化
func (s *LinkedStack[T]) Init() *LinkedStack[T] {
	s.top = new(list.Node[T])
	s.size = 0
	return s
}

// 初始化
func (s *LinkedStack[T]) New() *LinkedStack[T] {
	return new(LinkedStack[T]).Init()
}

// Push
func (s *LinkedStack[T]) Push(value T) {
	node := &list.Node[T]{Value: value}
	node.Next = s.top.Next
	s.top.Next = node
	s.size++
}

// Pop
func (s *LinkedStack[T]) Pop() (T, bool) {
	var value T
	if s.size == 0 {
		return value, true
	}
	value = s.top.Next.Value
	s.top.Next = s.top.Next.Next
	s.size--
	return value, s.size == 0
}

// Top
func (s *LinkedStack[T]) Top() (T, bool) {
	var value T
	if s.size == 0 {
		return value, true
	}
	value = s.top.Next.Value
	return value, s.size == 0
}

// Size
func (s *LinkedStack[T]) Size() int {
	return s.size
}

// list 转slice
func (s *LinkedStack[T]) ToSlice() []T {
	if s.size == 0 {
		return nil
	}
	res := make([]T, 0, s.size)
	cur := s.top
	for cur.Next != nil {
		cur = cur.Next
		res = append(res, cur.Value)
	}
	return res
}
