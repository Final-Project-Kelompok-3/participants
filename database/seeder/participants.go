package seeder

import (
	"log"

	"github.com/Final-Project-Kelompok-3/participants/internal/model"
	"gorm.io/gorm"
)

func participantsSeeder(conn *gorm.DB) {

	var participants = []model.Participants{
		{Name: "John Seeder", Address: "Jl. Address1 Universe1", NISN: "12345678", FinalReportScore: 80, Email: "test1@email.com", FileRequirement: "Sertifikat1.pdf"},
		{Name: "Jack Great", Address: "Jl. Address2 Universe2", NISN: "23456789", FinalReportScore: 82, Email: "test2@email.com", FileRequirement: "Sertifikat2.pdf"},
		{Name: "Glenn Hopper", Address: "Jl. Address3 Universe3", NISN: "34567890", FinalReportScore: 83, Email: "test2@email.com", FileRequirement: "Sertifikat3.pdf"},
	}

	if err := conn.Create(&participants).Error; err != nil {
		log.Printf("cannot seed data participants, with error %v\n", err)
	}
	log.Println("success seed data participants")
}
