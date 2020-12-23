package controller

import (
	"go-basic/model/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) {
	userId := c.Query("userId")
	intUserId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil || intUserId <= 0 {
		genResult(c, 1, "invalid user id", nil)
		return
	} else {
		user, err := service.GetUser(userId)
		if err != nil {
			genResult(c, 1, err.Error(), nil)
			return
		}
		genResult(c, 0, "ok", user)
	}
}

func addUser(c *gin.Context) {
	_, ok := c.Get("name")
	if !ok {
		genResult(c, 1, "unknown user id", nil)
	} else {
		genResult(c, 0, "ok", nil)
	}
}
