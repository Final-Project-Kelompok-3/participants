package seeder

import (
	"log"

	"github.com/Final-Project-Kelompok-3/participants/internal/model"
	"gorm.io/gorm"
)

func schoolsSeeder(conn *gorm.DB) {

	var schools = []model.Schools{
		{LevelsID: 1, Name: "SD 1 Universe 1", Address: "Jl. Address 1 Universe 1"},
		{LevelsID: 2, Name: "SMP 1 Universe 1", Address: "Jl. Address 1 Universe 1"},
		{LevelsID: 3, Name: "SMA 1 Universe 1", Address: "Jl. Address 1 Universe 1"},
	}

	if err := conn.Create(&schools).Error; err != nil {
		log.Printf("cannot seed data schools, with error %v\n", err)
	}
	log.Println("success seed data schools")
}
