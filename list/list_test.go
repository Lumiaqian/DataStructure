package list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLPush(t *testing.T) {
	Convey("Test LPush", t, func() {
		var linkListA, linkListB List[int]
		linkListA = *linkListA.InitList()
		linkListB = *linkListB.InitList()
		linkListA.LPush(1, 2, 3, 4, 5)
		linkListB.LPush(5, 4, 3, 2, 1)
		Convey("The result of LPush:5,4,3,2,1", func() {
			So(linkListA.ToSlice(), ShouldResemble, []int{5, 4, 3, 2, 1})
		})
		Convey("The result of LPush:1,2,3,4,5", func() {
			So(linkListB.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
		})
	})
}

func TestRPush(t *testing.T) {
	Convey("Test RPush", t, func() {
		var linkListA, linkListB List[int]
		linkListA = *linkListA.InitList()
		linkListB = *linkListB.InitList()
		linkListA.RPush(1, 2, 3, 4, 5)
		linkListB.RPush(5, 4, 3, 2, 1)
		Convey("The result of RPush:1,2,3,4,5", func() {
			So(linkListA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
		})
		Convey("The result of RPush:5,4,3,2,1", func() {
			So(linkListB.ToSlice(), ShouldResemble, []int{5, 4, 3, 2, 1})
		})
	})
}

func TestIsEmpty(t *testing.T) {
	Convey("Test IsEmpty ", t, func() {
		var linkListA, linkListB List[int]
		linkListA = *linkListA.InitList()
		linkListB = *linkListB.InitList()
		linkListB.LPush(5, 4, 3, 2, 1)

		Convey("The result of IsEmpty should be true", func() {
			So(linkListA.IsEmpty(), ShouldEqual, true)
		})

		Convey("The result of IsEmpty should be false", func() {
			So(linkListB.IsEmpty(), ShouldEqual, false)
		})

	})
}

func TestLen(t *testing.T) {
	Convey("Test Len ", t, func() {
		var linkListA, linkListB List[int]
		linkListA = *linkListA.InitList()
		linkListB = *linkListB.InitList()
		linkListA.RPush(1, 2)
		Convey("The reslut of Len should be 2", func() {
			So(linkListA.Len(), ShouldEqual, 2)
		})
		Convey("The reslut of Len should be 0", func() {
			So(linkListB.Len(), ShouldEqual, 0)
		})
	})
}

func TestGetElem(t *testing.T) {
	Convey("Test GetElem ", t, func() {
		var linkListA List[int]
		var linkListB List[string]
		linkListA = *linkListA.InitList()
		linkListB = *linkListB.InitList()
		linkListA.AddElem(5)
		Convey("The reslut of GetElem should be 5", func() {
			So(linkListA.GetElem(1), ShouldEqual, 5)
		})
		Convey("The reslut of GetElem should be 0", func() {
			So(linkListA.GetElem(0), ShouldEqual, 0)
		})
		Convey("The reslut of GetElem should be empty string", func() {
			So(linkListB.GetElem(0), ShouldBeEmpty)
		})
	})
}

func TestInsert(t *testing.T) {
	Convey("Test Insert", t, func() {
		var linkListA, linkListB List[int]
		linkListA = *linkListA.InitList()
		linkListB = *linkListB.InitList()
		linkListA.RPush(1, 2, 3, 4)
		Convey("The result of Insert:1,2,3,4,5", func() {
			linkListA.Insert(5, 5)
			So(linkListA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
		})
		Convey("The result of Insert:1,2,5,3,4", func() {
			linkListA.Insert(3, 5)
			So(linkListA.ToSlice(), ShouldResemble, []int{1, 2, 5, 3, 4})
		})
		Convey("The result of Insert:5,1,2,3,4", func() {
			linkListA.Insert(1, 5)
			So(linkListA.ToSlice(), ShouldResemble, []int{5, 1, 2, 3, 4})
		})
		Convey("The result of Insert out-of-bound-index is  empty", func() {
			linkListB.Insert(2, 1)
			So(linkListB.ToSlice(), ShouldBeEmpty)
		})
		Convey("The result of Insert Invalid-index is  empty", func() {
			linkListB.Insert(0, 1)
			So(linkListB.ToSlice(), ShouldBeEmpty)
		})
	})
}

func TestAddElem(t *testing.T) {
	Convey("Test AddElem", t, func() {
		var linkListA, linkListB List[int]
		linkListA = *linkListA.InitList()
		linkListB = *linkListB.InitList()
		linkListA.RPush(1, 2, 3, 4)
		linkListA.AddElem(5)
		Convey("The result of AddElem:5,1,2,3,4", func() {
			So(linkListA.ToSlice(), ShouldResemble, []int{5, 1, 2, 3, 4})
		})
		Convey("The result of AddElem in a empty list is 1", func() {
			linkListB.AddElem(1)
			So(linkListB.ToSlice(), ShouldResemble, []int{1})
		})
	})
}

func TestAppendElem(t *testing.T) {
	Convey("Test AppendElem", t, func() {
		var linkListA, linkListB List[int]
		linkListA = *linkListA.InitList()
		linkListB = *linkListB.InitList()
		linkListA.RPush(1, 2, 3, 4)
		linkListA.AppendElem(5)
		Convey("The result of AppendElem:1,2,3,4,5", func() {
			So(linkListA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
		})
		Convey("The result of AppendElem in a empty list is 1", func() {
			linkListB.AppendElem(1)
			So(linkListB.ToSlice(), ShouldResemble, []int{1})
		})
	})
}

func TestLocateElem(t *testing.T) {
	Convey("Test LocateElem", t, func() {
		var linkListA List[int]
		var linkListB List[string]
		linkListA = *linkListA.InitList()
		linkListB = *linkListB.InitList()
		linkListA.RPush(1, 2, 3, 4, 5)
		linkListB.RPush("a", "b", "c", "d", "e")
		Convey("The result of LocateElem(param:3) is true", func() {
			So(linkListA.LocateElem(func(value int) bool { return 3 == value }), ShouldBeTrue)
		})
		Convey("The result of LocateElem(param:5) is true", func() {
			So(linkListA.LocateElem(func(value int) bool { return 5 == value }), ShouldBeTrue)
		})
		Convey("The result of LocateElem(param:0) is true", func() {
			So(linkListA.LocateElem(func(value int) bool { return 0 == value }), ShouldBeFalse)
		})
		Convey("The result of LocateElem(param:c) is true", func() {
			So(linkListB.LocateElem(func(value string) bool { return "c" == value }), ShouldBeTrue)
		})
		Convey("The result of LocateElem(param:e) is true", func() {
			So(linkListB.LocateElem(func(value string) bool { return "e" == value }), ShouldBeTrue)
		})
		Convey("The result of LocateElem(param:f) is true", func() {
			So(linkListB.LocateElem(func(value string) bool { return "f" == value }), ShouldBeFalse)
		})
	})
}

func TestFindElem(t *testing.T) {
	Convey("Test FindElem", t, func() {
		var linkListA List[int]
		var linkListB List[string]
		linkListA = *linkListA.InitList()
		linkListB = *linkListB.InitList()
		linkListA.RPush(1, 2, 3, 4, 5)
		linkListB.RPush("a", "b", "c", "d", "e")
		Convey("The result of FindElem(param:3) is 3", func() {
			So(linkListA.FindElem(func(value int) bool { return 3 == value }), ShouldEqual, 3)
		})
		Convey("The result of FindElem(param:5) is 5", func() {
			So(linkListA.FindElem(func(value int) bool { return 5 == value }), ShouldEqual, 5)
		})
		Convey("The result of FindElem(param:6) is 0", func() {
			So(linkListA.FindElem(func(value int) bool { return 6 == value }), ShouldEqual, 0)
		})
		Convey("The result of FindElem(param:c) is 3", func() {
			So(linkListB.FindElem(func(value string) bool { return "c" == value }), ShouldEqual, 3)
		})
		Convey("The result of FindElem(param:e) is 3", func() {
			So(linkListB.FindElem(func(value string) bool { return "e" == value }), ShouldEqual, 5)
		})
		Convey("The result of FindElem(param:f) is 0", func() {
			So(linkListB.FindElem(func(value string) bool { return "f" == value }), ShouldEqual, 0)
		})
	})

}

func TestDelElem(t *testing.T) {
	Convey("Test DelElem", t, func() {
		var linkListA List[int]
		linkListA = *linkListA.InitList()
		linkListA.RPush(1, 2, 3, 4, 5)
		Convey("The result of DelElem(param:3) :1,2,4,5", func() {
			linkListA.DelElem(3)
			So(linkListA.ToSlice(), ShouldResemble, []int{1, 2, 4, 5})
		})
		Convey("The result of DelElem(param:1) :2,3,4,5", func() {
			linkListA.DelElem(1)
			So(linkListA.ToSlice(), ShouldResemble, []int{2, 3, 4, 5})
		})
		Convey("The result of DelElem(param:5) :1,2,3,4", func() {
			linkListA.DelElem(5)
			So(linkListA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4})
		})
		Convey("The result of DelElem(param:0) :1,2,3,4,5", func() {
			linkListA.DelElem(0)
			So(linkListA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
		})
		Convey("The result of DelElem(param:6) :1,2,3,4,5", func() {
			linkListA.DelElem(6)
			So(linkListA.ToSlice(), ShouldResemble, []int{1, 2, 3, 4, 5})
		})
	})
}
