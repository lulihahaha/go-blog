package service

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 生成post id
	p.ID = int64(snowflake.GenID())

	// 保存到数据库
	return mysql.CreatePost(p)

}

func GetPostByID(pid int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Error(err))
		return
	}

	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(post.AuthorID)", zap.Error(err))
		return
	}

	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID)", zap.Error(err))
		return
	}

	data = &models.ApiPostDetail{
		user.Username,
		post,
		community,
	}

	return
}
