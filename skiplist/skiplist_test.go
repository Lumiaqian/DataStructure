package skiplist

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSkipList(t *testing.T) {
	Convey("Test SkipList", t, func() {
		sl := NewSkipList("///")
		sl.Insert("lumia", 3)
		sl.Insert("qian", 6)
		sl.Insert("cao", 7)
		sl.Insert("yu", 9)
		sl.Insert("list", 12)
		sl.Insert("tree", 19)
		sl.Insert("trie", 21)
		sl.Insert("queue", 25)
		sl.Insert("stack", 17)
		Convey("The result of Find is true", func() {
			sl.Insert("heap", 16)
			So(sl.Find("heap", 16), ShouldBeTrue)
			So(sl.Find("tree", 19), ShouldBeTrue)
			So(sl.Find("queue", 25), ShouldBeTrue)
		})
		Convey("The result of Find is false", func() {
			So(sl.Find("heap", 16), ShouldBeFalse)
			So(sl.Find("lumia", 3), ShouldBeTrue)
			sl.Delete("lumia", 3)
			So(sl.Find("lumia", 3), ShouldBeFalse)

		})
	})
}
