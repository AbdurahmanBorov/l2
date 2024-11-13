package main

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func getAnagram(words []string) map[string][]string {
	myMap := make(map[string][]string)
	result := make(map[string][]string)

	wordsSet := make(map[string]struct{})

	for _, word := range words {
		word = strings.ToLower(word)
		if _, exists := wordsSet[word]; exists {
			continue
		}
		wordsSet[word] = struct{}{}
		sortedWord := sortString(word)

		myMap[sortedWord] = append(myMap[sortedWord], word)
	}

	for _, v := range myMap {
		if len(v) > 1 {
			result[v[0]] = v
		}
	}

	for _, v := range result {
		sort.Strings(v)
	}

	return result
}
