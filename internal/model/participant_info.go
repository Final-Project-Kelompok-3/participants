package model

import (
	"time"

	"gorm.io/gorm"
)

type ParticipantInfo struct {
	Model

	NISN             string `json:"nisn" gorm:"size:40;not null;unique"`
	FinalReportScore int    `json:"final_report_score" gorm:"size:3;not null"`
	Email            string `json:"email" validate:"email" gorm:"size:50;not null;unique"`
	FileRequirement  string `json:"file_requirement" gorm:"size:200;not null;unique"`
}

// BeforeCreate is a method for struct Role
// gorm call this method before they execute query
func (u *ParticipantInfo) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct Role
// gorm call this method before they execute query
func (u *ParticipantInfo) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
