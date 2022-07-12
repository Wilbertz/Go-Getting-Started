package main

import "fmt"

func main() {
	firstName := "Harald"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Tanja"
	fmt.Println(ptr, *ptr)
}
