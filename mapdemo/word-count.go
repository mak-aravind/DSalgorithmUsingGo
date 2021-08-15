package mapdemo

import (
	"strings"
)

func WordCount(sentence string) map[string]int {
	var wordsMappedToCount map[string]int = make(map[string]int)
	for _, word := range strings.Fields(sentence) {
		if count, wordAvailable := wordsMappedToCount[word]; wordAvailable {
			wordsMappedToCount[word] = count + 1
		} else {
			wordsMappedToCount[word] = 1
		}
	}
	return wordsMappedToCount
}
