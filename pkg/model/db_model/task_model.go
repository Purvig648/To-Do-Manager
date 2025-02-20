package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskName        string
	TaskDescription string
	TaskDeadline    time.Time
	TaskStatus      string
	UserID          uint
	User            User `gorm:"foreignKey:UserID"`
}
