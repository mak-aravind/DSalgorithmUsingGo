package wordcount

import (
	"strings"
)

func WordCount(sentence string) map[string]int {
	const ONE int = 1
	var wordsMappedToCount map[string]int = make(map[string]int)
	for _, word := range strings.Fields(sentence) {
		if count, wordAvailable := wordsMappedToCount[word]; wordAvailable {
			wordsMappedToCount[word] = count + ONE
		} else {
			wordsMappedToCount[word] = ONE
		}
	}
	return wordsMappedToCount
}
