package main

import "fmt"

func test() {
	type identity string
	var name identity = "Jane Doe"
	age := 25

	fmt.Println(name, age)
}
