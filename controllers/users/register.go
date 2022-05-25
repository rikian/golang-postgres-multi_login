package users

import (
	"golang/entity/request"

	"gorm.io/gorm"
)

func Register(db *gorm.DB, user *request.Tb_users) bool {
	result := db.Create(user)

	return result.Error == nil
}