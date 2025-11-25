package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itssiddhant/taskmanager/internal/models"
	"gorm.io/gorm"
)

type TaskHandler struct {
	DB *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDVal, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	task.ID = 0
	task.UserID = userID

	if err := h.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	userIDVal, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	var tasks []models.Task
	if err := h.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")

	userIDVal, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	var task models.Task
	if err := h.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delete(updates, "user_id")
	delete(updates, "id")

	if err := h.DB.Model(&task).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")

	userIDVal, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, ok := userIDVal.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	var task models.Task
	if err := h.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
