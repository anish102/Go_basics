package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	name := person{firstName: "Anish", lastName: "Karki"}
	fmt.Println(name)
}
