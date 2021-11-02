package service

import (
	"errors"
	"fmt"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

var (
	ErrorUserExist = errors.New("用户已存在")
)

func SignUp(p *models.ParamSingUp) (err error) {
	// 判断用户是否存在
	var exist bool
	exist, err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	if exist {
		return ErrorUserExist
	}
	// 生成UID
	userID := snowflake.GenID()

	// 密码加密

	// 构造一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 保存进数据库
	err = mysql.InsertUser(user)
	if err != nil {
		fmt.Sprintf("err: %v/n", err)
		return
	}
	return
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err := mysql.Login(user); err != nil {
		return "", err
	}

	// 生成JWT的Token
	aToken, _, err := jwt.GenToken(user.UserID, user.Username)
	return aToken, err
}
