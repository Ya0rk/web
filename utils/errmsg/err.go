package errmsg

import "fmt"

const (
	SUCCESS = 200
	ERROR   = 500

	// 登录
	ERROR_PASSWORD_WRONG       = 410
	ERROR_USER_NOT_EXIST       = 411
	ERROR_USERNAME_NULL        = 412
	ERROR_PASSWORD_NULL        = 413
	ERROR_TOKEN_NOT_EXIST      = 414
	ERROR_TOKEN_WRONG          = 415
	ERROR_TOKEN_EXPIRED        = 416
	ERROR_TOKEN_TYPE_WRONG     = 417
	ERROR_EMAIL_NOT_EXIST      = 418
	ERROR_VERIFICATIONCODE_LEN = 419

	// 注册
	ERROR_EMAIL_USED    = 420
	ERROR_USERNAME_USED = 421
	ERROR_USERNAME_LEN  = 422
	ERROR_PASSWORD_LEN  = 423
	ERROR_EMAIL_TYPE    = 424

	ERROR_VERIFICATIONCODE = 425

	ERROR_CARD_MSG = 426
	ERROR_PARAM    = 427

	//服务器错误
	ERROR_SEND_VERIFICATION_CODE = 510
	ERROR_RECOVER_PASSWD_FAIL    = 511
)

var msg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ERROR_PASSWORD_WRONG:       "密码错误",
	ERROR_USER_NOT_EXIST:       "用户不存在,请先注册",
	ERROR_USERNAME_NULL:        "用户名不能为空",
	ERROR_PASSWORD_NULL:        "密码不能为空",
	ERROR_TOKEN_NOT_EXIST:      "token不存在",
	ERROR_TOKEN_WRONG:          "token不正确",
	ERROR_TOKEN_EXPIRED:        "token过期",
	ERROR_TOKEN_TYPE_WRONG:     "token格式错误",
	ERROR_EMAIL_NOT_EXIST:      "该邮箱不存在，请先注册",
	ERROR_VERIFICATIONCODE_LEN: "验证码长度该为6",
	ERROR_VERIFICATIONCODE:     "验证码已过期或已过期，请重新生成验证码",
	ERROR_CARD_MSG:             "请按照格式填写名片信息",
	ERROR_PARAM:                "不能小于0",

	ERROR_EMAIL_USED:    "该邮箱已注册",
	ERROR_USERNAME_USED: "该用户名已存在",
	ERROR_USERNAME_LEN:  "用户名长度只能在4-20",
	ERROR_PASSWORD_LEN:  "密码长度只能在8-24",
	ERROR_EMAIL_TYPE:    "邮箱格式不正确",

	ERROR_SEND_VERIFICATION_CODE: "验证码发送失败",
	ERROR_RECOVER_PASSWD_FAIL:    "密码修改失败",
}

func GetErrMsg(code int) string {
	m, ok := msg[code]
	if ok {
		return m
	}
	return fmt.Sprintf("不存在该状态码：%d", code)
}
