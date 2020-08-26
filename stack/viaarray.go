package stack

import "fmt"

var stack [100]int
var size = 100
var top = -1

func push(value int) bool {
	if top == size {
		fmt.Println("Stack is overflow")
		return false
	} else {
		top++
		stack[top] = value
		return true
	}

}

func pop() bool {
	if top == -1 {
		fmt.Println("Stack is empty")
		return false
	} else {
		return true
	}
}
