package main

import (
	"fmt"
	"math"
)

func main() {
	var pi64 = math.Pi
	fmt.Println(pi64) // Prints the value of Pi with 64-bit precision
    var pi32 float32 = math.Pi
    fmt.Println(pi32) // Prints the value of Pi with 32-bit precision
    third := 1.0 / 3
    fmt.Println(third)  // Outputs: 0.3333333333333333
    fmt.Printf("%f\n", third)  // Outputs: 0.333333
    fmt.Printf("%.3f\n", third)  // Outputs: 0.333
    fmt.Printf("%4.2f\n", third)  // Outputs: 0.33
    fmt.Println(third)         // Outputs: 0.3333333333333333
    fmt.Printf("%v\n", third)  // Outputs: 0.3333333333333333
    fmt.Printf("%6.2f\n", third)     // Outputs:  0.33 (width is 6, precision is 2)
    fmt.Printf("%10.2f\n", third)    // Outputs:      0.33 (width is 10, precision is 2)
    fmt.Printf("%04.2f\n", third)    // Outputs: 00.33 (width is 4, precision is 2, padded with zeros)
    celsius := 21.0
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Print(fahrenheit, "º F")  // Outputs: 69.8

}