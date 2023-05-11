package shipstation

import (
	"encoding/json"
	"fmt"
)

type DeleteOrderResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
type OrdersResponse struct {
	Orders []Order `json:"orders"`
	Total  int     `json:"total"`
	Page   int     `json:"page"`
	Pages  int     `json:"pages"`
}

type Order struct {
	OrderID                   int                  `json:"orderId"`
	OrderNumber               string               `json:"orderNumber"`
	OrderKey                  string               `json:"orderKey"`
	OrderDate                 string               `json:"orderDate"`
	CreateDate                string               `json:"createDate"`
	ModifyDate                string               `json:"modifyDate"`
	PaymentDate               string               `json:"paymentDate"`
	ShipByDate                string               `json:"shipByDate"`
	OrderStatus               string               `json:"orderStatus"`
	CustomerID                int                  `json:"customerId"`
	CustomerUsername          string               `json:"customerUsername"`
	CustomerEmail             string               `json:"customerEmail"`
	BillTo                    Address              `json:"billTo"`
	ShipTo                    Address              `json:"shipTo"`
	Items                     []Item               `json:"items"`
	OrderTotal                float32              `json:"orderTotal"`
	AmountPaid                float32              `json:"amountPaid"`
	TaxAmount                 float32              `json:"taxAmount"`
	ShippingAmount            float32              `json:"shippingAmount"`
	CustomerNotes             string               `json:"customerNotes"`
	InternalNotes             string               `json:"float32ernalNotes"`
	Gift                      bool                 `json:"gift"`
	GiftMessage               string               `json:"giftMessage"`
	PaymentMethod             string               `json:"paymentMethod"`
	RequestedShippingSrvc     string               `json:"requestedShippingService"`
	CarrierCode               string               `json:"carrierCode"`
	ServiceCode               string               `json:"serviceCode"`
	PackageCode               string               `json:"packageCode"`
	Confirmation              string               `json:"confirmation"`
	ShipDate                  string               `json:"shipDate"`
	HoldUntilDate             string               `json:"holdUntilDate"`
	Weight                    Weight               `json:"weight"`
	Dimensions                string               `json:"dimensions"`
	InsuranceOptions          InsuranceOptions     `json:"insuranceOptions"`
	InternationalOptions      InternationalOptions `json:"float32ernationalOptions"`
	AdvancedOptions           AdvancedOptions      `json:"advancedOptions"`
	TagIDs                    []string             `json:"tagIds"`
	UserID                    string               `json:"userId"`
	ExternallyFulfilled       bool                 `json:"externallyFulfilled"`
	ExternallyFulfilledBy     string               `json:"externallyFulfilledBy"`
	ExternallyFulfilledByID   string               `json:"externallyFulfilledById"`
	ExternallyFulfilledByName string               `json:"externallyFulfilledByName"`
	LabelMessages             string               `json:"labelMessages"`
}

type Address struct {
	Name            string `json:"name"`
	Company         string `json:"company"`
	Street1         string `json:"street1"`
	Street2         string `json:"street2"`
	Street3         string `json:"street3"`
	City            string `json:"city"`
	State           string `json:"state"`
	PostalCode      string `json:"postalCode"`
	Country         string `json:"country"`
	Phone           string `json:"phone"`
	Residential     bool   `json:"residential"`
	AddressVerified string `json:"addressVerified"`
}

type Item struct {
	OrderItemID       float32  `json:"orderItemId"`
	LineItemKey       string   `json:"lineItemKey"`
	SKU               string   `json:"sku"`
	Name              string   `json:"name"`
	ImageURL          string   `json:"imageUrl"`
	Weight            Weight   `json:"weight"`
	Quantity          float32  `json:"quantity"`
	UnitPrice         float32  `json:"unitPrice"`
	TaxAmount         float32  `json:"taxAmount"`
	ShippingAmount    float32  `json:"shippingAmount"`
	WarehouseLocation string   `json:"warehouseLocation"`
	Options           []Option `json:"options"`
	ProductID         int      `json:"productId"`
	FulfillmentSKU    string   `json:"fulfillmentSku"`
	Adjustment        bool     `json:"adjustment"`
	UPC               string   `json:"upc"`
	CreateDate        string   `json:"createDate"`
	ModifyDate        string   `json:"modifyDate"`
}

type Weight struct {
	Value       float32 `json:"value"`
	Units       string  `json:"units"`
	WeightUnits float32 `json:"WeightUnits"`
}

type Option struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type InsuranceOptions struct {
	Provider       string  `json:"provider"`
	InsureShipment bool    `json:"insureShipment"`
	InsuredValue   float32 `json:"insuredValue"`
}

type InternationalOptions struct {
	Contents     string `json:"contents"`
	CustomsItems string `json:"customsItems"`
	NonDelivery  string `json:"nonDelivery"`
}

type AdvancedOptions struct {
	WarehouseID          string   `json:"warehouseId"`
	NonMachinable        bool     `json:"nonMachinable"`
	SaturdayDelivery     bool     `json:"saturdayDelivery"`
	ContainsAlcohol      bool     `json:"containsAlcohol"`
	MergedOrSplit        bool     `json:"mergedOrSplit"`
	MergedIDs            []string `json:"mergedIds"`
	ParentID             string   `json:"parentId"`
	StoreID              float32  `json:"storeId"`
	CustomField1         string   `json:"customField1"`
	CustomField2         string   `json:"customField2"`
	CustomField3         string   `json:"customField3"`
	Source               string   `json:"source"`
	BillToParty          string   `json:"billToParty"`
	BillToAccount        string   `json:"billToAccount"`
	BillToPostalCode     string   `json:"billToPostalCode"`
	BillToCountryCode    string   `json:"billToCountryCode"`
	BillToMyOtherAccount string   `json:"billToMyOtherAccount"`
}

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
