package linkedstack

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPush(t *testing.T) {
	Convey("Test LinkedStack Push", t, func() {
		var stack LinkedStack[int]
		stack = *stack.New()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		Convey("The result of Push:3,2,1", func() {
			So(stack.ToSlice(), ShouldResemble, []int{3, 2, 1})
		})
	})
}

func TestPop(t *testing.T) {
	Convey("Test LinkedStack Pop", t, func() {
		var stack LinkedStack[int]
		stack = *stack.New()
		Convey("The result of Pop:2、1,false", func() {
			stack.Push(1)
			stack.Push(2)
			stack.Push(3)
			value, falg := stack.Pop()
			So(stack.ToSlice(), ShouldResemble, []int{2, 1})
			So(value, ShouldEqual, 3)
			So(falg, ShouldBeFalse)
		})
		Convey("The result of Pop:empty,true", func() {
			value, falg := stack.Pop()
			So(stack.ToSlice(), ShouldBeEmpty)
			So(value, ShouldEqual, 0)
			So(falg, ShouldBeTrue)
			stack.Push(1)
			stack.Push(2)
			stack.Push(3)
			stack.Pop()
			stack.Pop()
			value, falg = stack.Pop()
			So(stack.ToSlice(), ShouldBeEmpty)
			So(value, ShouldEqual, 1)
			So(falg, ShouldBeTrue)
		})
	})
}

func TestTop(t *testing.T) {
	Convey("Test LinkedStack Top", t, func() {
		var stack LinkedStack[int]
		stack = *stack.New()
		var stringStack LinkedStack[string]
		stringStack = *stringStack.New()
		Convey("The result of Top:3,false", func() {
			stack.Push(1)
			stack.Push(2)
			stack.Push(3)
			value, falg := stack.Top()
			So(value, ShouldEqual, 3)
			So(falg, ShouldBeFalse)
			stringStack.Push("a")
			stringStack.Push("b")
			stringStack.Push("c")
			value2, falg2 := stringStack.Top()
			So(value2, ShouldEqual, "c")
			So(falg2, ShouldBeFalse)
		})
		Convey("The result of Top:empty,true", func() {
			value, falg := stack.Top()
			So(value, ShouldEqual, 0)
			So(falg, ShouldBeTrue)
			value2, falg2 := stringStack.Top()
			So(value2, ShouldEqual, "")
			So(falg2, ShouldBeTrue)
		})
	})
}

func TestSize(t *testing.T) {
	Convey("Test LinkedStack Size", t, func() {
		var stack LinkedStack[int]
		stack = *stack.New()
		Convey("The reslut of Size:0", func() {
			So(stack.Size(), ShouldEqual, 0)
		})
		Convey("The reslut of Size:1", func() {
			stack.Push(11)
			So(stack.Size(), ShouldEqual, 1)
		})
	})
}
