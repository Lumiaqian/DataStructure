package list

type DoublyNode[T any] struct {
	Value T
	Prev  *DoublyNode[T]
	Next  *DoublyNode[T]
}

type Node[T any] struct {
	Value T
	Next  *Node[T]
}
