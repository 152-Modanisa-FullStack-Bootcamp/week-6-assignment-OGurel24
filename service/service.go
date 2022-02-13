package service

import (
	"encoding/json"
	"fmt"
	"homework5/repository"
	"net/http"
	"os"
)

type configs struct {
	InitialBalanceAmount int `json:"initialBalanceAmount"`
	MinimumBalanceAmount int `json:"minimumBalanceAmount"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request, data []repository.User) {
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

func GetSpecificUser(w http.ResponseWriter, r *http.Request, username string, data []repository.User) {
	index, exist := isUserExist(username, data)

	if exist {
		user, err := json.Marshal(data[index])
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write(user)
		return
	}

	w.WriteHeader(404)
	_, err := fmt.Fprintf(w, "User could not be found!")
	if err != nil {
		w.WriteHeader(500)
	}
	return
}

func AddUser(w http.ResponseWriter, r *http.Request, username string, data []repository.User) []repository.User {
	_, exist := isUserExist(username, data)

	// If user is already exist return 417
	if exist {
		w.WriteHeader(417)
		_, err := fmt.Fprintf(w, "User is already exist")
		if err != nil {
			w.WriteHeader(500)
		}
		return data
	}

	// Otherwise add user
	newUser := repository.User{username, LoadConfig().InitialBalanceAmount}
	data = append(data, newUser)
	_, err := fmt.Fprintf(w, "User is added")
	if err != nil {
		w.WriteHeader(500)
	}

	return data
}

func UpdateUser(w http.ResponseWriter, r *http.Request, username string, data []repository.User) []repository.User {
	index, exist := isUserExist(username, data)

	//If user is not exist return 404
	if !exist {
		w.WriteHeader(404)
		_, err := fmt.Fprintf(w, "User could not be found!")
		if err != nil {
			w.WriteHeader(500)
			return data
		}

		return data
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

	if data[index].Balance+t.Balance > LoadConfig().MinimumBalanceAmount {
		data[index].Balance += t.Balance
		_, err = fmt.Fprintf(w, "Balance is changed successfully!")
		if err != nil {
			w.WriteHeader(500)
			fmt.Println(err)
			return data
		}
	}

	return data
}

// Helper function detects if the user is existed or not
func isUserExist(username string, data []repository.User) (int, bool) {
	for i, v := range data {
		if v.Name == username {
			return i, true
		}
	}
	return -1, false
}

func LoadConfig() configs {

	configData, err := os.ReadFile(".config/local.json")
	wd, _ := os.Getwd()
	fmt.Println(24, wd, 24)
	if err != nil {
		panic(err)
	}

	var currentConfigs configs
	err = json.Unmarshal(configData, &currentConfigs)

	return currentConfigs
}
