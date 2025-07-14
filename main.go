package main

import (
	"fmt"
	"strings"
	"regexp"
	"bufio"
	"os"
)

type WordsFreq struct {
	wordsFreq map[string]int
}

var storeWords = WordsFreq{
	wordsFreq: make(map[string]int),
}

func wordFreqCount(inputWord string) map[string]int {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	tempWords := reg.ReplaceAllString(inputWord, " ")
	tempWords = strings.ToLower(tempWords)

	words := strings.Fields(tempWords)
	for _, word := range words {
		storeWords.wordsFreq[word]++
	}
	return storeWords.wordsFreq
}

func palindCheck(inputWord string) bool {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	tempWords := reg.ReplaceAllString(inputWord, "")
	tempWords = strings.ToLower(tempWords)

	for i := 0; i < len(tempWords)/2; i++ {
		if tempWords[i] != tempWords[len(tempWords)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose one, A for Word Count, B for Palindrome Check: ")
	userChoice, _ := reader.ReadString('\n')
	userChoice = strings.TrimSpace(userChoice)

	for userChoice != "a" && userChoice != "b" && userChoice != "A" && userChoice != "B" {
		fmt.Print("Invalid choice. Please enter A or B: ")
		userChoice, _ = reader.ReadString('\n')
		userChoice = strings.TrimSpace(userChoice)
	}

	fmt.Print("Enter your text, press Enter to finish: ")
	inputWord, _ := reader.ReadString('\n')
	inputWord = strings.TrimSpace(inputWord)

	switch strings.ToUpper(userChoice) {
	case "A":
		wordCount := wordFreqCount(inputWord)
		fmt.Println("Word frequency count:")
		for word, count := range wordCount {
			fmt.Printf("%s: %d\n", word, count)
		}
	case "B":
		isPalindrome := palindCheck(inputWord)
		fmt.Printf("Is '%s' a palindrome? %t\n", inputWord, isPalindrome)
	}
}