package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/*
		countries := []string{"Angola", "China", "Nigeria", "Ghana", "Seychelles"}
		fmt.Println("Task 1: Print the countries")
		fmt.Println(countries)

		countries = slices.Insert(countries, 0, "Turkey")
		countries = slices.Insert(countries, 1, "Austria")
		countries = slices.Insert(countries, 4, "Brazil")
		countries = slices.Insert(countries, 6, "Lesotho")

		fmt.Println("\nTask 2: Insert countries the countries")
		for _, country := range countries {
			fmt.Printf("%s\n", country)
		}

		fmt.Println("\nTask 3: Print the last and the middle countries")
		fmt.Println("Last country:", countries[len(countries)-1])
		fmt.Println("Middle country:", countries[len(countries)/2])

		fmt.Println("\nTask 4: Print the countries with more than 8 letters")
		for _, country := range countries {
			if len(country) > 8 {
				fmt.Printf("%s\n", country)
			}
		}
	*/
	// Standard Input and Args

	ourArgs := os.Args[1:]

	if len(ourArgs) != 2 {
		fmt.Println(errors.New("Error: Only exactly 2 strings requrired"))
		return
	}
	inputFileName := ourArgs[0]
	outputFileName := ourArgs[1]

	data, readErr := os.ReadFile(inputFileName)
	if readErr != nil {
		fmt.Printf("Error: Error reading file %v", readErr)
		return
	}
	inputFileContent := string(data)
	fmt.Println(inputFileContent)

	writeErr := os.WriteFile(outputFileName, []byte(inputFileContent), 0774)
	if writeErr != nil {
		fmt.Printf("Error: writing file %v", writeErr)
		return
	}
}
