package main

import "fmt"

func main() {

	var userName string = "ooyelabi"
	var favouriteInt int = 100
	var phoneLength float64 = 6.39
	var isAfrican bool = true

	text := "My name is " + userName
	newNumber := 10 + 20 + favouriteInt
	newPhoneLength := 2 * phoneLength
	isNotAfrican := !isAfrican

	if phoneLength > 6.43 {
		fmt.Println("My phone is longer than the Apple Iphone 17 Pro Max.")
	} else if phoneLength < 6.43 {
		fmt.Println("The Apple Iphone 17 Pro Max is longer than my phone.")
	} else {
		fmt.Println("I think my phone is the Apple Iphone 17 Pro Max.")
	}

	fmt.Println(text)
	fmt.Println(newNumber)
	fmt.Println(newPhoneLength)
	fmt.Println(isNotAfrican)
}
