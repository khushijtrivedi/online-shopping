package controllers

import (
	"net/http"
	"online-shopping-api/models"

	"github.com/gin-gonic/gin"
)

func DeleteOrder(c *gin.Context) {
	orderID := c.Param("id") // Extract order ID from URL parameters

	if err := models.DeleteOrder(orderID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Order deleted"})
}
func GetOrder(c *gin.Context) {
	orderID := c.Param("id")
	order, err := models.GetOrderById(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func UpdateOrder(c *gin.Context) {
	orderID := c.Param("id")
	var updatedOrder models.Order // Assuming you have an Order model defined

	// Bind the incoming JSON to the updatedOrder struct
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Fetch the existing order (you might need to implement this)
	existingOrder, err := models.GetOrderByID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Update the fields (make sure you are updating the correct fields)
	existingOrder.ItemName = updatedOrder.ItemName // Assuming these fields exist
	existingOrder.Price = updatedOrder.Price
	existingOrder.Quantity = updatedOrder.Quantity

	// Save the updated order back to your data store
	if err := models.UpdateOrder(orderID, existingOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Order updated", "order": existingOrder})
}

func CreateOrderFromCart(c *gin.Context) {
	var request struct {
		UserID          string `json:"userId" binding:"required"`
		PaymentMethod   string `json:"paymentMethod" binding:"required"`
		ShippingAddress string `json:"shippingAddress" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	orders, err := models.CreateOrderFromCart(request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create orders from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Order created successfully", "orders": orders})
}
