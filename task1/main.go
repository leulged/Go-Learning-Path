package main

import (
	"fmt"
)

// Validates if the score is between 0 and 100
func check(score int) bool {
	return score >= 0 && score <= 100
}

// Calculates the average score from a map of course scores
func calculateAverage(grades map[string]int) float64 {
	if len(grades) == 0 {
		return 0
	}

	var total int
	for _, score := range grades {
		total += score
	}

	return float64(total) / float64(len(grades))
}

func main() {
	var name string
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)

	var subjTaken int
	fmt.Println("Please enter number of courses you have taken:")
	fmt.Scanln(&subjTaken)

	studentInfo := make(map[string]int)

	for i := 0; i < subjTaken; i++ {
		var course string

		// Course name validation loop
		for {
			fmt.Println("Please enter the course name:")
			fmt.Scanln(&course)

			if course == "" {
				fmt.Println("Course name cannot be empty. Try again.")
				continue
			}

			if _, exists := studentInfo[course]; exists {
				fmt.Println("You already entered this course. Try a different one.")
				continue
			}

			break
		}

		// Score input and validation loop
		var score int
		fmt.Println("Please enter the score you got for", course)
		fmt.Scanln(&score)

		if !check(score) {
			fmt.Println("Invalid score! Score must be between 0 and 100.")
			valid := false
			for attempts := 0; attempts < 3; attempts++ {
				fmt.Println("Please enter a valid score between 0 and 100:")
				fmt.Scanln(&score)
				if check(score) {
					valid = true
					break
				} else {
					fmt.Println("Invalid input. Try again.")
				}
			}
			if !valid {
				fmt.Println("Too many invalid attempts. Skipping this course.")
				continue
			}
		}

		// Store course and score
		studentInfo[course] = score
	}

	// Output the results
	fmt.Printf("\nStudent Name: %s\n", name)
	fmt.Println("Grades:")
	for subject, score := range studentInfo {
		fmt.Printf("  %s: %d\n", subject, score)
	}

	avg := calculateAverage(studentInfo)
	fmt.Printf("Average grade: %.2f\n", avg)
}
