package dbmodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username       string
	EmailID        string `gorm:"unique"`
	HashedPassword string
}
