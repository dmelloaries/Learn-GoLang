package main

import "fmt"

func main() {
	// Declaring a slice
	var numbers []int

	// Appending elements to a slice
	numbers = append(numbers, 10, 20, 30)
	fmt.Println("Slice after appending elements:", numbers)

	// Creating a slice with values
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("Names slice:", names)

	// Accessing elements
	fmt.Println("First element:", names[0])

	// Slicing an array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4] // Includes index 1, excludes index 4
	fmt.Println("Slice of array (1:4):", slice)

	// Slice length and capacity
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice)) // Capacity counts till the end of the array

	// Modifying slices
	slice[0] = 100
	fmt.Println("Modified slice:", slice)
	fmt.Println("Original array after modification:", arr)

	// Using make to create a slice
	newSlice := make([]int, 3, 5) // Length 3, Capacity 5
	newSlice[0] = 1
	newSlice[1] = 2
	newSlice[2] = 3
	fmt.Println("Slice created with make:", newSlice)

	// Appending beyond capacity
	newSlice = append(newSlice, 4, 5, 6)
	fmt.Println("Slice after appending beyond capacity:", newSlice)

	// Iterating over a slice
	fmt.Println("Iterating through slice:")
	for index, value := range newSlice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Copying slices
	copySlice := make([]int, len(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("Copied slice:", copySlice)
}
