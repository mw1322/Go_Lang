package main

import "fmt"

func main() {
	var languages = make(map[string]string)

	languages["rb"] = "RUBY"
	languages["js"] = "JavaScript"
	for key,value := range languages {
		 fmt.Printf("For key %v , value is %v\n",key,value);
	}
	// fmt.Println(languages["rb"])
}