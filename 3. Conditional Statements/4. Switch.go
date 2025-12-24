package main

import "fmt"

func main(){
	num := 6

	switch num {
	case 1: 
		fmt.Println("This is 1")
	case 2: 
		fmt.Println("This is 2")
	case 3:
		fmt.Println("This is 3")
	case 4,5: 
		fmt.Println("This is either 4 or 5")
	default: 
		fmt.Println("This number is not among 1 to 5")
	}
}