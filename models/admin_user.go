package models

import (
	"time"
	"github.com/iqysf/gorm"
)

type AdminUser struct {
	gorm.Model
	Username      string `gorm:"index"`
	Password      string `gorm:"index"`
	LastLoginTime time.Time
	LastLoginIP   string
}
