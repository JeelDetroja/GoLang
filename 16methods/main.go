package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")
	// no inheritace in golang; no super or parent

	jeel := User{"Jeel", "jeel@gmail.com", true, 20}
	fmt.Println(jeel)
	fmt.Printf("Jeel details are: %+v\n", jeel)
	fmt.Printf("Name : %v && Email : %v\n", jeel.Name, jeel.Email)
	jeel.GetStatus()
	jeel.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test.com"
	fmt.Println("Email of this user is: ", u.Email)
}
