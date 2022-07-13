package main

import (
	"fmt"
)

func main() {
	//u := models.User{
	//	ID:        2,
	//	FirstName: "Tricia",
	//	LastName:  "McMillan",
	//}
	//fmt.Println("Hello Playground")

	port := 3000
	startWebServer(port, 2)
}

func startWebServer(port int, numberOfRetries int) {
	fmt.Println("Starting server...")
	// Do something
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
}
