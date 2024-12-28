package model

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserCard struct {
	gorm.Model
	Avatar     string `json:"avatar" gorm:"type:varchar(255);not null" validate:"required,url" label:"头像"`        // 限制不能为空且必须是 URL
	Name       string `json:"name" gorm:"type:varchar(20);not null" validate:"required,min=4,max=20" label:"用户名"` // 必须有名字，且长度限制
	BlogSite   string `json:"blog_site" gorm:"type:varchar(255)" validate:"omitempty,url" label:"博客"`             // 可选且必须是 URL
	BriefIntro string `json:"brief_intro" gorm:"type:text" validate:"omitempty,max=50" label:"简介"`                // 可选，最大长度为50
}

// 创建一个全局的验证器实例
var validate = validator.New()

// Validate 函数用于触发验证检查
func (u *UserCard) Validate() error {
	// 使用 validator 库的 Struct() 方法来验证结构体
	if err := validate.Struct(u); err != nil {
		// 如果验证失败，返回错误
		return err
	}
	return nil
}
