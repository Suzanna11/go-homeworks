package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("This program counts words in the given string")
	fmt.Println("Enter the string")
	

	reader := bufio.NewReader(os.Stdin)
     text, _ := reader.ReadString('\n')

	var wordcount int
	wordcount = 0

	for _, char := range text {
		if char == ' ' || char == '\n' || char == '\t' {
			wordcount++
		}

	}

	

	fmt.Printf("number of words in the string is %d", wordcount)

}
