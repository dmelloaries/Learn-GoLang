package main

import "fmt"

func main() {
	// Declaring a map
	var studentScores map[string]int

	// Initializing the map
	studentScores = make(map[string]int)

	// Adding key-value pairs
	studentScores["Alice"] = 90
	studentScores["Bob"] = 85
	studentScores["Charlie"] = 88

	// Printing the map
	fmt.Println("Student Scores:", studentScores)

	// Accessing elements
	fmt.Println("Alice's score:", studentScores["Alice"])

	// Checking if a key exists
	score, exists := studentScores["David"]
	if exists {
		fmt.Println("David's score:", score)
	} else {
		fmt.Println("David not found in the map")
	}

	// Shorter declaration
	countryCodes := map[string]string{
		"US": "United States",
		"IN": "India",
		"JP": "Japan",
	}
	fmt.Println("Country Codes:", countryCodes)

	// Updating a value
	studentScores["Bob"] = 95
	fmt.Println("Updated Scores:", studentScores)

	// Deleting a key-value pair
	delete(studentScores, "Charlie")
	fmt.Println("Scores after deletion:", studentScores)

	// Iterating over a map
	fmt.Println("Iterating through student scores:")
	for name, score := range studentScores {
		fmt.Printf("Name: %s, Score: %d\n", name, score)
	}

	// Map length
	fmt.Println("Number of students:", len(studentScores))
}
