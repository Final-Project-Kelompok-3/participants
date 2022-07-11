package dto

type CreateRegistrationRequirementsRequest struct {
	RegistrationPeriodsID uint `json:"registration_period_id" validate:"required"`
	FileRequirementsID    uint `json:"file_requirement_id" validate:"required"`
}

type UpdateRegistrationRequirementsRequest struct {
	RegistrationPeriodsID *uint `json:"registration_period_id" validate:"required"`
	FileRequirementsID    *uint `json:"file_requirement_id" validate:"required"`
}

type RegistrationRequirementsResponse struct {
	ID                    uint `json:"id"`
	RegistrationPeriodsID uint `json:"registration_period_id" validate:"required"`
	FileRequirementsID    uint `json:"file_requirement_id" validate:"required"`
}
