package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/models"
	"web_app/service"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreatePostHandler failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("GetCurrentUser failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	if err := service.CreatePost(p); err != nil {
		zap.L().Error("service.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, CodeSuccess)
}

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := service.GetPostByID(pid)
	if err != nil {
		zap.L().Error("service.GetPostByID(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
