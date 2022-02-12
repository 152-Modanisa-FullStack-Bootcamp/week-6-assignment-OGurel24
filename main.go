package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

var data = []user{
	{ID: 1, Name: "onur", Balance: 999},
	{ID: 2, Name: "ugur", Balance: 10},
}

func main() {

	http.HandleFunc("/", getData)
	_ = http.ListenAndServe(":8080", nil)

}

func getData(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI

	if r.Method == "PUT" {
		addUser(w, r, url[1:])
		return
	}

	if r.Method == "POST" {
		updateUser(w, r, url[1:])
		return
	}

	//should work with only GET
	if r.Method != "GET" {
		w.WriteHeader(501)
		_, err := fmt.Fprintf(w, "Invalid request method")
		if err != nil {
			panic(err)
		}
		return
	}

	if url != "/" {
		getSpecificData(w, r, url[1:])
		return
	}

	users, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(users)
	if err != nil {
		panic(err)
	}
}

func getSpecificData(w http.ResponseWriter, r *http.Request, username string) {
	fmt.Println(username)
	for _, v := range data {
		if username == v.Name {
			user, _ := json.Marshal(v)
			_, err := w.Write(user)
			if err != nil {
				panic(err)
			}

			return
		}
	}
	w.WriteHeader(404)
	fmt.Fprintf(w, "User could not be found!")
	return
}

func addUser(w http.ResponseWriter, r *http.Request, username string) {
	for _, v := range data {
		if v.Name == username {
			w.WriteHeader(417)
			_, err := fmt.Fprintf(w, "User is already exist")
			if err != nil {
				panic(err)
			}

			return
		}
	}

	newUser := user{len(data), username, 24}
	data = append(data, newUser)
	_, err := fmt.Fprintf(w, "User is added")
	if err != nil {
		panic(err)
	}

	return
}

func updateUser(w http.ResponseWriter, r *http.Request, username string) {
	type updateInfo struct {
		Balance int `json:"balance"`
	}

	for i, v := range data {
		if v.Name == username {
			decoder := json.NewDecoder(r.Body)
			var t updateInfo
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			data[i].Balance += t.Balance
			_, err = fmt.Fprintf(w, "Balance is changed successfully!")
			if err != nil {
				panic(err)
			}
			return
		}
	}

	w.WriteHeader(404)
	_, err := fmt.Fprintf(w, "User could not be found!")
	if err != nil {
		panic(err)
	}
	return
}
