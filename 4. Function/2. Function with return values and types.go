package main

import "fmt"

// Syntax:
// func functionName(parameter type) returnType {
//     // function body
//     return value
// }

// func add(num1 int, num2 int) int{
// 	sum := num1 + num2
// 	return sum
// }

// func main(){
// 	a := 10
// 	b := 20

// 	fmt.Println("The sum is :", add(a, b))
// }

// multiple return value:
func getNumbers(num1 int, num2 int) (int, int){
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func main(){
	a := 10
	b := 20

	p,q := getNumbers(a,b)
	fmt.Print(p, q)
}