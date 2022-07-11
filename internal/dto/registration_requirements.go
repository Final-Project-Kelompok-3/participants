package dto

type CreateRegistrationRequirementsRequest struct {
	RegistrationPeriodsID uint `json:"registration_periods_id" validate:"required"`
	FileRequirementsID    uint `json:"file_requirements_id" validate:"required"`
}

type UpdateRegistrationRequirementsRequest struct {
	RegistrationPeriodsID *uint `json:"registration_periods_id"`
	FileRequirementsID    *uint `json:"file_requirements_id"`
}

type RegistrationRequirementsResponse struct {
	ID                    uint `json:"id"`
	RegistrationPeriodsID uint `json:"registration_periods_id" validate:"required"`
	FileRequirementsID    uint `json:"file_requirements_id" validate:"required"`
}
