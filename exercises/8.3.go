package exercises

type Node[T any] struct {
	Value T
	next  *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

type List[T any] struct {
	len  int
	head *Node[T]
	tail *Node[T]
}

func (l *List[T]) Add(val T) {
	if l.len == 0 {
		l.head = &Node[T]{
			Value: val,
		}
		l.tail = l.head
		l.len += 1
		return
	}
	node := &Node[T]{
		Value: val,
	}
	l.tail.next = node
	l.tail = node
	l.len += 1
}

func (l *List[T]) Front() *Node[T] {
	return l.head
}

func (l *List[T]) Len() int {
	return l.len
}

func NewList[T any]() List[T] {
	return List[T]{
		len: 0,
	}
}
