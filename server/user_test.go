package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/devkaare/web-store/model"
)

var testUser = &model.User{
	Email:    "johndoe@gmail.com",
	Password: "strongpassword123",
}

func TestCreateUser(t *testing.T) {
	setup()

	rawData := url.Values{}

	rawData.Add("email", testUser.Email)
	rawData.Add("password", testUser.Password)

	rawData.Add("api_key", testApiKey)

	urlStr := fmt.Sprintf("http://localhost:%d/users?%v", port, rawData.Encode())

	req, err := http.NewRequest("POST", urlStr, nil)
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

	urlStr := fmt.Sprintf("http://localhost:%d/users/%d", port, testUser.UserID)

	req, err = http.NewRequest("GET", urlStr, nil)
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

	rawData := url.Values{}

	rawData.Add("email", testUser.Email)
	rawData.Add("password", testUser.Password)

	rawData.Add("api_key", testApiKey)

	urlStr := fmt.Sprintf("http://localhost:%d/users/%d?%v", port, testUser.UserID, rawData.Encode())

	req, err := http.NewRequest("PUT", urlStr, nil)
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

	urlStr := fmt.Sprintf("http://localhost:%d/users", port)

	req, err = http.NewRequest("GET", urlStr, nil)
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

	rawData := url.Values{}

	rawData.Add("api_key", testApiKey)

	urlStr := fmt.Sprintf("http://localhost:%d/users/%d?%v", port, testUser.UserID, rawData.Encode())

	req, err = http.NewRequest("DELETE", urlStr, nil)
	if err != nil {
		t.Fatalf("TestDeleteUser: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Result().StatusCode != http.StatusOK {
		t.Fatalf("TestDeleteUser: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}
}
