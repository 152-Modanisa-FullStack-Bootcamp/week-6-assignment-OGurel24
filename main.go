package main

import (
	"fmt"
	"homework5/repository"
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
			service.GetSpecificUser(w, r, url[1:], repository.Data)
			return
		}
		service.GetAllUsers(w, r, repository.Data)
		return
	}

	// Put Redirection
	if r.Method == "PUT" {
		repository.Data = service.AddUser(w, r, url[1:], repository.Data)
		return
	}

	// Post Redirection
	if r.Method == "POST" {
		repository.Data = service.UpdateUser(w, r, url[1:], repository.Data)
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

