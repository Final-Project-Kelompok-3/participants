package model

import (
	"time"

	"gorm.io/gorm"
)

type Schools struct {
	Model

	LevelsID uint   `json:"levels_id" gorm:"not null"`
	Name     string `json:"name" gorm:"size:255;not null"`
	Address  string `json:"address" gorm:"size:255;not null"`

	// Levels Levels
}

// BeforeCreate iså a method for struct Role
// gorm call this method before they execute query
func (u *Schools) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct Role
// gorm call this method before they execute query
func (u *Schools) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
