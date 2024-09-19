package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to files in goLang")
	content := "Hello World from goLang"

	file,err := os.Create("./newFile.txt")

	if err != nil {
		panic(err)
	}

	length,err := io.WriteString(file,content)
	if err != nil {
		panic(err)
	}

    fmt.Println("Length is : ",length)
	defer file.Close()
	readFile("./newFile.txt")
}

func readFile(FileName string)  {
	dataByte , err := ioutil.ReadFile(FileName);
	if err != nil {
		panic(err)
	}
	fmt.Println("Text Data inside the file is \n",string(dataByte))
}

