package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://loc.dev"

func main() {
	fmt.Println("LOC web request")

	responce, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("responce type %T\n", responce)

	defer responce.Body.Close()

	databytes, err := ioutil.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
