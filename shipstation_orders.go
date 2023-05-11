package shipstation

import (
	"encoding/json"
	"fmt"
)

func (s *ShipStation) AddOrder(orderData []byte) (*Order, error) {
	url := fmt.Sprintf("%s/orders/createorder", s.baseURL)
	resp, err := s.sendRequest("POST", url, "create/update order", orderData)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
		return nil, err
	}

	// Parse the response
	var order Order
	err = json.NewDecoder(resp.Body).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *ShipStation) GetOrder(orderID int) (*Order, error) {
	url := fmt.Sprintf("%s/orders/%d", s.baseURL, orderID)

	resp, err := s.sendRequest("GET", url, "retrieve order", nil)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Parse the response
	var order Order
	err = json.NewDecoder(resp.Body).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// GetOrders retrieves an all orders from ShipStation
func (s *ShipStation) GetOrders() (*OrdersResponse, error) {
	url := fmt.Sprintf("%s/orders", s.baseURL)

	resp, err := s.sendRequest("GET", url, "retrieve orders", nil)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	// Parse the response
	var ordersResponse OrdersResponse
	err = json.NewDecoder(resp.Body).Decode(&ordersResponse)
	if err != nil {
		return nil, err
	}
	return &ordersResponse, nil
}

// DeleteOrder delete an order from ShipStation based on the order number.
func (s *ShipStation) DeleteOrder(orderID int) (*DeleteOrderResponse, error) {
	url := fmt.Sprintf("%s/orders/%d", s.baseURL, orderID)

	resp, err := s.sendRequest("DELETE", url, "retrieve order", nil)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	// Parse the response
	var deleteOrderResponse DeleteOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&deleteOrderResponse)
	if err != nil {
		return nil, err
	}

	return &deleteOrderResponse, nil
}
