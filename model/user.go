package model

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=20" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Email    string `gorm:"type:varchar(255);not null" json:"email" validate:"required,max=255" label:"邮箱"`
}

type UserRegister struct {
	gorm.Model
	Username         string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=20" label:"用户名"`
	Password         string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Email            string `gorm:"type:varchar(255);not null" json:"email" validate:"required,max=255" label:"邮箱"`
	VerificationCode string `gorm:"type:char(6);not null" json:"verificationCode" validate:"required,regexp=^[0-9]{6}$" label:"验证码"`
}

type UserEmailLogin struct {
	gorm.Model
	Email            string `gorm:"type:varchar(255);not null" json:"email" validate:"required,max=255" label:"邮箱"`
	VerificationCode string `gorm:"type:char(6);not null" json:"verificationCode" validate:"required,regexp=^[0-9]{6}$" label:"验证码"`
}

type UserEmail struct {
	gorm.Model
	Email string `gorm:"type:varchar(255);not null" json:"email" validate:"required,max=255" label:"邮箱"`
}

type RecoverPasswd struct {
	gorm.Model
	Email            string `gorm:"type:varchar(255);not null" json:"email" validate:"required,max=255" label:"邮箱"`
	Password         string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	VerificationCode string `gorm:"type:char(6);not null" json:"verificationCode" validate:"required,regexp=^[0-9]{6}$" label:"验证码"`
}

/*
密码加密 ： scrypt加密方法
https://pkg.go.dev/golang.org/x/crypto/scrypt
*/

// 钩子函数,在写入数据库之前加密 密码和邮箱
func (u *User) BeforeSave(*gorm.DB) error {
	u.Password = ScryptPw(u.Password)
	u.Email = ScryptPw(u.Email)
	return nil
}
func (u *User) BeforeUpdate(*gorm.DB) error {
	u.Password = ScryptPw(u.Password)
	return nil
}

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{92, 42, 7, 68, 66, 5, 224, 37}
	// 进行加密处理
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	// 编码为 base64
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
