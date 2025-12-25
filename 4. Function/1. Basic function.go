package main

import "fmt"

// Syntax:
// func functionName(parameter type) {
//     // function body
// }

func add(num1 int, num2 int){
	sum := num1 + num2
	fmt.Println("Sum is :", sum)
}

func main(){
	a := 10
	b := 20
	
	add(a, b)
}