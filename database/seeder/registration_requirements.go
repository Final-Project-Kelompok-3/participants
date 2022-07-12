package seeder

import (
	"log"

	"github.com/Final-Project-Kelompok-3/participants/internal/model"
	"gorm.io/gorm"
)

func registrationRequirementsSeeder(conn *gorm.DB) {

	var registration_requirements = []model.RegistrationRequirements{
		{RegistrationPeriodsID: 1, FileRequirementsID: 1},
		{RegistrationPeriodsID: 2, FileRequirementsID: 2},
		{RegistrationPeriodsID: 3, FileRequirementsID: 3},
	}

	if err := conn.Create(&registration_requirements).Error; err != nil {
		log.Printf("cannot seed data registration_requirements, with error %v\n", err)
	}
	log.Println("success seed data registration_requirements")
}
