package main

import "fmt"

func main() {
	var manish = User{"Manish", "waliamanish122@gmail.com", true, 16}

	 fmt.Printf("%+v\n", manish)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}