package main

import "fmt"

func main() {
	var firstName *string = new(string)
	*firstName = "Harald"
	fmt.Print(*firstName)
}
