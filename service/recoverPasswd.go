package service

import (
	"web/model"
	"web/utils/errmsg"
)

func RecoverPasswd(email string, newPassword string) int {
	var user model.User
	// 创建一个映射，更新 `passwd` 字段
	updates := map[string]interface{}{
		"password": model.ScryptPw(newPassword),
	}
	err := db.Model(&user).Where("email = ?", model.ScryptPw(email)).Updates(updates).Error
	if err != nil {
		println(err.Error())
		return errmsg.ERROR_RECOVER_PASSWD_FAIL
	}
	return errmsg.SUCCESS
}
