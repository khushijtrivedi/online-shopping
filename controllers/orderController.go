package controllers

import (
	"net/http"
	"online-shopping-api/models"

	"github.com/gin-gonic/gin"
)

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

	var updatedOrder models.Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order data"})
		return
	}

	updatedOrder.ID = orderID

	err := models.UpdateOrder(orderID, updatedOrder)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Order updated", "order": updatedOrder})
}
