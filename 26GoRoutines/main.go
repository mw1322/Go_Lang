package main

import (
	"fmt"
	"time"
)

// func main() {
// 	go greeter("Hello")
// 	greeter("World")
// }

// func greeter(message string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(message, "iteration:", i)
// 	}
// }



type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}
func main() {
	// Creating 3 email objects with different dates
	emails := [3]email{
		{"Email 1", time.Date(2019, 11, 10, 0, 0, 0, 0, time.UTC)},  // Older than 2020
		{"Email 2", time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC)},    // Newer than 2020
		{"Email 3", time.Date(2018, 7, 20, 0, 0, 0, 0, time.UTC)},   // Older than 2020
	}

	// Check the age of the emails
	isOld := checkEmailAge(emails)

	// Print the results
	for i, email := range emails {
		fmt.Printf("Is '%s' old? %v\n", email.body, isOld[i])
	}
}