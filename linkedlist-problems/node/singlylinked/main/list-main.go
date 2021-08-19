package main

import (
	"fmt"

	"github.com/mak-aravind/DSalgorithmUsingGo/linkedlist-problems/node/singlylinked"
)

func main() {
	listOfValueToAdd := []int{45, 67, 80, 90}
	fmt.Println("List to add:", listOfValueToAdd)
	list := singlylinked.List{}

	for _, value := range listOfValueToAdd {
		list.Insert(value)
	}
	list.Print()
}
