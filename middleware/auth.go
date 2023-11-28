package middleware

import (
	"strconv"

	"github.com/Bobby-P-dev/FinalProject3_kel7/database"
	"github.com/Bobby-P-dev/FinalProject3_kel7/helpers"
	"github.com/Bobby-P-dev/FinalProject3_kel7/models"
	"github.com/Bobby-P-dev/FinalProject3_kel7/utils/error_utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifToken, err := helpers.VerifToken(c)
		_ = verifToken

		if err != nil {
			errr := error_utils.NewUnauthorized("unauthorized")
			c.AbortWithStatusJSON(errr.Status(), errr)
			return
		}
		c.Set("userData", verifToken)
		c.Next()
	}
}
func RoleAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			err := error_utils.NewNotFoundError("data not found")
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}
		if user.Role != "admin" {
			err := error_utils.NewUnauthorized("only admin can access")
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func TaskAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		taskId, err := strconv.Atoi(c.Param("taskId"))

		if err != nil {
			err := error_utils.NewBadRequest("invalid parameter")
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Task := models.Task{}

		err = db.Select("user_id").First(&Task, taskId).Error

		if err != nil {
			err := error_utils.NewNotFoundError("data not found")
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if int(Task.UserID) != int(userID) {
			err := error_utils.NewUnauthorized("Cannot access data")
			c.AbortWithStatusJSON(err.Status(), err)
			return
		}
		c.Next()
	}

}
