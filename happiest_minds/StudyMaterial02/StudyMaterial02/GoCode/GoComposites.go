
package main 

import "fmt"

//__________________________________________________________

func playWithTypesAndDefaultValues() {
	// Following Types Are Initialised To Default Values Of 0
	var a int 
	var a8 int8
	var a16 int16
	var a32 int32
	var a64 int64

	aa := 99
	fmt.Println("Value Of aa: ", aa )
	fmt.Printf("Type Of aa: %T \n", aa )

	var ua uint 
	var ua8 uint8
	var ua16 uint16
	var ua32 uint32
	var ua64 uint64

	fmt.Println(a, a8, a16, a32, a64)
	fmt.Println(ua, ua8, ua16, ua32, ua64)

	var f1 float32
	var f2 float64
	fmt.Println( f1, f2 )

	ff := 99.90
	fmt.Println("Value Of ff: ", ff)
	fmt.Printf("Type Of ff: %T \n", ff)

	var some string
	fmt.Println( some )

	greeting := "Good Morning!"
	fmt.Println("Value Of greeting: ", greeting)
	fmt.Printf("Type Of greeting: %T \n", greeting)

	var b bool
	fmt.Println( b )

	var c1 complex64
	fmt.Println( c1 )

	var c2 complex128
	fmt.Println( c2 )
}
	

//__________________________________________________________

func playWithArrays() {
	var a[4]int 

	fmt.Println( a )
	// range Operator Gives
	//		List of Tuples Having 2 Members
	//		First Member Will Be Index
	//		Second Member Will Be Value
	for i, value := range a {
		fmt.Println( i, value )
	}

	fmt.Println( a[0] )
	fmt.Println( a[1] )
	fmt.Println( a[ len( a ) - 1 ] )

// invalid argument: index 4 out of bounds [0:4]
	// fmt.Println( a[ len( a ) ] )	

	// _ Is A Special Variable
	//		Used For Ignoring i.e. Which You Don't Wants To Use
	for _, value := range a {
		fmt.Println( value )
	}

	for index := range a {
		fmt.Println( index, a[ index ] )
	}
}

//__________________________________________________________

func playWithArraysAgain() {
	var q [5]int = [5]int{ 10, 20, 30, 40, 50 }
	var r [5]int = [5]int{ 10, 20, 30 }

	fmt.Println("Array Length q: ", len( q ) )
	fmt.Println("Array Length r: ", len( r ) )
	fmt.Printf("Array q Type: %T \n ", q )
	fmt.Printf("Array r Type: %T \n", r )

	for index, value := range q {
		fmt.Println("At Index: ", index, " Value: ", value)
	}

	for index, value := range r {
		fmt.Println("At Index: ", index, " Value: ", value)
	}

	// ... Dots Means Size Deduced From Initialisation List

	s := [...]int{10, 20, 30, 99, 100, 111}
	fmt.Println("Array s: ", s )
	fmt.Println("Array Length s: ", len( s ) )
	fmt.Printf("Array q Type: %T \n ", s )
//	fmt.Printf("Array r Type: %T \n", r )

	for index, value := range s {
		fmt.Println("At Index: ", index, " Value: ", value )
	}

	aa := [2]int{ 10, 20 }
	bb := [...]int{ 10, 20 }
	cc := [2]int{ 10, 30 }

	fmt.Println("aa Equals bb: ", aa == bb )	
	fmt.Println("aa Equals cc: ", aa == cc )
	fmt.Println("cc Equals bb: ", cc == bb )

	aaa := [2]int{ 10, 20 }
	bbb := [3]int{ 10, 20 }
	fmt.Println( aaa, bbb )
	// fmt.Println("aaa == bbb : ", aaa == bbb )

	someArray := [...]int{ 0: 11, 3: 33, 7: 77 }
	fmt.Println("someArray : ", someArray )
	fmt.Println("Length Of someArray: ", len( someArray) )
	fmt.Printf("Type Of someArray: %T \n", someArray )

	someArrayAgain := [...]int{ 99: -555 }
	fmt.Println("someArray : ", someArrayAgain )
	fmt.Println("Length Of someArray: ", len( someArrayAgain ) )
	fmt.Printf("Type Of someArray: %T \n", someArrayAgain )

	// Implicit Type Casting Happening
	var some [4]float64 = [4]float64{ 99.90, 88.8, 111, }
	fmt.Println( "some Array: ", some )

	var ii int = 111
	var ff float64 = 99.99
	fmt.Println( ii, ff )
//	var result = ii + ff

	// var someAgain [3]string = [3]string{ "Ding", "Dong", 111 }
	// fmt.Println( "someAgain Array: ", someAgain )	
}

//__________________________________________________________

// Array Elements Are Pass By Values
func changeArray( some [10]int ) {
	fmt.Println("changeArray Function Called...")
	for index, _ := range some {
		some[ index ] = 111
	}
}

// Array Itself Are Also Pass By Value
func changeArrayAgain( some [10]int ) {
	fmt.Println("changeArray Function Called...")
	some = [10]int{11, 11, 11, 11, 11, 11, 11, 11, 11, 11}
}

// Array Pass By Reference
func changeArrayOnceAgain( someAddress *[10]int ) {
	for index, _ := range someAddress {
		someAddress[ index ] = 111
	}
}

// Array Pass By Reference Using Slice
func changeArrayOnceMore( someSlice []int ) {
	for index, _ := range someSlice {
		someSlice[index] = 222
	}
}

func playWithArraysOnceAgain() {
	var aa [10]int = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println("Value Of aa: ", aa )
	changeArray( aa )
	fmt.Println("Value Of aa: ", aa )

	fmt.Println("Value Of aa: ", aa )
	changeArrayAgain( aa )
	fmt.Println("Value Of aa: ", aa )	

	fmt.Println("Value Of aa: ", aa )
	changeArrayOnceAgain( &aa ) 	// Passing Address Of Array aa
	fmt.Println("Value Of aa: ", aa )	

	fmt.Println("Value Of aa: ", aa )
	// Here aa[ : ] Slice Over Array aa
	changeArrayOnceMore( aa[ : ] ) // Passing Slice Of Array aa
	fmt.Println("Value Of aa: ", aa )		
}


//__________________________________________________________

func playWithArrayAndSlices() {
	a := [...]int{ 10, 20, 30, 40, 50, 60, 70, 80, 90, 99 }
	fmt.Println("Array a: ", a)

	some1 := a[ 0 : 4 ]
	fmt.Println("Slice some1: ", some1 )

	some2 := a[ 4 : 9 ]
	fmt.Println("Slice some2: ", some2 )

	some3 := a[ : 5 ]
	fmt.Println("Slice some3: ", some3 )

	some4 := a[ 6 : ]
	fmt.Println("Slice some4: ", some4 )

	some5 := a[ : ]
	fmt.Println("Slice some5: ", some5 )
}

//__________________________________________________________

// Function Taking Slice As Argument
func reverse( s []int ) {
	for i, j := 0, len( s ) - 1 ; i < j ; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

func playWithArrayAndSlicesAgain() {
	a := [...]int{ 10, 20, 30, 40, 50, 60, 70, 80, 90, 99 }
	fmt.Println("Array a: ", a)

	some1 := a[ 0 : 4 ]
	fmt.Println("Array a: ", a )	
	reverse( some1 )
	fmt.Println("Array a: ", a )

	some2 := a[ 4 : 9 ]
	fmt.Println("Array a: ", a )	
	reverse( some2 )
	fmt.Println("Array a: ", a )

	some3 := a[ : ]
	fmt.Println("Array a: ", a )	
	reverse( some3 )
	fmt.Println("Array a: ", a )
}

// Function: playWithArrayAndSlicesAgain
// Array a:  [10 20 30 40 50 60 70 80 90 99]
// Array a:  [10 20 30 40 50 60 70 80 90 99]
// Array a:  [40 30 20 10 50 60 70 80 90 99]
// Array a:  [40 30 20 10 50 60 70 80 90 99]
// Array a:  [40 30 20 10 90 80 70 60 50 99]

/*
	some5 := a[ : ]
	fmt.Println("Array a: ", a )	
	doSomething( some5 )
	fmt.Println("Array a: ", a )
*/

//__________________________________________________________

// Function Taking 2 Slices As Argument
func printCommon( summer []string, quater []string ) {
	for _, summerMonth := range summer {
		for _, quaterMonth := range quater {
			if summerMonth == quaterMonth {
				fmt.Printf("%s Month Appear In Both!", summerMonth)
			}
		}
	}
}

// Following Both Signatures Are Equivalent
// func slicesEqual( x []string, y []string ) bool {
func slicesEqual( x, y []string ) bool {
	if len( x ) != len( y ) {
		return false 
	}

	for index := range x {
		if x[index] != y[index] {
			return false
		}
	}

	return true
}

func playWithArrayAndSlicesOnceAgain() {
	months := [...]string {
		1: "January", 2: "February", 3: "March", 4: "April",
		5: "May", 6: "June", 7: "July", 8: "August",
		9: "September", 10: "October", 11: "November", 12: "December",
	}

	quater2 := months[ 4 : 7] 
	summer := months[ 6 : 9 ]
	fmt.Println( months )
	fmt.Println( quater2 )
	fmt.Println( summer )
	printCommon( summer, quater2 )

	firstThree := months[ 1 : 4   ]
	quater1 := months[ 1 : 4 ]

	// invalid operation: firstThree == quater1 (slice can only be compared to nil)
	// fmt.Println("Slices Are Equal: ", firstThree == quater1 ) 
	result := slicesEqual( firstThree, quater1 )
	fmt.Println("Slices Are Equal: ", result ) 
}
//________________________________________________________________
// 				GO SLICES CONCEPTS
//________________________________________________________________

// Slices represent variable-length sequences whose elements all have the same type. 
// 	A slice type is written []T, where the elements have type T; 
// 		it looks like an array type without a size.

// A slice is a lightweight data structure that gives access to a subsequence 
// (or perhaps all) of the elements of an array, which is known as the 
// slice’s underlying array. 

// 	A slice has three components: a pointer, a length, and a capacity. 
// 		The pointer points to the first element of the array that is 
//          reachable through the slice, which is not necessarily the array’s 
//          first element. 
// 		The length is the number of slice elements; it can’t exceed the capacity, 
// 			which is usually the number of elements between 
// 			the start of the slice and the end of the underlying array. 
// 		The built-in functions len and cap return those values.


// The slice operator s[ i : j ], where 0 ≤ i ≤ j ≤ cap(s), 
// 	Creates a new slice that refers to elements i through j-1 of the sequence s, 
// 	which may be an array variable, a pointer to an array, or anotherslice. 
// 	The resulting slice has j-i elements. 
// 	If i is omitted,it’s 0,and if j isomitted, it’s len(s). 

// 	Thus the slice months[1:13] refers to the whole range of valid months, 
// 	as does the slice months[1:]; the slice months[:] refers to the whole array.


//__________________________________________________________

func playWithSlicesOnceMore() {
// The built-in function make creates a slice of a specified element type, 
// length, and capacity. The capacity argument may be omitted, 
// in which case the capacity equals the length.
// 		make([]T, len)
// 		make([]T, len, cap) // same as make([]T, cap)[:len]
	
	s := make( []string, 3 )
	fmt.Printf("Slice Type: %T \n", s)

	fmt.Println("Slice: ", s)
	fmt.Println("Slice Length and Capacity: ", len(s), cap(s) )

	s[0] = "Ding"
	s[1] = "Dong"
	s[2] = "Ting"

	fmt.Println("Slice: ", s)
	fmt.Println("Slice Length and Capacity: ", len(s), cap(s) )

	s = append( s, "Tong" )
	fmt.Println("Slice: ", s)
	fmt.Println("Slice Length and Capacity: ", len(s), cap(s) )
	s = append( s, "Ming", "Mong" )	
	fmt.Println("Slice: ", s)
	fmt.Println("Slice Length and Capacity: ", len(s), cap(s) )
	s = append( s, "Modi")

	fmt.Println("Slice: ", s)
	fmt.Println("Slice Length and Capacity: ", len(s), cap(s) )
}

//__________________________________________________________

func playWithArrayAndSlicesConcepts() {
	numbers := [...]int{ 10, 20, 30, 40, 50, 60 }
	numbersCopy := numbers
	fmt.Println(" Numbers Length Capacity: ", len( numbers ), cap( numbers ) )
	fmt.Println(" NumbersCopy Length Capacity: ", len( numbersCopy ), cap( numbersCopy ) )

	fmt.Println( "Numbers: ", numbers )
	fmt.Println( "NumbersCopy: ", numbersCopy )
	numbers[ 0 ] = 99
	fmt.Println( "Numbers: ", numbers )
	fmt.Println( "NumbersCopy: ", numbersCopy )

	some := numbers[ 0 : 3 ]
	fmt.Println( "Some   : ", some )
	fmt.Println(" Some Length Capacity: ", len( some ), cap( some ) )

	some = append( some, 99 )
	fmt.Println( "Numbers: ", numbers )
	fmt.Println(" Numbers Length Capacity: ", len( numbers ), cap( numbers ) )
	fmt.Println( "Some   : ", some )
	fmt.Println(" Some Length Capacity: ", len( some ), cap( some ) )

	someCopy := some
	fmt.Println( "Some   : ", some )
	fmt.Println( "SomeCopy   : ", someCopy )
	some[0] = 777
	fmt.Println( "Some   : ", some )
	fmt.Println( "SomeCopy   : ", someCopy )

	//________________________________________

	// ss := [...]int{ 10, 20, 30, 40 }
	ss := make( []int, 4 )
	//ss = append(ss, 10, 20, 30, 40)
	ss[0] = 10
	ss[1] = 20
	ss[2] = 30
	ss[3] = 40
	cc := make( []int, len(ss) )

	fmt.Println("SS :", ss )
	fmt.Println("CC : ", cc )
	// cannot use ss (variable of type [4]int) as []int value in assignment
	// Shallow Copy i.e. Reference Assignment
	cc = ss
	fmt.Println("SS :", ss )
	fmt.Println("CC : ", cc )

	ss[3] = 777
	fmt.Println("SS :", ss )
	fmt.Println("CC : ", cc )

	dd := make( []int, len(ss ))
	// Deep Copy: Full Copy Of The Object Get Created
	copy( dd, ss )
	fmt.Println("SS :", ss )
	fmt.Println("DD : ", dd )

	ss[2] = 666
	fmt.Println("SS :", ss )
	fmt.Println("DD : ", dd )
	
}


//__________________________________________________________
//__________________________________________________________
//__________________________________________________________
//__________________________________________________________

func main() {
	fmt.Println("\nFunction: playWithTypesAndDefaultValues")
	playWithTypesAndDefaultValues()
	
	fmt.Println("\nFunction: playWithArrays")
	playWithArrays()

	fmt.Println("\nFunction: playWithArraysAgain")
	playWithArraysAgain()

	fmt.Println("\nFunction: playWithArraysOnceAgain")
	playWithArraysOnceAgain()

	fmt.Println("\nFunction: playWithArrayAndSlices")
	playWithArrayAndSlices()

	fmt.Println("\nFunction: playWithArrayAndSlicesAgain")
	playWithArrayAndSlicesAgain()

	fmt.Println("\nFunction: playWithArrayAndSlicesOnceAgain")
	playWithArrayAndSlicesOnceAgain()

	fmt.Println("\nFunction: playWithSlicesOnceMore")
	playWithSlicesOnceMore()

	fmt.Println("\nFunction: playWithArrayAndSlicesConcepts")
	playWithArrayAndSlicesConcepts()

	// fmt.Println("\nFunction: ")
	// fmt.Println("\nFunction: ")
	// fmt.Println("\nFunction: ")
}