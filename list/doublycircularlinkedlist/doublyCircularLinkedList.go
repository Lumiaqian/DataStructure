package doublycircularlinkedlist

type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// 初始化Doubly Circular Linked List detail
func (l *List[T]) Init() *List[T] {
	l.head = new(Node[T])
	l.tail = new(Node[T])
	l.head.Next = l.tail
	l.tail.Prev = l.head
	l.head.Prev = l.tail
	l.tail.Next = l.head
	l.size = 0
	return l
}

// 初始化Doubly Circular Linked List
func (l *List[T]) New() *List[T] {
	return new(List[T]).Init()
}

// 获取链表长度
func (l *List[T]) Len() int {
	return l.size
}

// 判断链表是否为空
func (l *List[T]) IsEmpty() bool {
	if l.size == 0 {
		return true
	}
	return false
}

// 获取头节点
func (l *List[T]) Front() *Node[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.head.Next
}

// 获取尾节点
func (l *List[T]) Back() *Node[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.tail.Prev
}

// 获取指定索引上的元素
func (l *List[T]) Get(index int) *Node[T] {
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

// 插入指定位置
func (l *List[T]) Insert(index int, value T) {
	if index <= 0 || index > l.size+1 {
		return
	}
	node := &Node[T]{Value: value}
	if index == 1 {
		node.Next = l.head.Next
		node.Prev = l.head
		l.head.Next.Prev = node
		l.head.Next = node
		l.size++
		return
	}
	if index == l.size+1 {
		pre := l.tail.Prev
		pre.Next = node
		node.Prev = pre
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
	next.Prev = node
	node.Next = next
	cur.Next = node
	node.Prev = cur
	l.size++
}

// 从头部插入元素
func (l *List[T]) PushFront(value T) {
	node := &Node[T]{Value: value}
	node.Next = l.head.Next
	node.Prev = l.head
	l.head.Next.Prev = node
	l.head.Next = node
	l.size++
}

// 从尾部插入元素
func (l *List[T]) PushBack(value T) {
	node := &Node[T]{Value: value}
	l.tail.Prev.Next = node
	node.Prev = l.tail.Prev
	node.Next = l.tail
	l.tail.Prev = node
	l.size++
}

// LPush,类似redis lpush命令，往头位置插入
func (l *List[T]) LPush(values ...T) {
	for _, v := range values {
		l.PushFront(v)
	}
}

// RPush,类似redis rpush命令,往尾部位置插入
func (l *List[T]) RPush(values ...T) {
	for _, v := range values {
		l.PushBack(v)
	}
}

// 删除指定索引上的元素
func (l *List[T]) Del(index int) {
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
func (l *List[T]) Remove(fn func(value T) bool) {
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
func (l *List[T]) Locate(fn func(value T) bool) bool {
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
func (l *List[T]) ToSlice() []T {
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
