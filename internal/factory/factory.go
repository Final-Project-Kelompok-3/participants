package factory

import (
	"github.com/Final-Project-Kelompok-3/participants/internal/repository"
	"gorm.io/gorm"
)

type Factory struct {
	FileRequirementsRepository         repository.FileRequirements
	LevelsRepository                   repository.Levels
	ParticipantsInfoRepository         repository.ParticipantsInfo
	ParticipantsRepository             repository.Participants
	RegistrationPeriodsRepository      repository.RegistrationPeriods
	RegistrationPricesRepository       repository.RegistrationPrices
	RegistrationRequirementsRepository repository.RegistrationRequirements
	SchoolsRepository                  repository.Schools
}

func NewFactory(db *gorm.DB) *Factory {
	return &Factory{
		FileRequirementsRepository:         repository.NewFileRequirements(db),
		LevelsRepository:                   repository.NewLevels(db),
		ParticipantsInfoRepository:         repository.NewParticipantInfo(db),
		ParticipantsRepository:             repository.NewParticipants(db),
		RegistrationPeriodsRepository:      repository.NewRegistrationPeriods(db),
		RegistrationPricesRepository:       repository.NewRegistrationPrices(db),
		RegistrationRequirementsRepository: repository.NewRegistrationRequirements(db),
		SchoolsRepository:                  repository.NewSchools(db),
	}
}
