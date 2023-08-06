package main

import (
	"fmt"
)

func main() {
	var scan string
	fmt.Println("Some changes")
	fmt.Println("Print something")
	fmt.Scanln(&scan)
	fmt.Println("Your input was: " + scan)
}
