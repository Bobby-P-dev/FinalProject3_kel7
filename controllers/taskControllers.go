package controllers

import (
	"net/http"
	"strconv"

	"github.com/Bobby-P-dev/FinalProject3_kel7/database"
	"github.com/Bobby-P-dev/FinalProject3_kel7/helpers"
	"github.com/Bobby-P-dev/FinalProject3_kel7/models"
	"github.com/Bobby-P-dev/FinalProject3_kel7/utils/error_utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateTask(c *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := uint(userData["id"].(float64))
	Task := models.Task{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	//mendefaultkan nilai status menjadi false
	Task.Status = false
	Task.UserID = userID

	err := db.Create(&Task).Error

	if err != nil {
		errr := error_utils.NewBadRequest("failed to create task")
		c.JSON(errr.Status(), errr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          Task.ID,
		"title":       Task.Title,
		"status":      Task.Status,
		"description": Task.Description,
		"user_id":     Task.UserID,
		"category_id": Task.CategoryID,
		"created_at":  Task.CreatedAt,
	})
}

func GetTask(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	var Task []models.Task

	Tasks := models.Task{}
	userID := uint(userData["id"].(float64))

	Tasks.UserID = userID

	err := db.Preload("User").Find(&Task, Tasks).Error

	if err != nil {
		errr := error_utils.NewNotFoundError("task not found")
		c.JSON(errr.Status(), errr)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Task,
	})
}

func PutTask(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	taskId, _ := strconv.Atoi(c.Param("taskId"))

	Task := models.Task{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Task)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	Task.UserID = userID
	Task.ID = uint(taskId)
	Task.CategoryID = Task.ID

	err := db.Model(&Task).Where("id = ?", taskId).Updates(models.Task{
		Title: Task.Title, Description: Task.Description,
	}).Error

	if err != nil {
		err := error_utils.NewBadRequest("Failed to update task")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          Task.ID,
		"title":       Task.Title,
		"description": Task.Description,
		"status":      Task.Status,
		"user_id":     Task.UserID,
		"category_id": Task.CategoryID,
		"update_at":   Task.UpdatedAt,
	})
}

func PatchStatusTask(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	taskId, _ := strconv.Atoi(c.Param("taskId"))

	Task := models.Task{}

	userID := uint(userData["id"].(float64))

	Task.UserID = userID
	Task.ID = uint(taskId)
	Task.CategoryID = Task.ID

	err := db.First(&Task, taskId).Error

	if err != nil {
		err := error_utils.NewNotFoundError("task not found")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	var updateData struct {
		Status bool `json:"status"`
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&updateData)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	db.Model(&Task).Update("status", updateData.Status)

	c.JSON(http.StatusOK, gin.H{
		"id":          Task.ID,
		"title":       Task.Title,
		"description": Task.Description,
		"status":      Task.Status,
		"user_id":     Task.UserID,
		"category_id": Task.CategoryID,
		"update_at":   Task.UpdatedAt,
	})
}

func PatchCategoryTask(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	taskId, _ := strconv.Atoi(c.Param("taskId"))

	Task := models.Task{}

	userID := uint(userData["id"].(float64))

	Task.UserID = userID
	Task.ID = uint(taskId)

	err := db.First(&Task, taskId).Error

	if err != nil {
		err := error_utils.NewNotFoundError("task not found")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	var updateCategory struct {
		CategoryID int `json:"category_id"`
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&updateCategory)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	db.Model(&Task).Update("category_id", updateCategory.CategoryID)

	c.JSON(http.StatusOK, gin.H{
		"id":          Task.ID,
		"title":       Task.Title,
		"description": Task.Description,
		"status":      Task.Status,
		"user_id":     Task.UserID,
		"category_id": Task.CategoryID,
		"update_at":   Task.UpdatedAt,
	})
}

func DeleteteTask(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	taskId, _ := strconv.Atoi(c.Param("taskId"))

	Task := models.Task{}

	userID := uint(userData["id"].(float64))

	Task.UserID = userID
	Task.ID = uint(taskId)

	if err := db.First(&Task, taskId).Error; err != nil {
		err := error_utils.NewNotFoundError("task not found")
		c.JSON(err.Status(), err)
		return
	}

	err := db.Delete(&Task).Error
	if err != nil {
		err := error_utils.NewBadRequest("failed to delete task")
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your task has been succesfully deleted",
	})

}
