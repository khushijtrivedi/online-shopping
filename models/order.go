package models

import (
	"errors"
	"fmt"
	"sync"
)

// Order represents a single order.
type Order struct {
	ID       string  `json:"id"`
	ItemName string  `json:"item_name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// In-memory data store for orders.
var (
	orders         = make(map[string]Order)
	orderIDCounter int        // Counter to generate unique order IDs
	mu             sync.Mutex // To ensure thread-safe operations on orders
)

// GetOrderById retrieves an order by its ID.
func GetOrderById(id string) (Order, error) {
	mu.Lock()
	defer mu.Unlock()

	order, exists := orders[id]
	if !exists {
		return Order{}, errors.New("order not found")
	}
	return order, nil
}

// UpdateOrder updates an order in the data store.
func UpdateOrder(id string, updatedOrder Order) error {
	mu.Lock()
	defer mu.Unlock()

	_, exists := orders[id]
	if !exists {
		return errors.New("order not found")
	}

	orders[id] = updatedOrder
	return nil
}

// CreateOrder adds a new order to the in-memory store.
func CreateOrder(order Order) {
	mu.Lock()
	defer mu.Unlock()

	orders[order.ID] = order
}

// CreateOrderFromCart creates orders based on items in the user's cart.
func CreateOrderFromCart(userID string) ([]Order, error) {
	cart, err := GetUserCart(userID)
	if err != nil {
		return nil, err
	}

	var createdOrders []Order
	for itemID, quantity := range cart.Items {
		// Generate a unique order ID
		orderIDCounter++
		orderID := fmt.Sprintf("order-%d", orderIDCounter)

		order := Order{
			ID:       orderID, // Use a generated ID for the order
			ItemName: itemID,  // Replace with a proper item name lookup as needed
			Price:    10.0,    // Set the price based on actual item data
			Quantity: quantity,
		}

		// Create the order in the in-memory store.
		CreateOrder(order)
		createdOrders = append(createdOrders, order)
	}

	// Clear the cart after creating the order.
	err = UpdateUserCart(userID, Cart{Items: make(map[string]int)})
	if err != nil {
		return nil, err
	}

	return createdOrders, nil
}
