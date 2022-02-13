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

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestGetSpecificUsersWithGet(t *testing.T) {
	url := "http://localhost:8080/onur"
	method := "GET"


	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(body))
}
