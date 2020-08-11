package main

import "fmt"

func main() {
	fmt.Println("This program counts words in the given string")
	fmt.Println("Enter the string")
	var input string

	fmt.Scanf(" %[^\t\n]s", &input)

	var wordcount int
	wordcount = 0

	for _, char := range input {
		if char == ' ' || char == '\n' || char == '\t' {
			wordcount++
		}

	}

	wordcount++ // last character

	fmt.Printf("number of words in the string is %d", wordcount)

}
