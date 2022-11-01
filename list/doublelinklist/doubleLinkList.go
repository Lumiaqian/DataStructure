package doublelinklist

import "datastructure/list"

type DoubleLinkList[T any] struct {
	head *list.DoublyNode[T]
	tail *list.DoublyNode[T]
	size int
}

// 初始化双向链表
func (l *DoubleLinkList[T]) Init() *DoubleLinkList[T] {
	l.head = new(list.DoublyNode[T])
	l.tail = new(list.DoublyNode[T])
	l.head.Next = l.tail
	l.head.Prev = nil
	l.tail.Prev = l.head
	l.tail.Next = nil
	l.size = 0
	return l
}

// 新建双向链表
func (l *DoubleLinkList[T]) New() *DoubleLinkList[T] {
	return new(DoubleLinkList[T]).Init()
}

// 获取链表长度
func (l *DoubleLinkList[T]) Len() int {
	return l.size
}

// 判断链表是否为空
func (l *DoubleLinkList[T]) IsEmpty() bool {
	if l.size == 0 {
		return true
	}
	return false
}

// 获取链表头元素
func (l *DoubleLinkList[T]) Front() *list.DoublyNode[T] {
	if l.size == 0 {
		return nil
	}
	return l.head.Next
}

// 获取链表尾元素
func (l *DoubleLinkList[T]) Back() *list.DoublyNode[T] {
	if l.size == 0 {
		return nil
	}
	return l.tail.Prev
}

// 获取指定索引上的元素
func (l *DoubleLinkList[T]) Get(index int) *list.DoublyNode[T] {
	if index <= 0 || index > l.size {
		return nil
	}
	if index == 1 {
		return l.head.Next
	}
	if index == l.size {
		return l.tail.Prev
	}
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur
}

// 指定向指定位置插入元素
func (l *DoubleLinkList[T]) Insert(index int, value T) {
	if index <= 0 || index > l.size+1 {
		return
	}
	node := &list.DoublyNode[T]{Value: value}
	if index == 1 {
		node.Next = l.head.Next
		l.head.Next.Prev = node
		l.head.Next = node
		node.Prev = l.head
		l.size++
		return
	}
	if index == l.size+1 {
		l.tail.Prev.Next = node
		node.Prev = l.tail.Prev
		node.Next = l.tail
		l.tail.Prev = node
		l.size++
		return
	}
	cur := l.head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	next := cur.Next
	node.Next = next
	next.Prev = node
	cur.Next = node
	node.Prev = cur
	l.size++
}

// 从头部插入元素
func (l *DoubleLinkList[T]) PushFront(value T) {
	node := &list.DoublyNode[T]{Value: value}
	node.Next = l.head.Next
	node.Prev = l.head
	l.head.Next.Prev = node
	l.head.Next = node
	l.size++
}

// 从尾部插入元素
func (l *DoubleLinkList[T]) PushBack(value T) {
	node := &list.DoublyNode[T]{Value: value}
	l.tail.Prev.Next = node
	node.Prev = l.tail.Prev
	node.Next = l.tail
	l.tail.Prev = node
	l.size++
}

// LPush,类似redis lpush命令，往头位置插入
func (l *DoubleLinkList[T]) LPush(values ...T) {
	for _, v := range values {
		l.PushFront(v)
	}
}

// RPush,类似redis rpush命令,往尾部位置插入
func (l *DoubleLinkList[T]) RPush(values ...T) {
	for _, v := range values {
		l.PushBack(v)
	}
}

// 删除指定索引上的元素
func (l *DoubleLinkList[T]) Del(index int) {
	if index <= 0 || index > l.size {
		return
	}
	cur := l.head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	cur.Next.Next.Prev = cur
	cur.Next = cur.Next.Next
	l.size--
}

// 删除指定元素
func (l *DoubleLinkList[T]) Remove(fn func(value T) bool) {
	if l.size == 0 {
		return
	}
	cur := l.head
	for cur.Next != l.tail {
		cur = cur.Next
		if fn(cur.Value) {
			pre := cur.Prev
			cur.Next.Prev = pre
			pre.Next = cur.Next
			l.size--
		}
	}
}

// 按值查找，判断是否包含这个value
func (l *DoubleLinkList[T]) Locate(fn func(value T) bool) bool {
	if l.size == 0 {
		return false
	}
	cur := l.head
	for cur.Next != l.tail {
		cur = cur.Next
		if fn(cur.Value) {
			return true
		}
	}
	return false
}

// list转slice
func (l *DoubleLinkList[T]) ToSlice() []T {
	if l.size == 0 {
		return nil
	}
	res := make([]T, 0, l.size)
	cur := l.head
	for cur.Next != l.tail {
		cur = cur.Next

		res = append(res, cur.Value)
	}
	return res
}
