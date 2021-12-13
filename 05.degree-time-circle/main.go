package main

import (
	"fmt"
)

func main() {

	var number int = parseNumber()
	var time string = formatDegree(number)

	fmt.Printf(time)
}

func formatDegree(number int) string {
	var minutes int = (number * 60) / 30
	var numberHours int = minutes / 60
	var numberMinutes int = minutes % 60
	var time = fmt.Sprintf("It is %d hours %d minutes.", numberHours, numberMinutes)

	return time
}

func parseNumber() int {
	var number int

	fmt.Scan(&number)

	return number
}
