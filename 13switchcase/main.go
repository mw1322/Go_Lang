package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang : ")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)
	// fmt.Println(time.Now())
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4")		
	case 5:
		fmt.Println("Dice value is 5")	
	case 6:
		fmt.Println("Dice value is 6")				
	}
}