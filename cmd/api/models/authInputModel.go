package models

import "gorm.io/gorm"

type AuthInput struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}
