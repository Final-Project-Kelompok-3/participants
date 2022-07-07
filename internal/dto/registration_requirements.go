package dto

type CreateRegistrationRequirementsRequest struct {
	RegistrationPeriodsID int `json:"registration_period_id" validate:"required"`
	FileRequirementsID    int `json:"file_requirement_id" validate:"required"`
}

type UpdateRegistrationRequirementsRequest struct {
	RegistrationPeriodsID int `json:"registration_period_id" validate:"required"`
	FileRequirementsID    int `json:"file_requirement_id" validate:"required"`
}

type RegistrationRequirementsResponse struct {
	ID                    int `json:"id"`
	RegistrationPeriodsID int `json:"registration_period_id" validate:"required"`
	FileRequirementsID    int `json:"file_requirement_id" validate:"required"`
}
