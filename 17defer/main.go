package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("Hello")
	myDefer()
}
//hello 43210 world
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}