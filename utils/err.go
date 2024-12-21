package utils

import "fmt"

const (
	SUCCESS = 200
	ERROR   = 500

	// 登录
	ERROR_PASSWORD_WRONG = 1001
	ERROR_USER_NOT_EXIST = 1002

	// 注册
	ERROR_PHONE_EXIST   = 2001
	ERROR_USERNAME_USED = 2002
)

var msg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",

	ERROR_PHONE_EXIST:   "该号码已存在",
	ERROR_USERNAME_USED: "该用户名已存在",
}

func GetErrMsg(code int) string {
	m, ok := msg[code]
	if ok {
		return m
	}
	return fmt.Sprintf("不存在该状态码：%d", code)
}
