package model

import (
	"time"

	"gorm.io/gorm"
)

type RegistrationPrices struct {
	Model

	RegistrationFee int `json:"registration_fee" gorm:"size:10;not null"`
	SystemChange    int `json:"system_change" gorm:"size:10;not null;unique"`
	Total           int `json:"total" gorm:"size:10;not null;unique"`
}

// BeforeCreate is a method for struct Role
// gorm call this method before they execute query
func (u *RegistrationPrices) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct Role
// gorm call this method before they execute query
func (u *RegistrationPrices) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
