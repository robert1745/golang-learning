package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for pizza:")

	// comma ok || err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, your rating is:, ", input)
}
