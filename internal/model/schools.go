package model

import (
	"time"

	"gorm.io/gorm"
)

type Schools struct {
	Model

	LevelsID int    `json:"levels_id" gorm:"size:3;not null;unique"`
	Name     string `json:"name" gorm:"size:255;not null;unique"`
	Address  string `json:"address" gorm:"size:255;not null"`

	Levels Levels `gorm:"references:Levels;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
