package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"regexp"
)

func main() {
	http.HandleFunc("/square/", squareHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func squareHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("squareHandler called.")
	fmt.Println(r)

	array := make([]byte, r.ContentLength)
	r.Body.Read(array)
	fmt.Println("Body:", string(array))
	strBody := string(array)

	matched, err := regexp.MatchString(`[a-z]=[0-9]+`, r.Form.Get("word=number"))
	if matched==false {
		fmt.Println(err)
	}

	split := strings.Split(strBody, "=")
	number, _ := strconv.Atoi(split[1])
	number = number * number
	w.Write([]byte(strconv.Itoa(number)))
}
