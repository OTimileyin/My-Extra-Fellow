package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "Tim, Oye"
	nameSlice := strings.Fields(name)

	fmt.Println(len(nameSlice))

}
