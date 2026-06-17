package main

import "fmt"

func main(){
	// boolean
	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	// string
	var name [3]string
	name[0] = "golang"
	fmt.Println(name)

	// to declare it in single line
	friends := [3]string {"Bob", "Alice", "Tom"}
	fmt.Println(friends) 

	// 2D arrays
	numbers := [2][2]int {{1,2},{3,4}}
	fmt.Println(numbers)

	
	/*
	Benifints of using arrays : 
	- fixed size, that is predictable
	- Memory optimization
	- Constant time access
	*/
}