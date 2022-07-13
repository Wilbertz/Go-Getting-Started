package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	u.ID = 1
	u.FirstName = "Harald"
	u.LastName = "Wilbertz"
	fmt.Println(u)
	fmt.Println(u.FirstName)

	u2 := user{ID: 1, FirstName: "Arthur", LastName: "Dent"}
	fmt.Println(u2)
}
