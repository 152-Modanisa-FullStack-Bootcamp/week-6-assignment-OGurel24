package service_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"homework5/repository"
	"homework5/service"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllUsersWithGet(t *testing.T) {
	// Mock DB
	mockDB := make([]repository.User, len(repository.Data))
	copy(mockDB, repository.Data)

	//Send request
	req, err := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	service.GetAllUsers(w, req, mockDB)
	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	//Read response
	data, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	//Write returned data to variable
	var users []repository.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return
	}

	// Assert status code, response length and after whole data
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, 5, len(users))
	assert.Equal(t, mockDB, users)
}

func TestGetAllUsersWithPost(t *testing.T) {
	// Mock DB
	mockDB := make([]repository.User, len(repository.Data))
	copy(mockDB, repository.Data)

	//Send request
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	service.GetAllUsers(w, req, mockDB)
	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetSpecificUsersWithGet(t *testing.T) {
	// Mock DB
	mockDB := make([]repository.User, len(repository.Data))
	copy(mockDB, repository.Data)

	// Send request
	req := httptest.NewRequest(http.MethodGet, "/Onur", nil)
	w := httptest.NewRecorder()
	service.GetSpecificUser(w, req, "Onur", mockDB)
	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(res.Body)

	//Read response
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Write returned data to variable
	var user repository.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Assert status code, response length and after whole data
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, mockDB[0], user)
}

func TestAddUserSuccess(t *testing.T) {
	// Mock DB
	mockDB := make([]repository.User, len(repository.Data))
	copy(mockDB, repository.Data)
	mockDBInitialLength := len(mockDB)

	newUser := repository.User{
		Name: "onur2"}

	// Send request
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/%s", newUser.Name), nil)
	w := httptest.NewRecorder()
	mockDB = service.AddUser(w, req, newUser.Name, mockDB)
	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(res.Body)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "User is added", string(data))
	assert.Equal(t, mockDBInitialLength+1, len(mockDB))
	assert.Equal(t, mockDB[mockDBInitialLength].Name, newUser.Name)
	assert.Equal(t, mockDB[mockDBInitialLength].Balance, service.LoadConfig().InitialBalanceAmount)

}

func TestAddUserFail(t *testing.T) {
	//Tries to add exist user and gets 417

	// Mock DB
	mockDB := make([]repository.User, len(repository.Data))
	copy(mockDB, repository.Data)
	mockDBInitialLength := len(mockDB)

	newUser := repository.User{
		Name: "Onur"}

	// Send request
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/%s", newUser.Name), nil)
	w := httptest.NewRecorder()
	mockDB = service.AddUser(w, req, newUser.Name, mockDB)
	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(res.Body)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.Equal(t, 417, res.StatusCode)
	assert.Equal(t, "User is already exist", string(data))
	assert.Equal(t, mockDBInitialLength, len(mockDB))
}
