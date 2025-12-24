package main

import "fmt"

func main(){

	// Using:  >, <, >=, <=, ==
	age := 12

	if age >= 18{
		fmt.Println("You are an adult and you can vote.")
	}else if age >= 12 {
		fmt.Println("You are a teenager and you cannot vote.")
	}else{
		fmt.Println("You are only a kid and you cannot vote.")
	}

	// Using : &&, ||, !

	gender := "male"
	age_ := 21
	gender = "female"
	age_ = 17

	if (gender=="male" && age_>=21){
		fmt.Println("You are an adult male and eligible to be married.")
	}else if(gender=="female" && age_>=18){
		fmt.Println("You are an adult female and eligible to be married.")
	}else{
		fmt.Println("You are not eligible to be married.")
	}

	isPassed := false

	if(!isPassed){  // '!' converts the bool value 
		fmt.Println("You are passed")
	}
}