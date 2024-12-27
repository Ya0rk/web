package errmsg

import "fmt"

const (
	SUCCESS = 200
	ERROR   = 500

	// 登录
	ERROR_PASSWORD_WRONG = 410
	ERROR_USER_NOT_EXIST = 411
	ERROR_USERNAME_NULL  = 412
	ERROR_PASSWORD_NULL  = 413

	// 注册
	ERROR_EMAIL_USED    = 420
	ERROR_USERNAME_USED = 421
	ERROR_USERNAME_LEN  = 422
	ERROR_PASSWORD_LEN  = 423
	ERROR_EMAIL         = 424
)

var msg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在,请先注册",
	ERROR_USERNAME_NULL:  "用户名不能为空",
	ERROR_PASSWORD_NULL:  "密码不能为空",

	ERROR_EMAIL_USED:    "该邮箱已注册",
	ERROR_USERNAME_USED: "该用户名已存在",
	ERROR_USERNAME_LEN:  "用户名长度只能在4-20",
	ERROR_PASSWORD_LEN:  "密码长度只能在8-24",
	ERROR_EMAIL:         "邮箱格式不正确",
}

func GetErrMsg(code int) string {
	m, ok := msg[code]
	if ok {
		return m
	}
	return fmt.Sprintf("不存在该状态码：%d", code)
}
