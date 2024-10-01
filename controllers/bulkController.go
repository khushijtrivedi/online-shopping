package controllers

import (
	"net/http"
	"online-shopping-api/models"

	"github.com/gin-gonic/gin"
)

// BulkAddToCart allows users to add multiple items to the cart in one request.
func BulkAddToCart(c *gin.Context) {
	// Extract user ID from request (assuming a query param or header).
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Bind the list of item IDs from the request body.
	var itemIDs []string
	if err := c.ShouldBindJSON(&itemIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Retrieve the user's cart from the model.
	cart, err := models.GetUserCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve cart"})
		return
	}

	// Add each item to the cart.
	for _, itemID := range itemIDs {
		err := cart.AddItem(itemID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// Save the updated cart back to the in-memory store.
	err = models.UpdateUserCart(userID, cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Items added to cart", "cart": cart})
}
