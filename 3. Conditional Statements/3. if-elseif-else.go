package main

import "fmt"

func main(){
	age := 12;

	if age > 18{
		fmt.Println("You are an adult and you can vote.");
	}else if age>10 {
		fmt.Println("You are a teenager and you cannot vote.");
	}else{
		fmt.Println("You are only a kid and you cannot vote.")
	}
}