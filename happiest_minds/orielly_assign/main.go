package main

import "fmt"

func half(n int) (int, bool) {
	halfValue := n / 2
	isEven := n%2 == 0
	return halfValue, isEven
}

func greatest(numbers ...int) int {
    if len(numbers) == 0 {
        panic("No numbers provided")
    }

    max := numbers[0]
    for _, num := range numbers[1:] {
        if num > max {
            max = num
        }
    }
    return max
}
func makeOddGenerator() func() int {
    i := -1
    return func() int {
        i += 2
        return i
    }
}
func fib(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }
    return fib(n-1) + fib(n-2)
}

func causePanic() {
    panic("Something went wrong!")
}

func riskyFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    
    fmt.Println("About to panic...")
    panic("Runtime panic")
    fmt.Println("This will not be executed")
}
func square(x *float64) *float64{
    *x = *x * *x
	return x
}
// Swap function that returns swapped values
func swap(x, y int) (int, int) {
    return y, x
}

func main() {
	// 2. Question : 
	halfValue, isEven := half(4)
	fmt.Println(halfValue,isEven)

	//3rd Question : 
	var arr = []int{1,3,5,4,3}
	fmt.Println(greatest(arr...))

	// 4. Question
	fmt.Println(makeOddGenerator()())

	// 5. Question
	fmt.Println(fib(0))  // Output: 0
	fmt.Println(fib(1))  // Output: 1
	fmt.Println(fib(5))  // Output: 5
	fmt.Println(fib(10)) // Output: 55

	// 6. Question
	// causePanic() // The program will stop and print "Something went wrong!"
	riskyFunction()
    fmt.Println("Program continues after recovery")

	// 7. Question
	// x := 42
	// p := &x // p holds the memory address of x
	// fmt.Println(p) // Output: the memory address of x

	// 8. Question
	// x := 42
	// p := &x   // p is a pointer to x
	// *p = 100  // assigns the value 100 to the memory location p points to
	// fmt.Println(x)  // Output: 100 (x is now updated to 100)

	// 9. Question
	// p := new(int)   // p is a pointer to a newly allocated int
	// fmt.Println(*p) // Output: 0 (the zero value for int)
	// *p = 100        // assign a value to the allocated memory
	// fmt.Println(*p) // Output: 100

	// 10. Question
	//  x := 1.5
    // fmt.Println(*square(&x))
	// 11. Question
	x := 1
    y := 2

    fmt.Println("Before swap: x =", x, ", y =", y)

    // Swap the values of x and y
    x, y = swap(x, y)

    fmt.Println("After swap: x =", x, ", y =", y)
	// 6. Question
	// 6. Question
	// 6. Question

}