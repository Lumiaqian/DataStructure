package queue

const defaultVolume = 10

type ArrayQueue[T any] struct {
	elems []T
}

// 初始化队列
func (s *ArrayQueue[T]) Init() *ArrayQueue[T] {
	s.elems = make([]T, 0, defaultVolume)
	return s
}

// 初始化队列
func (s *ArrayQueue[T]) New() *ArrayQueue[T] {
	return new(ArrayQueue[T]).Init()
}

// 将数据放入队列尾部
func (s *ArrayQueue[T]) Put(e T) {
	s.elems = append(s.elems, e)
}

// 取出对头元素并删除，第二返回值表示当前队是否为空
func (s *ArrayQueue[T]) Pop() (T, bool) {
	var value T
	if len(s.elems) == 0 {
		return value, true
	}
	value = s.elems[0]
	s.elems = s.elems[1:]
	return value, len(s.elems) == 0
}

// 取出队头元素，第二返回值表示当前队是否为空
func (s *ArrayQueue[T]) Top() (T, bool) {
	var value T
	if len(s.elems) == 0 {
		return value, true
	}
	value = s.elems[0]
	return value, len(s.elems) == 0
}

// 获取队大小
func (s *ArrayQueue[T]) Size() int {
	return len(s.elems)
}
