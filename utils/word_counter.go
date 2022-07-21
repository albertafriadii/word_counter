package utils

import (
	"log"
	"regexp"
	"strings"

	"github.com/detectlanguage/detectlanguage-go"
)

// to replace all symbol or not alphabet
func cleanText(text string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	checkError("Cannot clean text", err)

	proccessedWord := reg.ReplaceAllString(text, "")
	return proccessedWord
}

// checking error
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// to count the word
func WordCount(s string) map[string]int {
	count := make(map[string]int)
	client := detectlanguage.New("ad9372aeeb7754aadb0d59132d8ffd6b")

	detections, err := client.Detect(s)

	if err != nil {
		log.Fatal(err.Error())
	}

	confirmWord := detections[0].Language
	words := strings.Fields(s)

	for _, word := range words {
		cleanWord := cleanText(strings.ToLower(word))
		if len(words) < 3 || len(words) > 10000 {
			log.Fatal("Must more than 3 words or less than 10000 words")
		}

		if confirmWord == "en" {
			count[cleanWord]++
		}
	}

	return count
}
