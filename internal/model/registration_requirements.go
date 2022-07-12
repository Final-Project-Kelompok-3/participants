package model

import (
	"time"

	"gorm.io/gorm"
)

type RegistrationRequirements struct {
	Model
	RegistrationPeriodsID uint `json:"registration_periods_id" gorm:"not null"`
	FileRequirementsID    uint `json:"file_requirements_id" gorm:"not null"`
	// RegistrationPeriods   RegistrationPeriods
	// FileRequirements      FileRequirements
}

// BeforeCreate is a method for struct Role
// gorm call this method before they execute query
func (u *RegistrationRequirements) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct Role
// gorm call this method before they execute query
func (u *RegistrationRequirements) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
