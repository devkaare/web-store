package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/devkaare/web-store/model"
)

var testUser = &model.User{
	Email:    "johndoe@gmail.com",
	Password: "strongpassword123",
}

func TestCreateUser(t *testing.T) {
	setup()

	apiUrl := fmt.Sprintf("http://localhost:%d", port)
	resource := "/users/"

	rawData := url.Values{}

	rawData.Set("email", fmt.Sprintf("updated%s", testUser.Email))
	rawData.Set("password", fmt.Sprintf("updated%s", testUser.Password))

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(rawData.Encode()))
	if err != nil {
		t.Fatalf("TestCreateUser: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(respRec, req)

	if respRec.Result().StatusCode != http.StatusOK {
		t.Fatalf("TestCreateUser: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}

	var result model.User

	d := json.NewDecoder(respRec.Result().Body)
	if err := d.Decode(&result); err != nil {
		t.Fatalf("TestCreateUser: %v", err)
	}

	testUser.UserID = result.UserID
}

func TestGetUserByUserID(t *testing.T) {
	setup()

	req, err = http.NewRequest("GET", fmt.Sprintf("/users/%d", testUser.UserID), nil)
	if err != nil {
		t.Fatalf("TestGetUserByUserID: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Result().StatusCode != http.StatusOK {
		t.Fatalf("TestGetUserByUserID: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}

	result := respRec.Result().Body
	data, err := io.ReadAll(result)
	if err != nil {
		t.Fatalf("TestGetUserByUserID: %v", err)
	}

	fmt.Printf("[+] Successfully found user: %v\n", string(data))
}

func TestUpdateUser(t *testing.T) {
	setup()

	apiUrl := fmt.Sprintf("http://localhost:%d", port)
	resource := fmt.Sprintf("/users/%d", testUser.UserID)
	rawData := url.Values{}
	rawData.Set("email", testUser.Email)
	rawData.Set("password", testUser.Password)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	req, err := http.NewRequest("PUT", urlStr, strings.NewReader(rawData.Encode()))
	if err != nil {
		t.Fatalf("TestUpdateUser: %v", err)
	}

	// req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(respRec, req)

	if respRec.Result().StatusCode != http.StatusOK {
		t.Fatalf("TestUpdateUser: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}
}

func TestGetUsers(t *testing.T) {
	setup()

	req, err = http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("TestGetUsers: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Result().StatusCode != http.StatusOK {
		t.Fatalf("TestGetUsers: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}

	result := respRec.Result().Body
	data, err := io.ReadAll(result)
	if err != nil {
		t.Fatalf("TestGetUsers: %v", err)
	}

	fmt.Printf("[+] Successfully got users: %v\n", string(data))
}

func TestDeleteUser(t *testing.T) {
	setup()

	req, err = http.NewRequest("DELETE", fmt.Sprintf("/users/%d", testUser.UserID), nil)
	if err != nil {
		t.Fatalf("TestDeleteUser: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Result().StatusCode != http.StatusOK {
		t.Fatalf("TestDeleteUser: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}
}
