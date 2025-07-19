package main

import (
	"belajar-go/identities"
	"fmt"
)

func main() {
	var identitas identities.Identity
	name := "John Doe"
	address := "123 Main St"
	age := 30
	identitas.Identities(name, address, age)
	fmt.Println(identitas)
}
