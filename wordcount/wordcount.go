package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This program counts words in the given string")
	fmt.Println("Enter the string")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	automata(text)

	var wordcount = 0
	for _, char := range text {
		if char == '\n' {
			wordcount++
		}
	}
	wordcount++ //last character
	fmt.Printf("number of words in the string is %d", wordcount)

}

const q0 = 0 // there is no input
const q1 = 1

func automata(text string) int {

	var wordcount int
	var q = q0
	for _, char := range text {

		switch q {
		case q0:
			if char == ' ' {
				q = q1
			}
		case q1:
			if char == ' ' {
				q = q0
			}
			wordcount++
		}
	}
	
	return wordcount
}
