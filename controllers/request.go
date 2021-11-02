package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const ContextUsername = "username"

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCurrentUser 获取当前用户登录ID
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUsername)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
