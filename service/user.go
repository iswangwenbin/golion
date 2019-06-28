package service

import (
	"github.com/iswangwenbin/golion/model"
	"github.com/iswangwenbin/golion/pkg/databasex"
	"github.com/jinzhu/gorm"
)

func GetOneUser() *model.User {
	user := new(model.User)
	user.UserId = 10001
	user.Username = "George"
	return user
}

func GetByUserId(uid int) (model.User, error) {
	db := databasex.GetDatabase()
	user := model.User{}
	if err := db.Debug().Where("user_id = ?", uid).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return user, nil
	}
	return user, nil
}
