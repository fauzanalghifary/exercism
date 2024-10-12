package linkedlist

type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func NewList(elements ...interface{}) *List {
	list := &List{}
	for _, e := range elements {
		list.Push(e)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.next = l.Head
		l.Head.prev = newNode
		l.Head = newNode
	}
}

func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v}

	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.prev = l.Tail
		l.Tail.next = newNode
		l.Tail = newNode
	}
}

func (l *List) Shift() (interface{}, error) {
	val := l.Head.Value
	if l.Head.next == nil {
		l.Head = nil
		l.Tail = nil
		return val, nil
	}

	l.Head = l.Head.next
	l.Head.prev = nil

	return val, nil
}

func (l *List) Pop() (interface{}, error) {
	val := l.Tail.Value

	if l.Tail.prev == nil {
		l.Head = nil
		l.Tail = nil
		return val, nil
	}

	l.Tail = l.Tail.prev
	l.Tail.next = nil

	return val, nil
}

func (l *List) Reverse() {
	for n := l.Head; n != nil; n = n.prev {
		n.prev, n.next = n.next, n.prev
	}
	l.Head, l.Tail = l.Tail, l.Head
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	return l.Tail
}
