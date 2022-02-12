package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func init() {
	loadConfig()
}

type user struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type configs struct {
	InitialBalanceAmount int `json:"initialBalanceAmount"`
	MinimumBalanceAmount int `json:"minimumBalanceAmount"`
}

var data = []user{
	{Name: "onur", Balance: 999},
	{Name: "ugur", Balance: 10},
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Collect all users and handle error if there is any
	users, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
	}

	// Return all users and handle error if there is any
	_, err = w.Write(users)
	if err != nil {
		w.WriteHeader(500)
	}
}

func GetSpecificUser(w http.ResponseWriter, r *http.Request, username string) {
	index, exist := isUserExist(username)

	if exist {
		user, _ := json.Marshal(data[index])
		_, err := w.Write(user)
		if err != nil {
			w.WriteHeader(500)
		}

		return
	}

	w.WriteHeader(404)
	_, err := fmt.Fprintf(w, "User could not be found!")
	if err != nil {
		w.WriteHeader(500)
	}
	return
}

func AddUser(w http.ResponseWriter, r *http.Request, username string) {
	_, exist := isUserExist(username)

	// If user is already exist return 417
	if exist {
		w.WriteHeader(417)
		_, err := fmt.Fprintf(w, "User is already exist")
		if err != nil {
			w.WriteHeader(500)
		}
	}

	// Otherwise add user
	newUser := user{username, loadConfig().InitialBalanceAmount}
	data = append(data, newUser)
	_, err := fmt.Fprintf(w, "User is added")
	if err != nil {
		w.WriteHeader(500)
	}

	return
}

func UpdateUser(w http.ResponseWriter, r *http.Request, username string) {
	index, exist := isUserExist(username)

	//If user is not exist return 404
	if !exist {
		w.WriteHeader(404)
		_, err := fmt.Fprintf(w, "User could not be found!")
		if err != nil {
			w.WriteHeader(500)
		}

		return
	}

	// Otherwise create a balance update info struct and update user balance
	type updateInfo struct {
		Balance int `json:"balance"`
	}

	decoder := json.NewDecoder(r.Body)
	var t updateInfo
	err := decoder.Decode(&t)

	if err != nil {
		w.WriteHeader(500)
	}

	if data[index].Balance+t.Balance > loadConfig().MinimumBalanceAmount {
		data[index].Balance += t.Balance
		_, err = fmt.Fprintf(w, "Balance is changed successfully!")
		if err != nil {
			w.WriteHeader(500)
		}
	}

	return
}

// Helper function detects if the user is existed or not
func isUserExist(username string) (int, bool) {

	for i, v := range data {
		if v.Name == username {
			return i, true
		}
	}
	return -1, false
}

func loadConfig() configs {

	configData, err := os.ReadFile(".config/local.json")
	if err != nil {
		panic(err)
	}

	var currentConfigs configs
	err = json.Unmarshal(configData, &currentConfigs)

	return currentConfigs
}
