package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter the rating of Your Hotel : ");

	input , _ := reader.ReadString('\n');
	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),64);

	if err != nil{
		fmt.Println(err);
	} else {
		fmt.Println("New Rating of hotel : ", numRating + 1);
	}
}