package model

import (
	"time"

	"gorm.io/gorm"
)

type RegistrationRequirements struct {
	Model
	RegistrationPeriodsID int                 `json:"registration_period_id" gorm:"size:3;not null"`
	FileRequirementsID    int                 `json:"file_requirement_id" gorm:"size:3;not null;unique"`
	RegistrationPeriods   RegistrationPeriods `gorm:"references:RegistrationPeriods;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FileRequirements      FileRequirements    `gorm:"references:FileRequirements;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
