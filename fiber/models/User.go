package models

import "gorm.io/gorm"

// User is a model for user
type User struct {
	gorm.Model
	UID      string `json:"uid" example:"1"`
	USERNAME string `json:"username" example:"Alice"`
	GENDER   string `json:"gender" example:"Male, Female"`
}
