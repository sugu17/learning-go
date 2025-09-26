package list

type Node[T comparable] struct {
	Value T
	next  *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

type List[T comparable] struct {
	len  int
	head *Node[T]
	tail *Node[T]
}

func NewList[T comparable]() List[T] {
	return List[T]{
		len: 0,
	}
}

func (l *List[T]) Front() *Node[T] {
	return l.head
}

func (l *List[T]) Len() int {
	return l.len
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

func (l *List[T]) Insert(index int, val T) bool {
	if index == 0 {
		node := &Node[T]{
			Value: val,
			next:  nil,
		}
		if l.head == nil {
			l.head = &Node[T]{
				Value: val,
				next:  node,
			}
			l.tail = l.head
			l.len += 1
			return true
		}
		prev := l.head.next
		node.next = prev
		l.head.next = node
		l.len += 1
		return true
	}
	for it, idx := l.Front(), 0; it != nil; it, idx = it.Next(), idx+1 {
		if idx != index {
			continue
		}
		node := &Node[T]{
			Value: val,
			next:  nil,
		}
		if it.next == nil {
			l.tail.next = node
			l.tail = node
		} else {
			prev := it.next
			node.next = prev
			it.next = node
		}
		l.len += 1
		return true
	}
	return false
}

func (l *List[T]) Index(val T) int {
	for it, index := l.Front(), 0; it != nil; it, index = it.Next(), index+1 {
		if it.Value == val {
			return index
		}
	}
	return -1
}
