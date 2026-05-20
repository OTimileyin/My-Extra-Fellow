package main

import "testing"

func TestParseInputContent(tester *testing.T) {
	fileContent, err := ParseInputContent("text.txt")
	if err != nil {
		tester.Errorf("Test Failed: %s", err)
	}

	if fileContent != "Tim" {
		tester.Errorf("Expected, file Content is : %s", fileContent)
	}

}
