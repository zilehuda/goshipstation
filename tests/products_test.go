package shipstation_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	shipstation "github.com/zilehuda/goshipstation"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProducts(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request parameters if needed
		assert.Equal(t, "/products", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		// Return a sample response
		productResponse := shipstation.ProductsResponse{
			Total: 1,
			Page:  1,
			Pages: 0,
			Products: []shipstation.Product{
				{ProductID: 123, SKU: "SKU123"},
				{ProductID: 456, SKU: "SKU456"},
			},
		}

		responseData, _ := json.Marshal(productResponse)

		w.WriteHeader(http.StatusOK)
		w.Write(responseData)
	}))
	defer server.Close()

	// Create a new instance of the ShipStation client
	client := shipstation.NewShipStation("your-api-key", "your-api-secret")
	client.SetBaseURL(server.URL)

	// Call the GetProducts method
	productsResponse, err := client.GetProducts()

	// Assert the expected behavior
	assert.NoError(t, err)
	assert.NotNil(t, productsResponse)
	assert.Len(t, productsResponse.Products, 2)
	assert.Equal(t, 123, productsResponse.Products[0].ProductID)
	assert.Equal(t, "SKU123", productsResponse.Products[0].SKU)
	assert.Equal(t, 456, productsResponse.Products[1].ProductID)
	assert.Equal(t, "SKU456", productsResponse.Products[1].SKU)
}

func TestGetProduct(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request parameters if needed
		assert.Equal(t, "/products/123", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		// Return a sample response
		product := &shipstation.Product{ProductID: 123, SKU: "SKU123"}
		responseData, _ := json.Marshal(product)

		w.WriteHeader(http.StatusOK)
		w.Write(responseData)
	}))
	defer server.Close()

	// Create a new instance of the ShipStation client
	client := shipstation.NewShipStation("your-api-key", "your-api-secret")
	client.SetBaseURL(server.URL)

	// Call the GetProduct method
	product, err := client.GetProduct(123)

	// Assert the expected behavior
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, 123, product.ProductID)
	assert.Equal(t, "SKU123", product.SKU)
}

func TestUpdateProduct(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request parameters if needed
		assert.Equal(t, "/products/123", r.URL.Path)
		assert.Equal(t, "PUT", r.Method)

		// Verify the request payload if needed
		updateData, _ := ioutil.ReadAll(r.Body)
		assert.Equal(t, []byte(`{"name":"New Name"}`), updateData)

		// Return a sample response
		responseData := []byte(`{"message": "Product updated successfully"}`)

		w.WriteHeader(http.StatusOK)
		w.Write(responseData)
	}))
	defer server.Close()

	// Create a new instance of the ShipStation client
	client := shipstation.NewShipStation("your-api-key", "your-api-secret")
	client.SetBaseURL(server.URL)

	// Prepare the update data
	updateData := []byte(`{"name":"New Name"}`)

	// Call the UpdateProduct method
	updateResponse, err := client.UpdateProduct(123, updateData)

	// Assert the expected behavior
	assert.NoError(t, err)
	assert.NotNil(t, updateResponse)
	assert.Equal(t, "Product updated successfully", updateResponse.Message)
}
