package main

import "fmt"

func printSomething(){
	fmt.Println("Something is being printed...")
}

func sayHello(name string){
	fmt.Println("Hello, It's me", name)
}

func main(){
	printSomething()
	sayHello("Nahid")
}