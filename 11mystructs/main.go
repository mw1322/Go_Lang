package main

import "fmt"

func main() {
	var manish = User{"Manish", "waliamanish122@gmail.com", true, 16}

	fmt.Println(manish)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}