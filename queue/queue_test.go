package queue

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPut(t *testing.T) {
	Convey("Test ArrayQueue Put", t, func() {
		var queue ArrayQueue[int]
		queue = *queue.New()
		queue.Put(1)
		queue.Put(2)
		queue.Put(3)
		Convey("The result of Put:1,2,3", func() {
			So(queue.elems, ShouldResemble, []int{1, 2, 3})
		})
	})
}

func TestPop(t *testing.T) {
	Convey("Test ArrayQueue Pop", t, func() {
		var queue ArrayQueue[int]
		queue = *queue.New()
		Convey("The result of Pop:2、3,false", func() {
			queue.Put(1)
			queue.Put(2)
			queue.Put(3)
			value, falg := queue.Pop()
			So(queue.elems, ShouldResemble, []int{2, 3})
			So(value, ShouldEqual, 1)
			So(falg, ShouldBeFalse)
		})
		Convey("The result of Pop:empty,true", func() {
			value, falg := queue.Pop()
			So(queue.elems, ShouldBeEmpty)
			So(value, ShouldEqual, 0)
			So(falg, ShouldBeTrue)
			queue.Put(1)
			queue.Put(2)
			queue.Put(3)
			queue.Pop()
			queue.Pop()
			value, falg = queue.Pop()
			So(queue.elems, ShouldBeEmpty)
			So(value, ShouldEqual, 3)
			So(falg, ShouldBeTrue)
		})
	})
}

func TestTop(t *testing.T) {
	Convey("Test ArrayQueue Top", t, func() {
		var queue ArrayQueue[int]
		queue = *queue.New()
		var stringqueue ArrayQueue[string]
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
	Convey("Test ArrayQueue Size", t, func() {
		var queue ArrayQueue[int]
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
