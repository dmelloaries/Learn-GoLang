package main

import "fmt"

func main() {
	// Declaring an array
	var numbers [5]int

	// Assigning values to array elements
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// Printing the array
	fmt.Println("Array elements:", numbers)

	// Accessing elements by index
	fmt.Println("First element:", numbers[0])
	fmt.Println("Last element:", numbers[4])

	// Shorter array declaration
	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println("Names array:", names)

	// Array length
	fmt.Println("Length of names array:", len(names))

	// Iterating through an array with a for loop
	fmt.Println("Iterating through the array:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Iterating with range
	fmt.Println("Using range to iterate:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Multidimensional array
	matrix := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("2D Array (Matrix):", matrix)

	// Accessing elements in a 2D array
	fmt.Println("Element at (1, 1):", matrix[1][1])
}
