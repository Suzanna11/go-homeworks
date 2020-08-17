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

	
	countingWords(text)

}

func countingWords(text string) bool {
	var wordcount = 0
	var state bool
	for _, char := range text {
		if char == ' ' || char == '\n' || char == '\t' {
			state = false
		}else if state == false {
				state = true
				wordcount++
			}
	}
	
	fmt.Printf("number of words in the string is %d", wordcount)
	return state
	}


