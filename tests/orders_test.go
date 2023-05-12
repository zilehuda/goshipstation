package shipstation_test

import (
	"encoding/json"
	shipstation "github.com/zilehuda/goshipstation"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOrder(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request parameters if needed
		assert.Equal(t, "/orders/createorder", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		// Return a sample response
		order := &shipstation.Order{OrderID: 123, OrderNumber: "ABC123"}
		responseData, _ := json.Marshal(order)

		w.WriteHeader(http.StatusOK)
		w.Write(responseData)
	}))
	defer server.Close()

	// Create a new instance of the ShipStation client
	client := shipstation.NewShipStation("your-api-key", "your-api-secret")
	client.SetBaseURL(server.URL)

	// Create a sample order payload
	orderData := []byte(`{"orderNumber": "ABC123"}`)

	// Call the AddOrder method
	order, err := client.AddOrder(orderData)

	// Assert the expected behavior
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, 123, order.OrderID)
	assert.Equal(t, "ABC123", order.OrderNumber)
}

func TestGetOrder(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request parameters if needed
		assert.Equal(t, "/orders/123", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		// Return a sample response
		order := &shipstation.Order{OrderID: 123, OrderNumber: "ABC123"}
		responseData, _ := json.Marshal(order)

		w.WriteHeader(http.StatusOK)
		w.Write(responseData)
	}))
	defer server.Close()

	// Create a new instance of the ShipStation client
	client := shipstation.NewShipStation("your-api-key", "your-api-secret")
	client.SetBaseURL(server.URL)

	// Call the GetOrder method
	order, err := client.GetOrder(123)

	// Assert the expected behavior
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, 123, order.OrderID)
	assert.Equal(t, "ABC123", order.OrderNumber)
}

func TestGetOrders(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request parameters if needed
		assert.Equal(t, "/orders", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		// Return a sample response
		orderResponse := shipstation.OrdersResponse{
			Page:  1,
			Total: 1,
			Orders: []shipstation.Order{
				{OrderID: 123, OrderNumber: "ABC123"},
				{OrderID: 456, OrderNumber: "DEF456"},
			},
		}

		responseData, _ := json.Marshal(orderResponse)

		w.WriteHeader(http.StatusOK)
		w.Write(responseData)
	}))
	defer server.Close()

	// Create a new instance of the ShipStation client
	client := shipstation.NewShipStation("your-api-key", "your-api-secret")
	client.SetBaseURL(server.URL)

	// Call the GetOrders method
	ordersResponse, err := client.GetOrders()

	// Assert the expected behavior
	assert.NoError(t, err)
	assert.NotNil(t, ordersResponse)
	assert.Len(t, ordersResponse.Orders, 2)
	assert.Equal(t, 123, ordersResponse.Orders[0].OrderID)
	assert.Equal(t, "ABC123", ordersResponse.Orders[0].OrderNumber)
	assert.Equal(t, 456, ordersResponse.Orders[1].OrderID)
	assert.Equal(t, "DEF456", ordersResponse.Orders[1].OrderNumber)
}

func TestDeleteOrder(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request parameters if needed
		assert.Equal(t, "/orders/123", r.URL.Path)
		assert.Equal(t, "DELETE", r.Method)

		// Return a sample response
		responseData := []byte(`{"message": "Order deleted successfully"}`)

		w.WriteHeader(http.StatusOK)
		w.Write(responseData)
	}))
	defer server.Close()

	// Create a new instance of the ShipStation client
	client := shipstation.NewShipStation("your-api-key", "your-api-secret")
	client.SetBaseURL(server.URL)

	// Call the DeleteOrder method
	deleteResponse, err := client.DeleteOrder(123)

	// Assert the expected behavior
	assert.NoError(t, err)
	assert.NotNil(t, deleteResponse)
	assert.Equal(t, "Order deleted successfully", deleteResponse.Message)
}
