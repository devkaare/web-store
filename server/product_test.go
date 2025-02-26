package server

import "github.com/devkaare/web-store/model"

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"strings"
// 	"testing"
// )

var testProduct = model.Product{
	Name:      "testshirt",
	Price:     10,
	Sizes:     []byte(`{"sizes": {"small", "medium", "large", "extra_large"}}`),
	ImagePath: "./views/assets/images/shirt.png",
}

// func TestCreateUser(t *testing.T) {
// 	setup()
//
// 	apiUrl := fmt.Sprintf("http://localhost:%d", port)
// 	resource := "/users/"
// 	rawData := url.Values{}
// 	rawData.Set("email", "willsmithspersonalemail@gmail.com")
// 	rawData.Set("password", "supersecret123")
//
// 	u, _ := url.ParseRequestURI(apiUrl)
// 	u.Path = resource
// 	urlStr := u.String()
//
// 	req, err := http.NewRequest("POST", urlStr, strings.NewReader(rawData.Encode()))
// 	if err != nil {
// 		t.Fatalf("TestCreateUser: %v", err)
// 	}
//
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//
// 	r.ServeHTTP(respRec, req)
//
// 	if respRec.Result().StatusCode != http.StatusOK {
// 		t.Fatalf("TestCreateUser: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
// 	}
// }
//
// func TestGetUserByUserID(t *testing.T) {
// 	setup()
//
// 	req, err = http.NewRequest("GET", fmt.Sprintf("/users/%d", testUserID), nil)
// 	if err != nil {
// 		t.Fatalf("TestGetUserByUserID: %v", err)
// 	}
//
// 	r.ServeHTTP(respRec, req)
//
// 	if respRec.Result().StatusCode != http.StatusOK {
// 		t.Fatalf("TestGetUserByUserID: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
// 	}
//
// 	result := respRec.Result().Body
// 	data, err := io.ReadAll(result)
// 	if err != nil {
// 		t.Fatalf("TestGetUserByUserID: %v", err)
// 	}
//
// 	fmt.Printf("[+] Successfully found user: %v\n", string(data))
// }
//
// func TestUpdateUser(t *testing.T) {
// 	setup()
//
// 	apiUrl := fmt.Sprintf("http://localhost:%d", port)
// 	resource := fmt.Sprintf("/users/%d", testUserID)
// 	rawData := url.Values{}
// 	rawData.Set("email", "coolwillsmith@gmail.com")
// 	rawData.Set("password", "supersecret123")
//
// 	u, _ := url.ParseRequestURI(apiUrl)
// 	u.Path = resource
// 	urlStr := u.String()
//
// 	req, err := http.NewRequest("PUT", urlStr, strings.NewReader(rawData.Encode()))
// 	if err != nil {
// 		t.Fatalf("TestUpdateUser: %v", err)
// 	}
//
// 	// req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//
// 	r.ServeHTTP(respRec, req)
//
// 	if respRec.Result().StatusCode != http.StatusOK {
// 		t.Fatalf("TestUpdateUser: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
// 	}
// }
//
// func TestGetUsers(t *testing.T) {
// 	setup()
//
// 	req, err = http.NewRequest("GET", "/users", nil)
// 	if err != nil {
// 		t.Fatalf("TestGetUsers: %v", err)
// 	}
//
// 	r.ServeHTTP(respRec, req)
//
// 	if respRec.Result().StatusCode != http.StatusOK {
// 		t.Fatalf("TestGetUsers: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
// 	}
//
// 	result := respRec.Result().Body
// 	data, err := io.ReadAll(result)
// 	if err != nil {
// 		t.Fatalf("TestGetUsers: %v", err)
// 	}
//
// 	fmt.Printf("[+] Successfully got users: %v\n", string(data))
// }
//
// func TestDeleteUser(t *testing.T) {
// 	setup()
//
// 	req, err = http.NewRequest("DELETE", fmt.Sprintf("/users/%d", testUserID), nil)
// 	if err != nil {
// 		t.Fatalf("TestDeleteUser: %v", err)
// 	}
//
// 	r.ServeHTTP(respRec, req)
//
// 	if respRec.Result().StatusCode != http.StatusOK {
// 		t.Fatalf("TestDeleteUser: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
// 	}
// }
