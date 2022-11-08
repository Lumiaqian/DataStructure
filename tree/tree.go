package tree

import (
	"datastructure/queue"
	"datastructure/util"
	"math"

	"golang.org/x/exp/constraints"
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

// 获取二叉树深度
func (t *BinaryTree[T]) Depth() int {
	return t.root.depth()
}

func (n *Node[T]) depth() int {
	if n == nil {
		return 0
	}
	return util.Max(n.Left.depth(), n.Right.depth()) + 1
}

/*
*
*满二叉树：如果二叉树中除了叶子结点，每个结点的度都是2，则是满二叉树
*性质：
*	满二叉树中第 i 层的节点数为 2^i - 1 个。
*	深度为k的满二叉树必有 2k-1 个节点 ，叶子数为 2k-1。
*	满二叉树中不存在度为 1 的节点，每一个分支点中都两棵深度相同的子树，且叶子节点都在最底层。
*	具有 n 个节点的满二叉树的深度为 log2(n+1)。
*
* 判断方法： 左子树为满二叉树且右子树为满二叉树
 */
func (t *BinaryTree[T]) IsFBT() bool {
	return t.root.isFBT()
}

func (n *Node[T]) isFBT() bool {
	if n == nil {
		return true
	}
	if !n.Left.isFBT() {
		return false
	}
	if (n.Left == nil && n.Right != nil) || (n.Left != nil && n.Right == nil) {
		return false
	}
	return n.Right.isFBT()
}

/*
* 完全二叉树 : 如果二叉树中除去最后一层节点为满二叉树，且最后一层的结点依次从左到右分布，则此二叉树被称为完全二叉树。
* 判断方法：
*1.按照层次遍历
*2.任意一个节点有右子节点无左子节点，则为非完全二叉树
*3.在条件1不违规的情况下，如果遇到左右子节点不齐全的情况下，后续节点皆为叶子节点，如果后续存在非叶子节点，则为非完全二叉树。
 */
func (t *BinaryTree[T]) IsCBT() bool {
	if t.root == nil {
		return true
	}
	var queue queue.ArrayQueue[Node[T]]
	var leaf = false //是否遇到左右子节点不齐全的情况
	queue = *queue.New()
	queue.Put(*t.root)
	for queue.Size() > 0 {
		cur, _ := queue.Pop()
		l := cur.Left
		r := cur.Right
		//如果遇到了不双全的节点后，又发现当前节点有子节点 或者 有右节点但无左节点
		if (leaf && (l != nil || r != nil)) || (l == nil && r != nil) {
			return false
		}
		if l != nil {
			queue.Put(*l)
		}
		if r != nil {
			queue.Put(*r)
		}
		if l == nil || r == nil {
			leaf = true
		}
	}
	return true
}

/*
* 平衡二叉树 ：对于任何一个子树来说，左树高度和右树高度差不超过1
* 判断方法：
* A && B && C
* A: 左子树为平衡二叉树
* B: 右子树为平衡二叉树
* C: abs(左高-右高) <= 1
 */
type Data struct {
	IsBalanced bool
	Height     int
}

func (t *BinaryTree[T]) IsAVL() bool {
	var check func(*Node[T]) *Data
	check = func(n *Node[T]) *Data {
		if n == nil {
			return &Data{
				IsBalanced: true,
				Height:     0,
			}
		}
		left := check(n.Left)
		right := check(n.Right)
		hight := util.Max(left.Height, right.Height) + 1
		isBalance := left.IsBalanced && right.IsBalanced && math.Abs(float64(left.Height-right.Height)) <= 1
		return &Data{
			IsBalanced: isBalance,
			Height:     hight,
		}
	}
	return check(t.root).IsBalanced
}

/*
* 二叉搜索树：每一个子树，左节点比根节点小，右节点比根节点大的二叉树，二叉树的节点无重复值
*判断方法：
*  中序（递归中序或者非递归中序均可）遍历二叉树，如果是升序的即为搜索二叉树。
 */
func IsBST[T constraints.Ordered](root *Node[T], value T) (bool, T) {
	if root == nil {
		return true, value
	}
	left, _ := IsBST(root.Left, value)
	if !left {
		return false, value
	}
	if root.Value <= value {
		return false, value
	}
	value = root.Value
	return IsBST(root.Right, value)
}
