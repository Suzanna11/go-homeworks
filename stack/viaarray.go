package main

import "fmt"


var stack[100]int
var size = 100
var top = -1
func Push(value int)bool {
	if top == size {
		fmt.Println("Stack is overflow")
		return false
	}else{
		top++
		stack[top] = value
		return true
	}

	}

func Pop() (int, bool) {
	if top == -1 {
		fmt.Println("Stack is empty")
		return false
	}else{
		return stack[top--]
	}
}
	
