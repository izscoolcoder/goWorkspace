package main

import (
	"fmt"
	mymath "main/cool/cool"
)

func main() {
	var strings [5]string

	strings[0] = "Apple"
	strings[1] = "Orange"
	strings[2] = "Banana"
	strings[3] = "Grape"
	strings[4] = "Plum"

	fmt.Printf("\n=> Iterate over array\n")
	for i, fruit := range strings {
		fmt.Println(i, fruit)
	}

	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}

	fmt.Println(mymath.Add(64, 32))
}
