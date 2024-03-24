package model

import (
	"MyGram/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for an User
type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age          int           `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,range(9|100)~Age has to be above 8 years old"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL;" json:"photos"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL;" json:"comments"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL;" json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.HashPass(u.Password)
	err = nil
	return
}
