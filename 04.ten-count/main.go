package main

import (
	"fmt"
	"strconv"
)

func main() {

	var number int = parseNumber()
	var tenNumberCount int = getNumberTenCount(number)

	fmt.Printf(strconv.Itoa(tenNumberCount))
}

func parseNumber() int {
	var number int

	fmt.Scan(&number)

	return number
}

func getNumberTenCount(number int) int {
	var stringNumber = strconv.Itoa(number)

	if isOneDigitNumber(number) {
		return 0
	}

	if isTwoDigitStringNumber(stringNumber) {
		return parseTwoDigitTenCount(number)
	}

	return getPrelastStringNumber(stringNumber)
}

func isTwoDigitStringNumber(stringNumber string) bool {
	var numberLength int = len(stringNumber) - 1

	return numberLength == 1
}

func isOneDigitNumber(number int) bool {
	return number < 10
}

func parseTwoDigitTenCount(number int) int {
	return number / 10
}

func getPrelastStringNumber(stringNumber string) int {
	var prelastPosition = len(stringNumber) - 2
	var prelastNumber = string(stringNumber[prelastPosition])
	var value, _ = strconv.Atoi(prelastNumber)

	return value
}
