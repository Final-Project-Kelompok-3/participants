package seeder

import (
	"log"

	"github.com/Final-Project-Kelompok-3/participants/internal/model"
	"gorm.io/gorm"
)

func registrationPricesSeeder(conn *gorm.DB) {

	var registration_prices = []model.RegistrationPrices{
		{RegistrationFee: 100, SystemChange: 1, Total: 1234567},
		{RegistrationFee: 200, SystemChange: 2, Total: 2345678},
		{RegistrationFee: 300, SystemChange: 3, Total: 3456789},
	}

	if err := conn.Create(&registration_prices).Error; err != nil {
		log.Printf("cannot seed data registration_prices, with error %v\n", err)
	}
	log.Println("success seed data registration_prices")
}
