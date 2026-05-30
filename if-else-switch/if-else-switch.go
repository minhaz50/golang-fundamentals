package main 

import "fmt"

func main () {
	// example of if-else
	// number := 2
	// if (number > 5){
	// 	fmt.Println("Number is greater then 5")
	// }else{
	// 	fmt.Println("Number is less then 5")
	// }

	// example of if,else-if, else statement
	/*
	age := 17

	if (age > 18) {
		fmt.Println("You are eligible for married.")
	}else if (age < 18) {
		fmt.Println("You are not eligible for married.")
	} else if (age == 18) {
		fmt.Println("You are just a teenager, not eligible for married.")
	}
	*/

	day := 3

	switch day {
	case 1:
		fmt.Println("sunday")
	case 2: 
		fmt.Println("Monday")
	case 3: 
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	default: 
		fmt.Println("Another day")

	}

}