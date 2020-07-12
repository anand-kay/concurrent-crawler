package indexedwords

import "strings"

var indexedWords = make(map[string]map[string]bool)

func SetIndexedWords(words []string, url string) {
	for _, word := range words {
		word = strings.ToLower(word)

		tempMap := indexedWords[word]

		if tempMap == nil {
			tempMap = make(map[string]bool)
			tempMap[url] = true

			indexedWords[word] = tempMap
		}

		tempMap[url] = true

	}
}

func GetIndexedWords() *map[string]map[string]bool {
	return &indexedWords
}
