package main

import (
	"log"
	"net/http"
	"reddynn/controller"
)

func main() {
	http.HandleFunc("/", controller.Welcome)
	http.HandleFunc("/signup", controller.Signup)
	http.HandleFunc("/signin", controller.Signin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
