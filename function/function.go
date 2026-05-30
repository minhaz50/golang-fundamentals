
// Introduction of Functions
// Example : Sum of two numbers

/*
package main

import "fmt"

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}

func main () {
	a := 10
	b := 20

	add(a,b)
	add(5,7)
}
*/

/*
// Function with Returns values and types

package main

import "fmt"

func add (num1 int, num2 int) int{
	sum := num1 + num2
	fmt.Println("This sum is inside funciton: ",sum)
	return sum
}

func main() {
	a := 10
	b := 20

	result := add(a,b)
	fmt.Println("main function return value ", result)
}
*/

package main

import "fmt"

func printWelcomeMessage(){
	fmt.Println("Welcome to the Application.")
}

func getUserName() string {
	var name string
	fmt.Println("Enter Your name: ")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int) {
	var num1 int
	var num2 int 
	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)

	return num1, num2
}

func add (num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func displayMessage (name string, sum int){
	fmt.Println("Hello, ", name)
	fmt.Println("Summation: ", sum)
} 

func goodByeMessage() {
	fmt.Println("Thank you for using the application")
	fmt.Println("Goodbye")
}

func factorial(n int) int {
	if (n == 0){
		return 1
	}
	return n * factorial(n-1)
}

func main(){
	// Print Welcome message
	printWelcomeMessage()

	// Get a name as input
	name := getUserName()

	// Get  number as input 
	num1, num2 := getTwoNumbers()

	// Add numbers 
	sum := add(num1, num2)

	// display result
	displayMessage(name, sum)

	// Print a goodby message
	goodByeMessage()

	fmt.Println(factorial(5))

}

