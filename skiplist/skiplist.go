package skiplist

import (
	"fmt"
	"math"
	"math/rand"

	"golang.org/x/exp/constraints"
)

const (
	//最高层数
	MAX_LEVEL   = 16
	probability = 0.25
)

type SkipList[T constraints.Ordered] struct {
	head   *Node[T]
	level  int
	length int
}

/**
 * 跳表的节点，每个节点记录了当前节点数据和所在层数数据
 */
type Node[T constraints.Ordered] struct {
	value T   // 跳表保存的值
	score int // 排序
	level int // 层高
	/**
	 * 表示当前节点位置的下一个节点所有层的数据，从上层切换到下层，就是数组下标-1，
	 * forwards[3]表示当前节点在第三层的下一个节点。
	 */
	forward []*Node[T]
}

// 新建链表节点
func newNode[T constraints.Ordered](value T, score, level int) *Node[T] {
	node := &Node[T]{
		value:   value,
		score:   score,
		level:   level,
		forward: make([]*Node[T], level),
	}
	for i := range node.forward {
		node.forward[i] = new(Node[T])
	}
	return node
}

// 新建跳表
func NewSkipList[T constraints.Ordered](value T) *SkipList[T] {
	head := newNode(value, math.MinInt32, MAX_LEVEL)
	return &SkipList[T]{
		head:   head,
		level:  1,
		length: 0,
	}
}

// 随机层数
func randomLevel() int {
	var level int = 1
	for level < MAX_LEVEL {
		if rand.Float64() < probability {
			break
		}
		level++
	}
	return level
}

// 往跳表中插入数据
func (sl *SkipList[T]) Insert(value T, score int) {
	level := randomLevel()
	newNode := newNode(value, score, level)
	cur := sl.head
	updated := make([]*Node[T], level)
	fmt.Printf("randomLevel:%d \n", level)
	fmt.Printf("updated length:%d \n", len(updated))
	fmt.Printf("cur forward length:%d \n", len(cur.forward))
	fmt.Printf("newNode:%+v \n", newNode)
	fmt.Printf("cur : %+v \n", cur)
	fmt.Printf("cur forward[6] : %+v \n", cur.forward[6])
	for i := level - 1; i >= 0; i-- {
		for cur.forward[i] != nil {
			if cur.forward[i].value == value {
				return
			} else if cur.forward[i].score > score {
				updated[i] = cur
				break
			}
			cur = cur.forward[i]
		}
		if cur.forward[i] == nil {
			updated[i] = cur
		}
	}
	for i := 0; i < level; i++ {
		newNode.forward[i] = updated[i].forward[i]
		updated[i].forward[i] = newNode
	}

	if sl.level < level {
		sl.level = level
	}
	sl.length++
}

// 查找
func (sl *SkipList[T]) Find(value T, score int) bool {
	if sl == nil {
		return false
	}
	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil {
			if cur.forward[i].score == score && cur.forward[i].value == value {
				return true
			} else if cur.forward[i].score > score {
				break
			}
			cur = cur.forward[i]
		}
	}
	return false
}

// 删除
func (sl *SkipList[T]) Delete(value T, score int) {
	if sl == nil {
		return
	}
	cur := sl.head
	updated := make([]*Node[T], sl.level)
	// 查找每层得节点
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forward[i] != nil {
			for ; cur.forward[i] != nil && cur.forward[i].score < score; cur = cur.forward[i] {

			}
			updated[i] = cur
		}
	}
	// 遍历每层删除节点
	for i := 0; i < sl.level; i++ {
		if updated[i].forward[i] != nil && updated[i].forward[i].score == score && updated[i].forward[i].value == value {
			updated[i].forward[i] = updated[i].forward[i].forward[i]
		}
	}
	// 更新当前最大层数
	for i := sl.level - 1; i > 0; i-- {
		if sl.head.forward[i] == nil {
			sl.level--
		}
	}
	sl.length--
}

func (sl *SkipList[T]) print() {
	for i := sl.level; i >= 0; i-- {
		p := sl.head
		for ; p != nil; p = p.forward[i] {
			fmt.Printf("%v->", p.value)
		}
		fmt.Printf("nil\n")
	}
}
