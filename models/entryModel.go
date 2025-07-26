package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	TimeStamp time.Time
	Status    bool
}
