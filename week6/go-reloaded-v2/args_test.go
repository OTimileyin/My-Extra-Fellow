package main

import (
	"testing"
)

func TestParseArgs1(tester *testing.T) {

	input, output, err := ParseArgs([]string{"Tim", "Oye"})

	if err != nil {
		tester.Errorf("Test Failed: %s", err)
	}
	if input != "Tim" {
		tester.Errorf("Expected input is: %s", input)
	}
	if output != "Oye" {
		tester.Errorf("Expected output is: %s", output)
	}
}

// func TestParseArgs2(tester *testing.T) {

// 	input, output, err := ParseArgs([]string{"Tim", "Oye", "Timo"})

// 	if err != nil {
// 		tester.Errorf("Test Failed: %s", err)
// 	}
// 	if input != "Tim" {
// 		tester.Errorf("Expected input is: %s", input)
// 	}
// 	if output != "Oye" {
// 		tester.Errorf("Expected output is: %s", output)
// 	}

// }
