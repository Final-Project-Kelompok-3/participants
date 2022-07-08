package model

import (
	"time"

	"gorm.io/gorm"
)

type Levels struct {
	Model

	Name string `json:"name" gorm:"size:100;not null;unique"`
}

// BeforeCreate iså a method for struct Role
// gorm call this method before they execute query
func (u *Levels) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct Role
// gorm call this method before they execute query
func (u *Levels) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
