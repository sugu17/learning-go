package binarytree

const (
	Smaller = -1
	Equal   = 0
	Greater = 1
)

type OrderableFunc[T any] func(a, b T) int

type Node[T any] struct {
	val   T
	left  *Node[T]
	right *Node[T]
}

func (n *Node[T]) Add(val T, orderer OrderableFunc[T]) *Node[T] {
	if n == nil {
		return &Node[T]{
			val: val,
		}
	}
	switch r := orderer(val, n.val); r {
	case Smaller:
		n.left = n.left.Add(val, orderer)
	case Greater:
		n.right = n.right.Add(val, orderer)
	}
	return n
}

func (n *Node[T]) Contains(val T, orderer OrderableFunc[T]) bool {
	if n == nil {
		return false
	}
	switch r := orderer(val, n.val); r {
	case Smaller:
		return n.left.Contains(val, orderer)
	case Greater:
		return n.right.Contains(val, orderer)
	}
	return true
}

type Tree[T any] struct {
	root    *Node[T]
	orderer OrderableFunc[T]
}

func (t *Tree[T]) Add(val T) {
	t.root = t.root.Add(val, t.orderer)
}

func (t *Tree[T]) Contains(val T) bool {
	return t.root.Contains(val, t.orderer)
}

func NewTree[T any](orderer OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		orderer: orderer,
	}
}
