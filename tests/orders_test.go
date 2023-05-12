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
