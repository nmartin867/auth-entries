package models

import (
	"gorm.io/gorm"
	"time"
)

type AuthEntry struct {
	gorm.Model
	TimeStamp time.Time
	Status    bool
}
