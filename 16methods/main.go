package main

import "fmt"

func main() {
	var manish = User{"Manish", "waliamanish122@gmail.com", true, 16}

	fmt.Println(manish)
	manish.GetStatus()
	manish.NewMail()
	fmt.Printf("%+v\n",manish)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active or not :",u.Status)
}
func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Your new Email id : ",u.Email)
}