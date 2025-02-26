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

var (
	port       = 3000
	testUserID = 12
)

func checkUser(userID uint32) error {
	req, err = http.NewRequest("GET", "/users", nil)
	if err != nil {
		return fmt.Errorf("checkUser: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusOK {
		return fmt.Errorf("checkUser: expected: %d, received: %d", http.StatusOK, respRec.Code)
	}

	result := respRec.Result().Body
	data, err := io.ReadAll(result)

	if err != nil {
		return fmt.Errorf("checkUser: %v", err)
	}
	var users []model.User

	if err := json.Unmarshal(data, &users); err != nil {
		return fmt.Errorf("checkUser: %v", err)
	}

	for _, u := range users {
		if u.UserID == userID {
			fmt.Printf("Successfully found user: %v\n", u)
			return nil
		}
	}

	return fmt.Errorf("checkUser: user with user_id: %d does not exist", userID)
}

func TestCreateUser(t *testing.T) {
	setup()

	apiUrl := fmt.Sprintf("http://localhost:%d", port)
	resource := "/users/"
	rawData := url.Values{}
	rawData.Set("email", "willsmithspersonalemail@gmail.com")
	rawData.Set("password", "supersecret123")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	fmt.Println(strings.NewReader(rawData.Encode()))
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(rawData.Encode()))
	if err != nil {
		t.Fatalf("TestCreateUser: %v", err)
	}

	// req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(respRec, req)

	fmt.Println(respRec.Result().Status)

	TestGetUsers(t)
}

func TestGetUsers(t *testing.T) {
	setup()

	req, err = http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("TestGetUsers: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusOK {
		t.Fatalf("TestGetUsers: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}

	result := respRec.Result().Body
	data, err := io.ReadAll(result)
	if err != nil {
		t.Fatalf("TestGetUsers: %v", err)
	}

	fmt.Println(string(data))
}

func TestGetUserByUserID(t *testing.T) {
	setup()

	req, err = http.NewRequest("GET", fmt.Sprintf("/users/%d", testUserID), nil)
	if err != nil {
		t.Fatalf("TestGetUserByUserID: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusOK {
		t.Fatalf("TestGetUserByUserID: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}

	result := respRec.Result().Body
	data, err := io.ReadAll(result)
	if err != nil {
		t.Fatalf("TestGetUserByUserID: %v", err)
	}

	fmt.Println(string(data))
}

func TestUpdateUser(t *testing.T) {
	setup()

	apiUrl := fmt.Sprintf("http://localhost:%d", port)
	resource := fmt.Sprintf("/users/%d", testUserID)
	rawData := url.Values{}
	rawData.Set("email", "coolwillsmith@gmail.com")
	rawData.Set("password", "supersecret123")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	fmt.Println(strings.NewReader(rawData.Encode()))
	req, err := http.NewRequest("PUT", urlStr, strings.NewReader(rawData.Encode()))
	if err != nil {
		t.Fatalf("TestUpdateUser: %v", err)
	}

	// req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(respRec, req)

	fmt.Println(respRec.Result().Status)

	TestGetUsers(t)
}

func TestDeleteUser(t *testing.T) {
	setup()

	req, err = http.NewRequest("DELETE", fmt.Sprintf("/users/%d", testUserID), nil)
	if err != nil {
		t.Fatalf("TestDeleteUser: %v", err)
	}

	r.ServeHTTP(respRec, req)

	fmt.Println(respRec.Result().Status)

	TestGetUsers(t)
}
