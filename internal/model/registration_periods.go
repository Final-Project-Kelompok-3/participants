package model

import (
	"time"

	"gorm.io/gorm"
)

type RegistrationPeriods struct {
	Model

	SchoolsID            int       `json:"schools_id" gorm:"size:5;not null;unique"`
	RegistrationPricesID int       `json:"registration_prices_id" gorm:"size:5;not null;unique"`
	StartDate            time.Time `json:"start_date" gorm:"size:30;not null"`
	EndDate              time.Time `json:"end_date" gorm:"size:30;not null"`
	Description          string    `json:"description" gorm:"size:300;not null"`

	Schools            Schools            `gorm:"foreignKey:schools_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RegistrationPrices RegistrationPrices `gorm:"foreignKey:registration_prices_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// BeforeCreate is a method for struct Role
// gorm call this method before they execute query
func (u *RegistrationPeriods) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct Role
// gorm call this method before they execute query
func (u *RegistrationPeriods) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
