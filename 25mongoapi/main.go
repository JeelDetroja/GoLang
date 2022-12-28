package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/jeeldetroja/GoLang/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("server is getting started....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listning at port 4000 ....")
}
