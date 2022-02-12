package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var data = []int{1, 2, 3}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getData(w http.ResponseWriter, r *http.Request) {
	for _, value := range data {

		fmt.Fprintf(w, strconv.Itoa(value))
	}
}

func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/data", getData)
	_ = http.ListenAndServe(":8080", nil)

}
