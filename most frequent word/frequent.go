package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	wordCount(text)

}

func wordCount(str string) map[string]int {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word]++
		} else {
			counts[word] = 1

		}
		
	}

	return counts

}
