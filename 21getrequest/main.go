package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("welcome to get request in golang")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	responce, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	fmt.Println("Status code : ", responce.StatusCode)
	fmt.Println("Status code : ", responce.ContentLength)

	content, _ := ioutil.ReadAll(responce.Body)

	fmt.Println(string(content))

}
