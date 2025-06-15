package main

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) InsertAtEnd(data int16) {
	newnode := &Node{data: data, next: nil}
	if l.Head == nil {
		l.Head = newnode
	} else {
		var temp *Node = l.Head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newnode
	}
}

func add(head1 *Node, head2 *Node) *Node {
	list3 := &LinkedList{Head: nil}
	a := head1
	b := head2
	for a != nil {
		list3.InsertAtEnd(a.data + b.data)
		a = a.next
		b = b.next
	}
	return list3.Head
}

func Do() {
	var list1 LinkedList = LinkedList{Head: nil}
	list1.InsertAtEnd(10)
	list1.InsertAtEnd(15)
	list1.InsertAtEnd(30)
	list1.InsertAtEnd(20)
	var list2 LinkedList = LinkedList{Head: nil}
	list2.InsertAtEnd(1)
	list2.InsertAtEnd(25)
	list2.InsertAtEnd(36)
	list2.InsertAtEnd(4)
	resHead := add(list1.Head, list2.Head)
	Display(resHead)
}
