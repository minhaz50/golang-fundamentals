package main
 
import "fmt"

func main() {
	var name string = "Minha"
	price := 20.40
	var numbers = [3] int {10,20, 30}
	values := []int{1,2,3}
	var isActive bool = true
	fmt.Println(isActive,name,numbers,values)
	fmt.Printf("Name: %s, Numbers: %d, price: %.2f, isActive: %t\n", name, numbers, price, isActive)
	fmt.Printf("Type: %T\n",values)
}