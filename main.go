package main

import (
	"embed"
	"html/template"
	"net/http"
	"online-shopping-api/controllers"
	"online-shopping-api/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

// Embed templates directory
//
//go:embed templates/*
var templatesFS embed.FS

func main() {
	r := gin.Default()

	tmpl := template.Must(template.ParseFS(templatesFS, "templates/*.html"))
	r.SetHTMLTemplate(tmpl)

	// Custom Middleware for logging requests
	r.Use(middleware.RequestLogger())

	// User routes
	r.POST("/register", controllers.RegisterUser)

	// Items routes
	r.GET("/items", controllers.GetItems)
	r.POST("/cart/bulk", controllers.BulkAddToCart)

	// Order routes
	r.GET("/order/:id", controllers.GetOrder)
	r.PUT("/order/:id", controllers.UpdateOrder)

	// Custom error handling
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"ErrorMessage": "Page Not Found",
		})
	})

	// Set custom server configuration
	server := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
