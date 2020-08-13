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

}

const q0 = 0 // there is no input
const q1 = 1

func automata(text string) int {
	var wordcount = 0

	var q = q0
	for _, char := range text {

		switch q {
		case q0:
			if char == ' ' {
				wordcount++
				q = q1
			}
		case q1:
			if char == ' ' {
				q = q0
			}
		}
	}

	fmt.Printf("number of words in the string is %d", wordcount)
	return wordcount
}
