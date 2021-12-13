package main

import (
	"fmt"
)

func main() {

	var number int = parseNumber()
	var time string = formatTime(number)

	fmt.Printf(time)
}

func formatTime(number int) string {
	var numberHours int = number / 60
	var numberMinutes int = number % 60
	var time = fmt.Sprintf("It is %d hours %d minutes.", numberHours, numberMinutes)

	return time
}

func parseNumber() int {
	var number int

	fmt.Scan(&number)

	return number
}
