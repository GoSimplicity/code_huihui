package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;size:50;index" json:"username"`
	Password string `gorm:"not null;size:255" json:"-"`
	Email    string `gorm:"unique;not null;size:100;index" json:"email"`
	Phone    string `gorm:"unique;not null;size:20;index" json:"phone"`
	Avatar   string `gorm:"size:255;default:''" json:"avatar"`
	Age      int    `gorm:"default:20" json:"age"`
	Gender   int8   `gorm:"default:1;comment:'1:男,2:女'" json:"gender"`
	Status   int8   `gorm:"default:1;comment:'1:正常,2:禁用'" json:"status"`
}

func (User) TableName() string {
	return "ch_user"
}
