package main

import "fmt"

// Data types:
//  1. int
//  2. float32
//  3. bool
//  4. string

func main(){

	var a int = 10;
	var b = 7;
	c := 9;

	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);

	var variable = "This is a variable";  // string
	variable = 99;       // number type can't be used as the variable is string type
	variable = "Hello";     // string
	fmt.Println(variable);

	n := 45;
	n := 78;   // Can't use ':' second time for upadating a variable 
	n = 78;       // Right approach : without using the ':'
	fmt.Println(n);
	
	const pI = 3.1416;
	pI = 3.14159;  // Can't re-assign cause pI is const 
	fmt.Println(pI);

	var x  = 66;
	fmt.Println(x);

	x = 67;
	fmt.Println(x);

	x = 71;  // It won't effect, because x is already printed

}