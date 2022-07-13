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
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	// Do something
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	return port, nil
}
