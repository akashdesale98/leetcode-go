package main

import "fmt"

type Node struct {
	data   int
	next   *Node
	length int
}

func main() {
	a := []int{4, 2, 1, 6, 7, 3}

	head := BuildLL(a)
	// PrintLL(head)
	// fmt.Println("length", FindLLLength(head))
	// head = Delete(head, 4)
	// PrintLL(head)
	// head = Insert(head, 7, 8)
	// PrintLL(head)
	head = ReverseIterative(head)
	PrintLL(head)
}

func BuildLL(a []int) *Node {
	head := &Node{
		data: a[0],
	}

	curr := head
	for i := 1; i < len(a); i++ {
		temp := &Node{
			data: a[i],
		}
		curr.next = temp
		curr = temp
	}

	return head
}

func PrintLL(curr *Node) {
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}

func FindLLLength(curr *Node) int {
	i := 0
	for curr != nil {
		i++
		curr = curr.next
	}

	return i
}

func CheckIfPresent(curr *Node, data int) bool {
	for curr != nil {
		if curr.data == data {
			return true
		}
		curr = curr.next
	}

	return false
}

func Insert(head *Node, pos, data int) *Node {
	if head == nil {
		if pos == 1 {
			return &Node{data: data}
		}
		return head
	}

	if pos == 1 {
		return &Node{data: data, next: head}
	}

	count := 1
	curr := head
	for curr != nil {
		count++
		if count == pos {
			temp := &Node{data: data, next: curr.next}
			curr.next = temp
			break
		}

		curr = curr.next
	}

	return head
}

func Delete(head *Node, data int) *Node {

	if head == nil {
		return nil
	}

	if head.data == data {
		// newHead := head.next
		// head.next = nil
		// return newHead
		return head.next
	}

	curr := head
	for curr.next != nil {
		if curr.next.data == data {
			curr.next = curr.next.next
			break
		}
		curr = curr.next
	}

	// curr, prev := head, &Node{}
	// for curr != nil {
	// 	if curr.data == data {
	// 		prev.next = prev.next.next
	// 		break
	// 	}
	// 	prev = curr
	// 	curr = curr.next
	// }

	return head
}

func ReverseIterative(head *Node) *Node {
	var prev *Node
	curr := head
	for curr != nil {
		next := curr.next // Store next node
		curr.next = prev  // Reverse current node's pointer
		prev = curr       // Move pointers one position ahead
		curr = next
	}
	return prev // Return the new head of the reversed list
}

func ReverseRecusive(head *Node) *Node {

	if head == nil || head.next == nil {
		return head
	}

	node := ReverseRecusive(head.next)

	front := head.next
	front.next = head
	head.next = nil

	return node

}
