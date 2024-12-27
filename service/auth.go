package service

import (
	"web/model"
	"web/utils/errmsg"
)

func Login(name string, password string) int {
	var user model.User
	db.Where("username = ?", name).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if model.ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}

func CreateUser(data *model.UserRegister) int {
	var user = model.User{
		Username: data.Username,
		Password: data.Password,
		Email:    data.Email,
	}
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}
