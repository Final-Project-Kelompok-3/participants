package seeder

import "github.com/Final-Project-Kelompok-3/participants/database"

func Seed() {

	conn := database.GetConnection()

	fileRequirementSeeder(conn)
	levelsSeeder(conn)
	// otherTableSeeder(conn)
}
