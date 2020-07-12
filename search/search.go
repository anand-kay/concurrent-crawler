package search

import (
	"strings"

	"crawler/indexedwords"
)

func SearchWords(searchKey string) map[string]bool {
	var indexedWords *map[string]map[string]bool = indexedwords.GetIndexedWords()

	return (*indexedWords)[strings.ToLower(searchKey)]
}
