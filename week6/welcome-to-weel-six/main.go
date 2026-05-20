package main

import (
	"fmt"
)

func main() {
	FirstFunction()
	fmt.Println(SecondFunction())
	Joiner("Study", "Practice", "Build")
	fmt.Println(RealJoiner("Study", "Practice", "Build"))

}

func FirstFunction() {
	fmt.Println("Hi")
}

// SecondFunction is a GodMode Powerful things that gives Millions
func SecondFunction() int {

	fmt.Println("Hi, I am the second function")
	return 5
}

func Joiner(s, t, u string) {
	concat := s + " -> " + t + " -> " + u
	//concat := s , " -> " , t , " -> " , u
	fmt.Println(concat)
}

func RealJoiner(s, t, u string) string {
	concat := s + " -> " + t + " -> " + u
	return concat
}
