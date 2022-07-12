package model

import (
	"time"

	"gorm.io/gorm"
)

type FileRequirements struct {
	Model

	Name          string `json:"name" gorm:"size:40;not null;unique"`
	FileExtension string `json:"file_extension" gorm:"size:5;not null"`
}

// BeforeCreate is a method for struct Role
// gorm call this method before they execute query
func (u *FileRequirements) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct Role
// gorm call this method before they execute query
func (u *FileRequirements) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
