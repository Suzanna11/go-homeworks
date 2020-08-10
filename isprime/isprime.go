package main

import "fmt"

func main() {
	var num int
	fmt.Println("This program checks if the number is prime number, or not")
	fmt.Println("Enter the number")
	n, err := fmt.Scan(&num)
	if n < 1 {
		fmt.Println("data input error", err)
		return
	}
	if !(num > 0) {
		fmt.Println("data input error, a natural number is required")
		return
	}

	isPrime := true
	if num == 0 || num == 1 {
		fmt.Printf("%d is not a prime number\n", num)
	} else {
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				fmt.Printf("%d is not a prime number\n", num)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf("%d is a prime number\n", num)
		}
	}
}