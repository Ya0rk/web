package service

import (
	"errors"
	"gorm.io/gorm"
	"web/model"
	"web/utils/errmsg"
)

type CardService interface {
	CreateCard(data model.UserCard) int
	GetCards(pageSize int, pageNum int) (card []model.UserCard, code int, num int64)
}

type cardService struct {
}

// 构造函数
func NewCardService() CardService {
	return &cardService{}
}

func (c *cardService) CreateCard(data model.UserCard) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

func (c *cardService) GetCards(pageSize int, pageNum int) (cards []model.UserCard, code int, num int64) {
	var userCards []model.UserCard
	var total int64

	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&userCards).Count(&total).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errmsg.ERROR, 0
	}
	return userCards, errmsg.SUCCESS, total
}
