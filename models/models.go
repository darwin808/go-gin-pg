package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Username string `json:"username" gorm:"text;not null;default:null`
	Email    string `json:"email" gorm:"text;not null;default:null`
	Password string `json:"password" gorm:"text;not null;default:null`
}
