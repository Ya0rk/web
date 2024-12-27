package service

import (
	"regexp"
	"web/model"
	"web/utils/errmsg"
)

func Check(data model.UserRegister) int {
	// 长度检查
	switch {
	case len(data.Username) < 4 || len(data.Username) > 20:
		return errmsg.ERROR_USERNAME_LEN // 用户名长度必须在4到20之间
	case len(data.Password) < 8 || len(data.Password) > 24:
		return errmsg.ERROR_PASSWORD_LEN
	case !IsValidEmail(data.Email):
		return errmsg.ERROR_EMAIL_TYPE
	}

	// 重名 + 邮箱已注册检查
	if ok := IsUsernameExist(data.Username); ok {
		return errmsg.ERROR_USERNAME_USED
	}

	if ok := IsEmailExist(data.Email); ok {
		return errmsg.ERROR_EMAIL_USED
	}

	return errmsg.SUCCESS
}

// 验证电子邮件格式
func IsValidEmail(email string) bool {
	// 正则表达式匹配基本的电子邮件格式
	// 这个正则表达式并不完美，但足以检查常见的错误
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// 检查邮箱是否存在
func IsEmailExist(email string) bool {
	var user model.User
	db.Select("id").Where("email = ?", model.ScryptPw(email)).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

// 判断用户是否存在
func IsUsernameExist(username string) bool {
	var user model.User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}
