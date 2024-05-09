package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) InsertFront(data int) {
	newNode := &Node{
		data: data,
		next: dll.head,
		prev: nil,
	}

	if dll.head != nil {
		dll.head.prev = newNode
	}

	dll.head = newNode

	if dll.tail == nil {
		dll.tail = newNode
	}
}

func (dll *DoublyLinkedList) InsertBack(data int) {
	newNode := &Node{
		data: data,
	}

	if dll.tail != nil {
		newNode.prev = dll.tail
		dll.tail.next = newNode
	}

	dll.tail = newNode
	if dll.head == nil {
		dll.head = newNode
	}
}

func (dll *DoublyLinkedList) Delete(data int) {
	curr := dll.head
	for curr != nil {
		if curr.data == data {
			if curr.prev != nil {
				curr.prev.next = curr.next
			} else {
				dll.head = curr.next
			}

			if curr.next != nil {
				curr.next.prev = curr.prev
			} else {
				dll.tail = curr.next
			}
			return
		}

		curr = curr.next
	}
}

func (dll *DoublyLinkedList) TraverseForward() {
	currentNode := dll.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (dll *DoublyLinkedList) TraverseBackward() {
	currentNode := dll.tail
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.prev
	}
	fmt.Println()
}

func main() {
	doublyList := NewDoublyLinkedList()

	// Insertion
	doublyList.InsertFront(5)
	doublyList.InsertFront(3)
	doublyList.InsertBack(7)
	doublyList.InsertBack(9)

	fmt.Println("Forward traversal:")
	doublyList.TraverseForward()

	fmt.Println("Backward traversal:")
	doublyList.TraverseBackward()

	// Deletion
	doublyList.Delete(7)

	fmt.Println("Forward traversal after deleting 7:")
	doublyList.TraverseForward()
}
