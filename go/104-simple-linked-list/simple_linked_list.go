package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func New(elements []int) *List {
	list := &List{}
	for _, e := range elements {
		list.Push(e)
	}
	return list
}

func (l *List) Size() int {
	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (l *List) Push(element int) {
	newNode := &Node{Value: element}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	newNode.prev = l.Tail
	l.Tail.next = newNode
	l.Tail = newNode
}

func (l *List) Pop() (int, error) {

	if l.Tail == nil {
		return 0, errors.New("empty list")
	}

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

func (l *List) Array() []int {
	var result []int
	fmt.Println(l.Head)
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.next
	}

	return result
}

func (l *List) Reverse() *List {
	reversedList := &List{}
	current := l.Tail
	for current != nil {
		reversedList.Push(current.Value)
		current = current.prev
	}

	return reversedList

}
