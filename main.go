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
	wordWithCountList := make([]Word, len(words))

	// kelimeleri ve "a" harfi sayılarını structta tutma
	for i, word := range words {
		countA := strings.Count(word, "a")
		wordWithCountList[i] = Word{
			Word:  word,
			Count: countA,
		}
	}

	// kaç tane "a" varsa ona göre sıralama eşitse ona göre sıralama yapılan yer
	sort.Slice(wordWithCountList, func(i, j int) bool {
		if wordWithCountList[i].Count == wordWithCountList[j].Count {
			return len(wordWithCountList[i].Word) > len(wordWithCountList[j].Word)
		}
		return wordWithCountList[i].Count > wordWithCountList[j].Count
	})

	// sıralanmıslari gosterme yapılıyor
	sortedWords := make([]string, len(words))
	for i, word := range wordWithCountList {
		sortedWords[i] = word.Word
	}

	return sortedWords
}

func main() {
	wordList := []string{"aasd", "a", "aab", "aaaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	sortedWords := sortByA(wordList)
	fmt.Println(sortedWords)
}
