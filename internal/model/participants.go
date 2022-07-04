package model

import (
	"time"

	"gorm.io/gorm"
)

type Participants struct {
	Model

	Name             string `json:"name" gorm:"size:200;not null;unique"`
	Address          string `json:"address" gorm:"size:200;not null;unique"`
	NISN             string `json:"nisn" gorm:"size:200;not null;unique"`
	FinalReportScore int    `json:"final_report_score" gorm:"size:200;not null;unique"`
	Email            string `json:"email" validate:"email" gorm:"size:200;not null;unique"`
	FileRequirement  string `json:"file_requirement" gorm:"size:200;not null;unique"`
}

// BeforeCreate is a method for struct Role
// gorm call this method before they execute query
func (u *Participants) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct Role
// gorm call this method before they execute query
func (u *Participants) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
