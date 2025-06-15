package main

import (
	"fmt"
)

type Node struct {
	data int16
	next *Node
}

var Head *Node = nil

func InsertAtEnd(data int16) {
	newnode := &Node{data: data, next: nil}
	if Head == nil {
		Head = newnode
	} else {
		var temp *Node = Head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newnode
	}
}

func InsertAtBegin(data int16) {
	newnode := &Node{data: data, next: nil}
	if Head == nil {
		Head = newnode
	} else {
		newnode.next = Head
		Head = newnode
	}
}

func Display(Head *Node) {
	var Temp *Node = Head
	for Temp != nil {
		fmt.Print(Temp.data, "->")
		Temp = Temp.next
	}
	fmt.Print("nil")
}

func main() {
	fmt.Println("Linked List")
	InsertAtEnd(10)
	InsertAtBegin(20)
	InsertAtEnd(23)
	InsertAtEnd(13)
	InsertAtBegin(5)
	Display(Head)
	fmt.Println()
	Do()
}

// Another example of a linked list implementation in Go
// This approach is very useful for creating multiplle linked lists very easily
// package main

// import "fmt"

// type Node struct {
//     data int
//     next *Node
// }

// type LinkedList struct {
//     head *Node
// }

// // Insert at end
// func (ll *LinkedList) Insert(value int) {
//     newNode := &Node{data: value}
//     if ll.head == nil {
//         ll.head = newNode
//         return
//     }
//     current := ll.head
//     for current.next != nil {
//         current = current.next
//     }
//     current.next = newNode
// }

// // Print all elements
// func (ll *LinkedList) Display() {
//     current := ll.head
//     for current != nil {
//         fmt.Print(current.data, " -> ")
//         current = current.next
//     }
//     fmt.Println("nil")
// }

// func main() {
//     list := LinkedList{}
//     list.Insert(10)
//     list.Insert(20)
//     list.Insert(30)
//     list.Display()
// }
