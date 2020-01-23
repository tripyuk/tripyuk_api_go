package model

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name              string         `gorm:"column:name" json:"name"`
	Email             string         `gorm:"column:email" json:"email"`
	Password          string         `gorm:"column:password" json:"password"`
	RememberToken     string         `gorm:"column:remember_token" json:"remember_token"`
	LoginAt           mysql.NullTime `gorm:"column:login_at" json:"login_at"`
	IsVerified        bool           `gorm:"column:is_verified" json:"is_verified"`
	VerificationToken string         `gorm:"column:verification_token" json:"verification_token"`
	VerifiedAt        mysql.NullTime `gorm:"column:verified_at" json:"verified_at"`
	Handphone         string         `gorm:"column:handphone; size:15" json:"handphone"`
}

func (User) TableName() string {
	return "users"
}
