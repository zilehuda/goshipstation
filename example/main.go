package main

import (
	"fmt"
	"log"

	shipstation "github.com/zilehuda/goshipstation"
)

func main() {
	apiKey := ""
	apiSecret := ""
	client := shipstation.NewShipStation(apiKey, apiSecret)

	// Call the GetOrders method
	err := getOrders(client)
	if err != nil {
		log.Fatal("Error retrieving orders:", err)
	}

	// Creating a new order
	var order *shipstation.Order
	order, err = addOrder(client)
	if err != nil {
		log.Fatal("Error retrieving orders:", err)
	}

	// Get a specific order by ID
	orderID := order.OrderID
	err = getOrderById(client, orderID)
	if err != nil {
		log.Fatal("Error retrieving order:", err)
	}
	// Delete specific order by ID
	err = deleteOrderById(client, order.OrderID)
	if err != nil {
		log.Fatal("Error in deleting order:", err)
	}
}

func addOrder(client *shipstation.ShipStation) (*shipstation.Order, error) {
	orderData := []byte(`
		{
		  "orderNumber": "TEST-ORDER-API-DOCS",
		  "orderDate": "2015-06-29T08:46:27.0000000",
		  "paymentDate": "2015-06-29T08:46:27.0000000",
		  "shipByDate": "2015-07-05T00:00:00.0000000",
		  "orderStatus": "awaiting_shipment",
		  "customerId": 37701499,
		  "customerUsername": "headhoncho@whitehouse.gov",
		  "customerEmail": "headhoncho@whitehouse.gov",
		  "billTo": {
			"name": "The President",
			"company": null,
			"street1": null,
			"street2": null,
			"street3": null,
			"city": null,
			"state": null,
			"postalCode": null,
			"country": null,
			"phone": null,
			"residential": null
		  },
		  "shipTo": {
			"name": "The President",
			"company": "US Govt",
			"street1": "1600 Pennsylvania Ave",
			"street2": "Oval Office",
			"street3": null,
			"city": "Washington",
			"state": "DC",
			"postalCode": "20500",
			"country": "US",
			"phone": "555-555-5555",
			"residential": true
		  },
		  "items": [
			{
			  "lineItemKey": "vd08-MSLbtx",
			  "sku": "ABC123",
			  "name": "Test item #1",
			  "imageUrl": null,
			  "weight": {
				"value": 24,
				"units": "ounces"
			  },
			  "quantity": 2,
			  "unitPrice": 99.99,
			  "taxAmount": 2.5,
			  "shippingAmount": 5,
			  "warehouseLocation": "Aisle 1, Bin 7",
			  "options": [
				{
				  "name": "Size",
				  "value": "Large"
				}
			  ],
			  "productId": 123456,
			  "fulfillmentSku": null,
			  "adjustment": false,
			  "upc": "32-65-98"
			},
			{
			  "lineItemKey": null,
			  "sku": "DISCOUNT CODE",
			  "name": "10% OFF",
			  "imageUrl": null,
			  "weight": {
				"value": 0,
				"units": "ounces"
			  },
			  "quantity": 1,
			  "unitPrice": -20.55,
			  "taxAmount": null,
			  "shippingAmount": null,
			  "warehouseLocation": null,
			  "options": [],
			  "productId": 123456,
			  "fulfillmentSku": "SKU-Discount",
			  "adjustment": true,
			  "upc": null
			}
		  ],
		  "amountPaid": 218.73,
		  "taxAmount": 5,
		  "shippingAmount": 10,
		  "customerNotes": "Please ship as soon as possible!",
		  "internalNotes": "Customer called and would like to upgrade shipping",
		  "gift": true,
		  "giftMessage": "Thank you!",
		  "paymentMethod": "Credit Card",
		  "requestedShippingService": "Priority Mail"
		}`)

	order, err := client.AddOrder(orderData)
	if err != nil {
		return nil, err
	}
	fmt.Println("New order created with orderId: ", order.OrderID)
	return order, nil
}

// Print the total number of orders
func getOrders(client *shipstation.ShipStation) error {
	ordersResponse, err := client.GetOrders()
	if err != nil {
		return err
	}
	fmt.Println("Total orders:", len(ordersResponse.Orders))
	return nil
}

// Print the details of a specific order by ID
func getOrderById(client *shipstation.ShipStation, orderID int) error {
	order, err := client.GetOrder(orderID)
	if err != nil {
		return err
	}

	fmt.Printf("Order ID: %d\n", order.OrderID)
	fmt.Printf("Order Number: %s\n", order.OrderNumber)
	// Print other order fields as needed

	return nil
}

func deleteOrderById(client *shipstation.ShipStation, orderID int) error {
	resp, err := client.DeleteOrder(orderID)
	if err != nil {
		return err
	}

	fmt.Printf("Status: %t\n", resp.Status)
	fmt.Printf("Message: %s\n", resp.Message)
	// Print other order fields as needed

	return nil
}
