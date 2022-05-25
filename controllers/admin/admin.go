package admin

import (
	"golang/entity/request"
	"golang/entity/response"

	// "golang/utils"
	"strings"
	"time"

	"gorm.io/gorm"
)

func Register(db *gorm.DB, user *request.Tb_users) bool {
	result := db.Create(user)

	return result.Error == nil
}

func Login(db *gorm.DB, user *request.Tb_users) (response.Tb_users, bool) {
	var status_user response.Tb_users
	result := db.Where("user_email = ? AND user_password = ?", user.User_email, user.User_password).First(&status_user)
	if result.Error != nil {
		return status_user, true
	}
	
	return status_user, false
}

func Logout(db *gorm.DB, id, email string) bool {
	result := db.Table("tb_users").Where("user_id = ? AND user_email = ?", id, email).Updates(response.Tb_users{User_session: "null", Lats_update: time.Now().Format("2006-01-02 15:04:05")})
	
	return result.Error == nil
}

func SaveTokenLogin(db *gorm.DB, tokenLogin, email, password string) bool {
	result := db.Table("tb_users").Where("user_email = ? AND user_password = ?", email, password).Updates(response.Tb_users{User_session: tokenLogin, Lats_update: time.Now().Format("2006-01-02 15:04:05")})

	return result.Error == nil
}

func GetUserAfterLogin(db *gorm.DB, id, token string) (response.Tb_users, bool) {
	var data_user response.Tb_users
	result := db.Where("user_id = ?", id).First(&data_user)
	if result.Error != nil {
		return data_user, true
	}

	if strings.Split(data_user.User_session, ".")[2] != token {
		return data_user, true
	}

	return data_user, false
	// isValidToken, isValidExpired := utils.DecryptToken(data_user.User_session)
	// log.Println(isValidExpired)
	// log.Println(isValidToken)
}