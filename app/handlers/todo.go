package handlers

import (
	"net/http"
	"strconv"

	"todo-service/config"
	"todo-service/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req models.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		UserID:      userID.(uint),
		Status:      models.TodoStatusPending,
	}

	if req.Priority == 0 {
		todo.Priority = 1 // Default priority
	}

	if err := config.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo created successfully",
		"todo":    todo.ToResponse(),
	})
}

func GetTodos(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Parse query parameters
	status := c.Query("status")
	priority := c.Query("priority")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := config.DB.Where("user_id = ?", userID)

	// Apply filters
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if priority != "" {
		query = query.Where("priority = ?", priority)
	}

	var todos []models.Todo
	var total int64

	// Count total records
	query.Model(&models.Todo{}).Count(&total)

	// Get paginated results
	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	// Convert to response format
	var todoResponses []models.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, todo.ToResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": todoResponses,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func GetTodo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	todoID := c.Param("id")

	var todo models.Todo
	if err := config.DB.Where("id = ? AND user_id = ?", todoID, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todo": todo.ToResponse(),
	})
}

func UpdateTodo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	todoID := c.Param("id")

	var todo models.Todo
	if err := config.DB.Where("id = ? AND user_id = ?", todoID, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var req models.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if provided
	if req.Title != nil {
		todo.Title = *req.Title
	}
	if req.Description != nil {
		todo.Description = *req.Description
	}
	if req.Status != nil {
		todo.Status = *req.Status
	}
	if req.Priority != nil {
		todo.Priority = *req.Priority
	}
	if req.DueDate != nil {
		todo.DueDate = req.DueDate
	}

	if err := config.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo updated successfully",
		"todo":    todo.ToResponse(),
	})
}

func DeleteTodo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	todoID := c.Param("id")

	var todo models.Todo
	if err := config.DB.Where("id = ? AND user_id = ?", todoID, userID).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := config.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted successfully",
	})
}

func GetTodoStats(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var totalTodos, completedTodos, pendingTodos int64

	config.DB.Model(&models.Todo{}).Where("user_id = ?", userID).Count(&totalTodos)
	config.DB.Model(&models.Todo{}).Where("user_id = ? AND status = ?", userID, models.TodoStatusCompleted).Count(&completedTodos)
	config.DB.Model(&models.Todo{}).Where("user_id = ? AND status = ?", userID, models.TodoStatusPending).Count(&pendingTodos)

	c.JSON(http.StatusOK, gin.H{
		"stats": gin.H{
			"total":     totalTodos,
			"completed": completedTodos,
			"pending":   pendingTodos,
		},
	})
}
