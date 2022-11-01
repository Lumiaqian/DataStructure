package linkedqueue

import "datastructure/list"

type LinkedQueue[T any] struct {
	// 使用双向循环链表，头节点为队头，尾节点为队尾
	top  *list.DoublyNode[T]
	tail *list.DoublyNode[T]
	size int
}

// 初始化
func (s *LinkedQueue[T]) Init() *LinkedQueue[T] {
	s.top = new(list.DoublyNode[T])
	s.tail = new(list.DoublyNode[T])
	s.top.Next = s.tail
	s.tail.Prev = s.top
	s.top.Prev = s.tail
	s.tail.Next = s.top
	s.size = 0
	return s
}

// 初始化
func (q *LinkedQueue[T]) New() *LinkedQueue[T] {
	return new(LinkedQueue[T]).Init()
}

// Push
func (q *LinkedQueue[T]) Put(value T) {
	node := &list.DoublyNode[T]{Value: value}
	pre := q.tail.Prev
	node.Next = q.tail
	q.tail.Prev = node
	node.Prev = pre
	pre.Next = node
	q.size++
}

// Pop
func (q *LinkedQueue[T]) Pop() (T, bool) {
	var value T
	if q.size == 0 {
		return value, true
	}
	value = q.top.Next.Value
	next := q.top.Next
	q.top.Next = next.Next
	next.Next.Prev = q.top
	q.size--
	return value, q.size == 0
}

// Top
func (q *LinkedQueue[T]) Top() (T, bool) {
	var value T
	if q.size == 0 {
		return value, true
	}
	value = q.top.Next.Value
	return value, q.size == 0
}

// Size
func (q *LinkedQueue[T]) Size() int {
	return q.size
}

// list 转slice
func (q *LinkedQueue[T]) ToSlice() []T {
	if q.size == 0 {
		return nil
	}
	res := make([]T, 0, q.size)
	cur := q.top
	for cur.Next != q.tail {
		cur = cur.Next
		res = append(res, cur.Value)
	}
	return res
}
