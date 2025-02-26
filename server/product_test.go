package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/devkaare/web-store/model"
)

var testProduct = model.Product{
	Name:      "testshirt",
	Price:     10,
	Sizes:     []byte(`{"sizes": {"Small", "Medium", "Large", "Extra Large"}}`),
	ImagePath: "./views/assets/images/shirt.png",
}

func TestCreateProduct(t *testing.T) {
	setup()

	apiUrl := fmt.Sprintf("http://localhost:%d", port)
	resource := "/products/"

	rawData := url.Values{}

	rawData.Set("name", testProduct.Name)
	rawData.Set("price", fmt.Sprintf("%d", testProduct.Price))
	rawData.Set("sizes", string(testProduct.Sizes))
	rawData.Set("imagePath", testProduct.ImagePath)

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
		t.Fatalf("TestCreateProduct: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}
}

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
