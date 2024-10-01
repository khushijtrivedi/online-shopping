package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the start time of the request.
		startTime := time.Now()

		// Process request.
		c.Next()

		// Determine color based on HTTP method.
		methodColor := getMethodColor(c.Request.Method)

		// Log the HTTP method in color.
		log.Printf("%s%s\033[0m", methodColor, c.Request.Method)

		// Log route, status code, and time taken.
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()

		// Log format: [Method] [Route] [Status] [Duration]
		log.Printf("Route: %s | Status: %d | Duration: %v", c.Request.URL.Path, statusCode, duration)
	}
}

// getMethodColor returns the ANSI color code based on the HTTP method.
func getMethodColor(method string) string {
	switch method {
	case "GET":
		return "\033[34m" // Blue for GET.
	case "POST":
		return "\033[32m" // Green for POST.
	case "PUT":
		return "\033[33m" // Yellow for PUT.
	case "DELETE":
		return "\033[31m" // Red for DELETE.
	default:
		return "\033[36m" // Cyan for other methods.
	}
}
