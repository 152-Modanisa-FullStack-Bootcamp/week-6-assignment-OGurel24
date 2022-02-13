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
	req, err := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	service.GetAllUsers(w, req)
	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	data, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var users []repository.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return
	}

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, 2, len(users))

}

func TestGetAllUsersWithPost(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	service.GetAllUsers(w, req)
	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)
	data, _ := ioutil.ReadAll(res.Body)

	fmt.Println("og")
	fmt.Println(string(data))
	fmt.Println(res.StatusCode)

	fmt.Println("og")

}

func TestGetSpecificUsersWithGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	service.GetSpecificUser(w, req, "onur")
	res := w.Result()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	data, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	if res.StatusCode != 200 {
		err.Error()
	}

	user := repository.User{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		return
	}

	assert.Equal(t, 999, user.Balance)
	assert.Equal(t, "onur", user.Name)

}
