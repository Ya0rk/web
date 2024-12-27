package service

import (
	"web/model"
	"web/utils/errmsg"
)

// 密码方式登录验证
func LoginByPasswd(name string, password string) int {
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

// 邮箱验证码方式登录验证
func LoginByEmail(email string, verificationCode string) int {
	// 判断邮箱是否存在
	if ok := IsEmailExist(email); !ok {
		return errmsg.ERROR_EMAIL_NOT_EXIST
	}
	// 判断验证码是否正确
	if ok := NewEmailService().VerifyVerificationCode(email, verificationCode); !ok {
		return errmsg.ERROR_VERIFICATIONCODE
	}

	return errmsg.SUCCESS
}

// 注册创建用户，使用邮箱方式
func CreateUser(data *model.UserRegister) int {
	var user = model.User{
		Username: data.Username,
		Password: data.Password,
		Email:    data.Email,
	}
	err := db.Create(&user).Error
	if err != nil {
		println(err.Error())
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}
