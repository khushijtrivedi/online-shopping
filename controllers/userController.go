package controllers

import (
	"net/http"
	"online-shopping-api/models"
	"online-shopping-api/validators"

	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=8"`
}

func RegisterUser(c *gin.Context) {
	var form RegisterForm

	// Bind the incoming form data to the RegisterForm struct.
	if err := c.ShouldBind(&form); err != nil {
		c.AsciiJSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// Validate password strength.
	if !validators.ValidatePasswordStrength(form.Password) {
		c.AsciiJSON(http.StatusBadRequest, gin.H{"error": "Password too weak"})
		return
	}

	// Register user using the model function.
	userID, err := models.RegisterUser(form.Email, form.Password) // Update to receive user ID
	if err != nil {
		c.AsciiJSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.AsciiJSON(http.StatusOK, gin.H{"status": "User registered successfully", "userID": userID}) // Include user ID in response
}
