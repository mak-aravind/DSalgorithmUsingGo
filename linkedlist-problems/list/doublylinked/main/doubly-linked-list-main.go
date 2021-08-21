package main

import (
	"fmt"

	"github.com/mak-aravind/DSalgorithmUsingGo/linkedlist-problems/list/doublylinked"
)

func main() {
	listOfValueToAddUsingHead := []int{45, 67, 80, 90}
	fmt.Println("List to add using head:", listOfValueToAddUsingHead)
	list := doublylinked.List{}
	for _, value := range listOfValueToAddUsingHead {
		list.InsertUsingHead(value)
	}
	listOfValueToAddAtTail := []int{145, 267, 380, 490}
	fmt.Println("List to add using tail:", listOfValueToAddAtTail)
	for _, value := range listOfValueToAddAtTail {
		list.InsertAtTail(value)
	}

	list.Print()
}
