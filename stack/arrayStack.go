package stack

const defaultVolume = 10

type ArrayStack[T any] struct {
	elems []T
}

// 初始化栈
func (s *ArrayStack[T]) Init() *ArrayStack[T] {
	s.elems = make([]T, 0, defaultVolume)
	return s
}

// 初始化栈
func (s *ArrayStack[T]) New() *ArrayStack[T] {
	return new(ArrayStack[T]).Init()
}

// 将元素放入栈顶
func (s *ArrayStack[T]) Push(e T) {
	s.elems = append(s.elems, e)
}

// 取出栈顶元素并删除，第二返回值表示当前栈是否为空
func (s *ArrayStack[T]) Pop() (T, bool) {
	var value T
	if len(s.elems) == 0 {
		return value, true
	}
	value = s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return value, len(s.elems) == 0
}

// 取出栈顶元素，第二返回值表示当前栈是否为空
func (s *ArrayStack[T]) Top() (T, bool) {
	var value T
	if len(s.elems) == 0 {
		return value, true
	}
	value = s.elems[len(s.elems)-1]
	return value, len(s.elems) == 0
}

// 获取栈大小
func (s *ArrayStack[T]) Size() int {
	return len(s.elems)
}
