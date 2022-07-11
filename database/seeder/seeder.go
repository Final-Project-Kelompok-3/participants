package seeder

import "github.com/Final-Project-Kelompok-3/participants/database"

func Seed() {

	conn := database.GetConnection()

	fileRequirementSeeder(conn)
	levelsSeeder(conn)
	participantInfoSeeder(conn)
	participantsSeeder(conn)
	registrationPeriodsSeeder(conn)
	registrationPricesSeeder(conn)
	registrationRequirementsSeeder(conn)
	schoolsSeeder(conn)

	// otherTableSeeder(conn)

	// &model.FileRequirements{},
	// &model.Levels{},
	// &model.ParticipantInfo{},
	// &model.Participants{},
	// // &model.Schools{},
	// &model.RegistrationPrices{},

	// // &model.RegistrationPeriods{},
	// // &model.RegistrationRequirements{},
}
