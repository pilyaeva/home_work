package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func Top10(sourceText string) []string {
	re := regexp.MustCompile(`\s+`)
	textWithoutExtraSpaces := re.ReplaceAllString(sourceText, " ")

	if textWithoutExtraSpaces == "" {
		return []string{}
	}

	words := strings.Split(textWithoutExtraSpaces, " ")
	repeats := make(map[string]int)

	for _, word := range words {
		repeats[word]++
	}

	var sortedData []kv

	for k, v := range repeats {
		sortedData = append(sortedData, kv{Key: k, Value: v})
	}

	sort.Slice(sortedData, func(i, j int) bool {
		if sortedData[i].Value == sortedData[j].Value {
			return sortedData[i].Key < sortedData[j].Key
		}

		return sortedData[i].Value > sortedData[j].Value
	})

	var mostFrequencyWordsCount int

	if len(sortedData) >= 10 {
		mostFrequencyWordsCount = 10
	} else {
		mostFrequencyWordsCount = len(sortedData)
	}

	mostFrequencyWords := make([]string, mostFrequencyWordsCount)

	for i := 0; i < mostFrequencyWordsCount; i++ {
		mostFrequencyWords[i] = sortedData[i].Key
	}

	return mostFrequencyWords
}
