package main

import "fmt"

const n = 10000000
const m = 10
var globalArray [n]int
var globalSum int

func main() {
	for i := 0; i < n; i++{
		globalArray[i] = 1
	}

	for i :=0; i < n; i++{
		go subsum(i * n/m)
	}
	subsum(0)

	fmt.Println("sum=", globalSum)

}

func subsum(start int) {
	for i := start; i < start + n/m; i++{
		globalSum  += globalArray[i]
	}

}