package service

import (
	"encoding/json"
	"fmt"
	"homework5/service/repository"
	"net/http"
	"os"
)

type configs struct {
	InitialBalanceAmount int `json:"initialBalanceAmount"`
	MinimumBalanceAmount int `json:"minimumBalanceAmount"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Collect all users and handle error if there is any
	users, err := json.Marshal(repository.Data)
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
		user, _ := json.Marshal(repository.Data[index])
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
	newUser := repository.User{username, loadConfig().InitialBalanceAmount}
	repository.Data = append(repository.Data, newUser)
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

	if repository.Data[index].Balance+t.Balance > loadConfig().MinimumBalanceAmount {
		repository.Data[index].Balance += t.Balance
		_, err = fmt.Fprintf(w, "Balance is changed successfully!")
		if err != nil {
			w.WriteHeader(500)
		}
	}

	return
}

// Helper function detects if the user is existed or not
func isUserExist(username string) (int, bool) {

	for i, v := range repository.Data {
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
