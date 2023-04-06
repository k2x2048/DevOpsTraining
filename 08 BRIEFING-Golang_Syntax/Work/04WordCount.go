package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

//strings.Count(str, substr)

func WordsOccurrences(Phrase string) map[string]int {
	Phrase = strings.Replace(Phrase,"."," ",-1)
	Phrase = strings.Replace(Phrase,","," ",-1)
	Phrase = strings.Replace(Phrase,";"," ",-1)
	Phrase = strings.Replace(Phrase,"'"," ",-1)
	Phrase = strings.Replace(Phrase,"\""," ",-1)
	
	//fmt.Println(Phrase)
	
	WordsCount := make(map[string]int)
	
	//Convert Phrase to array of words
	AllWordsArray := strings.Fields(Phrase)
	
	for _, word := range AllWordsArray {
		_, ExistInArray := WordsCount[word]
		if ExistInArray {
			WordsCount[word] += 1
		} else {
			WordsCount[word] = 1
		}
	}
	return WordsCount
}

func main() {
	var TextLine string
	fmt.Print("Write a phrase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	TextLine = scanner.Text()

	//Print Words occurences
	for Word, Occur := range WordsOccurrences(TextLine) {
		fmt.Println(Word,"x",Occur)
	}
}
