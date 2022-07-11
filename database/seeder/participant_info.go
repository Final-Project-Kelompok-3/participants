package seeder

import (
	"log"

	"github.com/Final-Project-Kelompok-3/participants/internal/model"
	"gorm.io/gorm"
)

func participantInfoSeeder(conn *gorm.DB) {

	var participant_info = []model.ParticipantInfo{
		{NISN: "12345678", FinalReportScore: 80, Email: "test1@email.com", FileRequirement: "Sertifikat1.pdf"},
		{NISN: "23456789", FinalReportScore: 82, Email: "test2@email.com", FileRequirement: "Sertifikat2.pdf"},
		{NISN: "34567890", FinalReportScore: 83, Email: "test2@email.com", FileRequirement: "Sertifikat3.pdf"},
	}

	if err := conn.Create(&participant_info).Error; err != nil {
		log.Printf("cannot seed data participant_info, with error %v\n", err)
	}
	log.Println("success seed data participant_info")
}
