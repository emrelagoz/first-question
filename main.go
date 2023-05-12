package main

import (
	"fmt"
	"sort"
	"strings"
)

type Word struct {
	Word  string
	Count int
}

func sortByA(words []string) []string {
	wordList2 := make([]Word, len(words))

	// kelimeleri ve "a" harfi sayılarını structta tutma
	for i, word := range words {
		countA := strings.Count(word, "a")
		wordList2[i] = Word{
			Word:  word,
			Count: countA,
		}
	}

	// kaç tane "a" varsa ona göre sıralama eşitse ona göre sıralama yapılan yer
	sort.Slice(wordList2, func(i, j int) bool {
		if wordList2[i].Count == wordList2[j].Count {
			return len(wordList2[i].Word) > len(wordList2[j].Word)
		}
		return wordList2[i].Count > wordList2[j].Count
	})

	// sıralanmıslari gosterme yapılıyor
	sortedWords := make([]string, len(words))
	for i, word := range wordList2 {
		sortedWords[i] = word.Word
	}

	return sortedWords
}

func main() {
	wordList := []string{"aasd", "a", "aab", "aaaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	sortedWords := sortByA(wordList)
	fmt.Println(sortedWords)
}
