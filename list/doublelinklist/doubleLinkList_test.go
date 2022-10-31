package doublelinklist

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLPush(t *testing.T) {
	Convey("Test LPush", t, func() {
		var listA DoubleLinkList[int]
		var listB DoubleLinkList[string]
		listA = *listA.New()
		listB = *listB.New()
		listA.LPush(1, 2, 3, 4, 5)
		listB.LPush("a", "b", "c", "d", "e")
		Convey("The result of LPush:5,4,3,2,1", func() {
			So(listA.ToSlice(), ShouldResemble, []int{5, 4, 3, 2, 1})
		})
		Convey("The result of LPush:e,d,c,b,a", func() {
			So(listB.ToSlice(), ShouldResemble, []string{"e", "d", "c", "b", "a"})
		})
	})
}

func TestRPush(t *testing.T) {
	Convey("Test RPush", t, func() {
		var listA DoubleLinkList[int]
		var listB DoubleLinkList[string]
		listA = *listA.New()
		listB = *listB.New()
		listA.RPush(1, 2, 3, 4, 5)
		listB.RPush("a", "b", "c", "d", "e")
		Convey("The result of RPush:1,2,3,4,5", func() {
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
		})
		Convey("The result of RPush:a,b,c,d,e", func() {
			So(listB.ToSlice(), ShouldResemble, []string{"a", "b", "c", "d", "e"})
		})
	})
}

func TestIsEmpty(t *testing.T) {
	Convey("Test IsEmpty", t, func() {
		var listA DoubleLinkList[int]
		var listB DoubleLinkList[string]
		listA = *listA.New()
		listB = *listB.New()
		listA.LPush(1, 2, 3, 4, 5)
		Convey("The result of IsEmpty:false", func() {
			So(listA.IsEmpty(), ShouldBeFalse)
		})
		Convey("The result of IsEmpty:true", func() {
			So(listB.IsEmpty(), ShouldBeTrue)
		})
	})
}

func TestLen(t *testing.T) {
	Convey("Test Len", t, func() {
		var listA DoubleLinkList[int]
		var listB DoubleLinkList[string]
		listA = *listA.New()
		listB = *listB.New()
		listA.LPush(1, 2, 3, 4, 5)
		Convey("The result of Len:5", func() {
			So(listA.Len(), ShouldEqual, 5)
		})
		Convey("The result of Len:0", func() {
			So(listB.Len(), ShouldEqual, 0)
		})
	})
}

func TestFront(t *testing.T) {
	Convey("Test Front", t, func() {
		var listA DoubleLinkList[int]
		listA = *listA.New()
		Convey("The result of Front:nil", func() {
			So(listA.Front(), ShouldBeNil)
		})
		Convey("The result of Front:10", func() {
			listA.PushFront(10)
			So(listA.Front().Value, ShouldEqual, 10)
		})
	})
}
func TestBack(t *testing.T) {
	Convey("Test Back", t, func() {
		var listA DoubleLinkList[int]
		listA = *listA.New()
		Convey("The result of Back:nil", func() {
			So(listA.Back(), ShouldBeNil)
		})
		Convey("The result of Back:10", func() {
			listA.LPush(1, 2, 3, 4, 5)
			listA.PushBack(10)
			So(listA.Back().Value, ShouldEqual, 10)
		})
	})
}

func TestGet(t *testing.T) {
	Convey("Test Get", t, func() {
		var listA DoubleLinkList[int]
		listA = *listA.New()
		Convey("The result of Get:nil", func() {
			So(listA.Get(0), ShouldBeNil)
			So(listA.Get(1), ShouldBeNil)
		})
		Convey("The result of Get:10,30,50", func() {
			listA.RPush(10, 20, 30, 40, 50)
			So(listA.Get(1).Value, ShouldEqual, 10)
			So(listA.Get(3).Value, ShouldEqual, 30)
			So(listA.Get(5).Value, ShouldEqual, 50)
		})
	})
}

func TestInsert(t *testing.T) {
	Convey("Test Insert", t, func() {
		var listA DoubleLinkList[int]
		listA = *listA.New()
		listA.RPush(1, 2, 3, 4, 5)
		Convey("The result of Insert:1,2,3,4,5,6", func() {
			listA.Insert(6, 6)
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5, 6})
		})
		Convey("The result of Insert:6,1,2,3,4,5", func() {
			listA.Insert(1, 6)
			So(listA.ToSlice(), ShouldResemble, []int{6, 1, 2, 3, 4, 5})
		})
		Convey("The result of Insert:1,2,6,3,4,5", func() {
			listA.Insert(3, 6)
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 6, 3, 4, 5})
		})
		Convey("The result of Insert:1,2,3,4,5", func() {
			listA.Insert(0, 6)
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
			listA.Insert(7, 6)
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
		})
	})
}

func TestDel(t *testing.T) {
	Convey("Test Del", t, func() {
		var listA DoubleLinkList[int]
		listA = *listA.New()
		listA.RPush(1, 2, 3, 4, 5)
		Convey("The result of Del:1,2,3,4,5", func() {
			listA.Del(0)
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
			So(listA.size, ShouldEqual, 5)
			listA.Del(6)
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
			So(listA.size, ShouldEqual, 5)
		})
		Convey("The result of Del:2,3,4,5", func() {
			listA.Del(1)
			So(listA.ToSlice(), ShouldResemble, []int{2, 3, 4, 5})
			So(listA.size, ShouldEqual, 4)
		})
		Convey("The result of Del:1,2,4,5", func() {
			listA.Del(3)
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 4, 5})
			So(listA.size, ShouldEqual, 4)
		})
		Convey("The result of Del:1,2,3,4", func() {
			listA.Del(5)
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4})
			So(listA.size, ShouldEqual, 4)
		})
	})
}

func TestRemove(t *testing.T) {
	Convey("Test Remove", t, func() {
		var listA DoubleLinkList[int]
		listA = *listA.New()
		Convey("The result of Remove:empty", func() {
			listA.Remove(func(value int) bool { return 1 == value })
			So(listA.ToSlice(), ShouldBeEmpty)
			So(listA.size, ShouldEqual, 0)
		})
		Convey("The result of Remove:1,2,3,4,5", func() {
			listA.RPush(1, 2, 3, 4, 5)
			listA.Remove(func(value int) bool { return 0 == value })
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
			So(listA.size, ShouldEqual, 5)
			listA.Remove(func(value int) bool { return 6 == value })
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
			So(listA.size, ShouldEqual, 5)
		})
		Convey("The result of Remove:2,3,4,5", func() {
			listA.RPush(1, 2, 3, 4, 5)
			listA.Remove(func(value int) bool { return 1 == value })
			So(listA.ToSlice(), ShouldResemble, []int{2, 3, 4, 5})
			So(listA.size, ShouldEqual, 4)
		})
		Convey("The result of Remove:1,2,4,5", func() {
			listA.RPush(1, 2, 3, 4, 5)
			listA.Remove(func(value int) bool { return 3 == value })
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 4, 5})
			So(listA.size, ShouldEqual, 4)
		})
		Convey("The result of Remove:1,2,3,4", func() {
			listA.RPush(1, 2, 3, 4, 5)
			listA.Remove(func(value int) bool { return 5 == value })
			So(listA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4})
			So(listA.size, ShouldEqual, 4)
		})
	})
}

func TestLocate(t *testing.T) {
	Convey("Test Locate", t, func() {
		var listA DoubleLinkList[int]
		listA = *listA.New()
		Convey("The result of Locate:false", func() {
			So(listA.Locate(func(value int) bool { return 1 == value }), ShouldBeFalse)
			listA.RPush(1, 2, 3, 4, 5)
			So(listA.Locate(func(value int) bool { return 0 == value }), ShouldBeFalse)
		})
		Convey("The result of Locate:true", func() {
			listA.RPush(10, 20, 30, 40, 50)
			So(listA.Locate(func(value int) bool { return 30 == value }), ShouldBeTrue)
		})
	})
}
