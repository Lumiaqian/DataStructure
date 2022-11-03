package tree

import (
	"datastructure/queue"
)

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type BinaryTree[T any] struct {
	root *Node[T]
}

func (t *BinaryTree[T]) Init() *BinaryTree[T] {
	t.root = new(Node[T])
	t.root.Left = nil
	t.root.Right = nil
	return t
}

func (t *BinaryTree[T]) New() *BinaryTree[T] {
	return new(BinaryTree[T]).Init()
}

// 遍历，1-前序；2-中序；3-后序；4-层次
func (t *BinaryTree[T]) ToSlice(opt int) []T {
	if t.root == nil {
		return nil
	}
	res := []T{}
	switch opt {
	case 1:
		var toSliceByPre func(root *Node[T])
		toSliceByPre = func(root *Node[T]) {
			if root == nil {
				return
			}
			res = append(res, root.Value)
			toSliceByPre(root.Left)
			toSliceByPre(root.Right)
		}
		toSliceByPre(t.root)
	case 2:
		var toSliceByIn func(root *Node[T])
		toSliceByIn = func(root *Node[T]) {
			if root == nil {
				return
			}
			toSliceByIn(root.Left)
			res = append(res, root.Value)
			toSliceByIn(root.Right)
		}
		toSliceByIn(t.root)
	case 3:
		var toSliceByPost func(root *Node[T])
		toSliceByPost = func(root *Node[T]) {
			if root == nil {
				return
			}
			toSliceByPost(root.Left)
			toSliceByPost(root.Right)
			res = append(res, root.Value)
		}
		toSliceByPost(t.root)
	case 4:
		var toSliceByBFS func(root *Node[T])
		toSliceByBFS = func(root *Node[T]) {
			if root == nil {
				return
			}
			var queue queue.ArrayQueue[Node[T]]
			queue = *queue.New()
			//入队
			queue.Put(*root)
			for queue.Size() > 0 {
				//出队
				cur, _ := queue.Pop()
				res = append(res, cur.Value)
				if cur.Left != nil {
					queue.Put(*cur.Left)
				}
				if cur.Right != nil {
					queue.Put(*cur.Right)
				}
			}

		}
		toSliceByBFS(t.root)
	}
	return res
}

// 前+中序遍历数组构建二叉树
func BuildByPreIn[T comparable](pre, in []T) *Node[T] {
	if len(pre) == 0 {
		return nil
	}
	//1.根节点赋值
	root := &Node[T]{Value: pre[0]}
	//2.获取根节点在中序遍历数组的index
	var i int
	for index, v := range in {
		if v == pre[0] {
			i = index
			break
		}
	}
	//3.递归
	root.Left = BuildByPreIn(pre[1:i+1], in[:i])
	root.Right = BuildByPreIn(pre[i+1:], in[i+1:])
	return root
}

// 后+中序遍历数组构建二叉树
func BuildByPostIn[T comparable](post, in []T) *Node[T] {
	if len(post) == 0 {
		return nil
	}
	//初始化根节点
	root := &Node[T]{Value: post[len(post)-1]}
	//2.获取根节点在中序遍历数组中的index
	var i int
	for index, v := range in {
		if v == post[len(post)-1] {
			i = index
			break
		}
	}
	//3.递归
	root.Left = BuildByPostIn(post[:i], in[:i+1])
	root.Right = BuildByPostIn(post[i:len(post)-1], in[i+1:])
	return root
}
