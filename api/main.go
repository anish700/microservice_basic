package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var MySigninKey = []byte(os.Getenv("SECRET_KEY"))

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(w, "Secret information from server")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
	
		}
	})
}
func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Printf("server is running at PORT :", "9001")
	handleRequests()
}
