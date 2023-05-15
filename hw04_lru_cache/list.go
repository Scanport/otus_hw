package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	size int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newNode := &ListItem{Value: v, Next: l.head, Prev: nil}
	if l.head != nil {
		l.head.Prev = newNode
	} else {
		l.tail = newNode
	}
	l.head = newNode
	l.size++
	return newNode
}

func (l *list) PushBack(v interface{}) *ListItem {
	newNode := &ListItem{Value: v, Next: nil, Prev: l.tail}
	if l.tail != nil {
		l.tail.Next = newNode
	} else {
		l.head = newNode
	}
	l.tail = newNode
	l.size++
	return newNode
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}
	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.head {
		return
	}
	// Remove the node from its current position
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	// Move the node to the front
	i.Prev = nil
	i.Next = l.head
	if l.head != nil {
		l.head.Prev = i
	} else {
		l.tail = i
	}
	l.head = i
}
