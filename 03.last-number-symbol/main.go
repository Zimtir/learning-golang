package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number = parseNumber()
	var lastNumber = getLastNumber(number)

	fmt.Print(lastNumber)
}

func parseNumber() int {
	var number int

	fmt.Scan(&number)

	return number
}

func getLastNumber(number int) int {
	var stringNumber = strconv.Itoa(number)
	var stringLastNumber = string(stringNumber[len(stringNumber)-1])
	var lastNumber, _ = strconv.Atoi(stringLastNumber)

	return lastNumber
}
