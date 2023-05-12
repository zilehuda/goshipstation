package shipstation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ProductsResponse represents the response structure for retrieving all products.
type ProductsResponse struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	Pages    int       `json:"pages"`
}

// Product represents a single product.
type Product struct {
	Aliases                 interface{}     `json:"aliases"`
	ProductID               int             `json:"productId"`
	SKU                     string          `json:"sku"`
	Name                    string          `json:"name"`
	Price                   float64         `json:"price"`
	DefaultCost             float64         `json:"defaultCost"`
	Length                  int             `json:"length"`
	Width                   int             `json:"width"`
	Height                  int             `json:"height"`
	WeightOz                int             `json:"weightOz"`
	InternalNotes           interface{}     `json:"internalNotes"`
	FulfillmentSKU          string          `json:"fulfillmentSku"`
	CreateDate              string          `json:"createDate"`
	ModifyDate              string          `json:"modifyDate"`
	Active                  bool            `json:"active"`
	ProductCategory         ProductCategory `json:"productCategory"`
	ProductType             interface{}     `json:"productType"`
	WarehouseLocation       string          `json:"warehouseLocation"`
	DefaultCarrierCode      string          `json:"defaultCarrierCode"`
	DefaultServiceCode      string          `json:"defaultServiceCode"`
	DefaultPackageCode      string          `json:"defaultPackageCode"`
	DefaultIntlCarrierCode  string          `json:"defaultIntlCarrierCode"`
	DefaultIntlServiceCode  string          `json:"defaultIntlServiceCode"`
	DefaultIntlPackageCode  string          `json:"defaultIntlPackageCode"`
	DefaultConfirmation     string          `json:"defaultConfirmation"`
	DefaultIntlConfirmation string          `json:"defaultIntlConfirmation"`
	CustomsDescription      interface{}     `json:"customsDescription"`
	CustomsValue            interface{}     `json:"customsValue"`
	CustomsTariffNo         interface{}     `json:"customsTariffNo"`
	CustomsCountryCode      interface{}     `json:"customsCountryCode"`
	NoCustoms               interface{}     `json:"noCustoms"`
	Tags                    []Tag           `json:"tags"`
}

// ProductCategory represents the category of a product.
type ProductCategory struct {
	CategoryID int    `json:"categoryId"`
	Name       string `json:"name"`
}

// Tag represents a product tag.
type Tag struct {
	TagID int    `json:"tagId"`
	Name  string `json:"name"`
}

// UpdateProductResponse represents the response structure for updating a product.
type UpdateProductResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *ShipStation) GetProducts() (*ProductsResponse, error) {
	url := fmt.Sprintf("%s/products", s.baseURL)
	resp, err := s.sendRequest("GET", url, "retrieve products", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to retrieve products. Status code: %d and error message: %s", resp.StatusCode, string(body))
	}

	// Parse the response
	var productsResponse ProductsResponse
	err = json.NewDecoder(resp.Body).Decode(&productsResponse)
	if err != nil {
		return nil, err
	}

	return &productsResponse, nil
}

func (s *ShipStation) GetProduct(productID int) (*Product, error) {
	url := fmt.Sprintf("%s/products/%d", s.baseURL, productID)

	resp, err := s.sendRequest("GET", url, "retrieve product", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to retrieve product. Status code: %d and error message: %s", resp.StatusCode, string(body))
	}

	// Parse the response
	var product Product
	err = json.NewDecoder(resp.Body).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *ShipStation) UpdateProduct(productID int, updateData []byte) (*UpdateProductResponse, error) {
	url := fmt.Sprintf("%s/products/%d", s.baseURL, productID)

	resp, err := s.sendRequest("PUT", url, "update product", updateData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to update product. Status code: %d and error message: %s", resp.StatusCode, string(body))
	}

	// Parse the response
	var updateProductResponse UpdateProductResponse
	err = json.NewDecoder(resp.Body).Decode(&updateProductResponse)
	if err != nil {
		return nil, err
	}

	return &updateProductResponse, nil
}
