package main

import (
	"fmt"
	"os"
)

func main() {

	myName := "Tim"

	for idx, char := range myName {
		fmt.Println(idx, "---", string(char))

	}

	//Assignment 2
	name := "Tim Oyelabi"
	userName := strings.Fields(strings.ToLower(name))
	firstName := string(userName[0][0])
	secondName := string(userName[1])

	email := firstName + secondName + "@01edu.net"
	fmt.Println(email)

}
