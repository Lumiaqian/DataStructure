package linkedqueue

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPut(t *testing.T) {
	Convey("Test LinkedQueue Put", t, func() {
		var queue LinkedQueue[int]
		queue = *queue.New()
		queue.Put(1)
		queue.Put(2)
		queue.Put(3)
		Convey("The result of Put:1,2,3", func() {
			So(queue.ToSlice(), ShouldResemble, []int{1, 2, 3})
		})
	})
}

func TestPop(t *testing.T) {
	Convey("Test LinkedQueue Pop", t, func() {
		var queue LinkedQueue[int]
		queue = *queue.New()
		Convey("The result of Pop:2、3,false", func() {
			queue.Put(1)
			queue.Put(2)
			queue.Put(3)
			value, falg := queue.Pop()
			So(queue.ToSlice(), ShouldResemble, []int{2, 3})
			So(value, ShouldEqual, 1)
			So(falg, ShouldBeFalse)
		})
		Convey("The result of Pop:empty,true", func() {
			value, falg := queue.Pop()
			So(queue.ToSlice(), ShouldBeEmpty)
			So(value, ShouldEqual, 0)
			So(falg, ShouldBeTrue)
			queue.Put(1)
			queue.Put(2)
			queue.Put(3)
			queue.Pop()
			queue.Pop()
			value, falg = queue.Pop()
			So(queue.ToSlice(), ShouldBeEmpty)
			So(value, ShouldEqual, 3)
			So(falg, ShouldBeTrue)
		})
	})
}

func TestTop(t *testing.T) {
	Convey("Test LinkedQueue Top", t, func() {
		var queue LinkedQueue[int]
		queue = *queue.New()
		var stringqueue LinkedQueue[string]
		stringqueue = *stringqueue.New()
		Convey("The result of Top:1,false", func() {
			queue.Put(1)
			queue.Put(2)
			queue.Put(3)
			value, falg := queue.Top()
			So(value, ShouldEqual, 1)
			So(falg, ShouldBeFalse)
			stringqueue.Put("a")
			stringqueue.Put("b")
			stringqueue.Put("c")
			value2, falg2 := stringqueue.Top()
			So(value2, ShouldEqual, "a")
			So(falg2, ShouldBeFalse)
		})
		Convey("The result of Top:empty,true", func() {
			value, falg := queue.Top()
			So(value, ShouldEqual, 0)
			So(falg, ShouldBeTrue)
			value2, falg2 := stringqueue.Top()
			So(value2, ShouldEqual, "")
			So(falg2, ShouldBeTrue)
		})
	})
}

func TestSize(t *testing.T) {
	Convey("Test LinkedQueue Size", t, func() {
		var queue LinkedQueue[int]
		queue = *queue.New()
		Convey("The reslut of Size:0", func() {
			So(queue.Size(), ShouldEqual, 0)
		})
		Convey("The reslut of Size:1", func() {
			queue.Put(11)
			So(queue.Size(), ShouldEqual, 1)
		})
	})
}
