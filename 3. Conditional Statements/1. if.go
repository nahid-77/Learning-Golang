package main

import "fmt"

func main(){
	age := 19

	if age > 18{
		fmt.Println("You can vote.")
	}

	name := "admin"

	if(name == "admin"){
		fmt.Println("You can enter the admin server.")
	}
}