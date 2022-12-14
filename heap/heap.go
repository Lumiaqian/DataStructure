package heap

import (
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	elems []T
	isMin bool // ture 为小根堆，false为大根堆
}

// 获取堆的容量大小
func (h *Heap[T]) Len() int {
	return len(h.elems)
}

// 交换索引i，j位置的值
func (h *Heap[T]) Swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

// 比较i，j位置的值
func (h *Heap[T]) Less(i, j int) bool {
	if h.isMin {
		return h.elems[i] < h.elems[j]
	}
	return h.elems[i] > h.elems[j]
}

// 上浮
func (h *Heap[T]) up(i int) {
	for {
		f := (i - 1) / 2 //父节点的位置
		if i == f || h.Less(f, i) {
			break
		}
		h.Swap(f, i)
		i = f
	}
}

// 添加元素
func (h *Heap[T]) Push(value T) {
	h.elems = append(h.elems, value)
	h.up(len(h.elems) - 1)
}

// 下沉
func (h *Heap[T]) down(i0, n int) {
	i := i0
	for {
		child := i*2 + 1
		if child >= n || child < 0 {
			break
		}
		other := child + 1
		if other < n && h.Less(other, child) {
			child = other
		}

		if !h.Less(child, i) {
			break
		}
		h.Swap(i, child)
		i = child
	}

}

// Pop
func (h *Heap[T]) Pop() T {
	n := h.Len() - 1
	h.Swap(0, n)
	h.down(0, n)
	h.elems = h.elems[:h.Len()-1]
	return h.elems[h.Len()-1]
}

// 建堆
func (h *Heap[T]) Init() {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}
