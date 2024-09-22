
package main

import "fmt"
import "os"

//______________________________________________
// Go Function 
//		Takes No Arguments and Doens't Return Anything
func helloWorld() {
	fmt.Println("Hello World!!!")
}

//______________________________________________

// Constant
const boildingF = 212.0
var ff = 99
var cc = 88

func fToC(f float64) float64 {
	return ( f  - 32 ) *  5 / 9 
}

func playWithConstantAndVariables() {
	//Variables
	var f  = boildingF
	var c  = ( f  - 32 ) * 5 / 9 

	fmt.Printf("\nBoiling Point: %g Farenheit", f)
	fmt.Printf("\nBoiling Point: %g Centrigrade", c)

	var boildingC = fToC( boildingF )
	fmt.Printf("\nBoiling Point: %g Centrigrade", boildingC)

	const first, second = 99.99, 88.88
	fmt.Printf("\nFirst : %g, Second : %g", first, second)

	// boildingF = 8888
	// first = 888
	fmt.Printf("\nVariables Outside: %d %d",ff, cc)
}

//______________________________________________

type Celsius  	  float64
type Fahrenheit   float64

const (
	BoilingC  			Celsius = 100.0
	FreezingC 			Celsius = 0.0
	AbosulteZeroC 		Celsius = -273.15
	FreezingF 			Fahrenheit = 0.05
)

func playWithTypes() {
	// var ff = 11.11
	// var gg = 22.22
	var ff, gg = 11.11, 22.22
	var rr = ff + gg
	fmt.Println("Result: ", rr )

	var ffAgain Celsius = 11.11
	var ggAgain Fahrenheit = 22.22

	fmt.Println("Values Again: ", ffAgain, ggAgain )
	// Compilation Error
	// invalid operation: ffAgain + ggAgain (mismatched types Celsius and Fahrenheit)
	// var rrAgain = ffAgain + ggAgain
	var rrAgain = float64( ffAgain )  + float64( ggAgain )
	fmt.Println("Result: ", rrAgain)
	
	var someInt int64  = 99
	var someFloat float64 = 11.11
	// In Go Lang
	//		Most Of The Places Implicit Type Conversion Doesn't Allowed	
	// Compilation Error
	// invalid operation: someInt + someFloat (mismatched types int64 and float64)
	// var something = someInt + someFloat
	// In Go Lang: Explicit Type Conversion To Be Done
	var something = float64( someInt ) + someFloat
	fmt.Println( "Explicit Conversion: ", something )
}

func FToC( f Fahrenheit ) Celsius {
	return Celsius( ( f - 32 ) * 5 / 9 )
}

func CToF( c Celsius ) Fahrenheit {
	return Fahrenheit( c * 9 / 5 + 32 )
}

func playWithTypesAgain() {
	var result = FToC( FreezingF )
	fmt.Println("Result : ", result )

	var resultAgain = CToF( FreezingC )	
	fmt.Println("resultAgain : ", resultAgain )

	var resultOnceAgain = FToC( 212.15 )
	fmt.Println( resultOnceAgain )
}


//______________________________________________

func playWithCommandLineArguments() {
	var someString, seperator string 	
	seperator = "\n"
	var args = os.Args
	
	for i := 0 ; i < len( args ) ; i++ {
		someString =  someString + seperator + args[i]
	}

	fmt.Println( someString )

}

//______________________________________________


//______________________________________________
//______________________________________________
//______________________________________________

func main() {
	fmt.Println("\nFunction: helloWorld")
	helloWorld()

	fmt.Println("\nFunction: playWithConstantAndVariables")
	playWithConstantAndVariables()

	fmt.Println("\nFunction: playWithTypes")
	playWithTypes()

	fmt.Println("\nFunction: playWithTypesAgain")
	playWithTypesAgain()

	fmt.Println("\nFunction: playWithCommandLineArguments")
	playWithCommandLineArguments()

	// fmt.Println("\nFunction: ")
	// fmt.Println("\nFunction: ")
	// fmt.Println("\nFunction: ")
	// fmt.Println("\nFunction: ")
}

// https://pkg.go.dev/
