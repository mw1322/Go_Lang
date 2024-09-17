package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	var name string 
	fmt.Println(name);
	fmt.Printf("Variable is of type: %T \n",name)

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter Rating for Hotel : ")
	input , _ := reader.ReadString('\n');
	fmt.Println("Your hotel rating is : ",input);
	// fmt.Println("hello from go");
}