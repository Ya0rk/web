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
	case !isValidEmail(data.Email):
		return errmsg.ERROR_EMAIL
	}

	// 重名 + 邮箱已注册检查
	var user model.User
	db.Select("id").Where("username = ?", data.Username).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}

	db.Select("id").Where("email = ?", model.ScryptPw(data.Email)).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_EMAIL_USED
	}

	return errmsg.SUCCESS
}

// 验证电子邮件格式
func isValidEmail(email string) bool {
	// 正则表达式匹配基本的电子邮件格式
	// 这个正则表达式并不完美，但足以检查常见的错误
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
