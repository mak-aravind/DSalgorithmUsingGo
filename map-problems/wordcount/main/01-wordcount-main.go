package main

import (
	"fmt"

	"github.com/mak-aravind/DSalgorithmUsingGo/map-problems/wordcount"
)

func main() {
	listOfSentences := [3]string{
		"I am the god and I am the evil",
		"red shoes red toes white eyes white hair",
		"tall towers tall trees wow great wow fantastic",
	}

	for _, sentence := range listOfSentences {
		fmt.Println(sentence)
		fmt.Println(wordcount.WordCount(sentence))
	}
}
