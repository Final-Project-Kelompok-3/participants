package model

import (
	"time"

	"gorm.io/gorm"
)

type RegistrationRequirements struct {
	Model
	RegistrationPeriodID int `json:"final_report_score" gorm:"size:3;not null"`
	FileRequirementID    int `json:"nisn" gorm:"size:40;not null;unique"`
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
