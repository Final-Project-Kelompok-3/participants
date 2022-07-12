package seeder

import (
	"log"

	"github.com/Final-Project-Kelompok-3/participants/internal/model"
	"gorm.io/gorm"
)

func levelsSeeder(conn *gorm.DB) {

	var levels = []model.Levels{
		{Name: "SD"},
		{Name: "SMP"},
		{Name: "SMA"},
	}

	if err := conn.Create(&levels).Error; err != nil {
		log.Printf("cannot seed data levels, with error %v\n", err)
	}
	log.Println("success seed data levels")
}
