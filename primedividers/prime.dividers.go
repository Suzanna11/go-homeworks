package main

import "fmt"


func main(){
	var num int
	n, err := fmt.Scan(&num)
	if n < 1 {
		fmt.Println("data input error", err)
		return
	}
	if !(num > 0) {
		fmt.Println("data input error, a natural number is required")
		return
	}
	if num == 1 {
		fmt.Println(1)
	}

	k:=nextPrime(1)
for; k <= num ; k = nextPrime(k){
	for; num%k == 0; num /= k{
		fmt.Print(k, "\t" )

	}
}


}
func isPrime(num int)bool{

	if num == 0 || num == 1 {
		return false
	}
	
	for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				return false
			}
		}	
		return true
	}

func nextPrime(num int)int {
	for num = num+1 ; isPrime(num) == false; num++ {
	}
	return num

}
