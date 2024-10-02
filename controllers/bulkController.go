package controllers

import (
	"net/http"
	"online-shopping-api/models"

	"github.com/gin-gonic/gin"
)

// BulkAddToCart allows users to add multiple items to the cart in one request.
func BulkAddToCart(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var request struct {
		ItemIDs []string `json:"item_ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	cart, err := models.GetUserCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve cart"})
		return
	}

	for _, itemID := range request.ItemIDs {
		err := cart.AddItem(itemID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	err = models.UpdateUserCart(userID, cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Items added to cart", "cart": cart})
}
