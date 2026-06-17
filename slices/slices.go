package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// usefull methods
func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println("uninitialized slice is nil : ",nums == nil)

	// we initialized size as 2. but we can use size more then 2. 
	var number = make([]int, 0, 5)
	fmt.Println("initialized slice is not nil : ",number == nil)
	// capacity -> maximum numbers of elements can fit
	
	number = append(number, 1)
	number = append(number, 2)
	number = append(number, 3)
	number = append(number, 4)
	number = append(number, 5)
	number = append(number, 6)
	fmt.Println(number)
	fmt.Println("It will show the capacity : ",cap(number))


	// slice operator

	var num  = []int {1,2,3,4,5,6,7,8}
	fmt.Println(num[2:])

	// slice compareing

	var num1 = []int {1,2}
	var num2 = []int {1,5}

	fmt.Println(slices.Equal(num1, num2))


}