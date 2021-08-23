package main

import (
	"fmt"

	"github.com/mak-aravind/DSalgorithmUsingGo/linkedlist-problems/leetcode/list"
)

func main() {
	doublyLinkedList := list.Constructor()
	doublyLinkedList.AddAtHead(3)
	doublyLinkedList.AddAtTail(5)
	doublyLinkedList.AddAtTail(7)
	fmt.Println("Size:", list.Constructor().GetSize())
	doublyLinkedList.AddAtIndex(1, 6)
	data := doublyLinkedList.Get(0)
	fmt.Println("Recieved:", data)
	doublyLinkedList.TraveseFromHeadUsingNext()
	fmt.Println()
	doublyLinkedList.TraveseFromTailUsingPrevious()
	/*param_1 := obj.Get(index)
	obj.AddAtHead(val)
	obj.AddAtTail(val)
	obj.AddAtIndex(index, val)
	obj.DeleteAtIndex(index)*/
}
