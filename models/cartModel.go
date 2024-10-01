package models

import "errors"

// Cart represents a user's shopping cart.
type Cart struct {
	Items map[string]int `json:"items"` // Key is item ID, value is quantity.
}

// In-memory data store for user carts.
var userCarts = make(map[string]Cart)

// GetUserCart retrieves the cart for a specific user by user ID.
func GetUserCart(userID string) (Cart, error) {
	cart, exists := userCarts[userID]
	if !exists {
		// Initialize a new cart if the user has no existing cart.
		cart = Cart{Items: make(map[string]int)}
		userCarts[userID] = cart
	}
	return cart, nil
}

// UpdateUserCart saves the updated cart for a user by user ID.
func UpdateUserCart(userID string, cart Cart) error {
	userCarts[userID] = cart
	return nil
}

// AddItem adds an item to the user's cart.
func (c *Cart) AddItem(itemID string) error {
	// Check if the item exists in the inventory (replace with real check in production).
	if !ItemExists(itemID) {
		return errors.New("item not found in inventory")
	}

	// If the item is already in the cart, increment the quantity.
	if _, exists := c.Items[itemID]; exists {
		c.Items[itemID]++
	} else {
		// Otherwise, add the item with a quantity of 1.
		c.Items[itemID] = 1
	}

	return nil
}

// ItemExists checks if an item is available in inventory (dummy implementation).
func ItemExists(itemID string) bool {
	inventory := []string{"item1", "item2", "item3"} // Example inventory items.
	for _, item := range inventory {
		if item == itemID {
			return true
		}
	}
	return false
}
