package heap

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInit(t *testing.T) {
	Convey("Test heap Init", t, func() {
		heap := new(Heap[int])
		Convey("The result of minHeap Init is 1, 2, 7, 3, 5, 9, 10, 6", func() {
			heap.isMin = true
			heap.elems = []int{9, 3, 7, 6, 5, 1, 10, 2}
			heap.Init()
			So(heap.elems, ShouldResemble, []int{1, 2, 7, 3, 5, 9, 10, 6})
		})
		Convey("The result of maxHeap Init is 10, 6, 9, 3, 5, 1, 7, 2", func() {
			heap.isMin = false
			heap.elems = []int{9, 3, 7, 6, 5, 1, 10, 2}
			heap.Init()
			So(heap.elems, ShouldResemble, []int{10, 6, 9, 3, 5, 1, 7, 2})
		})
	})
}

func TestPush(t *testing.T) {
	Convey("Test heap Push", t, func() {
		heap := new(Heap[int])
		heap.isMin = true
		heap.elems = []int{9, 3, 7, 6, 5, 1, 10, 2}
		heap.Init()
		heap.Push(0)
		Convey("The result of Push is 0,1,7,2,5,9,10,6,3", func() {
			So(heap.elems, ShouldResemble, []int{0, 1, 7, 2, 5, 9, 10, 6, 3})
		})
	})
}

func TestPop(t *testing.T) {
	Convey("Test heap Pop", t, func() {
		heap := new(Heap[int])
		heap.isMin = true
		heap.elems = []int{9, 3, 7, 6, 5, 1, 10, 2}
		heap.Init()
		heap.Push(0)
		heap.Pop()
		Convey("The result of Pop is 1, 2, 7, 3, 5, 9, 10, 6", func() {
			So(heap.elems, ShouldResemble, []int{1, 2, 7, 3, 5, 9, 10, 6})
		})
	})
}
