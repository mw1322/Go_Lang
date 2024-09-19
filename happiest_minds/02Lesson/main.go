// My weight loss program.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
    // fmt.Print("My weight on the surface of Mars is ")
    // fmt.Print(149.0 * 0.3783) // Calculate weight on Mars
    // fmt.Print(" lbs, and I would be ")
    // fmt.Print(41 * 365 / 687) // Calculate age in Mars years
    // fmt.Print(" years old.")

    // const lightSpeed = 299792 // km/s (speed of light)
    // var distance = 56000000   // km (initial distance to Mars)
    
    // fmt.Println(distance / lightSpeed, "seconds") // Time for nearby Mars

    // distance = 401000000 // New distance when Mars is far from Earth
    // fmt.Println(distance / lightSpeed, "seconds") // Time for far Mars

    // fmt.Printf("%15v $%4v\n", "SpaceX", 94)
    // fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

    // const hoursPerDay, minutesPerHour = 24, 60

    //Group the variable
//     var (
//     distance = 56000000
//     speed = 100800
// )

    var num = rand.Intn(10) + 1 // Generates a random number between 1 and 10
    fmt.Println(num)
    
    num = rand.Intn(10) + 1     // Generates another random number between 1 and 10
    fmt.Println(num)
    var distance = rand.Intn(345000001) + 56000000 // Random distance from 56,000,000 to 401,000,000 km
    fmt.Println(distance, "km")

}


