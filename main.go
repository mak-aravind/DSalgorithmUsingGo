package main

import (
	"fmt"

	"example.com/go-demo-1/mapdemo"
)

func main() {
	listOfSentences := [3]string{
		"I am the god and I am the evil",
		"red shoes red toes white eyes white hair",
		"tall towers tall trees wow great wow fantastic",
	}

	for _, sentence := range listOfSentences {
		fmt.Println(sentence)
		fmt.Println(mapdemo.WordCount(sentence))
	}
}
