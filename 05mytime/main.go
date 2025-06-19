package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to MyTime!")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("2006-01-02 15:04:05 Monday")) // fixed way of formatting
}
