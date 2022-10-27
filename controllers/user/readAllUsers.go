package user

import (
	"api/register-login/orm"

	"github.com/gin-gonic/gin"
	"net/http"
)

func ReadAll(c *gin.Context) {
	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "User read success",
		"userId":  users,
	})
}
