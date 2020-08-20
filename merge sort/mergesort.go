package main

import "fmt"

func main() {
	array1 := [...]int{2,3,5}
	array2 := [...]int{8,6,9}
	var array3 [300]int
	fmt.Println("This program merges given two sorted arrays into one sorted array")
	var x=0
	var y=0
	var k=0
	for x<len(array1) && y<len(array2){
		if (array1[x] < array2[y]) {
			array3[k] = array1[x]
			x++
		}else{
			array3[k]=array2[y]
			y++
		}
		k++
	}
	if x>=len(array1){
		for y<len(array2){
			array3[k]=array2[y]
			y++
			k++
		}
	}

	if y>=len(array2){
		for x<len(array1){
			array3[k]=array1[x]
			x++
			k++
		}
	}
	fmt.Printf("Merged array is ")
	for x:=0; x<len(array1)+len(array2); x++{
		fmt.Printf("%d", array3[x])
	}
}