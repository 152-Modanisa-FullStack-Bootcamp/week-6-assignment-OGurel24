package main

import (
	"fmt"
	"homework5/service"
	"net/http"
)

func main() {

	http.HandleFunc("/", mainController)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)

}

func mainController(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI

	// Get Redirection
	if r.Method == "GET" {
		if url != "/" {
			service.GetSpecificUser(w, r, url[1:])
			return
		}
		service.GetAllUsers(w, r)
		return
	}

	// Put Redirection
	if r.Method == "PUT" {
		service.AddUser(w, r, url[1:])
		return
	}

	// Post Redirection
	if r.Method == "POST" {
		service.UpdateUser(w, r, url[1:])
		return
	}

	// Not implemented, return 501
	w.WriteHeader(501)
	_, err := fmt.Fprintf(w, "Invalid request method")
	if err != nil {
		panic(err)
	}

	return
}
