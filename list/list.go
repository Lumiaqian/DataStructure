package list

import "fmt"

type List[T any] struct {
	Head *Node[T]
}

func (l *List[T]) InitList() *List[T] {
	l.Head = nil
	return l
}

// 判断链表是否为空
func (l *List[T]) IsEmpty() bool {
	if l.Head == nil {
		return true
	}
	return false
}

// 获取链表的长度
func (l *List[T]) Len() int {
	if l.IsEmpty() {
		return 0
	}
	var size int = 0
	cur := l.Head
	for cur != nil {
		cur = cur.Next
		size++
	}
	return size
}

// 获取指定位置的元素,下标从1开始
func (l *List[T]) GetElem(index int) T {
	var value T
	if index <= 0 || index > l.Len() {
		return value
	}
	pre := l.Head
	for j := 0; j < index-1; j++ {
		if pre != nil {
			pre = pre.Next
		}
	}
	value = pre.Value
	return value
}

// 插入元素，下标从1开始
func (l *List[T]) Insert(index int, value T) {
	if index <= 0 || index > l.Len()+1 {
		return
	}
	pre := l.Head
	node := &Node[T]{Value: value}
	if index == 1 {
		node.Next = l.Head
		l.Head = node
	} else {
		for j := 1; j < index-1; j++ {
			pre = pre.Next
		}
		next := pre.Next
		pre.Next = node
		node.Next = next
	}
}

// 在头部插入元素
func (l *List[T]) AddElem(value T) {
	node := &Node[T]{Value: value}
	if l.IsEmpty() {
		l.Head = node
		return
	}
	node.Next = l.Head
	l.Head = node
}

// 在尾部插入元素
func (l *List[T]) AppendElem(value T) {
	node := &Node[T]{Value: value}
	if l.IsEmpty() {
		l.Head = node
		return
	}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
}

// 按值查找,判断是否包含这个value
func (l *List[T]) LocateElem(fn func(value T) bool) bool {
	if l.IsEmpty() {
		return false
	}
	cur := l.Head
	if fn(cur.Value) {
		return true
	}
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
	cur := l.Head
	for cur.Next != nil {
		num++
		cur = cur.Next
		if fn(cur.Value) {
			return num
		}
	}
	return 0
}

// 删除指定索引位置上的元素
func (l *List[T]) DelElem(index int) {
	if index <= 0 || index > l.Len() {
		return
	}
	pre := l.Head
	if index == 1 {
		l.Head = pre.Next
		return
	}
	for j := 1; j < index-1; j++ {
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
}

// 打印链表
func (l *List[T]) ShowList() {
	if l.IsEmpty() {
		fmt.Println("")
		return
	}

	cur := l.Head
	fmt.Print(cur.Value)
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
	cur := l.Head
	res = append(res, cur.Value)
	for cur.Next != nil {
		cur = cur.Next
		res = append(res, cur.Value)
	}
	return res
}

// LPush,类似redis lpush命令，往头位置插入
func (l *List[T]) LPush(elems ...T) {
	for _, e := range elems {
		l.AddElem(e)
	}
}

// RPush,类似redis rpush命令,往尾部位置插入
func (l *List[T]) RPush(elems ...T) {
	for _, e := range elems {
		l.AppendElem(e)
	}
}
