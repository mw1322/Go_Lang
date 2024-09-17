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
 
	fmt.Println(highScores)
	var courses = []string{"reactjs","javascript","swift","python","ruby"};
	fmt.Println(courses);
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}