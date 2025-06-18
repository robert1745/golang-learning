package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")

	// var name string = "Alice"
	// fmt.Println("The name of person is : ", name)

	var age int = 30
	fmt.Println("The age of person is : ", age)

	var height float64 = 5.7
	fmt.Println("The height of person is : ", height)

	var isEmployed bool = true
	fmt.Println("Is the person employed? :", isEmployed)

	var hobbies []string = []string{"reading", "hiking", "coding"}
	fmt.Println("The hobbies of person are : ", hobbies)

	var name = "shyam"
	fmt.Println("The name of person is : ", name)

	// This syntax is valid inside fxn only not in global scope
	// there you need to use var keyword
	// i := 10 // This will not work in global scope
	i := 10
	fmt.Println("The value of i is : ", i)
	fmt.Printf("This variable is of type : %T \n ", i)

}
