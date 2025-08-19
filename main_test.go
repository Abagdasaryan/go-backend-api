package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(corsMiddleware())

	api := r.Group("/api/v1")
	{
		api.GET("/health", healthCheck)
		api.GET("/", welcome)
		api.GET("/echo/:message", echoMessage)
		api.POST("/data", createData)
		api.GET("/data", getData)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, Response{
			Message:   "Welcome to Go Backend API",
			Status:    "success",
			Timestamp: time.Now(),
		})
	})

	return r
}

func TestHealthCheck(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response HealthResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "healthy", response.Status)
	assert.NotEmpty(t, response.Uptime)
	assert.Equal(t, "1.0.0", response.Version)
}

func TestWelcome(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Welcome to the API!", response.Message)
	assert.Equal(t, "success", response.Status)
}

func TestEchoMessage(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/echo/testmessage", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Message echoed successfully", response.Message)
	assert.Equal(t, "success", response.Status)
}

func TestRootEndpoint(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Welcome to Go Backend API", response.Message)
	assert.Equal(t, "success", response.Status)
}
