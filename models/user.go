package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName   string `json:"full_name"`
	PassportID string `json:"passport_id"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Password   string `json:"password"`
}
