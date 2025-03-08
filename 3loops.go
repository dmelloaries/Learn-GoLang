package main

import (
	"fmt"
	"time"
)

func main() {

	// Basic switch case
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week!")
	case "Friday":
		fmt.Println("Weekend is near!")
	default:
		fmt.Println("It's a regular day.")
	}

	// Switch without an expression (like an if-else ladder)
	num := 7
	switch {
	case num < 5:
		fmt.Println("Number is less than 5")
	case num >= 5 && num <= 10:
		fmt.Println("Number is between 5 and 10")
	default:
		fmt.Println("Number is greater than 10")
	}

	// Switch with multiple cases
	color := "red"
	switch color {
	case "red", "blue":
		fmt.Println("Primary color")
	case "green":
		fmt.Println("Secondary color")
	default:
		fmt.Println("Unknown color")
	}

	// Type switch
	var value interface{}
	value = 42

	switch v := value.(type) {
	case int:
		fmt.Printf("Value is an int: %d\n", v)
	case string:
		fmt.Printf("Value is a string: %s\n", v)
	default:
		fmt.Println("Unknown type")
	}

	// Switch with fallthrough
	switch day {
	case "Monday":
		fmt.Println("It's Monday")
		fallthrough
	case "Tuesday":
		fmt.Println("It's Tuesday (fallthrough from Monday)")
	default:
		fmt.Println("It's another day")
	}

	// Switch with time
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}
}
