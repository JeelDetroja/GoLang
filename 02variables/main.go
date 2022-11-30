package main

import "fmt"

const LoginToken string = "kjndhcbvndqk1342" //Public

func main() {
	var username string = "Jeel"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type : %T \n", isLoggedin)

	var money int = 3443
	fmt.Println(money)
	fmt.Printf("Variable is of type : %T \n", money)

	var auto = "set auto variables"
	fmt.Println(auto)
	fmt.Printf("Variable is of type : %T \n", auto)

	withoutdeclare := 6677.74857
	fmt.Println(withoutdeclare)
	fmt.Printf("Variable is of type : %T \n", withoutdeclare)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)
}
