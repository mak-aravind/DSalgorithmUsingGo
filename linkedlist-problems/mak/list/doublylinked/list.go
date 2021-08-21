package doublylinked

import "fmt"

type Node struct {
	data     int
	next     *Node
	previous *Node
}

type List struct {
	length int
	head   *Node
	tail   *Node
}

func (list *List) InsertUsingHead(data int) {
	nodeToInsert := Node{}
	nodeToInsert.data = data

	if list.length == 0 {
		list.head = &nodeToInsert
		list.tail = &nodeToInsert
		list.length++
		return
	} else {
		insertPointer := list.head
		for i := 0; i < list.length; i++ {
			if insertPointer.next == nil {
				insertPointer.next = &nodeToInsert
				nodeToInsert.previous = insertPointer
				list.tail = &nodeToInsert
				list.length++
				return
			}
			insertPointer = insertPointer.next
		}
	}
}

func (list *List) InsertAtTail(data int) {
	nodeToInsert := Node{}
	nodeToInsert.data = data

	if list.length == 0 {
		list.head = &nodeToInsert
		list.tail = &nodeToInsert
		list.length++
		return
	} else {
		insertPointer := list.tail
		insertPointer.next = &nodeToInsert
		nodeToInsert.previous = insertPointer
		list.tail = &nodeToInsert
		list.length++
		return
	}
}

func (list List) Print() {
	if list.length == 0 {
		fmt.Println("List is empty")
	}
	traversePointer := list.head
	fmt.Print(traversePointer.data)
	for traversePointer.next != nil {
		fmt.Print(" -> ", traversePointer.next.data)
		traversePointer = traversePointer.next
	}
	traversePointer = list.tail
	fmt.Println()
	fmt.Print(traversePointer.data)
	for traversePointer.previous != nil {
		fmt.Print(" -> ", traversePointer.previous.data)
		traversePointer = traversePointer.previous
	}
}
