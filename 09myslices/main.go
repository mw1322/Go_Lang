package main

import "fmt"

func main() {
	var fruitList = []string{"apple","grapes"};
	fruitList = append(fruitList, "Mango","Banana");
	fmt.Println(fruitList)
	highScores := make([]int,4)
	highScores[0] = 123
	highScores[1] = 231
	// highScores[5] = 202
 
	// fmt.Println(highScores)
	// var courses = []string{"reactjs","javascript","swift","python","ruby"};
	// fmt.Println(courses);
	// var index int = 2
	// courses = append(courses[:index],courses[index+1:]...)
	 // Creating a slice
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Slice:", numbers)        // [1 2 3 4 5]
    fmt.Println("Length:", len(numbers))  // 5
    fmt.Println("Capacity:", cap(numbers)) // 5

    // Appending to a slice
    numbers = append(numbers, 6)
    fmt.Println("Slice after append:", numbers) // [1 2 3 4 5 6]
    fmt.Println("Length:", len(numbers))  // 5
    fmt.Println("Capacity:", cap(numbers)) // 5
    // Slicing the slice
    subSlice := numbers[1:4]
    fmt.Println("Sub-slice:", subSlice)    // [2 3 4]

    // Copying the slice
    newSlice := make([]int, len(numbers))
    copy(newSlice, numbers)
    fmt.Println("Copied slice:", newSlice) // [1 2 3 4 5 6]
}