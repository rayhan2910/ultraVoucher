package main

import (
	"fmt"
	"sort"
	"strings"
)

func filter(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func Group(arr []string) [][]string {
	anagrams := make(map[string][]string)

	for _, word := range arr {
		sortedWord := filter(word)
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	var result [][]string
	for _, v := range anagrams {
		result = append(result, v)
	}

	return result
}

func main() {
	arr := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	groups := Group(arr)
	for _, group := range groups {
		var groupWithQuotes []string
		for _, word := range group {
			groupWithQuotes = append(groupWithQuotes, fmt.Sprintf("'%s'", word))
		}
		fmt.Printf("[%s]\n", strings.Join(groupWithQuotes, ", "))
	}
}
