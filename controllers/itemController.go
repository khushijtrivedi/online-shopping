package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	category := c.Query("category")
	priceStr := c.Query("price")

	// Filtering logic for in-memory data (replace with DB query in production)
	items := []map[string]interface{}{
		{"name": "Item1", "category": "electronics", "price": 100},
		{"name": "Item2", "category": "books", "price": 20},
		{"name": "Item3", "category": "clothing", "price": 50},
	}

	filteredItems := []map[string]interface{}{}
	maxPrice := 0
	if priceStr != "" {
		maxPrice, _ = strconv.Atoi(priceStr)
	}

	for _, item := range items {
		if (category == "" || item["category"] == category) && (priceStr == "" || item["price"].(int) <= maxPrice) {
			filteredItems = append(filteredItems, item)
		}
	}

	c.JSON(http.StatusOK, gin.H{"items": filteredItems})
}
