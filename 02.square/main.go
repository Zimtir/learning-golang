package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)

	var square = getSquare(number)

	fmt.Println(square)
}

func getSquare(number int) int {
	return number * number
}
