package tree

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBuildByPreIn(t *testing.T) {
	Convey("Test BinaryTree BuildByPreIn", t, func() {
		pre := []int{3, 9, 20, 15, 7}
		in := []int{9, 3, 15, 20, 7}
		var tree BinaryTree[int]
		tree = *tree.New()
		tree.root = BuildByPreIn(pre, in)
		emptyPre := []int{}
		var emptyTree BinaryTree[int]
		emptyTree = *emptyTree.New()
		emptyTree.root = BuildByPreIn(emptyPre, in)
		Convey("The result of BuildByPreIn pre-order is :3,9,20,15,7", func() {
			So(tree.ToSlice(1), ShouldResemble, []int{3, 9, 20, 15, 7})
		})
		Convey("The result of BuildByPreIn in-order is :9,3,15,20,7", func() {
			So(tree.ToSlice(2), ShouldResemble, []int{9, 3, 15, 20, 7})
		})
		Convey("The result of BuildByPreIn post-order is :9, 15, 7, 20, 3", func() {
			So(tree.ToSlice(3), ShouldResemble, []int{9, 15, 7, 20, 3})
		})
		Convey("The result of BuildByPreIn bfs-order is :3,9,20,15,7", func() {
			So(tree.ToSlice(4), ShouldResemble, []int{3, 9, 20, 15, 7})
		})
		Convey("The result of BuildByPreIn(empty) bfs-order is :empty", func() {
			So(emptyTree.ToSlice(4), ShouldBeEmpty)
		})
	})
}

func TestBuildByPostIn(t *testing.T) {
	Convey("Test BinaryTree BuildByPostIn", t, func() {
		post := []int{9, 15, 7, 20, 3}
		in := []int{9, 3, 15, 20, 7}
		var tree BinaryTree[int]
		tree = *tree.New()
		tree.root = BuildByPostIn(post, in)
		emptyPost := []int{}
		var emptyTree BinaryTree[int]
		emptyTree = *emptyTree.New()
		emptyTree.root = BuildByPostIn(emptyPost, in)
		Convey("The result of BuildByPostIn pre-order is :3,9,20,15,7", func() {
			So(tree.ToSlice(1), ShouldResemble, []int{3, 9, 20, 15, 7})
		})
		Convey("The result of BuildByPostIn in-order is :9,3,15,20,7", func() {
			So(tree.ToSlice(2), ShouldResemble, []int{9, 3, 15, 20, 7})
		})
		Convey("The result of BuildByPostIn post-order is :9, 15, 7, 20, 3", func() {
			So(tree.ToSlice(3), ShouldResemble, []int{9, 15, 7, 20, 3})
		})
		Convey("The result of BuildByPostIn bfs-order is :3,9,20,15,7", func() {
			So(tree.ToSlice(4), ShouldResemble, []int{3, 9, 20, 15, 7})
		})
		Convey("The result of BuildByPostIn(empty) bfs-order is :empty", func() {
			So(emptyTree.ToSlice(4), ShouldBeEmpty)
		})
	})
}

func TestDepth(t *testing.T) {
	Convey("Test BinaryTree Depth", t, func() {
		pre := []int{3, 9, 20, 15, 7}
		in := []int{9, 3, 15, 20, 7}
		var tree BinaryTree[int]
		Convey("The result of Depth is 3", func() {
			tree = *tree.New()
			tree.root = BuildByPreIn(pre, in)
			So(tree.Depth(), ShouldEqual, 3)
		})
		Convey("The result of Depth is 0", func() {
			So(tree.Depth(), ShouldEqual, 0)
		})
	})
}

func TestIsFBT(t *testing.T) {
	Convey("Test BinaryTree IsFBT", t, func() {
		pre1 := []int{3, 9, 8, 20, 15, 7}
		in1 := []int{8, 9, 3, 15, 20, 7}
		pre2 := []int{1, 2, 4, 5, 3, 6, 7}
		in2 := []int{4, 2, 5, 1, 6, 3, 7}
		var not_fbt, fbt BinaryTree[int]
		Convey("The result of IsFBT is fasle", func() {
			not_fbt = *not_fbt.New()
			not_fbt.root = BuildByPreIn(pre1, in1)
			So(not_fbt.ToSlice(3), ShouldResemble, []int{8, 9, 15, 7, 20, 3})
			So(not_fbt.IsFBT(), ShouldBeFalse)
		})
		Convey("The result of IsFBT is true", func() {
			fbt = *fbt.New()
			fbt.root = BuildByPreIn(pre2, in2)
			So(fbt.ToSlice(3), ShouldResemble, []int{4, 5, 2, 6, 7, 3, 1})
			So(fbt.IsFBT(), ShouldBeTrue)
		})
	})
}

func TestIsCBT(t *testing.T) {
	Convey("Test BinaryTree IsCBT", t, func() {
		pre1 := []int{3, 9, 8, 20, 15, 7}
		in1 := []int{8, 9, 3, 15, 20, 7}
		pre2 := []int{1, 2, 4, 5, 3, 6, 7}
		in2 := []int{4, 2, 5, 1, 6, 3, 7}
		var not_cbt, cbt BinaryTree[int]
		Convey("The result of IsCBT is fasle", func() {
			not_cbt = *not_cbt.New()
			not_cbt.root = BuildByPreIn(pre1, in1)
			So(not_cbt.ToSlice(3), ShouldResemble, []int{8, 9, 15, 7, 20, 3})
			So(not_cbt.IsCBT(), ShouldBeFalse)
		})
		Convey("The result of IsCBT is true", func() {
			cbt = *cbt.New()
			cbt.root = BuildByPreIn(pre2, in2)
			So(cbt.ToSlice(3), ShouldResemble, []int{4, 5, 2, 6, 7, 3, 1})
			So(cbt.IsCBT(), ShouldBeTrue)
		})
	})
}

func TestIsAVL(t *testing.T) {
	Convey("Test BinaryTree IsAVL", t, func() {
		pre1 := []int{3, 9, 20, 15, 7}
		in1 := []int{9, 3, 15, 20, 7}
		pre2 := []int{1, 2, 3, 4, 4, 3, 2}
		in2 := []int{4, 3, 4, 2, 4, 1, 2}
		var avl, not_avl BinaryTree[int]
		Convey("The result of IsAVL is true", func() {
			avl = *avl.New()
			avl.root = BuildByPreIn(pre1, in1)
			So(avl.ToSlice(3), ShouldResemble, []int{9, 15, 7, 20, 3})
			So(avl.IsAVL(), ShouldBeFalse)
		})
		Convey("The result of IsAVL is false", func() {
			not_avl = *not_avl.New()
			not_avl.root = BuildByPreIn(pre2, in2)
			So(not_avl.ToSlice(3), ShouldResemble, []int{4, 4, 3, 3, 2, 2, 1})
			So(not_avl.IsAVL(), ShouldBeFalse)
		})
	})
}
