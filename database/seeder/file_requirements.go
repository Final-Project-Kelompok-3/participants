package seeder

import (
	"log"

	"github.com/Final-Project-Kelompok-3/participants/internal/model"
	"gorm.io/gorm"
)

func fileRequirementSeeder(conn *gorm.DB) {

	var file_requirements = []model.FileRequirements{
		{Name: "Sertifikat1", FileExtension: "pdf"},
		{Name: "Sertifikat2", FileExtension: "pdf"},
		{Name: "Sertifikat3", FileExtension: "pdf"},
	}

	if err := conn.Create(&file_requirements).Error; err != nil {
		log.Printf("cannot seed data file_requirements, with error %v\n", err)
	}
	log.Println("success seed data file_requirements")
}