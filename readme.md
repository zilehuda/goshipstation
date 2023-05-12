# goShipstation

goShipstation is a Go package that provides a convenient way to interact with the ShipStation API. It allows you to manage orders and products within your ShipStation account.

## Installation

To install goShipstation, use the go get command:

```
go get github.com/zilehuda/goshipstation
```

## Usage

Import the goShipstation package in your Go code:

```go
import "github.com/zilehuda/goshipstation"
```

Create a new ShipStation client by providing your ShipStation API credentials:

```go
client := shipstation.NewClient("YOUR_API_KEY", "YOUR_API_SECRET")
```

### Orders

#### AddOrder

AddOrder method allows you to create a new order in ShipStation.

```go
orderData := []byte(`{"orderId": 123, "orderNumber": "ABC123", ...}`)

order, err := client.AddOrder(orderData)
if err != nil {
    // handle error
}

// Use the order object
fmt.Println("Order ID:", order.OrderID)
fmt.Println("Order Number:", order.OrderNumber)
// ...
```

#### GetOrder

GetOrder method retrieves a specific order from ShipStation by its order ID.

```go
orderID := 123

order, err := client.GetOrder(orderID)
if err != nil {
    // handle error
}

// Use the order object
fmt.Println("Order ID:", order.OrderID)
fmt.Println("Order Number:", order.OrderNumber)
// ...
```

#### GetOrders

GetOrders method retrieves all orders from ShipStation.

```go
ordersResponse, err := client.GetOrders()
if err != nil {
    // handle error
}

// Iterate over the orders
for _, order := range ordersResponse.Orders {
    fmt.Println("Order ID:", order.OrderID)
    fmt.Println("Order Number:", order.OrderNumber)
    // ...
}
```

#### DeleteOrder

DeleteOrder method allows you to delete an order from ShipStation by its order ID.

```go
orderID := 123

deleteResponse, err := client.DeleteOrder(orderID)
if err != nil {
    // handle error
}

if deleteResponse.Status {
    fmt.Println("Order deleted successfully")
} else {
    fmt.Println("Failed to delete order:", deleteResponse.Message)
}
```

### Products

#### GetProducts

GetProducts method retrieves all products from ShipStation.

```go
productsResponse, err := client.GetProducts()
if err != nil {
    // handle error
}

// Iterate over the products
for _, product := range productsResponse.Products {
    fmt.Println("Product ID:", product.ProductID)
    fmt.Println("Product Name:", product.Name)
    // ...
}
```

#### GetProduct

GetProduct method retrieves a specific product from ShipStation by its product ID.

```go
productID := 123

product, err := client.GetProduct(productID)
if err != nil {
    // handle error
}

// Use the product object
fmt.Println("Product ID:", product.ProductID)
fmt.Println("Product Name:", product.Name)
// ...
```

#### UpdateProduct

UpdateProduct method allows you to update a product in ShipStation.

```go
productID := 123
updateData := []byte(`{"name": "New Product Name", ...}`)

updateResponse, err := client.UpdateProduct(productID, updateData)
if err != nil {
    // handle error
}

if updateResponse.Success {
    fmt.Println("Product updated successfully")
} else {
    fmt.Println("Failed to update product:", updateResponse.Message)
}
```

## Contribution

Contributions to goShipstation are welcome! If you find a bug or want to suggest an improvement, please open an issue or submit a pull request on the [GitHub repository](https://github.com/zilehuda/goshipstation).

Before making a contribution, please ensure that you:

1. Have a clear understanding of the proposed changes or enhancements.
2. Follow the existing coding style and conventions in the project.
3. Write clear and concise commit messages.
4. Test your changes thoroughly.

### How to Contribute

1. Fork the goShipstation repository to your own GitHub account.
2. Clone the forked repository to your local machine.
3. Create a new branch for your changes:
   ```
   git checkout -b my-feature
   ```
4. Make your desired changes and additions.
5. Run the tests to ensure that everything is functioning correctly.
6. Commit your changes with descriptive commit messages:
   ```
   git commit -m "Add new feature: ..."
   ```
7. Push your branch to your forked repository:
   ```
   git push origin my-feature
   ```
8. Open a pull request on the main goShipstation repository and provide a clear description of your changes.

Once your contribution is submitted, it will be reviewed by the project maintainers. They may provide feedback or request further changes. Thank you for your valuable contribution!

## License

goShipstation is released under the [MIT License](https://opensource.org/licenses/MIT). Please refer to the [LICENSE](https://github.com/zilehuda/goShipstation/blob/main/LICENSE) file for more information.

---

Feel free to customize and adjust the contribution guidelines and license information according to your project's needs.