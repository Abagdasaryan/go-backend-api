package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Response represents a standard API response
type Response struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
	Status    string      `json:"status"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
	Version   string    `json:"version"`
}

var startTime = time.Now()

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create router
	r := gin.Default()

	// Add middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(corsMiddleware())

	// Routes
	api := r.Group("/api/v1")
	{
		api.GET("/health", healthCheck)
		api.GET("/", welcome)
		api.GET("/echo/:message", echoMessage)
		api.POST("/data", createData)
		api.GET("/data", getData)
	}

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, Response{
			Message:   "Welcome to Go Backend API",
			Status:    "success",
			Timestamp: time.Now(),
		})
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// Middleware
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Handlers
func healthCheck(c *gin.Context) {
	uptime := time.Since(startTime)

	c.JSON(200, HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Uptime:    uptime.String(),
		Version:   "1.0.0",
	})
}

func welcome(c *gin.Context) {
	c.JSON(200, Response{
		Message:   "Welcome to the API!",
		Status:    "success",
		Timestamp: time.Now(),
		Data: map[string]interface{}{
			"endpoints": []string{
				"GET /api/v1/health - Health check",
				"GET /api/v1/ - Welcome message",
				"GET /api/v1/echo/:message - Echo a message",
				"POST /api/v1/data - Create data",
				"GET /api/v1/data - Get all data",
			},
		},
	})
}

func echoMessage(c *gin.Context) {
	message := c.Param("message")

	c.JSON(200, Response{
		Message:   "Message echoed successfully",
		Status:    "success",
		Timestamp: time.Now(),
		Data: map[string]interface{}{
			"echo":   message,
			"length": len(message),
		},
	})
}

// Simple in-memory storage for demo purposes
var dataStore = make(map[string]interface{})

func createData(c *gin.Context) {
	var requestData map[string]interface{}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, Response{
			Message:   "Invalid JSON data",
			Status:    "error",
			Timestamp: time.Now(),
		})
		return
	}

	// Generate a simple ID
	id := time.Now().Format("20060102150405")
	dataStore[id] = requestData

	c.JSON(201, Response{
		Message:   "Data created successfully",
		Status:    "success",
		Timestamp: time.Now(),
		Data: map[string]interface{}{
			"id":   id,
			"data": requestData,
		},
	})
}

func getData(c *gin.Context) {
	c.JSON(200, Response{
		Message:   "Data retrieved successfully",
		Status:    "success",
		Timestamp: time.Now(),
		Data:      dataStore,
	})
}
