package factory

import (
	"github.com/Final-Project-Kelompok-3/participants/internal/repository"
	"gorm.io/gorm"
)

type Factory struct {
	// RoleRepository         repository.Role
	// UserRepository         repository.User
	ParticipantsRepository repository.Participants
}

func NewFactory(db *gorm.DB) *Factory {
	return &Factory{
		// RoleRepository:         repository.NewRole(db),
		// UserRepository:         repository.NewUser(db),
		ParticipantsRepository: repository.NewParticipants(db),
	}
}
