package main

import (
	"fmt"
)

// Multiply x2
// Increment +100

func main() {
	fmt.Println("Enter a number")

	var number = parseNumber()
	var functions = []func(number int) int{multiply, increment}

	var composition = compose(functions, number)

	fmt.Println(composition)
}

func compose(functions []func(number int) int, number int) int {
	for _, function := range functions {
		number = function(number)
	}

	return number
}

func multiply(number int) int {
	return number * 2
}

func increment(number int) int {
	return number + 100
}

func parseNumber() int {
	var input int

	fmt.Scan(&input)

	return input
}
