package main
 
import (
    "fmt"
	"strings"
)

func main() {
	text := "hello hh hh hello hello"
	for index,element := range wordCount(text){
		fmt.Println(index, "-", element )
	}
}

func wordCount(str string) map[string]int {
    text := strings.Fields(str)
    counts := make(map[string]int)
    for _, word := range text {
        _, ok := counts[word]
        if ok {
            counts[word] ++
        } else {
            counts[word] = 1
		}
	
	}
	
    return counts
}