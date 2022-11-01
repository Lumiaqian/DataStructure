package list

import "fmt"

type List[T any] struct {
	head *Node[T]
	size int
}

func (l *List[T]) New() *List[T] {
	return new(List[T]).InitList()
}

func (l *List[T]) InitList() *List[T] {
	l.head = new(Node[T])
	l.size = 0
	return l
}

// 判断链表是否为空
func (l *List[T]) IsEmpty() bool {
	if l.size == 0 {
		return true
	}
	return false
}

// 获取链表的长度
func (l *List[T]) Len() int {
	return l.size
}

// 获取指定位置的元素,下标从1开始
func (l *List[T]) GetElem(index int) *Node[T] {
	if index <= 0 || index > l.Len() {
		return nil
	}
	cur := l.head
	for j := 0; j < index; j++ {
		cur = cur.Next
	}
	return cur
}

// 插入元素，下标从1开始
func (l *List[T]) Insert(index int, value T) {
	if index <= 0 || index > l.Len()+1 {
		return
	}
	cur := l.head
	node := &Node[T]{Value: value}
	if index == 1 {
		node.Next = cur.Next
		cur.Next = node
	} else {
		for j := 1; j < index; j++ {
			cur = cur.Next
		}
		next := cur.Next
		cur.Next = node
		node.Next = next
	}
	l.size++
}

// 在头部插入元素
func (l *List[T]) PushFront(value T) {
	node := &Node[T]{Value: value}
	if l.IsEmpty() {
		l.head.Next = node
		l.size++
		return
	}
	node.Next = l.head.Next
	l.head.Next = node
	l.size++
}

// 在尾部插入元素
func (l *List[T]) PushBack(value T) {
	node := &Node[T]{Value: value}
	if l.IsEmpty() {
		l.head.Next = node
		l.size++
		return
	}
	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
	l.size++
}

// 按值查找,判断是否包含这个value
func (l *List[T]) LocateElem(fn func(value T) bool) bool {
	if l.IsEmpty() {
		return false
	}
	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
		if fn(cur.Value) {
			return true
		}
	}
	return false
}

// 按值查找并返回第一个下标，不存在返回0
func (l *List[T]) FindElem(fn func(value T) bool) int {
	var num int = 1
	if l.IsEmpty() {
		return 0
	}
	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
		if fn(cur.Value) {
			return num
		}
		num++
	}
	return 0
}

// 删除指定索引位置上的元素
func (l *List[T]) DelElem(index int) {
	if index <= 0 || index > l.Len() {
		return
	}
	cur := l.head
	if index == 1 {
		l.head.Next = cur.Next.Next
		l.size--
		return
	}
	for j := 1; j < index; j++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	l.size--
}

// 打印链表
func (l *List[T]) ShowList() {
	if l.IsEmpty() {
		fmt.Println("")
		return
	}

	cur := l.head
	//fmt.Print(cur.Value)
	for i := 0; i < l.Len(); i++ {
		cur = cur.Next
		if cur != nil {
			fmt.Print(" -> ")
			fmt.Print(cur.Value)
		}
	}
	fmt.Println("")
}

// list 转slice
func (l *List[T]) ToSlice() []T {
	if l.IsEmpty() {
		return nil
	}
	res := make([]T, 0, l.Len())
	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
		res = append(res, cur.Value)
	}
	return res
}

// LPush,类似redis lpush命令，往头位置插入
func (l *List[T]) LPush(elems ...T) {
	for _, e := range elems {
		l.PushFront(e)
	}
}

// RPush,类似redis rpush命令,往尾部位置插入
func (l *List[T]) RPush(elems ...T) {
	for _, e := range elems {
		l.PushBack(e)
	}
}
