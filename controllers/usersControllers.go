package controllers

import (
	"net/http"

	"github.com/Bobby-P-dev/FinalProject3_kel7/database"
	"github.com/Bobby-P-dev/FinalProject3_kel7/helpers"
	"github.com/Bobby-P-dev/FinalProject3_kel7/models"
	"github.com/Bobby-P-dev/FinalProject3_kel7/utils/error_utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	appJSON = "application/json"
)

func CreatAcount(c *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	err := db.Create(&User).Error

	if err != nil {
		errr := error_utils.NewBadRequest("failed to create account")
		c.JSON(errr.Status(), errr)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":         User.ID,
		"full_name":  User.FullName,
		"email":      User.Email,
		"created_at": User.CreatedAt,
	})
}

func Login(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	password = User.Password

	err := db.Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		errr := error_utils.NewUnauthorized("invalid email / password")
		c.JSON(errr.Status(), errr)
		return
	}
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		errr := error_utils.NewUnauthorized("invalid email / password")
		c.JSON(errr.Status(), errr)
		return
	}
	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func PutAcount(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_ = contentType

	User := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	User.ID = userID

	err := db.Model(&User).Updates(models.User{
		FullName: User.FullName, Email: User.Email,
	}).First(&User).Error

	if err != nil {
		err := error_utils.NewBadRequest("failed to update account")
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"fullname":  User.FullName,
		"update_at": User.UpdatedAt,
	})
}

func DelletAcount(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	User := models.User{}

	userID := uint(userData["id"].(float64))

	User.ID = userID

	if err := db.First(&User).Error; err != nil {
		err := error_utils.NewNotFoundError("acccount not found")
		c.JSON(err.Status(), err)
		return
	}

	err := db.Delete(&User).Error
	if err != nil {
		err := error_utils.NewBadRequest("failed to delete account")
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been succesfully deleted",
	})
}
