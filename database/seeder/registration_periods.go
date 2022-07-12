package seeder

import (
	"log"
	"time"

	"github.com/Final-Project-Kelompok-3/participants/internal/model"
	"gorm.io/gorm"
)

func registrationPeriodsSeeder(conn *gorm.DB) {
	var layout string = "2006-01-02"
	start_date_from_str, err := time.Parse(layout, "2022-07-08")
	if err != nil {
		log.Println(err)
	}
	end_date_from_str, err := time.Parse(layout, "2022-07-10")
	if err != nil {
		log.Println(err)
	}
	var registration_periods = []model.RegistrationPeriods{
		{SchoolsID: 1, RegistrationPricesID: 1, StartDate: start_date_from_str, EndDate: end_date_from_str, Description: "Sertifikat1.pdf"},
		{SchoolsID: 2, RegistrationPricesID: 2, StartDate: start_date_from_str, EndDate: end_date_from_str, Description: "Sertifikat2.pdf"},
		{SchoolsID: 3, RegistrationPricesID: 3, StartDate: start_date_from_str, EndDate: end_date_from_str, Description: "Sertifikat3.pdf"},
	}

	if err := conn.Create(&registration_periods).Error; err != nil {
		log.Printf("cannot seed data registration_periods, with error %v\n", err)
	}
	log.Println("success seed data registration_periods")
}
