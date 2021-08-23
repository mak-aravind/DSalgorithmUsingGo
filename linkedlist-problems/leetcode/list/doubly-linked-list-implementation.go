package list

import "fmt"

const FORWARD string = "forward"
const BACKWARD string = "backward"

type Node struct {
	data     int
	next     *Node
	previous *Node
}
type DoublyLinkedList struct {
	size int
	head *Node
	tail *Node
}

/** Initialize your data structure here. */
func Constructor() DoublyLinkedList {
	doublyLinkedList := DoublyLinkedList{
		size: 0,
		head: &Node{},
		tail: &Node{},
	}
	return doublyLinkedList
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (list *DoublyLinkedList) Get(index int) int {
	if list.size < 0 || list.size < index {
		return -1
	} else {
		if index < list.size/2 {
			traversePointer := list.head
			for i := 0; i <= index; i++ {
				if i == index {
					return traversePointer.data
				}
				traversePointer = traversePointer.next
			}
		} else {
			traversePointer := list.tail
			for i := list.size - 1; i <= index; i-- {
				if i == index {
					return traversePointer.data
				}
				traversePointer = traversePointer.previous
			}
		}
		return -1
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (list *DoublyLinkedList) AddAtIndex(index int, value int) {
	if index == list.size-1 {
		list.AddAtTail(value)
	}
	node := Node{}
	node.data = value
	if list.size < 0 || list.size < index {
		return
	} else {
		if index < list.size/2 {
			traversePointer := list.head
			for i := 0; i <= index; i++ {
				if i == index {
					node.previous = traversePointer
					traversePointer.next = &node
				}
				traversePointer = traversePointer.next
			}
		} else {
			traversePointer := list.tail
			for i := list.size - 1; i >= index; i-- {
				if i == index {
					node.previous = traversePointer
					traversePointer.next = &node
				}
				traversePointer = traversePointer.previous
			}
		}
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (list *DoublyLinkedList) AddAtHead(value int) {
	node := Node{}
	node.data = value
	if list.size == 0 {
		list.head = &node
		list.tail = &node
		list.size++
		return
	} else {
		insertionPointer := list.head
		for i := 0; i < list.size; i++ {
			if insertionPointer.next == nil {
				insertionPointer.next = &node
				node.previous = insertionPointer
				list.tail = &node
				list.size++
				return
			}
			insertionPointer = insertionPointer.next
		}
	}
}

/** Append a node of value val to the last element of the linked list. */
func (list *DoublyLinkedList) AddAtTail(value int) {
	node := Node{}
	node.data = value
	if list.size == 0 {
		list.head = &node
		list.tail = &node
		list.size++
		return
	} else {
		insertionPointer := list.tail
		insertionPointer.next = &node
		node.previous = insertionPointer
		list.tail = &node
		list.size++
		return
	}

}

func (list DoublyLinkedList) TraveseFromHeadUsingNext() {
	if list.head == nil {
		fmt.Println("The List is Empty")
	} else {
		list.PointerTraversal(list.head, FORWARD)
	}
}

func (list DoublyLinkedList) TraveseFromTailUsingPrevious() {
	if list.head == nil {
		fmt.Println("The List is Empty")
	} else {
		list.PointerTraversal(list.tail, BACKWARD)
	}
}

func (list *DoublyLinkedList) PointerTraversal(traversePointer *Node, traverseDirection string) {
	fmt.Print(traversePointer.data)
	if traverseDirection == FORWARD {
		for traversePointer.next != nil {
			fmt.Print(" -> ", traversePointer.next.data)
			traversePointer = traversePointer.next
		}
	} else if traverseDirection == BACKWARD {
		for traversePointer.previous != nil {
			fmt.Print(" -> ", traversePointer.previous.data)
			traversePointer = traversePointer.previous
		}
	}
}

func (list DoublyLinkedList) GetSize() int {
	return list.size
}
