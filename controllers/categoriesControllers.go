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

func CreateCategories(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	Category := models.Category{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Category)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	err := db.Create(&Category).Error

	if err != nil {
		errr := error_utils.NewBadRequest("failed to create category")
		c.JSON(errr.Status(), errr)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id"         : Category.ID,
		"type"       : Category.Type,
		"created_at" : Category.CreatedAt,
	})
}

func GetteCategories(c *gin.Context) {
	db := database.GetDB()

	var Category []models.Category

	Categorys := models.Category{}

	err := db.Preload("Tasks").Find(&Category, Categorys).Error

	if err != nil {
		errr := error_utils.NewNotFoundError("task not found")
		c.JSON(errr.Status(), errr)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Category,
	})
}

func PatchCategories(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	categoryId, err := strconv.Atoi(c.Param("categoryId"))

	if err != nil {
		err := error_utils.NewBadRequest("invalid parameter")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	Category := models.Category{}

	Category.ID = uint(categoryId)

	err = db.First(&Category, categoryId).Error

	if err != nil {
		err := error_utils.NewNotFoundError("category not found")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	var updateData struct {
		Type string `json:"type"`
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&updateData)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	db.Model(&Category).Update("type", updateData.Type)

	c.JSON(http.StatusOK, gin.H{
		"id"         : Category.ID,
		"type"       : Category.Type,
		"created_at" : Category.CreatedAt,
	})
}

func DeleteCategories(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	categoryId, err := strconv.Atoi(c.Param("categoryId"))

	if err != nil {
		err := error_utils.NewBadRequest("invalid parameter")
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	Category := models.Category{}

	userID := uint(userData["id"].(float64))

	Category.ID = userID
	Category.ID = uint(categoryId)

	if err := db.First(&Category, categoryId).Error; err != nil {
		err := error_utils.NewNotFoundError("category not found")
		c.JSON(err.Status(), err)
		return
	}

	err = db.Delete(&Category).Error
	if err != nil {
		err := error_utils.NewBadRequest("failed to delete category")
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Category has been succesfully deleted",
	})
}
