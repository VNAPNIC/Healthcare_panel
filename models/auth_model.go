package model

import (
	_ "gorm.io/gorm"
	"serverhealthcarepanel/utils"
	"time"
)

type Auth struct {
	BaseModel
	RoleId     uint      `gorm:"DEFAULT:0;NOT NULL;" json:"role_id"`
	Status     int       `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"status"`
	Username   string    `gorm:"Size:20;uniqueIndex;NOT NULL;" json:"user_name"`
	Password   string    `gorm:"Size:50;NOT NULL;" json:"-"`
	LoggedInAt time.Time `gorm:"type:datetime" json:"logged_in_at"`
	Role       Role      `gorm:"references:RoleId;"`
}

func (Auth) TableName() string {
	return TablePrefix + "auth"
}

func CreateUser(auth Auth) error {
	res := db.Create(&auth)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func CheckAuth(username string, password string) (error, bool, Auth) {
	var auth Auth

	res := db.Where(Auth{
		Username: username,
		Password: utils.EncodeUserPassword(password),
	}).Preload("Role").First(&auth)

	if err := res.Error; err != nil {
		return err, false, Auth{}
	}

	if auth.ID <= 0 {
		return nil, false, Auth{}
	}

	return nil, true, auth
}
