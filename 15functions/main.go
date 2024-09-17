package main

import "fmt"

func main() {
	values := []int{1,2,3,4,5,6}
	proRes,myMessage := proAdder(values...)
	fmt.Println(proRes)
	fmt.Println(myMessage)
	fmt.Println(adder(23, 34))
}

func proAdder(values ...int) (int,string) {
	total := 0
	for _,val := range values {
		total += val
	}
	return total , "Hi pro result function"
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}