package models

import (
	"time"

	"gorm.io/gorm"
)

type TodoStatus string

const (
	TodoStatusPending   TodoStatus = "pending"
	TodoStatusCompleted TodoStatus = "completed"
)

type Todo struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null" binding:"required,max=200"`
	Description string         `json:"description" gorm:"type:text"`
	Status      TodoStatus     `json:"status" gorm:"default:'pending'" binding:"omitempty,oneof=pending completed"`
	Priority    int            `json:"priority" gorm:"default:1" binding:"omitempty,min=1,max=5"`
	DueDate     *time.Time     `json:"due_date,omitempty"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	User        User           `json:"user,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateTodoRequest struct {
	Title       string     `json:"title" binding:"required,max=200"`
	Description string     `json:"description" binding:"omitempty,max=1000"`
	Priority    int        `json:"priority" binding:"omitempty,min=1,max=5"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

type UpdateTodoRequest struct {
	Title       *string     `json:"title,omitempty" binding:"omitempty,max=200"`
	Description *string     `json:"description,omitempty" binding:"omitempty,max=1000"`
	Status      *TodoStatus `json:"status,omitempty" binding:"omitempty,oneof=pending completed"`
	Priority    *int        `json:"priority,omitempty" binding:"omitempty,min=1,max=5"`
	DueDate     *time.Time  `json:"due_date,omitempty"`
}

type TodoResponse struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
	Priority    int        `json:"priority"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	UserID      uint       `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (t *Todo) ToResponse() TodoResponse {
	return TodoResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Priority:    t.Priority,
		DueDate:     t.DueDate,
		UserID:      t.UserID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
