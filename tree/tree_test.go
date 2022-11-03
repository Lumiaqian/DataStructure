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
