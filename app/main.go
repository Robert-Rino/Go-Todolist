package main

import (
	"log"
	"net/http"

	"todo-service/config"
	"todo-service/handlers"
	"todo-service/middleware"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Todo Service is running",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")

	// Public routes (no authentication required)
	auth := v1.Group("/auth")
	{
		auth.POST("/register", handlers.RegisterUser)
		auth.POST("/login", handlers.LoginUser)
	}

	// Public user routes
	users := v1.Group("/users")
	{
		users.GET("/:id", handlers.GetUserProfile)
	}

	// Protected routes (authentication required)
	protected := v1.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Current user routes
		protected.GET("/me", handlers.GetCurrentUser)

		// Todo routes
		todos := protected.Group("/todos")
		{
			todos.POST("/", handlers.CreateTodo)
			todos.GET("/", handlers.GetTodos)
			todos.GET("/stats", handlers.GetTodoStats)
			todos.GET("/:id", handlers.GetTodo)
			todos.PUT("/:id", handlers.UpdateTodo)
			todos.DELETE("/:id", handlers.DeleteTodo)
		}
	}

	return r
}

func main() {
	// Connect to database
	config.ConnectDatabase()

	// Setup router
	r := setupRouter()

	// Start server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
