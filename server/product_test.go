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

var testProduct = &model.Product{
	Name:      "testshirt",
	Price:     10,
	Sizes:     `["small", "medium", "large", "extra large"]`,
	ImagePath: "",
}

func TestCreateProduct(t *testing.T) {
	setup()

	rawData := url.Values{}

	rawData.Add("name", testProduct.Name)
	rawData.Add("price", fmt.Sprintf("%d", testProduct.Price))
	rawData.Add("sizes", testProduct.Sizes)
	rawData.Add("image_path", testProduct.ImagePath)
	rawData.Add("api_key", testApiKey)

	urlStr := fmt.Sprintf("http://localhost:%d/products?%v", port, rawData.Encode())

	req, err := http.NewRequest("POST", urlStr, nil)
	if err != nil {
		t.Fatalf("TestCreateProduct: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(respRec, req)

	if respRec.Result().StatusCode != http.StatusOK {
		t.Fatalf("TestCreateProduct: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}

	var result model.Product

	d := json.NewDecoder(respRec.Result().Body)
	if err := d.Decode(&result); err != nil {
		t.Fatalf("TestCreateProduct: %v", err)
	}

	testProduct.ProductID = result.ProductID
}

func TestGetProductByProductID(t *testing.T) {
	setup()

	urlStr := fmt.Sprintf("http://localhost:%d/products/%d", port, testProduct.ProductID)

	req, err = http.NewRequest("GET", urlStr, nil)
	if err != nil {
		t.Fatalf("TestGetProductByProductID: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Result().StatusCode != http.StatusOK {
		t.Fatalf("TestGetProductByProductID: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}

	result := respRec.Result().Body
	data, err := io.ReadAll(result)
	if err != nil {
		t.Fatalf("TestGetProductByProductID: %v", err)
	}

	fmt.Printf("[+] Successfully found product: %v\n", string(data))
}

func TestUpdateProduct(t *testing.T) {
	setup()

	rawData := url.Values{}

	rawData.Add("name", fmt.Sprintf("%s [SALE]", testProduct.Name))
	rawData.Add("price", fmt.Sprintf("%d", testProduct.Price-5))
	rawData.Add("sizes", testProduct.Sizes)
	rawData.Add("image_path", testProduct.ImagePath)

	rawData.Add("api_key", testApiKey)

	urlStr := fmt.Sprintf("http://localhost:%d/products/%d?%v", port, testProduct.ProductID, rawData.Encode())

	req, err := http.NewRequest("PUT", urlStr, nil)
	if err != nil {
		t.Fatalf("TestUpdateProduct: %v", err)
	}

	// req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(respRec, req)

	if respRec.Result().StatusCode != http.StatusOK {
		t.Fatalf("TestUpdateProduct: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}
}

func TestGetProducts(t *testing.T) {
	setup()

	urlStr := fmt.Sprintf("http://localhost:%d/products", port)

	req, err = http.NewRequest("GET", urlStr, nil)
	if err != nil {
		t.Fatalf("TestGetProducts: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Result().StatusCode != http.StatusOK {
		t.Fatalf("TestGetProducts: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}

	result := respRec.Result().Body
	data, err := io.ReadAll(result)
	if err != nil {
		t.Fatalf("TestGetProducts: %v", err)
	}

	fmt.Printf("[+] Successfully got products: %v\n", string(data))
}

func TestDeleteProduct(t *testing.T) {
	setup()

	rawData := url.Values{}

	rawData.Add("api_key", testApiKey)

	urlStr := fmt.Sprintf("http://localhost:%d/products/%d?%v", port, testProduct.ProductID, rawData.Encode())

	req, err = http.NewRequest("DELETE", urlStr, nil)
	if err != nil {
		t.Fatalf("TestDeleteProduct: %v", err)
	}

	r.ServeHTTP(respRec, req)

	if respRec.Result().StatusCode != http.StatusOK {
		t.Fatalf("TestDeleteProduct: \"expected: %v, received: %v\"", http.StatusOK, respRec.Code)
	}
}
