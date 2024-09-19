package main

import (
	"fmt"
	"math/rand"
)
var era = "AD"  // Wide scope variable


func main() {
	// var command = "walk outside the cave"
	// var walkOutside = strings.Contains(command, "outside")

	// fmt.Println("You find yourself in a dimly lit cavern.")
	// fmt.Println("Do you walk outside the cave?")

	// if walkOutside {
	//     fmt.Println("You step outside the cave into the bright sunlight.")
	// } else {
	//     fmt.Println("You stay in the cave, pondering your next move.")
	// }
	// fmt.Println("There is a sign near the entrance that reads 'No Minors'.")

	// var age = 41
	// var minor = age < 18

	// fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	// var command = "go east"
	// //comparisons can be made between strings based on lexicographical (alphabetical) order, where strings are compared character by character based on their Unicode values
	// if command == "go east" {
	//     fmt.Println("You head further up the mountain.")
	// } else if command == "go inside" {
	//     fmt.Println("You enter the cave where you live out the rest of your life.")
	// } else {
	//     fmt.Println("Didn't quite get that.")
	// }
	// var room = "cave"

	// if room == "cave" {
	//     fmt.Println("You find yourself in a dimly lit cavern.")
	// } else if room == "entrance" {
	//     fmt.Println("There is a cavern entrance here and a path to the east.")
	// } else if room == "mountain" {
	//     fmt.Println("There is a cliff here. A path leads west down the mountain.")
	// } else {
	//     fmt.Println("Everything is white.")
	// }
	//  fmt.Println("The year is 2100, should you leap?")

	// var year = 2100
	// // Determine if it's a leap year
	// var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	// if leap {
	//     fmt.Println("Look before you leap!")
	// } else {
	//     fmt.Println("Keep your feet on the ground.")
	// }
	// fmt.Println("There is a cavern entrance here and a path to the east.")
	// var command = "go inside"
	// switch command {
	// case "go east":
	// 	fmt.Println("You head further up the mountain.")
	// case "enter cave", "go inside":
	// 	fmt.Println("You find yourself in a dimly lit cavern.")
	// case "read sign":
	// 	fmt.Println("The sign reads 'No Minors'.")
	// default:
	// 	fmt.Println("Didn't quite get that.")
	// }
	// var count = 10
	// for count > 0 {
	//     fmt.Println(count)
	//     time.Sleep(time.Second)
	//     count--
	// }
	// fmt.Println("Liftoff!")
	// var degrees = 0
	// for {
	//     fmt.Println(degrees)
	//     degrees++
	//     if degrees >= 360 {
	//         degrees = 0
	//         if rand.Intn(2) == 0 {
	//             break
	//         }
	//     }
	// }
	// var globalVar = "I am global"

	// if true {
	//     var localVar = "I am local"
	//     fmt.Println(globalVar) // Accessible here
	//     fmt.Println(localVar)  // Accessible here
	// }

	// fmt.Println(globalVar) // Accessible here
	// // fmt.Println(localVar) // Not accessible here, will cause a compilation error
	// var count = 0 // Function scope

	// for count < 10 {
	// 	var num = rand.Intn(10) + 1 // Block scope inside the for loop
	// 	fmt.Println(num)
	// 	count++
	// }
	// num is not accessible here
	var count = 10 // Using var keyword
	// count := 10    // Short declaration
	fmt.Println(count)
	// Short declaration within a for loop
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}
	// count is out of scope here
	year := 2018  // Narrow scope variable

	switch month := rand.Intn(12) + 1; month {
	case 2:
		day := rand.Intn(28) + 1  // Narrow scope variable
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1  // Narrow scope variable
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1  // Narrow scope variable
		fmt.Println(era, year, month, day)
	}

}
