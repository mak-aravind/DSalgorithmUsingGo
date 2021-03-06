package singlylinked

import (
	"fmt"

	"github.com/mak-aravind/DSalgorithmUsingGo/performance"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	length int
	head   *Node
	tail   *Node
}

func (list *List) InsertUsingHead(data int) {
	defer performance.Duration(performance.Track("InsertUsingHead"))
	node := Node{}
	node.data = data

	if list.length == 0 {
		list.head = &node
		list.tail = &node
		list.length++
		return
	}
	insertPointer := list.head
	for i := 0; i < list.length; i++ {
		if insertPointer.next == nil {
			insertPointer.next = &node
			list.length++
			list.tail = &node
			return
		}
		insertPointer = insertPointer.next
	}
}

func (list *List) InsertAtTail(data int) {
	defer performance.Duration(performance.Track("InsertAtTail"))
	node := Node{}
	node.data = data

	if list.length == 0 {
		list.head = &node
		list.tail = &node
		list.length++
		return
	} else {
		insertPointer := list.tail
		insertPointer.next = &node
		list.tail = &node
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
}
