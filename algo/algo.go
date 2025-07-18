package algo

import (
	"fmt"
	"strings"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6}
	result := EvenNumberCount(numbers)
	fmt.Printf("Number of even numbers: %d \n", result)

	words := []string{"go", "is", "fun", "go", "go", "fun"}
	wordCount := wordCount(words)
	fmt.Println(wordCount)

	firstWord := "silent"
	secondWord := "listen"
	answer := CheckIfAnagrams(firstWord, secondWord)
	fmt.Printf("The words are anagrams: %t\n", answer)
}

func EvenNumberCount(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	evenNumbers := 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			evenNumbers = evenNumbers + 1
		}
	}
	return evenNumbers
}

func wordCount(words []string) map[string]int {
	wordCount := make(map[string]int)

	for number, word := range words {
		if _, exists := wordCount[word]; exists {
			number++
		} else {
			number = 1
		}
	}
	return wordCount
}

func CheckIfAnagrams(firstWord, secondWord string) bool {
	//check they are the same length, otherwise not anagrams
	if len(firstWord) != len(secondWord) {
		return false
	}

	//make sure everything is lower case
	firstWord = strings.ToLower(firstWord)
	secondWord = strings.ToLower(secondWord)

	//make a map of the letters and their occurence in the first word
	firstWordMap := make(map[rune]int)
	for value, letter := range firstWord {
		if _, exists := firstWordMap[letter]; exists {
			value++
		} else {
			value = 1
		}
	}

	//make a map of the letters and their occurence in the second word
	secondWordMap := make(map[rune]int)
	for value, letter := range secondWord {
		if _, exists := secondWordMap[letter]; exists {
			value++
		} else {
			value = 1
		}
	}

	//compare the maps
	if len(firstWordMap) != len(secondWordMap) {
		return false
	}

	for index, value1 := range firstWordMap {
		value2, exists := secondWordMap[index]
		if !exists || value1 != value2 {
			return false
		}
	}
	return true
}