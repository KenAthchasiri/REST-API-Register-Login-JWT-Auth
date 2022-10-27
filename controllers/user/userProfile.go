package user

import (
	"api/register-login/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user orm.User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "User read success",
		"userId":  user,
	})
}